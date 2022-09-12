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

    public sealed class ContactPointDiscordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of a custom avatar image to use. Defaults to ``.
        /// </summary>
        [Input("avatarUrl")]
        public Input<string>? AvatarUrl { get; set; }

        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        /// <summary>
        /// The templated content of the message. Defaults to ``.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

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
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The discord webhook URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Whether to use the bot account's plain username instead of "Grafana." Defaults to `false`.
        /// </summary>
        [Input("useDiscordUsername")]
        public Input<bool>? UseDiscordUsername { get; set; }

        public ContactPointDiscordArgs()
        {
        }
        public static new ContactPointDiscordArgs Empty => new ContactPointDiscordArgs();
    }
}
