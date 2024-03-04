// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// Manages Grafana Alerting contact points.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana/next/alerting/fundamentals/contact-points/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#contact-points)
    /// 
    /// This resource requires Grafana 9.1.0 or later.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myContactPoint = new Grafana.ContactPoint("myContactPoint", new()
    ///     {
    ///         Emails = new[]
    ///         {
    ///             new Grafana.Inputs.ContactPointEmailArgs
    ///             {
    ///                 Addresses = new[]
    ///                 {
    ///                     "one@company.org",
    ///                     "two@company.org",
    ///                 },
    ///                 DisableResolveMessage = false,
    ///                 Message = "{{ len .Alerts.Firing }} firing.",
    ///                 SingleEmail = true,
    ///                 Subject = "{{ template \"default.title\" .}}",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/contactPoint:ContactPoint contact_point_name {{contact_point_name}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/contactPoint:ContactPoint")]
    public partial class ContactPoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A contact point that sends notifications to other Alertmanager instances.
        /// </summary>
        [Output("alertmanagers")]
        public Output<ImmutableArray<Outputs.ContactPointAlertmanager>> Alertmanagers { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to DingDing.
        /// </summary>
        [Output("dingdings")]
        public Output<ImmutableArray<Outputs.ContactPointDingding>> Dingdings { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications as Discord messages
        /// </summary>
        [Output("discords")]
        public Output<ImmutableArray<Outputs.ContactPointDiscord>> Discords { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to an email address.
        /// </summary>
        [Output("emails")]
        public Output<ImmutableArray<Outputs.ContactPointEmail>> Emails { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Google Chat.
        /// </summary>
        [Output("googlechats")]
        public Output<ImmutableArray<Outputs.ContactPointGooglechat>> Googlechats { get; private set; } = null!;

        /// <summary>
        /// A contact point that publishes notifications to Apache Kafka topics.
        /// </summary>
        [Output("kafkas")]
        public Output<ImmutableArray<Outputs.ContactPointKafka>> Kafkas { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to LINE.me.
        /// </summary>
        [Output("lines")]
        public Output<ImmutableArray<Outputs.ContactPointLine>> Lines { get; private set; } = null!;

        /// <summary>
        /// The name of the contact point.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to OpsGenie.
        /// </summary>
        [Output("opsgenies")]
        public Output<ImmutableArray<Outputs.ContactPointOpsgeny>> Opsgenies { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to PagerDuty.
        /// </summary>
        [Output("pagerduties")]
        public Output<ImmutableArray<Outputs.ContactPointPagerduty>> Pagerduties { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Pushover.
        /// </summary>
        [Output("pushovers")]
        public Output<ImmutableArray<Outputs.ContactPointPushover>> Pushovers { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to SensuGo.
        /// </summary>
        [Output("sensugos")]
        public Output<ImmutableArray<Outputs.ContactPointSensugo>> Sensugos { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Slack.
        /// </summary>
        [Output("slacks")]
        public Output<ImmutableArray<Outputs.ContactPointSlack>> Slacks { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Microsoft Teams.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<Outputs.ContactPointTeam>> Teams { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Telegram.
        /// </summary>
        [Output("telegrams")]
        public Output<ImmutableArray<Outputs.ContactPointTelegram>> Telegrams { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Threema.
        /// </summary>
        [Output("threemas")]
        public Output<ImmutableArray<Outputs.ContactPointThreema>> Threemas { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
        /// </summary>
        [Output("victorops")]
        public Output<ImmutableArray<Outputs.ContactPointVictorop>> Victorops { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to Cisco Webex.
        /// </summary>
        [Output("webexes")]
        public Output<ImmutableArray<Outputs.ContactPointWebex>> Webexes { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
        /// </summary>
        [Output("webhooks")]
        public Output<ImmutableArray<Outputs.ContactPointWebhook>> Webhooks { get; private set; } = null!;

        /// <summary>
        /// A contact point that sends notifications to WeCom.
        /// </summary>
        [Output("wecoms")]
        public Output<ImmutableArray<Outputs.ContactPointWecom>> Wecoms { get; private set; } = null!;


        /// <summary>
        /// Create a ContactPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContactPoint(string name, ContactPointArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/contactPoint:ContactPoint", name, args ?? new ContactPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContactPoint(string name, Input<string> id, ContactPointState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/contactPoint:ContactPoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContactPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContactPoint Get(string name, Input<string> id, ContactPointState? state = null, CustomResourceOptions? options = null)
        {
            return new ContactPoint(name, id, state, options);
        }
    }

    public sealed class ContactPointArgs : global::Pulumi.ResourceArgs
    {
        [Input("alertmanagers")]
        private InputList<Inputs.ContactPointAlertmanagerArgs>? _alertmanagers;

        /// <summary>
        /// A contact point that sends notifications to other Alertmanager instances.
        /// </summary>
        public InputList<Inputs.ContactPointAlertmanagerArgs> Alertmanagers
        {
            get => _alertmanagers ?? (_alertmanagers = new InputList<Inputs.ContactPointAlertmanagerArgs>());
            set => _alertmanagers = value;
        }

        [Input("dingdings")]
        private InputList<Inputs.ContactPointDingdingArgs>? _dingdings;

        /// <summary>
        /// A contact point that sends notifications to DingDing.
        /// </summary>
        public InputList<Inputs.ContactPointDingdingArgs> Dingdings
        {
            get => _dingdings ?? (_dingdings = new InputList<Inputs.ContactPointDingdingArgs>());
            set => _dingdings = value;
        }

        [Input("discords")]
        private InputList<Inputs.ContactPointDiscordArgs>? _discords;

        /// <summary>
        /// A contact point that sends notifications as Discord messages
        /// </summary>
        public InputList<Inputs.ContactPointDiscordArgs> Discords
        {
            get => _discords ?? (_discords = new InputList<Inputs.ContactPointDiscordArgs>());
            set => _discords = value;
        }

        [Input("emails")]
        private InputList<Inputs.ContactPointEmailArgs>? _emails;

        /// <summary>
        /// A contact point that sends notifications to an email address.
        /// </summary>
        public InputList<Inputs.ContactPointEmailArgs> Emails
        {
            get => _emails ?? (_emails = new InputList<Inputs.ContactPointEmailArgs>());
            set => _emails = value;
        }

        [Input("googlechats")]
        private InputList<Inputs.ContactPointGooglechatArgs>? _googlechats;

        /// <summary>
        /// A contact point that sends notifications to Google Chat.
        /// </summary>
        public InputList<Inputs.ContactPointGooglechatArgs> Googlechats
        {
            get => _googlechats ?? (_googlechats = new InputList<Inputs.ContactPointGooglechatArgs>());
            set => _googlechats = value;
        }

        [Input("kafkas")]
        private InputList<Inputs.ContactPointKafkaArgs>? _kafkas;

        /// <summary>
        /// A contact point that publishes notifications to Apache Kafka topics.
        /// </summary>
        public InputList<Inputs.ContactPointKafkaArgs> Kafkas
        {
            get => _kafkas ?? (_kafkas = new InputList<Inputs.ContactPointKafkaArgs>());
            set => _kafkas = value;
        }

        [Input("lines")]
        private InputList<Inputs.ContactPointLineArgs>? _lines;

        /// <summary>
        /// A contact point that sends notifications to LINE.me.
        /// </summary>
        public InputList<Inputs.ContactPointLineArgs> Lines
        {
            get => _lines ?? (_lines = new InputList<Inputs.ContactPointLineArgs>());
            set => _lines = value;
        }

        /// <summary>
        /// The name of the contact point.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("opsgenies")]
        private InputList<Inputs.ContactPointOpsgenyArgs>? _opsgenies;

        /// <summary>
        /// A contact point that sends notifications to OpsGenie.
        /// </summary>
        public InputList<Inputs.ContactPointOpsgenyArgs> Opsgenies
        {
            get => _opsgenies ?? (_opsgenies = new InputList<Inputs.ContactPointOpsgenyArgs>());
            set => _opsgenies = value;
        }

        [Input("pagerduties")]
        private InputList<Inputs.ContactPointPagerdutyArgs>? _pagerduties;

        /// <summary>
        /// A contact point that sends notifications to PagerDuty.
        /// </summary>
        public InputList<Inputs.ContactPointPagerdutyArgs> Pagerduties
        {
            get => _pagerduties ?? (_pagerduties = new InputList<Inputs.ContactPointPagerdutyArgs>());
            set => _pagerduties = value;
        }

        [Input("pushovers")]
        private InputList<Inputs.ContactPointPushoverArgs>? _pushovers;

        /// <summary>
        /// A contact point that sends notifications to Pushover.
        /// </summary>
        public InputList<Inputs.ContactPointPushoverArgs> Pushovers
        {
            get => _pushovers ?? (_pushovers = new InputList<Inputs.ContactPointPushoverArgs>());
            set => _pushovers = value;
        }

        [Input("sensugos")]
        private InputList<Inputs.ContactPointSensugoArgs>? _sensugos;

        /// <summary>
        /// A contact point that sends notifications to SensuGo.
        /// </summary>
        public InputList<Inputs.ContactPointSensugoArgs> Sensugos
        {
            get => _sensugos ?? (_sensugos = new InputList<Inputs.ContactPointSensugoArgs>());
            set => _sensugos = value;
        }

        [Input("slacks")]
        private InputList<Inputs.ContactPointSlackArgs>? _slacks;

        /// <summary>
        /// A contact point that sends notifications to Slack.
        /// </summary>
        public InputList<Inputs.ContactPointSlackArgs> Slacks
        {
            get => _slacks ?? (_slacks = new InputList<Inputs.ContactPointSlackArgs>());
            set => _slacks = value;
        }

        [Input("teams")]
        private InputList<Inputs.ContactPointTeamArgs>? _teams;

        /// <summary>
        /// A contact point that sends notifications to Microsoft Teams.
        /// </summary>
        public InputList<Inputs.ContactPointTeamArgs> Teams
        {
            get => _teams ?? (_teams = new InputList<Inputs.ContactPointTeamArgs>());
            set => _teams = value;
        }

        [Input("telegrams")]
        private InputList<Inputs.ContactPointTelegramArgs>? _telegrams;

        /// <summary>
        /// A contact point that sends notifications to Telegram.
        /// </summary>
        public InputList<Inputs.ContactPointTelegramArgs> Telegrams
        {
            get => _telegrams ?? (_telegrams = new InputList<Inputs.ContactPointTelegramArgs>());
            set => _telegrams = value;
        }

        [Input("threemas")]
        private InputList<Inputs.ContactPointThreemaArgs>? _threemas;

        /// <summary>
        /// A contact point that sends notifications to Threema.
        /// </summary>
        public InputList<Inputs.ContactPointThreemaArgs> Threemas
        {
            get => _threemas ?? (_threemas = new InputList<Inputs.ContactPointThreemaArgs>());
            set => _threemas = value;
        }

        [Input("victorops")]
        private InputList<Inputs.ContactPointVictoropArgs>? _victorops;

        /// <summary>
        /// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
        /// </summary>
        public InputList<Inputs.ContactPointVictoropArgs> Victorops
        {
            get => _victorops ?? (_victorops = new InputList<Inputs.ContactPointVictoropArgs>());
            set => _victorops = value;
        }

        [Input("webexes")]
        private InputList<Inputs.ContactPointWebexArgs>? _webexes;

        /// <summary>
        /// A contact point that sends notifications to Cisco Webex.
        /// </summary>
        public InputList<Inputs.ContactPointWebexArgs> Webexes
        {
            get => _webexes ?? (_webexes = new InputList<Inputs.ContactPointWebexArgs>());
            set => _webexes = value;
        }

        [Input("webhooks")]
        private InputList<Inputs.ContactPointWebhookArgs>? _webhooks;

        /// <summary>
        /// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
        /// </summary>
        public InputList<Inputs.ContactPointWebhookArgs> Webhooks
        {
            get => _webhooks ?? (_webhooks = new InputList<Inputs.ContactPointWebhookArgs>());
            set => _webhooks = value;
        }

        [Input("wecoms")]
        private InputList<Inputs.ContactPointWecomArgs>? _wecoms;

        /// <summary>
        /// A contact point that sends notifications to WeCom.
        /// </summary>
        public InputList<Inputs.ContactPointWecomArgs> Wecoms
        {
            get => _wecoms ?? (_wecoms = new InputList<Inputs.ContactPointWecomArgs>());
            set => _wecoms = value;
        }

        public ContactPointArgs()
        {
        }
        public static new ContactPointArgs Empty => new ContactPointArgs();
    }

    public sealed class ContactPointState : global::Pulumi.ResourceArgs
    {
        [Input("alertmanagers")]
        private InputList<Inputs.ContactPointAlertmanagerGetArgs>? _alertmanagers;

        /// <summary>
        /// A contact point that sends notifications to other Alertmanager instances.
        /// </summary>
        public InputList<Inputs.ContactPointAlertmanagerGetArgs> Alertmanagers
        {
            get => _alertmanagers ?? (_alertmanagers = new InputList<Inputs.ContactPointAlertmanagerGetArgs>());
            set => _alertmanagers = value;
        }

        [Input("dingdings")]
        private InputList<Inputs.ContactPointDingdingGetArgs>? _dingdings;

        /// <summary>
        /// A contact point that sends notifications to DingDing.
        /// </summary>
        public InputList<Inputs.ContactPointDingdingGetArgs> Dingdings
        {
            get => _dingdings ?? (_dingdings = new InputList<Inputs.ContactPointDingdingGetArgs>());
            set => _dingdings = value;
        }

        [Input("discords")]
        private InputList<Inputs.ContactPointDiscordGetArgs>? _discords;

        /// <summary>
        /// A contact point that sends notifications as Discord messages
        /// </summary>
        public InputList<Inputs.ContactPointDiscordGetArgs> Discords
        {
            get => _discords ?? (_discords = new InputList<Inputs.ContactPointDiscordGetArgs>());
            set => _discords = value;
        }

        [Input("emails")]
        private InputList<Inputs.ContactPointEmailGetArgs>? _emails;

        /// <summary>
        /// A contact point that sends notifications to an email address.
        /// </summary>
        public InputList<Inputs.ContactPointEmailGetArgs> Emails
        {
            get => _emails ?? (_emails = new InputList<Inputs.ContactPointEmailGetArgs>());
            set => _emails = value;
        }

        [Input("googlechats")]
        private InputList<Inputs.ContactPointGooglechatGetArgs>? _googlechats;

        /// <summary>
        /// A contact point that sends notifications to Google Chat.
        /// </summary>
        public InputList<Inputs.ContactPointGooglechatGetArgs> Googlechats
        {
            get => _googlechats ?? (_googlechats = new InputList<Inputs.ContactPointGooglechatGetArgs>());
            set => _googlechats = value;
        }

        [Input("kafkas")]
        private InputList<Inputs.ContactPointKafkaGetArgs>? _kafkas;

        /// <summary>
        /// A contact point that publishes notifications to Apache Kafka topics.
        /// </summary>
        public InputList<Inputs.ContactPointKafkaGetArgs> Kafkas
        {
            get => _kafkas ?? (_kafkas = new InputList<Inputs.ContactPointKafkaGetArgs>());
            set => _kafkas = value;
        }

        [Input("lines")]
        private InputList<Inputs.ContactPointLineGetArgs>? _lines;

        /// <summary>
        /// A contact point that sends notifications to LINE.me.
        /// </summary>
        public InputList<Inputs.ContactPointLineGetArgs> Lines
        {
            get => _lines ?? (_lines = new InputList<Inputs.ContactPointLineGetArgs>());
            set => _lines = value;
        }

        /// <summary>
        /// The name of the contact point.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("opsgenies")]
        private InputList<Inputs.ContactPointOpsgenyGetArgs>? _opsgenies;

        /// <summary>
        /// A contact point that sends notifications to OpsGenie.
        /// </summary>
        public InputList<Inputs.ContactPointOpsgenyGetArgs> Opsgenies
        {
            get => _opsgenies ?? (_opsgenies = new InputList<Inputs.ContactPointOpsgenyGetArgs>());
            set => _opsgenies = value;
        }

        [Input("pagerduties")]
        private InputList<Inputs.ContactPointPagerdutyGetArgs>? _pagerduties;

        /// <summary>
        /// A contact point that sends notifications to PagerDuty.
        /// </summary>
        public InputList<Inputs.ContactPointPagerdutyGetArgs> Pagerduties
        {
            get => _pagerduties ?? (_pagerduties = new InputList<Inputs.ContactPointPagerdutyGetArgs>());
            set => _pagerduties = value;
        }

        [Input("pushovers")]
        private InputList<Inputs.ContactPointPushoverGetArgs>? _pushovers;

        /// <summary>
        /// A contact point that sends notifications to Pushover.
        /// </summary>
        public InputList<Inputs.ContactPointPushoverGetArgs> Pushovers
        {
            get => _pushovers ?? (_pushovers = new InputList<Inputs.ContactPointPushoverGetArgs>());
            set => _pushovers = value;
        }

        [Input("sensugos")]
        private InputList<Inputs.ContactPointSensugoGetArgs>? _sensugos;

        /// <summary>
        /// A contact point that sends notifications to SensuGo.
        /// </summary>
        public InputList<Inputs.ContactPointSensugoGetArgs> Sensugos
        {
            get => _sensugos ?? (_sensugos = new InputList<Inputs.ContactPointSensugoGetArgs>());
            set => _sensugos = value;
        }

        [Input("slacks")]
        private InputList<Inputs.ContactPointSlackGetArgs>? _slacks;

        /// <summary>
        /// A contact point that sends notifications to Slack.
        /// </summary>
        public InputList<Inputs.ContactPointSlackGetArgs> Slacks
        {
            get => _slacks ?? (_slacks = new InputList<Inputs.ContactPointSlackGetArgs>());
            set => _slacks = value;
        }

        [Input("teams")]
        private InputList<Inputs.ContactPointTeamGetArgs>? _teams;

        /// <summary>
        /// A contact point that sends notifications to Microsoft Teams.
        /// </summary>
        public InputList<Inputs.ContactPointTeamGetArgs> Teams
        {
            get => _teams ?? (_teams = new InputList<Inputs.ContactPointTeamGetArgs>());
            set => _teams = value;
        }

        [Input("telegrams")]
        private InputList<Inputs.ContactPointTelegramGetArgs>? _telegrams;

        /// <summary>
        /// A contact point that sends notifications to Telegram.
        /// </summary>
        public InputList<Inputs.ContactPointTelegramGetArgs> Telegrams
        {
            get => _telegrams ?? (_telegrams = new InputList<Inputs.ContactPointTelegramGetArgs>());
            set => _telegrams = value;
        }

        [Input("threemas")]
        private InputList<Inputs.ContactPointThreemaGetArgs>? _threemas;

        /// <summary>
        /// A contact point that sends notifications to Threema.
        /// </summary>
        public InputList<Inputs.ContactPointThreemaGetArgs> Threemas
        {
            get => _threemas ?? (_threemas = new InputList<Inputs.ContactPointThreemaGetArgs>());
            set => _threemas = value;
        }

        [Input("victorops")]
        private InputList<Inputs.ContactPointVictoropGetArgs>? _victorops;

        /// <summary>
        /// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
        /// </summary>
        public InputList<Inputs.ContactPointVictoropGetArgs> Victorops
        {
            get => _victorops ?? (_victorops = new InputList<Inputs.ContactPointVictoropGetArgs>());
            set => _victorops = value;
        }

        [Input("webexes")]
        private InputList<Inputs.ContactPointWebexGetArgs>? _webexes;

        /// <summary>
        /// A contact point that sends notifications to Cisco Webex.
        /// </summary>
        public InputList<Inputs.ContactPointWebexGetArgs> Webexes
        {
            get => _webexes ?? (_webexes = new InputList<Inputs.ContactPointWebexGetArgs>());
            set => _webexes = value;
        }

        [Input("webhooks")]
        private InputList<Inputs.ContactPointWebhookGetArgs>? _webhooks;

        /// <summary>
        /// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
        /// </summary>
        public InputList<Inputs.ContactPointWebhookGetArgs> Webhooks
        {
            get => _webhooks ?? (_webhooks = new InputList<Inputs.ContactPointWebhookGetArgs>());
            set => _webhooks = value;
        }

        [Input("wecoms")]
        private InputList<Inputs.ContactPointWecomGetArgs>? _wecoms;

        /// <summary>
        /// A contact point that sends notifications to WeCom.
        /// </summary>
        public InputList<Inputs.ContactPointWecomGetArgs> Wecoms
        {
            get => _wecoms ?? (_wecoms = new InputList<Inputs.ContactPointWecomGetArgs>());
            set => _wecoms = value;
        }

        public ContactPointState()
        {
        }
        public static new ContactPointState Empty => new ContactPointState();
    }
}