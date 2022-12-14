package grafana

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/grafana/terraform-provider-grafana/grafana"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumiverse/pulumi-grafana/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	grafanaPkg = "grafana"
	// modules:
	grafanaMod = "index" // the xyz module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// func boolRef(b bool) *bool {
// 	return &b
// }

// grafanaMember manufactures a type token for the Grafana package and the given module and type.
func grafanaMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(grafanaPkg + ":" + mod + ":" + mem)
}

// grafanaType manufactures a type token for the Grafana package and the given module and type.
func grafanaType(mod string, typ string) tokens.Type {
	return tokens.Type(grafanaMember(mod, typ))
}

// grafanaDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Grafana package and names the file by simply lower casing the data
// source's first character.
func grafanaDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return grafanaMember(mod+"/"+fn, res)
}

// grafanaType manufactures a standard resource token given a module and resource name.
// It automatically uses the Grafana package and names the file by simply lower casing the resource's
// first character.
func grafanaResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return grafanaType(mod+"/"+fn, res)
}

// func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
// 	return &license
// }

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(grafana.Provider(version.Version)())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "grafana",
		DisplayName: "Grafana",
		Publisher:   "Pulumiverse",
		LogoURL:     "",

		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing grafana cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "grafana", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumiverse/pulumi-grafana",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "grafana",
		Config: map[string]*tfbridge.SchemaInfo{
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_URL"},
				},
			},
			"auth": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_AUTH"},
				},
			},
			"http_headers": {},
			"retries": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_RETRIES"},
				},
			},
			"org_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ORG_ID"},
				},
			},
			"tls_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_KEY"},
				},
			},
			"tls_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_TLS_CERT"},
				},
			},
			"ca_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CA_CERT"},
				},
			},
			"insecure_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_INSECURE_SKIP_VERIFY"},
				},
			},
			"cloud_api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_API_KEY"},
				},
			},
			"cloud_api_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_CLOUD_API_URL"},
				},
			},
			"sm_access_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_SM_ACCESS_TOKEN"},
				},
			},
			"sm_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_SM_URL"},
				},
			},
			"store_dashboard_sha256": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_STORE_DASHBOARD_SHA256"},
				},
			},
			"oncall_access_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ONCALL_ACCESS_TOKEN"},
				},
			},
			"oncall_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_ONCALL_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"grafana_alert_notification": {
				Tok: grafanaResource(grafanaMod, "AlertNotification"),
				Docs: &tfbridge.DocInfo{
					Source: "alert_notification.md",
				},
			},
			"grafana_annotation": {
				Tok: grafanaResource(grafanaMod, "Annotation"),
			},
			"grafana_api_key": {
				Tok: grafanaResource(grafanaMod, "ApiKey"),
			},
			"grafana_builtin_role_assignment": {
				Tok: grafanaResource(grafanaMod, "BuiltinRoleAssignment"),
			},
			"grafana_cloud_api_key": {
				Tok: grafanaResource(grafanaMod, "CloudApiKey"),
			},
			"grafana_cloud_plugin_installation": {
				Tok: grafanaResource(grafanaMod, "CloudPluginInstallation"),
			},
			"grafana_cloud_stack": {
				Tok: grafanaResource(grafanaMod, "CloudStack"),
			},
			"grafana_contact_point": {
				Tok: grafanaResource(grafanaMod, "ContactPoint"),
			},
			"grafana_dashboard": {
				Tok: grafanaResource(grafanaMod, "Dashboard"),
			},
			"grafana_dashboard_permission": {
				Tok: grafanaResource(grafanaMod, "DashboardPermission"),
			},
			"grafana_data_source": {
				Tok: grafanaResource(grafanaMod, "DataSource"),
			},
			"grafana_data_source_permission": {
				Tok: grafanaResource(grafanaMod, "DataSourcePermission"),
			},
			"grafana_folder": {
				Tok: grafanaResource(grafanaMod, "Folder"),
			},
			"grafana_folder_permission": {
				Tok: grafanaResource(grafanaMod, "FolderPermission"),
			},
			"grafana_library_panel": {
				Tok: grafanaResource(grafanaMod, "LibraryPanel"),
			},
			"grafana_machine_learning_job": {
				Tok: grafanaResource(grafanaMod, "MachineLearningJob"),
			},
			"grafana_message_template": {
				Tok: grafanaResource(grafanaMod, "MessageTemplate"),
			},
			"grafana_mute_timing": {
				Tok: grafanaResource(grafanaMod, "MuteTiming"),
			},
			"grafana_notification_policy": {
				Tok: grafanaResource(grafanaMod, "NotificationPolicy"),
			},
			"grafana_oncall_escalation": {
				Tok: grafanaResource(grafanaMod, "OncallEscalation"),
			},
			"grafana_oncall_escalation_chain": {
				Tok: grafanaResource(grafanaMod, "OncallEscalationChain"),
			},
			"grafana_oncall_integration": {
				Tok: grafanaResource(grafanaMod, "OncallIntegration"),
			},
			"grafana_oncall_on_call_shift": {
				Tok: grafanaResource(grafanaMod, "OncallShift"),
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaResource(grafanaMod, "OutgoingWebhook"),
			},
			"grafana_oncall_route": {
				Tok: grafanaResource(grafanaMod, "OncallRoute"),
			},
			"grafana_oncall_schedule": {
				Tok: grafanaResource(grafanaMod, "OncallSchedule"),
			},
			"grafana_organization": {
				Tok: grafanaResource(grafanaMod, "Organization"),
			},
			"grafana_playlist": {
				Tok: grafanaResource(grafanaMod, "Playlist"),
			},
			"grafana_report": {
				Tok: grafanaResource(grafanaMod, "Report"),
			},
			"grafana_role": {
				Tok: grafanaResource(grafanaMod, "Role"),
			},
			"grafana_rule_group": {
				Tok: grafanaResource(grafanaMod, "RuleGroup"),
			},
			"grafana_service_account": {
				Tok: grafanaResource(grafanaMod, "ServiceAccount"),
			},
			"grafana_service_account_token": {
				Tok: grafanaResource(grafanaMod, "ServiceAccountToken"),
			},
			"grafana_synthetic_monitoring_check": {
				Tok: grafanaResource(grafanaMod, "SyntheticMonitoringCheck"),
			},
			"grafana_synthetic_monitoring_installation": {
				Tok: grafanaResource(grafanaMod, "SyntheticMonitoringInstallation"),
			},
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaResource(grafanaMod, "SyntheticMonitoringProbe"),
			},
			"grafana_team": {
				Tok: grafanaResource(grafanaMod, "Team"),
			},
			"grafana_team_external_group": {
				Tok: grafanaResource(grafanaMod, "ExternalGroup"),
			},
			"grafana_team_preferences": {
				Tok: grafanaResource(grafanaMod, "TeamPreferences"),
			},
			"grafana_user": {
				Tok: grafanaResource(grafanaMod, "User"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"grafana_cloud_ips": {
				Tok: grafanaDataSource(grafanaMod, "getCloudIps"),
			},
			"grafana_cloud_stack": {
				Tok: grafanaDataSource(grafanaMod, "getCloudStack"),
			},
			"grafana_dashboard": {
				Tok: grafanaDataSource(grafanaMod, "getDashboard"),
			},
			"grafana_dashboards": {
				Tok: grafanaDataSource(grafanaMod, "getDashboards"),
			},
			"grafana_folder": {
				Tok: grafanaDataSource(grafanaMod, "getFolder"),
			},
			"grafana_folders": {
				Tok: grafanaDataSource(grafanaMod, "getFolders"),
			},
			"grafana_library_panel": {
				Tok: grafanaDataSource(grafanaMod, "getLibraryPanel"),
			},
			"grafana_oncall_action": {
				Tok: grafanaDataSource(grafanaMod, "getOncallAction"),
			},
			"grafana_oncall_escalation_chain": {
				Tok: grafanaDataSource(grafanaMod, "getOncallEscalationChain"),
			},
			"grafana_oncall_outgoing_webhook": {
				Tok: grafanaDataSource(grafanaMod, "getOncallOutgoingWebhook"),
			},
			"grafana_oncall_schedule": {
				Tok: grafanaDataSource(grafanaMod, "getOncallSchedule"),
			},
			"grafana_oncall_slack_channel": {
				Tok: grafanaDataSource(grafanaMod, "getOncallSlackChannel"),
			},
			"grafana_oncall_team": {
				Tok: grafanaDataSource(grafanaMod, "getOncallTeam"),
			},
			"grafana_oncall_user": {
				Tok: grafanaDataSource(grafanaMod, "getOncallUser"),
			},
			"grafana_oncall_user_group": {
				Tok: grafanaDataSource(grafanaMod, "getOncallUserGroup"),
			},
			"grafana_organization": {
				Tok: grafanaDataSource(grafanaMod, "getOrganization"),
			},
			"grafana_synthetic_monitoring_probe": {
				Tok: grafanaDataSource(grafanaMod, "getSyntheticMonitoringProbe"),
			},
			"grafana_synthetic_monitoring_probes": {
				Tok: grafanaDataSource(grafanaMod, "getSyntheticMonitoringProbes"),
			},
			"grafana_team": {
				Tok: grafanaDataSource(grafanaMod, "getTeam"),
			},
			"grafana_user": {
				Tok: grafanaDataSource(grafanaMod, "getUser"),
				Docs: &tfbridge.DocInfo{
					Source: "user.md",
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", grafanaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				grafanaPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
