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

    public sealed class SyntheticMonitoringCheckSettingsTracerouteGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxHops")]
        public Input<int>? MaxHops { get; set; }

        [Input("maxUnknownHops")]
        public Input<int>? MaxUnknownHops { get; set; }

        [Input("ptrLookup")]
        public Input<bool>? PtrLookup { get; set; }

        public SyntheticMonitoringCheckSettingsTracerouteGetArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsTracerouteGetArgs Empty => new SyntheticMonitoringCheckSettingsTracerouteGetArgs();
    }
}
