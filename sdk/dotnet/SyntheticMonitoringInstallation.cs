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
    /// Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
    /// Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
    /// This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/installation/)
    /// * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Grafana = Lbrlabs_Pulumi.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var smStackCloudStack = new Grafana.CloudStack("smStackCloudStack", new()
    ///     {
    ///         Slug = "&lt;stack-slug&gt;",
    ///         RegionSlug = "us",
    ///     });
    /// 
    ///     var metricsPublish = new Grafana.CloudApiKey("metricsPublish", new()
    ///     {
    ///         Role = "MetricsPublisher",
    ///         CloudOrgSlug = "&lt;org-slug&gt;",
    ///     });
    /// 
    ///     var smStackSyntheticMonitoringInstallation = new Grafana.SyntheticMonitoringInstallation("smStackSyntheticMonitoringInstallation", new()
    ///     {
    ///         StackId = smStackCloudStack.Id,
    ///         MetricsInstanceId = smStackCloudStack.PrometheusUserId,
    ///         LogsInstanceId = smStackCloudStack.LogsUserId,
    ///         MetricsPublisherKey = metricsPublish.Key,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation")]
    public partial class SyntheticMonitoringInstallation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the logs instance to install SM on (stack's `logs_user_id` attribute).
        /// </summary>
        [Output("logsInstanceId")]
        public Output<int> LogsInstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the metrics instance to install SM on (stack's `prometheus_user_id` attribute).
        /// </summary>
        [Output("metricsInstanceId")]
        public Output<int> MetricsInstanceId { get; private set; } = null!;

        /// <summary>
        /// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
        /// </summary>
        [Output("metricsPublisherKey")]
        public Output<string> MetricsPublisherKey { get; private set; } = null!;

        /// <summary>
        /// Generated token to access the SM API.
        /// </summary>
        [Output("smAccessToken")]
        public Output<string> SmAccessToken { get; private set; } = null!;

        /// <summary>
        /// The ID of the stack to install SM on.
        /// </summary>
        [Output("stackId")]
        public Output<int> StackId { get; private set; } = null!;


        /// <summary>
        /// Create a SyntheticMonitoringInstallation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyntheticMonitoringInstallation(string name, SyntheticMonitoringInstallationArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, args ?? new SyntheticMonitoringInstallationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyntheticMonitoringInstallation(string name, Input<string> id, SyntheticMonitoringInstallationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyntheticMonitoringInstallation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyntheticMonitoringInstallation Get(string name, Input<string> id, SyntheticMonitoringInstallationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyntheticMonitoringInstallation(name, id, state, options);
        }
    }

    public sealed class SyntheticMonitoringInstallationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the logs instance to install SM on (stack's `logs_user_id` attribute).
        /// </summary>
        [Input("logsInstanceId", required: true)]
        public Input<int> LogsInstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the metrics instance to install SM on (stack's `prometheus_user_id` attribute).
        /// </summary>
        [Input("metricsInstanceId", required: true)]
        public Input<int> MetricsInstanceId { get; set; } = null!;

        /// <summary>
        /// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
        /// </summary>
        [Input("metricsPublisherKey", required: true)]
        public Input<string> MetricsPublisherKey { get; set; } = null!;

        /// <summary>
        /// The ID of the stack to install SM on.
        /// </summary>
        [Input("stackId", required: true)]
        public Input<int> StackId { get; set; } = null!;

        public SyntheticMonitoringInstallationArgs()
        {
        }
        public static new SyntheticMonitoringInstallationArgs Empty => new SyntheticMonitoringInstallationArgs();
    }

    public sealed class SyntheticMonitoringInstallationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the logs instance to install SM on (stack's `logs_user_id` attribute).
        /// </summary>
        [Input("logsInstanceId")]
        public Input<int>? LogsInstanceId { get; set; }

        /// <summary>
        /// The ID of the metrics instance to install SM on (stack's `prometheus_user_id` attribute).
        /// </summary>
        [Input("metricsInstanceId")]
        public Input<int>? MetricsInstanceId { get; set; }

        /// <summary>
        /// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
        /// </summary>
        [Input("metricsPublisherKey")]
        public Input<string>? MetricsPublisherKey { get; set; }

        /// <summary>
        /// Generated token to access the SM API.
        /// </summary>
        [Input("smAccessToken")]
        public Input<string>? SmAccessToken { get; set; }

        /// <summary>
        /// The ID of the stack to install SM on.
        /// </summary>
        [Input("stackId")]
        public Input<int>? StackId { get; set; }

        public SyntheticMonitoringInstallationState()
        {
        }
        public static new SyntheticMonitoringInstallationState Empty => new SyntheticMonitoringInstallationState();
    }
}
