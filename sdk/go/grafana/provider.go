// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the grafana package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
	Auth pulumi.StringPtrOutput `pulumi:"auth"`
	// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
	// `GRAFANA_CA_CERT` environment variable.
	CaCert pulumi.StringPtrOutput `pulumi:"caCert"`
	// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
	CloudApiKey pulumi.StringPtrOutput `pulumi:"cloudApiKey"`
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl pulumi.StringPtrOutput `pulumi:"cloudApiUrl"`
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken pulumi.StringPtrOutput `pulumi:"oncallAccessToken"`
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl pulumi.StringPtrOutput `pulumi:"oncallUrl"`
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken pulumi.StringPtrOutput `pulumi:"smAccessToken"`
	// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
	// correct value for each service region is cited in the [Synthetic Monitoring
	// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
	// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
	// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
	// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
	// each provider ensures all providers interact with the same SM API.
	SmUrl pulumi.StringPtrOutput `pulumi:"smUrl"`
	// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
	// `GRAFANA_TLS_CERT` environment variable.
	TlsCert pulumi.StringPtrOutput `pulumi:"tlsCert"`
	// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
	// environment variable.
	TlsKey pulumi.StringPtrOutput `pulumi:"tlsKey"`
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.Auth) {
		args.Auth = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_AUTH").(string))
	}
	if isZero(args.CaCert) {
		args.CaCert = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_CA_CERT").(string))
	}
	if isZero(args.CloudApiKey) {
		args.CloudApiKey = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_CLOUD_API_KEY").(string))
	}
	if isZero(args.CloudApiUrl) {
		args.CloudApiUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_CLOUD_API_URL").(string))
	}
	if isZero(args.InsecureSkipVerify) {
		args.InsecureSkipVerify = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "GRAFANA_INSECURE_SKIP_VERIFY").(bool))
	}
	if isZero(args.OncallAccessToken) {
		args.OncallAccessToken = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_ONCALL_ACCESS_TOKEN").(string))
	}
	if isZero(args.OncallUrl) {
		args.OncallUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_ONCALL_URL").(string))
	}
	if isZero(args.OrgId) {
		args.OrgId = pulumi.IntPtr(getEnvOrDefault(0, parseEnvInt, "GRAFANA_ORG_ID").(int))
	}
	if isZero(args.Retries) {
		args.Retries = pulumi.IntPtr(getEnvOrDefault(0, parseEnvInt, "GRAFANA_RETRIES").(int))
	}
	if isZero(args.SmAccessToken) {
		args.SmAccessToken = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_SM_ACCESS_TOKEN").(string))
	}
	if isZero(args.SmUrl) {
		args.SmUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_SM_URL").(string))
	}
	if isZero(args.StoreDashboardSha256) {
		args.StoreDashboardSha256 = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "GRAFANA_STORE_DASHBOARD_SHA256").(bool))
	}
	if isZero(args.TlsCert) {
		args.TlsCert = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_TLS_CERT").(string))
	}
	if isZero(args.TlsKey) {
		args.TlsKey = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_TLS_KEY").(string))
	}
	if isZero(args.Url) {
		args.Url = pulumi.StringPtr(getEnvOrDefault("", nil, "GRAFANA_URL").(string))
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:grafana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
	Auth *string `pulumi:"auth"`
	// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
	// `GRAFANA_CA_CERT` environment variable.
	CaCert *string `pulumi:"caCert"`
	// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
	CloudApiKey *string `pulumi:"cloudApiKey"`
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl *string `pulumi:"cloudApiUrl"`
	// Optional. HTTP headers mapping keys to values used for accessing the Grafana API. May alternatively be set via the
	// `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken *string `pulumi:"oncallAccessToken"`
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl *string `pulumi:"oncallUrl"`
	// The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
	// variable.
	OrgId *int `pulumi:"orgId"`
	// The amount of retries to use for Grafana API calls. May alternatively be set via the `GRAFANA_RETRIES` environment
	// variable.
	Retries *int `pulumi:"retries"`
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken *string `pulumi:"smAccessToken"`
	// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
	// correct value for each service region is cited in the [Synthetic Monitoring
	// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
	// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
	// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
	// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
	// each provider ensures all providers interact with the same SM API.
	SmUrl *string `pulumi:"smUrl"`
	// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
	StoreDashboardSha256 *bool `pulumi:"storeDashboardSha256"`
	// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
	// `GRAFANA_TLS_CERT` environment variable.
	TlsCert *string `pulumi:"tlsCert"`
	// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
	// environment variable.
	TlsKey *string `pulumi:"tlsKey"`
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
	Auth pulumi.StringPtrInput
	// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
	// `GRAFANA_CA_CERT` environment variable.
	CaCert pulumi.StringPtrInput
	// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
	CloudApiKey pulumi.StringPtrInput
	// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
	CloudApiUrl pulumi.StringPtrInput
	// Optional. HTTP headers mapping keys to values used for accessing the Grafana API. May alternatively be set via the
	// `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
	HttpHeaders pulumi.StringMapInput
	// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
	InsecureSkipVerify pulumi.BoolPtrInput
	// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
	OncallAccessToken pulumi.StringPtrInput
	// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
	OncallUrl pulumi.StringPtrInput
	// The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
	// variable.
	OrgId pulumi.IntPtrInput
	// The amount of retries to use for Grafana API calls. May alternatively be set via the `GRAFANA_RETRIES` environment
	// variable.
	Retries pulumi.IntPtrInput
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	SmAccessToken pulumi.StringPtrInput
	// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
	// correct value for each service region is cited in the [Synthetic Monitoring
	// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
	// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
	// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
	// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
	// each provider ensures all providers interact with the same SM API.
	SmUrl pulumi.StringPtrInput
	// Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
	StoreDashboardSha256 pulumi.BoolPtrInput
	// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
	// `GRAFANA_TLS_CERT` environment variable.
	TlsCert pulumi.StringPtrInput
	// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
	// environment variable.
	TlsKey pulumi.StringPtrInput
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
func (o ProviderOutput) Auth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Auth }).(pulumi.StringPtrOutput)
}

// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
// `GRAFANA_CA_CERT` environment variable.
func (o ProviderOutput) CaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCert }).(pulumi.StringPtrOutput)
}

// API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
func (o ProviderOutput) CloudApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudApiKey }).(pulumi.StringPtrOutput)
}

// Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
func (o ProviderOutput) CloudApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CloudApiUrl }).(pulumi.StringPtrOutput)
}

// A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
func (o ProviderOutput) OncallAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OncallAccessToken }).(pulumi.StringPtrOutput)
}

// An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
func (o ProviderOutput) OncallUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OncallUrl }).(pulumi.StringPtrOutput)
}

// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
func (o ProviderOutput) SmAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SmAccessToken }).(pulumi.StringPtrOutput)
}

// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
// correct value for each service region is cited in the [Synthetic Monitoring
// documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
// the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
// `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
// managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
// each provider ensures all providers interact with the same SM API.
func (o ProviderOutput) SmUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SmUrl }).(pulumi.StringPtrOutput)
}

// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
// `GRAFANA_TLS_CERT` environment variable.
func (o ProviderOutput) TlsCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TlsCert }).(pulumi.StringPtrOutput)
}

// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
// environment variable.
func (o ProviderOutput) TlsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TlsKey }).(pulumi.StringPtrOutput)
}

// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
