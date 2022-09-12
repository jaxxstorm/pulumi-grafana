// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Grafana.Outputs
{

    [OutputType]
    public sealed class OncallIntegrationTemplatesSlack
    {
        public readonly string? ImageUrl;
        public readonly string? Message;
        public readonly string? Title;

        [OutputConstructor]
        private OncallIntegrationTemplatesSlack(
            string? imageUrl,

            string? message,

            string? title)
        {
            ImageUrl = imageUrl;
            Message = message;
            Title = title;
        }
    }
}
