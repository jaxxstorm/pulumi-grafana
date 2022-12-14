// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Grafana.Inputs
{

    public sealed class ContactPointPagerdutyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The class or type of event, for example `ping failure`.
        /// </summary>
        [Input("class")]
        public Input<string>? Class { get; set; }

        /// <summary>
        /// The component being affected by the event.
        /// </summary>
        [Input("component")]
        public Input<string>? Component { get; set; }

        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        /// <summary>
        /// The group to which the provided component belongs to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The PagerDuty API key.
        /// </summary>
        [Input("integrationKey", required: true)]
        public Input<string> IntegrationKey { get; set; } = null!;

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        /// <summary>
        /// The PagerDuty event severity level. Default is `critical`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// The templated summary message of the event.
        /// </summary>
        [Input("summary")]
        public Input<string>? Summary { get; set; }

        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ContactPointPagerdutyArgs()
        {
        }
        public static new ContactPointPagerdutyArgs Empty => new ContactPointPagerdutyArgs();
    }
}
