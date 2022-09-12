// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/data_source/)
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
//			_, err = grafana.NewDataSource(ctx, "influxdb", &grafana.DataSourceArgs{
//				Type:         pulumi.String("influxdb"),
//				Url:          pulumi.String("http://influxdb.example.net:8086/"),
//				Username:     pulumi.String("myapp"),
//				Password:     pulumi.String("foobarbaz"),
//				DatabaseName: pulumi.Any(influxdb_database.Metrics.Name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewDataSource(ctx, "cloudwatch", &grafana.DataSourceArgs{
//				Type: pulumi.String("cloudwatch"),
//				JsonDatas: DataSourceJsonDataArray{
//					&DataSourceJsonDataArgs{
//						DefaultRegion: pulumi.String("us-east-1"),
//						AuthType:      pulumi.String("keys"),
//					},
//				},
//				SecureJsonDatas: DataSourceSecureJsonDataArray{
//					&DataSourceSecureJsonDataArgs{
//						AccessKey: pulumi.String("123"),
//						SecretKey: pulumi.String("456"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewDataSource(ctx, "prometheus", &grafana.DataSourceArgs{
//				Type: pulumi.String("prometheus"),
//				Url:  pulumi.String("https://aps-workspaces.eu-west-1.amazonaws.com/workspaces/ws-1234567890/"),
//				JsonDatas: DataSourceJsonDataArray{
//					&DataSourceJsonDataArgs{
//						HttpMethod:    pulumi.String("POST"),
//						Sigv4Auth:     pulumi.Bool(true),
//						Sigv4AuthType: pulumi.String("default"),
//						Sigv4Region:   pulumi.String("eu-west-1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewDataSource(ctx, "stackdriver", &grafana.DataSourceArgs{
//				Type: pulumi.String("stackdriver"),
//				JsonDatas: DataSourceJsonDataArray{
//					&DataSourceJsonDataArgs{
//						TokenUri:           pulumi.String("https://oauth2.googleapis.com/token"),
//						AuthenticationType: pulumi.String("jwt"),
//						DefaultProject:     pulumi.String("default-project"),
//						ClientEmail:        pulumi.String("client-email@default-project.iam.gserviceaccount.com"),
//					},
//				},
//				SecureJsonDatas: DataSourceSecureJsonDataArray{
//					&DataSourceSecureJsonDataArgs{
//						PrivateKey: pulumi.String("-----BEGIN PRIVATE KEY-----\nprivate-key\n-----END PRIVATE KEY-----\n"),
//					},
//				},
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
//	$ pulumi import grafana:index/dataSource:DataSource by_integer_id {{datasource id}}
//
// ```
//
// ```sh
//
//	$ pulumi import grafana:index/dataSource:DataSource by_uid {{datasource uid}}
//
// ```
type DataSource struct {
	pulumi.CustomResourceState

	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode pulumi.StringPtrOutput `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumi.BoolPtrOutput `pulumi:"basicAuthEnabled"`
	// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	BasicAuthPassword pulumi.StringPtrOutput `pulumi:"basicAuthPassword"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumi.StringPtrOutput `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders pulumi.StringMapOutput `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	JsonDataEncoded pulumi.StringPtrOutput `pulumi:"jsonDataEncoded"`
	// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	JsonDatas DataSourceJsonDataArrayOutput `pulumi:"jsonDatas"`
	// A unique name for the data source.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	SecureJsonDataEncoded pulumi.StringPtrOutput `pulumi:"secureJsonDataEncoded"`
	// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	SecureJsonDatas DataSourceSecureJsonDataArrayOutput `pulumi:"secureJsonDatas"`
	// The data source type. Must be one of the supported data source keywords.
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumi.StringPtrOutput `pulumi:"username"`
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
	opts = pkgResourceDefaultOpts(opts)
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
	// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	BasicAuthPassword *string `pulumi:"basicAuthPassword"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername *string `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName *string `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault *bool `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	JsonDatas []DataSourceJsonData `pulumi:"jsonDatas"`
	// A unique name for the data source.
	Name *string `pulumi:"name"`
	// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	Password *string `pulumi:"password"`
	// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	SecureJsonDatas []DataSourceSecureJsonData `pulumi:"secureJsonDatas"`
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
	AccessMode pulumi.StringPtrInput
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumi.BoolPtrInput
	// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	BasicAuthPassword pulumi.StringPtrInput
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumi.StringPtrInput
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumi.StringPtrInput
	// Custom HTTP headers
	HttpHeaders pulumi.StringMapInput
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumi.BoolPtrInput
	// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	JsonDataEncoded pulumi.StringPtrInput
	// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	JsonDatas DataSourceJsonDataArrayInput
	// A unique name for the data source.
	Name pulumi.StringPtrInput
	// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	Password pulumi.StringPtrInput
	// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	SecureJsonDataEncoded pulumi.StringPtrInput
	// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	SecureJsonDatas DataSourceSecureJsonDataArrayInput
	// The data source type. Must be one of the supported data source keywords.
	Type pulumi.StringPtrInput
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringPtrInput
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumi.StringPtrInput
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumi.StringPtrInput
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
	AccessMode *string `pulumi:"accessMode"`
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled *bool `pulumi:"basicAuthEnabled"`
	// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	BasicAuthPassword *string `pulumi:"basicAuthPassword"`
	// Basic auth username. Defaults to ``.
	BasicAuthUsername *string `pulumi:"basicAuthUsername"`
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName *string `pulumi:"databaseName"`
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault *bool `pulumi:"isDefault"`
	// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	JsonDatas []DataSourceJsonData `pulumi:"jsonDatas"`
	// A unique name for the data source.
	Name *string `pulumi:"name"`
	// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	Password *string `pulumi:"password"`
	// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	SecureJsonDatas []DataSourceSecureJsonData `pulumi:"secureJsonDatas"`
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
	AccessMode pulumi.StringPtrInput
	// Whether to enable basic auth for the data source. Defaults to `false`.
	BasicAuthEnabled pulumi.BoolPtrInput
	// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	BasicAuthPassword pulumi.StringPtrInput
	// Basic auth username. Defaults to ``.
	BasicAuthUsername pulumi.StringPtrInput
	// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
	DatabaseName pulumi.StringPtrInput
	// Custom HTTP headers
	HttpHeaders pulumi.StringMapInput
	// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
	IsDefault pulumi.BoolPtrInput
	// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	JsonDataEncoded pulumi.StringPtrInput
	// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	JsonDatas DataSourceJsonDataArrayInput
	// A unique name for the data source.
	Name pulumi.StringPtrInput
	// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
	Password pulumi.StringPtrInput
	// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
	SecureJsonDataEncoded pulumi.StringPtrInput
	// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	//
	// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
	SecureJsonDatas DataSourceSecureJsonDataArrayInput
	// The data source type. Must be one of the supported data source keywords.
	Type pulumi.StringInput
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringPtrInput
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	Url pulumi.StringPtrInput
	// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
	Username pulumi.StringPtrInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

// DataSourceArrayInput is an input type that accepts DataSourceArray and DataSourceArrayOutput values.
// You can construct a concrete instance of `DataSourceArrayInput` via:
//
//	DataSourceArray{ DataSourceArgs{...} }
type DataSourceArrayInput interface {
	pulumi.Input

	ToDataSourceArrayOutput() DataSourceArrayOutput
	ToDataSourceArrayOutputWithContext(context.Context) DataSourceArrayOutput
}

type DataSourceArray []DataSourceInput

func (DataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSource)(nil)).Elem()
}

func (i DataSourceArray) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return i.ToDataSourceArrayOutputWithContext(context.Background())
}

func (i DataSourceArray) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceArrayOutput)
}

// DataSourceMapInput is an input type that accepts DataSourceMap and DataSourceMapOutput values.
// You can construct a concrete instance of `DataSourceMapInput` via:
//
//	DataSourceMap{ "key": DataSourceArgs{...} }
type DataSourceMapInput interface {
	pulumi.Input

	ToDataSourceMapOutput() DataSourceMapOutput
	ToDataSourceMapOutputWithContext(context.Context) DataSourceMapOutput
}

type DataSourceMap map[string]DataSourceInput

func (DataSourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSource)(nil)).Elem()
}

func (i DataSourceMap) ToDataSourceMapOutput() DataSourceMapOutput {
	return i.ToDataSourceMapOutputWithContext(context.Background())
}

func (i DataSourceMap) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceMapOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
func (o DataSourceOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.AccessMode }).(pulumi.StringPtrOutput)
}

// Whether to enable basic auth for the data source. Defaults to `false`.
func (o DataSourceOutput) BasicAuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.BoolPtrOutput { return v.BasicAuthEnabled }).(pulumi.BoolPtrOutput)
}

// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to “.
//
// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
func (o DataSourceOutput) BasicAuthPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.BasicAuthPassword }).(pulumi.StringPtrOutput)
}

// Basic auth username. Defaults to “.
func (o DataSourceOutput) BasicAuthUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.BasicAuthUsername }).(pulumi.StringPtrOutput)
}

// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to “.
func (o DataSourceOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// Custom HTTP headers
func (o DataSourceOutput) HttpHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringMapOutput { return v.HttpHeaders }).(pulumi.StringMapOutput)
}

// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
func (o DataSourceOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// Serialized JSON string containing the json data. Replaces the jsonData attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
func (o DataSourceOutput) JsonDataEncoded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.JsonDataEncoded }).(pulumi.StringPtrOutput)
}

// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
//
// Deprecated: Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
func (o DataSourceOutput) JsonDatas() DataSourceJsonDataArrayOutput {
	return o.ApplyT(func(v *DataSource) DataSourceJsonDataArrayOutput { return v.JsonDatas }).(DataSourceJsonDataArrayOutput)
}

// A unique name for the data source.
func (o DataSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to “.
//
// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+.
func (o DataSourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
func (o DataSourceOutput) SecureJsonDataEncoded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.SecureJsonDataEncoded }).(pulumi.StringPtrOutput)
}

// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
//
// Deprecated: Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.
func (o DataSourceOutput) SecureJsonDatas() DataSourceSecureJsonDataArrayOutput {
	return o.ApplyT(func(v *DataSource) DataSourceSecureJsonDataArrayOutput { return v.SecureJsonDatas }).(DataSourceSecureJsonDataArrayOutput)
}

// The data source type. Must be one of the supported data source keywords.
func (o DataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique identifier. If unset, this will be automatically generated.
func (o DataSourceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The URL for the data source. The type of URL required varies depending on the chosen data source type.
func (o DataSourceOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// (Required by some data source types) The username to use to authenticate to the data source. Defaults to “.
func (o DataSourceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type DataSourceArrayOutput struct{ *pulumi.OutputState }

func (DataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSource)(nil)).Elem()
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) Index(i pulumi.IntInput) DataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSource {
		return vs[0].([]*DataSource)[vs[1].(int)]
	}).(DataSourceOutput)
}

type DataSourceMapOutput struct{ *pulumi.OutputState }

func (DataSourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSource)(nil)).Elem()
}

func (o DataSourceMapOutput) ToDataSourceMapOutput() DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) MapIndex(k pulumi.StringInput) DataSourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSource {
		return vs[0].(map[string]*DataSource)[vs[1].(string)]
	}).(DataSourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceInput)(nil)).Elem(), &DataSource{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceArrayInput)(nil)).Elem(), DataSourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceMapInput)(nil)).Elem(), DataSourceMap{})
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourceArrayOutput{})
	pulumi.RegisterOutputType(DataSourceMapOutput{})
}
