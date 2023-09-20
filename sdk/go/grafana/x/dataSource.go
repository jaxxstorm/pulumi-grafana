// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)
//
// The required arguments for this resource vary depending on the type of data
// source selected (via the 'type' argument).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"tokenUri":           "https://oauth2.googleapis.com/token",
//				"authenticationType": "jwt",
//				"defaultProject":     "default-project",
//				"clientEmail":        "client-email@default-project.iam.gserviceaccount.com",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"privateKey": "-----BEGIN PRIVATE KEY-----\nprivate-key\n-----END PRIVATE KEY-----\n",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = grafana.NewDataSource(ctx, "arbitrary-data", &grafana.DataSourceArgs{
//				Type:                  pulumi.String("stackdriver"),
//				JsonDataEncoded:       pulumi.String(json0),
//				SecureJsonDataEncoded: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON2, err := json.Marshal(map[string]interface{}{
//				"authType":          "default",
//				"basicAuthPassword": "mypassword",
//			})
//			if err != nil {
//				return err
//			}
//			json2 := string(tmpJSON2)
//			_, err = grafana.NewDataSource(ctx, "influxdb", &grafana.DataSourceArgs{
//				Type:              pulumi.String("influxdb"),
//				Url:               pulumi.String("http://influxdb.example.net:8086/"),
//				BasicAuthEnabled:  pulumi.Bool(true),
//				BasicAuthUsername: pulumi.String("username"),
//				DatabaseName:      pulumi.Any(influxdb_database.Metrics.Name),
//				JsonDataEncoded:   pulumi.String(json2),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON3, err := json.Marshal(map[string]interface{}{
//				"defaultRegion": "us-east-1",
//				"authType":      "keys",
//			})
//			if err != nil {
//				return err
//			}
//			json3 := string(tmpJSON3)
//			tmpJSON4, err := json.Marshal(map[string]interface{}{
//				"accessKey": "123",
//				"secretKey": "456",
//			})
//			if err != nil {
//				return err
//			}
//			json4 := string(tmpJSON4)
//			_, err = grafana.NewDataSource(ctx, "cloudwatch", &grafana.DataSourceArgs{
//				Type:                  pulumi.String("cloudwatch"),
//				JsonDataEncoded:       pulumi.String(json3),
//				SecureJsonDataEncoded: pulumi.String(json4),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON5, err := json.Marshal(map[string]interface{}{
//				"httpMethod":        "POST",
//				"prometheusType":    "Mimir",
//				"prometheusVersion": "2.4.0",
//			})
//			if err != nil {
//				return err
//			}
//			json5 := string(tmpJSON5)
//			tmpJSON6, err := json.Marshal(map[string]interface{}{
//				"basicAuthPassword": "password",
//			})
//			if err != nil {
//				return err
//			}
//			json6 := string(tmpJSON6)
//			_, err = grafana.NewDataSource(ctx, "prometheus", &grafana.DataSourceArgs{
//				Type:                  pulumi.String("prometheus"),
//				Url:                   pulumi.String("https://my-instances.com"),
//				BasicAuthEnabled:      pulumi.Bool(true),
//				BasicAuthUsername:     pulumi.String("username"),
//				JsonDataEncoded:       pulumi.String(json5),
//				SecureJsonDataEncoded: pulumi.String(json6),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import grafana:index/dataSource:DataSource by_integer_id {{datasource_id}} # To use the default provider org
//
// ```
//
// ```sh
//
//	$ pulumi import grafana:index/dataSource:DataSource by_uid {{datasource_uid}} # To use the default provider org
//
// ```
//
// ```sh
//
//	$ pulumi import grafana:index/dataSource:DataSource by_integer_id {{org_id}}:{{datasource_id}} # When "org_id" is set on the resource
//
// ```
//
// ```sh
//
//	$ pulumi import grafana:index/dataSource:DataSource by_uid {{org_id}}:{{datasource_uid}} # When "org_id" is set on the resource
//
// ```
type DataSource struct {
	pulumi.CustomResourceState

	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode pulumix.Output[*string] `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumix.Output[*bool] `pulumi:"basicAuthEnabled"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumix.Output[*string] `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumix.Output[*string] `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders pulumix.MapOutput[string] `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumix.Output[*bool] `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumix.Output[*string] `pulumi:"jsonDataEncoded"`
	// A unique name for the data source.
	Name pulumix.Output[string] `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumix.Output[*string] `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumix.Output[*string] `pulumi:"secureJsonDataEncoded"`
	// The data source type. Must be one of the supported data source keywords.
	Type pulumix.Output[string] `pulumi:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumix.Output[string] `pulumi:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumix.Output[*string] `pulumi:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumix.Output[*string] `pulumi:"username"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.HttpHeaders != nil {
		untypedSecretValue := pulumi.ToSecret(args.HttpHeaders.ToOutput(ctx.Context()).Untyped())
		args.HttpHeaders = pulumix.MustConvertTyped[map[string]string](untypedSecretValue)
	}
	if args.SecureJsonDataEncoded != nil {
		untypedSecretValue := pulumi.ToSecret(args.SecureJsonDataEncoded.ToOutput(ctx.Context()).Untyped())
		args.SecureJsonDataEncoded = pulumix.MustConvertTyped[*string](untypedSecretValue)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"httpHeaders",
		"secureJsonDataEncoded",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataSource
	err := ctx.RegisterResource("grafana:index/dataSource:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("grafana:index/dataSource:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode *string `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled *bool `pulumi:"basicAuthEnabled"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername *string `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName *string `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault *bool `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// A unique name for the data source.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// The data source type. Must be one of the supported data source keywords.
	Type *string `pulumi:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid *string `pulumi:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url *string `pulumi:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username *string `pulumi:"username"`
}

type DataSourceState struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode pulumix.Input[*string]
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumix.Input[*bool]
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumix.Input[*string]
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumix.Input[*string]
	// Custom HTTP headers
	HttpHeaders pulumix.Input[map[string]string]
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumix.Input[*bool]
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumix.Input[*string]
	// A unique name for the data source.
	Name pulumix.Input[*string]
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumix.Input[*string]
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumix.Input[*string]
	// The data source type. Must be one of the supported data source keywords.
	Type pulumix.Input[*string]
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumix.Input[*string]
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumix.Input[*string]
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumix.Input[*string]
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode *string `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled *bool `pulumi:"basicAuthEnabled"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername *string `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName *string `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault *bool `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// A unique name for the data source.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// The data source type. Must be one of the supported data source keywords.
	Type string `pulumi:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid *string `pulumi:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url *string `pulumi:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode pulumix.Input[*string]
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumix.Input[*bool]
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumix.Input[*string]
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumix.Input[*string]
	// Custom HTTP headers
	HttpHeaders pulumix.Input[map[string]string]
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumix.Input[*bool]
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumix.Input[*string]
	// A unique name for the data source.
	Name pulumix.Input[*string]
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumix.Input[*string]
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumix.Input[*string]
	// The data source type. Must be one of the supported data source keywords.
	Type pulumix.Input[string]
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumix.Input[*string]
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumix.Input[*string]
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumix.Input[*string]
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToOutput(ctx context.Context) pulumix.Output[DataSource] {
	return pulumix.Output[DataSource]{
		OutputState: o.OutputState,
	}
}

// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
func (o DataSourceOutput) AccessMode() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.AccessMode })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Whether to enable basic auth for the data source. Defaults to `false`.
func (o DataSourceOutput) BasicAuthEnabled() pulumix.Output[*bool] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*bool] { return v.BasicAuthEnabled })
	return pulumix.Flatten[*bool, pulumix.Output[*bool]](value)
}

// Basic auth username. Defaults to “.
func (o DataSourceOutput) BasicAuthUsername() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.BasicAuthUsername })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to “.
func (o DataSourceOutput) DatabaseName() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.DatabaseName })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Custom HTTP headers
func (o DataSourceOutput) HttpHeaders() pulumix.MapOutput[string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.MapOutput[string] { return v.HttpHeaders })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
func (o DataSourceOutput) IsDefault() pulumix.Output[*bool] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*bool] { return v.IsDefault })
	return pulumix.Flatten[*bool, pulumix.Output[*bool]](value)
}

// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
func (o DataSourceOutput) JsonDataEncoded() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.JsonDataEncoded })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// A unique name for the data source.
func (o DataSourceOutput) Name() pulumix.Output[string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[string] { return v.Name })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DataSourceOutput) OrgId() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.OrgId })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
func (o DataSourceOutput) SecureJsonDataEncoded() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.SecureJsonDataEncoded })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// The data source type. Must be one of the supported data source keywords.
func (o DataSourceOutput) Type() pulumix.Output[string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[string] { return v.Type })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Unique identifier. If unset, this will be automatically generated.
func (o DataSourceOutput) Uid() pulumix.Output[string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[string] { return v.Uid })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// The URL for the data source. The type of URL required varies depending on the chosen data source type.
func (o DataSourceOutput) Url() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.Url })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// (Required by some data source types) The username to use to authenticate to the data source. Defaults to “.
func (o DataSourceOutput) Username() pulumix.Output[*string] {
	value := pulumix.Apply[DataSource](o, func(v DataSource) pulumix.Output[*string] { return v.Username })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

func init() {
	pulumi.RegisterOutputType(DataSourceOutput{})
}
