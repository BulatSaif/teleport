/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/gravitational/form"
	"github.com/gravitational/log"
	"github.com/gravitational/roundtrip"
	"github.com/gravitational/session"
	"github.com/gravitational/teleport/lib/reversetunnel"
	"github.com/gravitational/teleport/lib/utils"
	"github.com/gravitational/trace"
	"github.com/julienschmidt/httprouter"
	"github.com/mailgun/ttlmap"
)

type MultiSiteHandler struct {
	httprouter.Router
	cfg       MultiSiteConfig
	auth      AuthHandler
	sites     *ttlmap.TtlMap
	templates map[string]*template.Template
}

type MultiSiteConfig struct {
	Tun        reversetunnel.Server
	AssetsDir  string
	AuthAddr   utils.NetAddr
	DomainName string
}

func NewMultiSiteHandler(cfg MultiSiteConfig) (*MultiSiteHandler, error) {
	lauth, err := NewLocalAuth(cfg.DomainName, []utils.NetAddr{cfg.AuthAddr})
	if err != nil {
		return nil, err
	}

	sites, err := ttlmap.NewMap(1024)
	if err != nil {
		return nil, err
	}

	h := &MultiSiteHandler{
		cfg:   cfg,
		auth:  lauth,
		sites: sites,
	}

	h.initTemplates(cfg.AssetsDir)

	// WEB views
	h.GET("/web/newuser/:token", h.newUser)
	h.POST("/web/finishnewuser", h.finishNewUser)
	h.GET("/web/login", h.login)
	h.GET("/web/loginerror", h.loginError)
	h.GET("/web/loginaftercreation", h.loginAfterCreation)
	h.GET("/web/logout", h.logout)
	h.POST("/web/auth", h.authForm)
	h.GET("/web/error", h.errorPage)

	h.GET("/", h.needsAuth(h.sitesIndex))
	h.GET("/web/sites", h.needsAuth(h.sitesIndex))

	// For ssh proxy
	h.POST("/sshlogin", h.loginSSHProxy)

	// Forward all requests to site handler
	sh := h.needsAuth(h.siteHandler)
	h.GET("/tun/:site/*path", sh)
	h.PUT("/tun/:site/*path", sh)
	h.POST("/tun/:site/*path", sh)
	h.DELETE("/tun/:site/*path", sh)

	// API views
	h.GET("/api/sites", h.needsAuth(h.handleGetSites))

	// Static assets
	h.Handler("GET", "/static/*filepath",
		http.FileServer(http.Dir(filepath.Join(cfg.AssetsDir, "assets"))))
	return h, nil
}

func (s *MultiSiteHandler) initTemplates(baseDir string) {
	tpls := []tpl{
		tpl{name: "login", include: []string{"assets/static/tpl/login.tpl", "assets/static/tpl/base.tpl"}},
		tpl{name: "error", include: []string{"assets/static/tpl/error.tpl", "assets/static/tpl/base.tpl"}},
		tpl{name: "newuser", include: []string{"assets/static/tpl/newuser.tpl", "assets/static/tpl/base.tpl"}},
		tpl{name: "sites", include: []string{"assets/static/tpl/sites.tpl", "assets/static/tpl/base.tpl"}},
		tpl{name: "site-servers", include: []string{"assets/static/tpl/site/servers.tpl", "assets/static/tpl/base.tpl"}},
		tpl{name: "site-events", include: []string{"assets/static/tpl/site/events.tpl", "assets/static/tpl/base.tpl"}},
	}
	s.templates = make(map[string]*template.Template)
	for _, t := range tpls {
		tpl := template.New(t.name)
		for _, i := range t.include {
			template.Must(tpl.ParseFiles(filepath.Join(baseDir, i)))
		}
		s.templates[t.name] = tpl
	}
}

func (h *MultiSiteHandler) String() string {
	return fmt.Sprintf("multi site")
}

func (h *MultiSiteHandler) newUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	token := p[0].Value
	user, QRImg, _, err := h.auth.NewUserForm(token)
	if err != nil {
		http.Redirect(w, r, ErrorPageLink("Signup link had expired"),
			http.StatusFound)
		return
	}

	base64QRImg := base64.StdEncoding.EncodeToString(QRImg)
	h.executeTemplate(w, "newuser", map[string]interface{}{
		"Token":    token,
		"Username": user,
		"QR":       base64QRImg})
}

func (h *MultiSiteHandler) finishNewUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var token, pass, pass2, hotpToken string

	err := form.Parse(r,
		form.String("token", &token, form.Required()),
		form.String("password", &pass, form.Required()),
		form.String("password_confirm", &pass2, form.Required()),
		form.String("hotp_token", &hotpToken, form.Required()),
	)

	if err != nil {
		http.Redirect(w, r, ErrorPageLink("Error: "+err.Error()),
			http.StatusFound)
		return
	}

	if pass != pass2 {
		http.Redirect(w, r, ErrorPageLink("Provided passwords mismatch"),
			http.StatusFound)
		return
	}

	err = h.auth.NewUserFinish(token, pass, hotpToken)
	if err != nil {
		if strings.Contains(err.Error(), "Wrong HOTP token") {
			http.Redirect(w, r, ErrorPageLink("Wrong HOTP token"),
				http.StatusFound)
		} else {
			http.Redirect(w, r, ErrorPageLink("Error: "+err.Error()),
				http.StatusFound)
		}
		return
	}

	http.Redirect(w, r, "/web/loginaftercreation", http.StatusFound)
}

func (h *MultiSiteHandler) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.executeTemplate(w, "login", nil)
}

func (h *MultiSiteHandler) loginError(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.executeTemplate(w, "login", map[string]interface{}{"ErrorString": "Wrong username or password or hotp token"})
}

func (h *MultiSiteHandler) loginAfterCreation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.executeTemplate(w, "login", map[string]interface{}{"InfoString": "Account was successfully created, you can login"})
}

func (h *MultiSiteHandler) errorPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	errorString := r.URL.Query().Get("message")
	h.executeTemplate(w, "error", map[string]interface{}{"ErrorString": errorString})
}

func ErrorPageLink(message string) string {
	return "/web/error?message=" + url.QueryEscape(message)
}

func (h *MultiSiteHandler) logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := session.ClearSession(w, h.cfg.DomainName); err != nil {
		log.Errorf("failed to clear session: %v", err)
		replyErr(w, http.StatusInternalServerError, fmt.Errorf("failed to logout"))
		return
	}
	http.Redirect(w, r, "/web/login", http.StatusFound)
}

func (h *MultiSiteHandler) authForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user, pass, hotpToken string

	err := form.Parse(r,
		form.String("username", &user, form.Required()),
		form.String("password", &pass, form.Required()),
		form.String("hotpToken", &hotpToken, form.Required()),
	)

	if err != nil {
		replyErr(w, http.StatusBadRequest, err)
		return
	}
	sid, err := h.auth.Auth(user, pass, hotpToken)
	if err != nil {
		log.Warningf("auth error: %v", err)
		http.Redirect(w, r, "/web/loginerror", http.StatusFound)
		return
	}
	if err := session.SetSession(w, h.cfg.DomainName, user, sid); err != nil {
		replyErr(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *MultiSiteHandler) loginSSHProxy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var credJSON string

	err := form.Parse(r,
		form.String("credentials", &credJSON, form.Required()),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(trace.Wrap(err).Error()))
		return
	}

	var cred SSHLoginCredentials
	if err := json.Unmarshal([]byte(credJSON), &cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(trace.Wrap(err).Error()))
		return
	}

	cert, err := h.auth.GetCertificate(cred)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(trace.Wrap(err).Error()))
		return
	}
	w.Write(cert)
}

func (s *MultiSiteHandler) siteEvents(w http.ResponseWriter, r *http.Request, p httprouter.Params, c Context) error {
	s.executeTemplate(w, "site-events", map[string]interface{}{"SiteName": p[0].Value})
	return nil
}

func (s *MultiSiteHandler) siteServers(w http.ResponseWriter, r *http.Request, p httprouter.Params, c Context) error {
	s.executeTemplate(w, "site-servers", map[string]interface{}{"SiteName": p[0].Value})
	return nil
}

func (s *MultiSiteHandler) sitesIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c Context) error {
	s.executeTemplate(w, "sites", nil)
	return nil
}

func (s *MultiSiteHandler) siteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, c Context) error {
	siteName := p[0].Value
	prefix := fmt.Sprintf("/tun/%v", siteName)
	i, ok := s.sites.Get(siteName)
	if !ok {
		tauth, err := NewTunAuth(s.auth, s.cfg.Tun, siteName)
		if err != nil {
			return err
		}
		i = NewSiteHandler(SiteHandlerConfig{
			Auth:      tauth,
			AssetsDir: s.cfg.AssetsDir,
			URLPrefix: prefix,
			NavSections: []NavSection{
				NavSection{
					Title: "Back to Portal",
					Icon:  "fa fa-arrow-circle-left",
					URL:   "/",
				},
			},
		})
		if err := s.sites.Set(siteName, i, 90); err != nil {
			return err
		}
	}
	sh := i.(http.Handler)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
	r.RequestURI = r.URL.String()
	log.Infof("siteHandler: %v %v", r.Method, r.URL)
	sh.ServeHTTP(w, r)
	return nil
}

func (h *MultiSiteHandler) handleGetSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c Context) error {
	roundtrip.ReplyJSON(w, http.StatusOK, sitesResponse(h.cfg.Tun.GetSites()))
	return nil
}

func (h *MultiSiteHandler) needsAuth(fn authHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Infof("getting cookie: %v", err)
			http.Redirect(w, r, "/web/login", http.StatusFound)
			return
		}
		d, err := session.DecodeCookie(cookie.Value)
		if err != nil {
			log.Warningf("failed to decode cookie '%v', err: %v", cookie.Value, err)
			http.Redirect(w, r, "/web/login", http.StatusFound)
			return
		}
		ctx, err := h.auth.ValidateSession(d.User, d.SID)
		if err != nil {
			log.Warningf("failed to validate session: %v", err)
			http.Redirect(w, r, "/web/login", http.StatusFound)
			return
		}
		if err := fn(w, r, p, ctx); err != nil {
			log.Errorf("error in handler: %v", err)
			roundtrip.ReplyJSON(
				w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func (h *MultiSiteHandler) executeTemplate(w http.ResponseWriter, name string, data interface{}) {
	tpl, ok := h.templates[name]
	if !ok {
		replyErr(w, http.StatusInternalServerError, fmt.Errorf("template '%v' not found", name))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Errorf("Execute template: %v", err)
		replyErr(w, http.StatusInternalServerError, fmt.Errorf("internal render error"))
	}
}

type Server struct {
	http.Server
}

func New(addr utils.NetAddr, cfg MultiSiteConfig) (*Server, error) {
	h, err := NewMultiSiteHandler(cfg)
	if err != nil {
		return nil, err
	}
	srv := &Server{}
	srv.Server.Addr = addr.Addr
	srv.Server.Handler = h
	return srv, nil
}

type site struct {
	Name          string    `json:"name"`
	LastConnected time.Time `json:"last_connected"`
	Status        string    `json:"status"`
}

func sitesResponse(rs []reversetunnel.RemoteSite) []site {
	out := make([]site, len(rs))
	for i := range rs {
		out[i] = site{
			Name:          rs[i].GetName(),
			LastConnected: rs[i].GetLastConnected(),
			Status:        rs[i].GetStatus(),
		}
	}
	return out
}

func CreateSignupLink(hostPort string, token string) string {
	return "http://" + hostPort + "/web/newuser/" + token
}

type authHandle func(http.ResponseWriter, *http.Request, httprouter.Params, Context) error
