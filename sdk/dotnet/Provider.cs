// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Grafana
{
    /// <summary>
    /// The provider type for the grafana package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [GrafanaResourceType("pulumi:providers:grafana")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        /// </summary>
        [Output("auth")]
        public Output<string?> Auth { get; private set; } = null!;

        /// <summary>
        /// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        /// `GRAFANA_CA_CERT` environment variable.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
        /// </summary>
        [Output("cloudApiKey")]
        public Output<string?> CloudApiKey { get; private set; } = null!;

        /// <summary>
        /// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        /// </summary>
        [Output("cloudApiUrl")]
        public Output<string?> CloudApiUrl { get; private set; } = null!;

        /// <summary>
        /// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Output("oncallAccessToken")]
        public Output<string?> OncallAccessToken { get; private set; } = null!;

        /// <summary>
        /// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        /// </summary>
        [Output("oncallUrl")]
        public Output<string?> OncallUrl { get; private set; } = null!;

        /// <summary>
        /// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Output("smAccessToken")]
        public Output<string?> SmAccessToken { get; private set; } = null!;

        /// <summary>
        /// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
        /// correct value for each service region is cited in the [Synthetic Monitoring
        /// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
        /// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
        /// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
        /// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
        /// each provider ensures all providers interact with the same SM API.
        /// </summary>
        [Output("smUrl")]
        public Output<string?> SmUrl { get; private set; } = null!;

        /// <summary>
        /// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        /// `GRAFANA_TLS_CERT` environment variable.
        /// </summary>
        [Output("tlsCert")]
        public Output<string?> TlsCert { get; private set; } = null!;

        /// <summary>
        /// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
        /// environment variable.
        /// </summary>
        [Output("tlsKey")]
        public Output<string?> TlsKey { get; private set; } = null!;

        /// <summary>
        /// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        /// `GRAFANA_CA_CERT` environment variable.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
        /// </summary>
        [Input("cloudApiKey")]
        public Input<string>? CloudApiKey { get; set; }

        /// <summary>
        /// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        /// </summary>
        [Input("cloudApiUrl")]
        public Input<string>? CloudApiUrl { get; set; }

        [Input("httpHeaders", json: true)]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Optional. HTTP headers mapping keys to values used for accessing the Grafana API. May alternatively be set via the
        /// `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        /// </summary>
        [Input("insecureSkipVerify", json: true)]
        public Input<bool>? InsecureSkipVerify { get; set; }

        /// <summary>
        /// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Input("oncallAccessToken")]
        public Input<string>? OncallAccessToken { get; set; }

        /// <summary>
        /// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        /// </summary>
        [Input("oncallUrl")]
        public Input<string>? OncallUrl { get; set; }

        /// <summary>
        /// The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
        /// variable.
        /// </summary>
        [Input("orgId", json: true)]
        public Input<int>? OrgId { get; set; }

        /// <summary>
        /// The amount of retries to use for Grafana API calls. May alternatively be set via the `GRAFANA_RETRIES` environment
        /// variable.
        /// </summary>
        [Input("retries", json: true)]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        /// </summary>
        [Input("smAccessToken")]
        public Input<string>? SmAccessToken { get; set; }

        /// <summary>
        /// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
        /// correct value for each service region is cited in the [Synthetic Monitoring
        /// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
        /// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
        /// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
        /// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
        /// each provider ensures all providers interact with the same SM API.
        /// </summary>
        [Input("smUrl")]
        public Input<string>? SmUrl { get; set; }

        /// <summary>
        /// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        /// </summary>
        [Input("storeDashboardSha256", json: true)]
        public Input<bool>? StoreDashboardSha256 { get; set; }

        /// <summary>
        /// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        /// `GRAFANA_TLS_CERT` environment variable.
        /// </summary>
        [Input("tlsCert")]
        public Input<string>? TlsCert { get; set; }

        /// <summary>
        /// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
        /// environment variable.
        /// </summary>
        [Input("tlsKey")]
        public Input<string>? TlsKey { get; set; }

        /// <summary>
        /// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProviderArgs()
        {
            Auth = Utilities.GetEnv("GRAFANA_AUTH");
            CaCert = Utilities.GetEnv("GRAFANA_CA_CERT");
            CloudApiKey = Utilities.GetEnv("GRAFANA_CLOUD_API_KEY");
            CloudApiUrl = Utilities.GetEnv("GRAFANA_CLOUD_API_URL");
            InsecureSkipVerify = Utilities.GetEnvBoolean("GRAFANA_INSECURE_SKIP_VERIFY");
            OncallAccessToken = Utilities.GetEnv("GRAFANA_ONCALL_ACCESS_TOKEN");
            OncallUrl = Utilities.GetEnv("GRAFANA_ONCALL_URL");
            OrgId = Utilities.GetEnvInt32("GRAFANA_ORG_ID");
            Retries = Utilities.GetEnvInt32("GRAFANA_RETRIES");
            SmAccessToken = Utilities.GetEnv("GRAFANA_SM_ACCESS_TOKEN");
            SmUrl = Utilities.GetEnv("GRAFANA_SM_URL");
            StoreDashboardSha256 = Utilities.GetEnvBoolean("GRAFANA_STORE_DASHBOARD_SHA256");
            TlsCert = Utilities.GetEnv("GRAFANA_TLS_CERT");
            TlsKey = Utilities.GetEnv("GRAFANA_TLS_KEY");
            Url = Utilities.GetEnv("GRAFANA_URL");
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
