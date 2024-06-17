package entitlements

type EntitlementKind string

// The EntitlementKind list should be 1:1 with the Features & FeatureStrings in salescenter/product/product.go,
// except CustomTheme which is dropped. CustomTheme entitlement only toggles the ability to "set" a theme;
// the value of that theme, if set, is stored and accessed outside of entitlements.
const (
	AccessLists            EntitlementKind = "AccessLists"
	AccessMonitoring       EntitlementKind = "AccessMonitoring"
	AccessRequests         EntitlementKind = "AccessRequests"
	App                    EntitlementKind = "App"
	CloudAuditLogRetention EntitlementKind = "CloudAuditLogRetention"
	DB                     EntitlementKind = "DB"
	Desktop                EntitlementKind = "Desktop"
	DeviceTrust            EntitlementKind = "DeviceTrust"
	ExternalAuditStorage   EntitlementKind = "ExternalAuditStorage"
	FeatureHiding          EntitlementKind = "FeatureHiding"
	HSM                    EntitlementKind = "HSM"
	Identity               EntitlementKind = "Identity"
	JoinActiveSessions     EntitlementKind = "JoinActiveSessions"
	K8s                    EntitlementKind = "K8s"
	MobileDeviceManagement EntitlementKind = "MobileDeviceManagement"
	OIDC                   EntitlementKind = "OIDC"
	OktaSCIM               EntitlementKind = "OktaSCIM"
	OktaUserSync           EntitlementKind = "OktaUserSync"
	Policy                 EntitlementKind = "Policy"
	SAML                   EntitlementKind = "SAML"
	SessionLocks           EntitlementKind = "SessionLocks"
	UpsellAlert            EntitlementKind = "UpsellAlert"
	UsageReporting         EntitlementKind = "UsageReporting"
)
