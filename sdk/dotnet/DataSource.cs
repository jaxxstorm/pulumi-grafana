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
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/data_source/)
    /// 
    /// The required arguments for this resource vary depending on the type of data
    /// source selected (via the 'type' argument).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Grafana = Lbrlabs_Pulumi.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var arbitrary_data = new Grafana.DataSource("arbitrary-data", new()
    ///     {
    ///         Type = "stackdriver",
    ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["tokenUri"] = "https://oauth2.googleapis.com/token",
    ///             ["authenticationType"] = "jwt",
    ///             ["defaultProject"] = "default-project",
    ///             ["clientEmail"] = "client-email@default-project.iam.gserviceaccount.com",
    ///         }),
    ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["privateKey"] = @"-----BEGIN PRIVATE KEY-----
    /// private-key
    /// -----END PRIVATE KEY-----
    /// ",
    ///         }),
    ///     });
    /// 
    ///     var influxdb = new Grafana.DataSource("influxdb", new()
    ///     {
    ///         Type = "influxdb",
    ///         Url = "http://influxdb.example.net:8086/",
    ///         Username = "myapp",
    ///         Password = "foobarbaz",
    ///         DatabaseName = influxdb_database.Metrics.Name,
    ///     });
    /// 
    ///     var cloudwatch = new Grafana.DataSource("cloudwatch", new()
    ///     {
    ///         Type = "cloudwatch",
    ///         JsonDatas = new[]
    ///         {
    ///             new Grafana.Inputs.DataSourceJsonDataArgs
    ///             {
    ///                 DefaultRegion = "us-east-1",
    ///                 AuthType = "keys",
    ///             },
    ///         },
    ///         SecureJsonDatas = new[]
    ///         {
    ///             new Grafana.Inputs.DataSourceSecureJsonDataArgs
    ///             {
    ///                 AccessKey = "123",
    ///                 SecretKey = "456",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var prometheus = new Grafana.DataSource("prometheus", new()
    ///     {
    ///         Type = "prometheus",
    ///         Url = "https://aps-workspaces.eu-west-1.amazonaws.com/workspaces/ws-1234567890/",
    ///         JsonDatas = new[]
    ///         {
    ///             new Grafana.Inputs.DataSourceJsonDataArgs
    ///             {
    ///                 HttpMethod = "POST",
    ///                 Sigv4Auth = true,
    ///                 Sigv4AuthType = "default",
    ///                 Sigv4Region = "eu-west-1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var stackdriver = new Grafana.DataSource("stackdriver", new()
    ///     {
    ///         Type = "stackdriver",
    ///         JsonDatas = new[]
    ///         {
    ///             new Grafana.Inputs.DataSourceJsonDataArgs
    ///             {
    ///                 TokenUri = "https://oauth2.googleapis.com/token",
    ///                 AuthenticationType = "jwt",
    ///                 DefaultProject = "default-project",
    ///                 ClientEmail = "client-email@default-project.iam.gserviceaccount.com",
    ///             },
    ///         },
    ///         SecureJsonDatas = new[]
    ///         {
    ///             new Grafana.Inputs.DataSourceSecureJsonDataArgs
    ///             {
    ///                 PrivateKey = @"-----BEGIN PRIVATE KEY-----
    /// private-key
    /// -----END PRIVATE KEY-----
    /// ",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/dataSource:DataSource by_integer_id {{datasource id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/dataSource:DataSource by_uid {{datasource uid}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/dataSource:DataSource")]
    public partial class DataSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
        /// </summary>
        [Output("accessMode")]
        public Output<string?> AccessMode { get; private set; } = null!;

        /// <summary>
        /// Whether to enable basic auth for the data source. Defaults to `false`.
        /// </summary>
        [Output("basicAuthEnabled")]
        public Output<bool?> BasicAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Output("basicAuthPassword")]
        public Output<string?> BasicAuthPassword { get; private set; } = null!;

        /// <summary>
        /// Basic auth username. Defaults to ``.
        /// </summary>
        [Output("basicAuthUsername")]
        public Output<string?> BasicAuthUsername { get; private set; } = null!;

        /// <summary>
        /// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        [Output("httpHeaders")]
        public Output<ImmutableDictionary<string, string>?> HttpHeaders { get; private set; } = null!;

        /// <summary>
        /// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// Serialized JSON string containing the json data. Replaces the json_data attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Output("jsonDataEncoded")]
        public Output<string?> JsonDataEncoded { get; private set; } = null!;

        /// <summary>
        /// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Output("jsonDatas")]
        public Output<ImmutableArray<Outputs.DataSourceJsonData>> JsonDatas { get; private set; } = null!;

        /// <summary>
        /// A unique name for the data source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Output("secureJsonDataEncoded")]
        public Output<string?> SecureJsonDataEncoded { get; private set; } = null!;

        /// <summary>
        /// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Output("secureJsonDatas")]
        public Output<ImmutableArray<Outputs.DataSourceSecureJsonData>> SecureJsonDatas { get; private set; } = null!;

        /// <summary>
        /// The data source type. Must be one of the supported data source keywords.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The URL for the data source. The type of URL required varies depending on the chosen data source type.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        /// <summary>
        /// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a DataSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSource(string name, DataSourceArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/dataSource:DataSource", name, args ?? new DataSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSource(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dataSource:DataSource", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// Get an existing DataSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSource Get(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSource(name, id, state, options);
        }
    }

    public sealed class DataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

        /// <summary>
        /// Whether to enable basic auth for the data source. Defaults to `false`.
        /// </summary>
        [Input("basicAuthEnabled")]
        public Input<bool>? BasicAuthEnabled { get; set; }

        /// <summary>
        /// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Input("basicAuthPassword")]
        public Input<string>? BasicAuthPassword { get; set; }

        /// <summary>
        /// Basic auth username. Defaults to ``.
        /// </summary>
        [Input("basicAuthUsername")]
        public Input<string>? BasicAuthUsername { get; set; }

        /// <summary>
        /// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Serialized JSON string containing the json data. Replaces the json_data attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Input("jsonDataEncoded")]
        public Input<string>? JsonDataEncoded { get; set; }

        [Input("jsonDatas")]
        private InputList<Inputs.DataSourceJsonDataArgs>? _jsonDatas;

        /// <summary>
        /// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Obsolete(@"Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.")]
        public InputList<Inputs.DataSourceJsonDataArgs> JsonDatas
        {
            get => _jsonDatas ?? (_jsonDatas = new InputList<Inputs.DataSourceJsonDataArgs>());
            set => _jsonDatas = value;
        }

        /// <summary>
        /// A unique name for the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Input("secureJsonDataEncoded")]
        public Input<string>? SecureJsonDataEncoded { get; set; }

        [Input("secureJsonDatas")]
        private InputList<Inputs.DataSourceSecureJsonDataArgs>? _secureJsonDatas;

        /// <summary>
        /// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Obsolete(@"Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.")]
        public InputList<Inputs.DataSourceSecureJsonDataArgs> SecureJsonDatas
        {
            get => _secureJsonDatas ?? (_secureJsonDatas = new InputList<Inputs.DataSourceSecureJsonDataArgs>());
            set => _secureJsonDatas = value;
        }

        /// <summary>
        /// The data source type. Must be one of the supported data source keywords.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The URL for the data source. The type of URL required varies depending on the chosen data source type.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public DataSourceArgs()
        {
        }
        public static new DataSourceArgs Empty => new DataSourceArgs();
    }

    public sealed class DataSourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method by which Grafana will access the data source: `proxy` or `direct`. Defaults to `proxy`.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

        /// <summary>
        /// Whether to enable basic auth for the data source. Defaults to `false`.
        /// </summary>
        [Input("basicAuthEnabled")]
        public Input<bool>? BasicAuthEnabled { get; set; }

        /// <summary>
        /// Basic auth password. Deprecated:Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Input("basicAuthPassword")]
        public Input<string>? BasicAuthPassword { get; set; }

        /// <summary>
        /// Basic auth username. Defaults to ``.
        /// </summary>
        [Input("basicAuthUsername")]
        public Input<string>? BasicAuthUsername { get; set; }

        /// <summary>
        /// (Required by some data source types) The name of the database to use on the selected data source server. Defaults to ``.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Custom HTTP headers
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Whether to set the data source as default. This should only be `true` to a single data source. Defaults to `false`.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Serialized JSON string containing the json data. Replaces the json_data attribute, this attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Input("jsonDataEncoded")]
        public Input<string>? JsonDataEncoded { get; set; }

        [Input("jsonDatas")]
        private InputList<Inputs.DataSourceJsonDataGetArgs>? _jsonDatas;

        /// <summary>
        /// (Required by some data source types). Deprecated: Use json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Obsolete(@"Use json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.")]
        public InputList<Inputs.DataSourceJsonDataGetArgs> JsonDatas
        {
            get => _jsonDatas ?? (_jsonDatas = new InputList<Inputs.DataSourceJsonDataGetArgs>());
            set => _jsonDatas = value;
        }

        /// <summary>
        /// A unique name for the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Required by some data source types) The password to use to authenticate to the data source. Deprecated: Use secure*json*data_encoded instead. It supports arbitrary JSON data, and therefore all attributes. This attribute is removed in Grafana 9.0+. Defaults to ``.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Serialized JSON string containing the secure json data. Replaces the secure*json*data attribute, this attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI.
        /// </summary>
        [Input("secureJsonDataEncoded")]
        public Input<string>? SecureJsonDataEncoded { get; set; }

        [Input("secureJsonDatas")]
        private InputList<Inputs.DataSourceSecureJsonDataGetArgs>? _secureJsonDatas;

        /// <summary>
        /// Deprecated: Use secure*json*data*encoded instead. It supports arbitrary JSON data, and therefore all attributes.
        /// </summary>
        [Obsolete(@"Use secure_json_data_encoded instead. It supports arbitrary JSON data, and therefore all attributes.")]
        public InputList<Inputs.DataSourceSecureJsonDataGetArgs> SecureJsonDatas
        {
            get => _secureJsonDatas ?? (_secureJsonDatas = new InputList<Inputs.DataSourceSecureJsonDataGetArgs>());
            set => _secureJsonDatas = value;
        }

        /// <summary>
        /// The data source type. Must be one of the supported data source keywords.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Unique identifier. If unset, this will be automatically generated.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The URL for the data source. The type of URL required varies depending on the chosen data source type.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// (Required by some data source types) The username to use to authenticate to the data source. Defaults to ``.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public DataSourceState()
        {
        }
        public static new DataSourceState Empty => new DataSourceState();
    }
}
