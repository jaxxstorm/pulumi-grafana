// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class SyntheticMonitoringCheckSettingsMultihttpEntryRequestBodyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content encoding of the body
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// The content type of the body
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The body payload
        /// </summary>
        [Input("payload")]
        public Input<string>? Payload { get; set; }

        public SyntheticMonitoringCheckSettingsMultihttpEntryRequestBodyArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsMultihttpEntryRequestBodyArgs Empty => new SyntheticMonitoringCheckSettingsMultihttpEntryRequestBodyArgs();
    }
}
