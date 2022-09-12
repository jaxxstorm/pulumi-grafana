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
    public sealed class ContactPointWebhook
    {
        /// <summary>
        /// Allows a custom authorization scheme - attaches an auth header with this value. Do not use in conjunction with basic auth parameters.
        /// </summary>
        public readonly string? AuthorizationCredentials;
        /// <summary>
        /// Allows a custom authorization scheme - attaches an auth header with this name. Do not use in conjunction with basic auth parameters.
        /// </summary>
        public readonly string? AuthorizationScheme;
        /// <summary>
        /// The username to use in basic auth headers attached to the request. If omitted, basic auth will not be used.
        /// </summary>
        public readonly string? BasicAuthPassword;
        /// <summary>
        /// The username to use in basic auth headers attached to the request. If omitted, basic auth will not be used.
        /// </summary>
        public readonly string? BasicAuthUser;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// The HTTP method to use in the request. Defaults to `POST`.
        /// </summary>
        public readonly string? HttpMethod;
        /// <summary>
        /// The maximum number of alerts to send in a single request. This can be helpful in limiting the size of the request body. The default is 0, which indicates no limit.
        /// </summary>
        public readonly int? MaxAlerts;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;
        /// <summary>
        /// The URL to send webhook requests to.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private ContactPointWebhook(
            string? authorizationCredentials,

            string? authorizationScheme,

            string? basicAuthPassword,

            string? basicAuthUser,

            bool? disableResolveMessage,

            string? httpMethod,

            int? maxAlerts,

            ImmutableDictionary<string, string>? settings,

            string? uid,

            string url)
        {
            AuthorizationCredentials = authorizationCredentials;
            AuthorizationScheme = authorizationScheme;
            BasicAuthPassword = basicAuthPassword;
            BasicAuthUser = basicAuthUser;
            DisableResolveMessage = disableResolveMessage;
            HttpMethod = httpMethod;
            MaxAlerts = maxAlerts;
            Settings = settings;
            Uid = uid;
            Url = url;
        }
    }
}
