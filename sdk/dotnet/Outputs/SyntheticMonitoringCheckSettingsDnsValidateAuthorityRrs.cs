// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana.Outputs
{

    [OutputType]
    public sealed class SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs
    {
        public readonly ImmutableArray<string> FailIfMatchesRegexps;
        public readonly ImmutableArray<string> FailIfNotMatchesRegexps;

        [OutputConstructor]
        private SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs(
            ImmutableArray<string> failIfMatchesRegexps,

            ImmutableArray<string> failIfNotMatchesRegexps)
        {
            FailIfMatchesRegexps = failIfMatchesRegexps;
            FailIfNotMatchesRegexps = failIfNotMatchesRegexps;
        }
    }
}
