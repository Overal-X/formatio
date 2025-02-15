package types

type CreateAppArgs struct {
	Name                  string `json:"name"`
	Url                   string `json:"url"`
	WebhookUrl            string `json:"webhook_url"`
	HookAttributes        string `json:"hook_attributes"`
	Description           string `json:"description"`
	CallbackUrl           string `json:"callback_url"`
	RequestOauthOnInstall bool   `json:"request_oauth_on_install"`
	Public                bool   `json:"public"`
	WebhookActive         bool   `json:"webhook_active"`
	RedirectUri           string `json:"redirect_uri"`
	SetupUrl              string `json:"setup_url"`
	SetupOnUpdate         bool   `json:"setup_on_update"`
	// Repository Permissions
	Actions              string `json:"actions"`
	Administration       string `json:"administration"`
	Checks               string `json:"checks"`
	Contents             string `json:"contents"`
	Deployments          string `json:"deployments"`
	Environments         string `json:"environments"`
	Issues               string `json:"issues"`
	Metadata             string `json:"metadata"`
	Packages             string `json:"packages"`
	Pages                string `json:"pages"`
	PullRequests         string `json:"pull_requests"`
	SecretScanningAlerts string `json:"secret_scanning_alerts"`
	Secrets              string `json:"secrets"`
	SecurityEvents       string `json:"security_events"`
	SingleFile           string `json:"single_file"`
	Statuses             string `json:"statuses"`
	VulnerabilityAlerts  string `json:"vulnerability_alerts"`
	Workflows            string `json:"workflows"`
	// Organization Permissions
	Members                       string `json:"members"`
	OrganizationAdmin             string `json:"organization_administration"`
	OrganizationHooks             string `json:"organization_hooks"`
	OrganizationPlan              string `json:"organization_plan"`
	OrganizationProjects          string `json:"organization_projects"`
	OrganizationSecrets           string `json:"organization_secrets"`
	OrganizationSelfHostedRunners string `json:"organization_self_hosted_runners"`
	OrganizationUserBlocking      string `json:"organization_user_blocking"`
	// User Permissions
	Email                string `json:"email"`
	FollowingUsers       string `json:"following"`
	GitSshKeys           string `json:"git_ssh_keys"`
	GpgKeys              string `json:"gpg_keys"`
	Interaction          string `json:"interaction_limits"`
	Profile              string `json:"profile"`
	StarredRepositories  string `json:"starring"`
	WatchingRepositories string `json:"watching"`
}

type CreateAppResult struct {
	AppUrl     string
	PrivateKey string
}

type HookAttributes struct {
	URL    string `json:"url"`              // Required. The URL of the webhook server.
	Active bool   `json:"active,omitempty"` // Optional. Defaults to true if not set.
}

type GitHubAppManifest struct {
	Name                  string            `json:"name"`                               // The name of the GitHub App.
	URL                   string            `json:"url"`                                // Required. The homepage of the GitHub App.
	HookAttributes        HookAttributes    `json:"hook_attributes,omitempty"`          // Webhook configuration.
	RedirectURL           string            `json:"redirect_url,omitempty"`             // Full URL to redirect to after registration.
	CallbackURLs          []string          `json:"callback_urls,omitempty"`            // List of callback URLs (max 10).
	SetupURL              string            `json:"setup_url,omitempty"`                // URL for additional setup after installation.
	Description           string            `json:"description,omitempty"`              // Description of the GitHub App.
	Public                bool              `json:"public,omitempty"`                   // Whether the app is public or private.
	DefaultEvents         []string          `json:"default_events,omitempty"`           // Events the GitHub App subscribes to.
	DefaultPermissions    map[string]string `json:"default_permissions,omitempty"`      // Permissions needed by the GitHub App.
	RequestOAuthOnInstall bool              `json:"request_oauth_on_install,omitempty"` // Request user to authorize the app upon install.
	SetupOnUpdate         bool              `json:"setup_on_update,omitempty"`          // Redirect users to setup_url after update.
	State                 string            `json:"state,omitempty"`                    // Used to protect against CSRF attacks.
}
