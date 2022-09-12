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

    public sealed class SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowMissing")]
        public Input<bool>? AllowMissing { get; set; }

        [Input("header", required: true)]
        public Input<string> Header { get; set; } = null!;

        [Input("regexp", required: true)]
        public Input<string> Regexp { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArgs Empty => new SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArgs();
    }
}
