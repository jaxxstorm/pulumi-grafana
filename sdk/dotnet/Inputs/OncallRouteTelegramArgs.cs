// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Grafana.Inputs
{

    public sealed class OncallRouteTelegramArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable notification in Telegram. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Telegram channel id. Alerts will be directed to this channel in Telegram.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public OncallRouteTelegramArgs()
        {
        }
        public static new OncallRouteTelegramArgs Empty => new OncallRouteTelegramArgs();
    }
}
