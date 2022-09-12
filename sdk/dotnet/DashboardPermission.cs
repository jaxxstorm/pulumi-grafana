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
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/dashboard_permissions/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using Pulumi;
    /// using Grafana = Lbrlabs_Pulumi.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var team = new Grafana.Team("team");
    /// 
    ///     var user = new Grafana.User("user", new()
    ///     {
    ///         Email = "user.name@example.com",
    ///     });
    /// 
    ///     var metrics = new Grafana.Dashboard("metrics", new()
    ///     {
    ///         ConfigJson = File.ReadAllText("grafana-dashboard.json"),
    ///     });
    /// 
    ///     var collectionPermission = new Grafana.DashboardPermission("collectionPermission", new()
    ///     {
    ///         DashboardId = metrics.DashboardId,
    ///         Permissions = new[]
    ///         {
    ///             new Grafana.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 Role = "Editor",
    ///                 Permission = "Edit",
    ///             },
    ///             new Grafana.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 TeamId = team.Id,
    ///                 Permission = "View",
    ///             },
    ///             new Grafana.Inputs.DashboardPermissionPermissionArgs
    ///             {
    ///                 UserId = user.Id,
    ///                 Permission = "Admin",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/dashboardPermission:DashboardPermission")]
    public partial class DashboardPermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the dashboard to apply permissions to.
        /// </summary>
        [Output("dashboardId")]
        public Output<int> DashboardId { get; private set; } = null!;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.DashboardPermissionPermission>> Permissions { get; private set; } = null!;


        /// <summary>
        /// Create a DashboardPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DashboardPermission(string name, DashboardPermissionArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/dashboardPermission:DashboardPermission", name, args ?? new DashboardPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DashboardPermission(string name, Input<string> id, DashboardPermissionState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/dashboardPermission:DashboardPermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DashboardPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DashboardPermission Get(string name, Input<string> id, DashboardPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new DashboardPermission(name, id, state, options);
        }
    }

    public sealed class DashboardPermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the dashboard to apply permissions to.
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<int> DashboardId { get; set; } = null!;

        [Input("permissions", required: true)]
        private InputList<Inputs.DashboardPermissionPermissionArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DashboardPermissionPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DashboardPermissionPermissionArgs>());
            set => _permissions = value;
        }

        public DashboardPermissionArgs()
        {
        }
        public static new DashboardPermissionArgs Empty => new DashboardPermissionArgs();
    }

    public sealed class DashboardPermissionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the dashboard to apply permissions to.
        /// </summary>
        [Input("dashboardId")]
        public Input<int>? DashboardId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.DashboardPermissionPermissionGetArgs>? _permissions;

        /// <summary>
        /// The permission items to add/update. Items that are omitted from the list will be removed.
        /// </summary>
        public InputList<Inputs.DashboardPermissionPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.DashboardPermissionPermissionGetArgs>());
            set => _permissions = value;
        }

        public DashboardPermissionState()
        {
        }
        public static new DashboardPermissionState Empty => new DashboardPermissionState();
    }
}
