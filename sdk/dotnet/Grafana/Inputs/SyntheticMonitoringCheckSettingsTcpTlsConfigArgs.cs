// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Inputs
{

    public sealed class SyntheticMonitoringCheckSettingsTcpTlsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        [Input("insecureSkipVerify")]
        public Input<bool>? InsecureSkipVerify { get; set; }

        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        public SyntheticMonitoringCheckSettingsTcpTlsConfigArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsTcpTlsConfigArgs Empty => new SyntheticMonitoringCheckSettingsTcpTlsConfigArgs();
    }
}
