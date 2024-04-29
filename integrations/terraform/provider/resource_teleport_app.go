// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2022 Gravitational, Inc.

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

package provider

import (
	"context"
	"fmt"

	apitypes "github.com/gravitational/teleport/api/types"
	
	"github.com/gravitational/teleport/integrations/lib/backoff"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jonboulle/clockwork"

	"github.com/gravitational/teleport/integrations/terraform/tfschema"
)

// resourceTeleportAppType is the resource metadata type
type resourceTeleportAppType struct{}

// resourceTeleportApp is the resource
type resourceTeleportApp struct {
	p Provider
}

// GetSchema returns the resource schema
func (r resourceTeleportAppType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaAppV3(ctx)
}

// NewResource creates the empty resource
func (r resourceTeleportAppType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTeleportApp{
		p: *(p.(*Provider)),
	}, nil
}

// Create creates the App
func (r resourceTeleportApp) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var err error
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	app := &apitypes.AppV3{}
	diags = tfschema.CopyAppV3FromTerraform(ctx, plan, app)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	
	appResource := app

	err = appResource.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error setting App defaults", trace.Wrap(err), "app"))
		return
	}

	id := appResource.Metadata.Name

	_, err = r.p.Client.GetApp(ctx, id)
	if !trace.IsNotFound(err) {
		if err == nil {
			existErr := fmt.Sprintf("App exists in Teleport. Either remove it (tctl rm app/%v)"+
				" or import it to the existing state (terraform import teleport_app.%v %v)", id, id, id)

			resp.Diagnostics.Append(diagFromErr("App exists in Teleport", trace.Errorf(existErr)))
			return
		}

		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(err), "app"))
		return
	}

	err = r.p.Client.CreateApp(ctx, appResource)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error creating App", trace.Wrap(err), "app"))
		return
	}
		
	// Not really an inferface, just using the same name for easier templating.
	var appI apitypes.Application
	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		appI, err = r.p.Client.GetApp(ctx, id)
		if trace.IsNotFound(err) {
			if bErr := backoff.Do(ctx); bErr != nil {
				resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(bErr), "app"))
				return
			}
			if tries >= r.p.RetryConfig.MaxTries {
				diagMessage := fmt.Sprintf("Error reading App (tried %d times) - state outdated, please import resource", tries)
				resp.Diagnostics.AddError(diagMessage, "app")
			}
			continue
		}
		break
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(err), "app"))
		return
	}

	appResource, ok := appI.(*apitypes.AppV3)
	if !ok {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Errorf("Can not convert %T to AppV3", appI), "app"))
		return
	}
	app = appResource

	diags = tfschema.CopyAppV3ToTerraform(ctx, app, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Attrs["id"] = types.String{Value: app.Metadata.Name}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read reads teleport App
func (r resourceTeleportApp) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state types.Object
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.String
	diags = req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	appI, err := r.p.Client.GetApp(ctx, id.Value)
	if trace.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(err), "app"))
		return
	}
	
	app := appI.(*apitypes.AppV3)
	diags = tfschema.CopyAppV3ToTerraform(ctx, app, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates teleport App
func (r resourceTeleportApp) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	app := &apitypes.AppV3{}
	diags = tfschema.CopyAppV3FromTerraform(ctx, plan, app)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	appResource := app


	if err := appResource.CheckAndSetDefaults(); err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating App", err, "app"))
		return
	}
	name := appResource.Metadata.Name

	appBefore, err := r.p.Client.GetApp(ctx, name)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", err, "app"))
		return
	}

	err = r.p.Client.UpdateApp(ctx, appResource)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating App", err, "app"))
		return
	}
		
	// Not really an inferface, just using the same name for easier templating.
	var appI apitypes.Application

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		appI, err = r.p.Client.GetApp(ctx, name)
		if err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", err, "app"))
			return
		}
		if appBefore.GetMetadata().Revision != appI.GetMetadata().Revision || false {
			break
		}

		if err := backoff.Do(ctx); err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(err), "app"))
			return
		}
		if tries >= r.p.RetryConfig.MaxTries {
			diagMessage := fmt.Sprintf("Error reading App (tried %d times) - state outdated, please import resource", tries)
			resp.Diagnostics.AddError(diagMessage, "app")
			return
		}
	}

	appResource, ok := appI.(*apitypes.AppV3)
	if !ok {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Errorf("Can not convert %T to AppV3", appI), "app"))
		return
	}
	diags = tfschema.CopyAppV3ToTerraform(ctx, app, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes Teleport App
func (r resourceTeleportApp) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.p.Client.DeleteApp(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error deleting AppV3", trace.Wrap(err), "app"))
		return
	}

	resp.State.RemoveResource(ctx)
}

// ImportState imports App state
func (r resourceTeleportApp) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	app, err := r.p.Client.GetApp(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading App", trace.Wrap(err), "app"))
		return
	}

	
	appResource := app.(*apitypes.AppV3)

	var state types.Object

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = tfschema.CopyAppV3ToTerraform(ctx, appResource, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	id := appResource.GetName()

	state.Attrs["id"] = types.String{Value: id}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
