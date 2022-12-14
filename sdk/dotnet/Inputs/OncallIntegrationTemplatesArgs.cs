// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana.Inputs
{

    public sealed class OncallIntegrationTemplatesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template for the key by which alerts are grouped.
        /// </summary>
        [Input("groupingKey")]
        public Input<string>? GroupingKey { get; set; }

        /// <summary>
        /// Template for sending a signal to resolve the Incident.
        /// </summary>
        [Input("resolveSignal")]
        public Input<string>? ResolveSignal { get; set; }

        /// <summary>
        /// Templates for Slack.
        /// </summary>
        [Input("slack")]
        public Input<Inputs.OncallIntegrationTemplatesSlackArgs>? Slack { get; set; }

        public OncallIntegrationTemplatesArgs()
        {
        }
        public static new OncallIntegrationTemplatesArgs Empty => new OncallIntegrationTemplatesArgs();
    }
}
