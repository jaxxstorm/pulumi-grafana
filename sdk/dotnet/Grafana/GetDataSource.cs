// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana
{
    public static class GetDataSource
    {
        /// <summary>
        /// Get details about a Grafana Datasource querying by either name, uid or ID
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Grafana = Lbrlabs.PulumiPackage.Grafana;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var prometheus = new Grafana.DataSource("prometheus", new()
        ///     {
        ///         Type = "prometheus",
        ///         Uid = "prometheus-ds-test-uid",
        ///         Url = "https://my-instance.com",
        ///         BasicAuthEnabled = true,
        ///         BasicAuthUsername = "username",
        ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["httpMethod"] = "POST",
        ///             ["prometheusType"] = "Mimir",
        ///             ["prometheusVersion"] = "2.4.0",
        ///         }),
        ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["basicAuthPassword"] = "password",
        ///         }),
        ///     });
        /// 
        ///     var fromName = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Name = prometheus.Name,
        ///     });
        /// 
        ///     var fromId = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Id = prometheus.Id,
        ///     });
        /// 
        ///     var fromUid = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Uid = prometheus.Uid,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataSourceResult> InvokeAsync(GetDataSourceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataSourceResult>("grafana:index/getDataSource:getDataSource", args ?? new GetDataSourceArgs(), options.WithDefaults());

        /// <summary>
        /// Get details about a Grafana Datasource querying by either name, uid or ID
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Grafana = Lbrlabs.PulumiPackage.Grafana;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var prometheus = new Grafana.DataSource("prometheus", new()
        ///     {
        ///         Type = "prometheus",
        ///         Uid = "prometheus-ds-test-uid",
        ///         Url = "https://my-instance.com",
        ///         BasicAuthEnabled = true,
        ///         BasicAuthUsername = "username",
        ///         JsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["httpMethod"] = "POST",
        ///             ["prometheusType"] = "Mimir",
        ///             ["prometheusVersion"] = "2.4.0",
        ///         }),
        ///         SecureJsonDataEncoded = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["basicAuthPassword"] = "password",
        ///         }),
        ///     });
        /// 
        ///     var fromName = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Name = prometheus.Name,
        ///     });
        /// 
        ///     var fromId = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Id = prometheus.Id,
        ///     });
        /// 
        ///     var fromUid = Grafana.GetDataSource.Invoke(new()
        ///     {
        ///         Uid = prometheus.Uid,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDataSourceResult> Invoke(GetDataSourceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataSourceResult>("grafana:index/getDataSource:getDataSource", args ?? new GetDataSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataSourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        [Input("uid")]
        public string? Uid { get; set; }

        public GetDataSourceArgs()
        {
        }
        public static new GetDataSourceArgs Empty => new GetDataSourceArgs();
    }

    public sealed class GetDataSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public GetDataSourceInvokeArgs()
        {
        }
        public static new GetDataSourceInvokeArgs Empty => new GetDataSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataSourceResult
    {
        /// <summary>
        /// The method by which Grafana will access the data source: `proxy` or `direct`.
        /// </summary>
        public readonly string AccessMode;
        /// <summary>
        /// Whether to enable basic auth for the data source.
        /// </summary>
        public readonly bool BasicAuthEnabled;
        /// <summary>
        /// Basic auth username.
        /// </summary>
        public readonly string BasicAuthUsername;
        /// <summary>
        /// (Required by some data source types) The name of the database to use on the selected data source server.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether to set the data source as default. This should only be `true` to a single data source.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
        /// </summary>
        public readonly string JsonDataEncoded;
        public readonly string Name;
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// The data source type. Must be one of the supported data source keywords.
        /// </summary>
        public readonly string Type;
        public readonly string Uid;
        /// <summary>
        /// The URL for the data source. The type of URL required varies depending on the chosen data source type.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// (Required by some data source types) The username to use to authenticate to the data source.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetDataSourceResult(
            string accessMode,

            bool basicAuthEnabled,

            string basicAuthUsername,

            string databaseName,

            string id,

            bool isDefault,

            string jsonDataEncoded,

            string name,

            string? orgId,

            string type,

            string uid,

            string url,

            string username)
        {
            AccessMode = accessMode;
            BasicAuthEnabled = basicAuthEnabled;
            BasicAuthUsername = basicAuthUsername;
            DatabaseName = databaseName;
            Id = id;
            IsDefault = isDefault;
            JsonDataEncoded = jsonDataEncoded;
            Name = name;
            OrgId = orgId;
            Type = type;
            Uid = uid;
            Url = url;
            Username = username;
        }
    }
}
