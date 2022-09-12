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
    public static class GetOncallUserGroup
    {
        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/user_groups/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleUserGroup = Grafana.GetOncallUserGroup.Invoke(new()
        ///     {
        ///         SlackHandle = "example_slack_handle",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOncallUserGroupResult> InvokeAsync(GetOncallUserGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOncallUserGroupResult>("grafana:index/getOncallUserGroup:getOncallUserGroup", args ?? new GetOncallUserGroupArgs(), options.WithDefaults());

        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/user_groups/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleUserGroup = Grafana.GetOncallUserGroup.Invoke(new()
        ///     {
        ///         SlackHandle = "example_slack_handle",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOncallUserGroupResult> Invoke(GetOncallUserGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOncallUserGroupResult>("grafana:index/getOncallUserGroup:getOncallUserGroup", args ?? new GetOncallUserGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOncallUserGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("slackHandle", required: true)]
        public string SlackHandle { get; set; } = null!;

        public GetOncallUserGroupArgs()
        {
        }
        public static new GetOncallUserGroupArgs Empty => new GetOncallUserGroupArgs();
    }

    public sealed class GetOncallUserGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("slackHandle", required: true)]
        public Input<string> SlackHandle { get; set; } = null!;

        public GetOncallUserGroupInvokeArgs()
        {
        }
        public static new GetOncallUserGroupInvokeArgs Empty => new GetOncallUserGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetOncallUserGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string SlackHandle;
        public readonly string SlackId;

        [OutputConstructor]
        private GetOncallUserGroupResult(
            string id,

            string slackHandle,

            string slackId)
        {
            Id = id;
            SlackHandle = slackHandle;
            SlackId = slackId;
        }
    }
}
