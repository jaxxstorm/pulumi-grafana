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
    public sealed class RolePermission
    {
        /// <summary>
        /// Specific action users granted with the role will be allowed to perform (for example: `users:read`)
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`)
        /// </summary>
        public readonly string? Scope;

        [OutputConstructor]
        private RolePermission(
            string action,

            string? scope)
        {
            Action = action;
            Scope = scope;
        }
    }
}
