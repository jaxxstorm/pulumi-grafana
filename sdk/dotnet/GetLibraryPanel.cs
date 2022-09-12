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
    public static class GetLibraryPanel
    {
        /// <summary>
        /// Data source for retrieving a single library panel by name or uid.
        /// </summary>
        public static Task<GetLibraryPanelResult> InvokeAsync(GetLibraryPanelArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLibraryPanelResult>("grafana:index/getLibraryPanel:getLibraryPanel", args ?? new GetLibraryPanelArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for retrieving a single library panel by name or uid.
        /// </summary>
        public static Output<GetLibraryPanelResult> Invoke(GetLibraryPanelInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLibraryPanelResult>("grafana:index/getLibraryPanel:getLibraryPanel", args ?? new GetLibraryPanelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLibraryPanelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the library panel.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The unique identifier (UID) of the library panel.
        /// </summary>
        [Input("uid")]
        public string? Uid { get; set; }

        public GetLibraryPanelArgs()
        {
        }
        public static new GetLibraryPanelArgs Empty => new GetLibraryPanelArgs();
    }

    public sealed class GetLibraryPanelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the library panel.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique identifier (UID) of the library panel.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public GetLibraryPanelInvokeArgs()
        {
        }
        public static new GetLibraryPanelInvokeArgs Empty => new GetLibraryPanelInvokeArgs();
    }


    [OutputType]
    public sealed class GetLibraryPanelResult
    {
        /// <summary>
        /// Timestamp when the library panel was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// Numerical IDs of Grafana dashboards containing the library panel.
        /// </summary>
        public readonly ImmutableArray<int> DashboardIds;
        /// <summary>
        /// Description of the library panel.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the folder where the library panel is stored.
        /// </summary>
        public readonly int FolderId;
        /// <summary>
        /// Name of the folder containing the library panel.
        /// </summary>
        public readonly string FolderName;
        /// <summary>
        /// Unique ID (UID) of the folder containing the library panel.
        /// </summary>
        public readonly string FolderUid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The JSON model for the library panel.
        /// </summary>
        public readonly string ModelJson;
        /// <summary>
        /// Name of the library panel.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The numeric ID of the library panel computed by Grafana.
        /// </summary>
        public readonly int OrgId;
        /// <summary>
        /// The numeric ID of the library panel computed by Grafana.
        /// </summary>
        public readonly int PanelId;
        /// <summary>
        /// Type of the library panel (eg. text).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique identifier (UID) of the library panel.
        /// </summary>
        public readonly string? Uid;
        /// <summary>
        /// Timestamp when the library panel was last modified.
        /// </summary>
        public readonly string Updated;
        /// <summary>
        /// Version of the library panel.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetLibraryPanelResult(
            string created,

            ImmutableArray<int> dashboardIds,

            string description,

            int folderId,

            string folderName,

            string folderUid,

            string id,

            string modelJson,

            string? name,

            int orgId,

            int panelId,

            string type,

            string? uid,

            string updated,

            int version)
        {
            Created = created;
            DashboardIds = dashboardIds;
            Description = description;
            FolderId = folderId;
            FolderName = folderName;
            FolderUid = folderUid;
            Id = id;
            ModelJson = modelJson;
            Name = name;
            OrgId = orgId;
            PanelId = panelId;
            Type = type;
            Uid = uid;
            Updated = updated;
            Version = version;
        }
    }
}
