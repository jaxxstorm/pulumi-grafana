// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages Grafana Alerting contact points.
//
// * [Official documentation](https://grafana.com/docs/grafana/next/alerting/fundamentals/contact-points/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#contact-points)
//
// This resource requires Grafana 9.1.0 or later.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.NewContactPoint(ctx, "myContactPoint", &grafana.ContactPointArgs{
//				Emails: grafana.ContactPointEmailArray{
//					&grafana.ContactPointEmailArgs{
//						Addresses: pulumi.StringArray{
//							pulumi.String("one@company.org"),
//							pulumi.String("two@company.org"),
//						},
//						DisableResolveMessage: pulumi.Bool(false),
//						Message:               pulumi.String("{{ len .Alerts.Firing }} firing."),
//						SingleEmail:           pulumi.Bool(true),
//						Subject:               pulumi.String("{{ template \"default.title\" .}}"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import grafana:index/contactPoint:ContactPoint contact_point_name {{contact_point_name}}
//
// ```
type ContactPoint struct {
	pulumi.CustomResourceState

	// A contact point that sends notifications to other Alertmanager instances.
	Alertmanagers ContactPointAlertmanagerArrayOutput `pulumi:"alertmanagers"`
	// A contact point that sends notifications to DingDing.
	Dingdings ContactPointDingdingArrayOutput `pulumi:"dingdings"`
	// Allow modifying the contact point from other sources than Terraform or the Grafana API.
	DisableProvenance pulumi.BoolPtrOutput `pulumi:"disableProvenance"`
	// A contact point that sends notifications as Discord messages
	Discords ContactPointDiscordArrayOutput `pulumi:"discords"`
	// A contact point that sends notifications to an email address.
	Emails ContactPointEmailArrayOutput `pulumi:"emails"`
	// A contact point that sends notifications to Google Chat.
	Googlechats ContactPointGooglechatArrayOutput `pulumi:"googlechats"`
	// A contact point that publishes notifications to Apache Kafka topics.
	Kafkas ContactPointKafkaArrayOutput `pulumi:"kafkas"`
	// A contact point that sends notifications to LINE.me.
	Lines ContactPointLineArrayOutput `pulumi:"lines"`
	// Name of the responder. Must be specified if username and id are empty.
	Name pulumi.StringOutput `pulumi:"name"`
	// A contact point that sends notifications to Grafana On-Call.
	Oncalls ContactPointOncallArrayOutput `pulumi:"oncalls"`
	// A contact point that sends notifications to OpsGenie.
	Opsgenies ContactPointOpsgenyArrayOutput `pulumi:"opsgenies"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// A contact point that sends notifications to PagerDuty.
	Pagerduties ContactPointPagerdutyArrayOutput `pulumi:"pagerduties"`
	// A contact point that sends notifications to Pushover.
	Pushovers ContactPointPushoverArrayOutput `pulumi:"pushovers"`
	// A contact point that sends notifications to SensuGo.
	Sensugos ContactPointSensugoArrayOutput `pulumi:"sensugos"`
	// A contact point that sends notifications to Slack.
	Slacks ContactPointSlackArrayOutput `pulumi:"slacks"`
	// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
	Sns ContactPointSnArrayOutput `pulumi:"sns"`
	// A contact point that sends notifications to Microsoft Teams.
	Teams ContactPointTeamArrayOutput `pulumi:"teams"`
	// A contact point that sends notifications to Telegram.
	Telegrams ContactPointTelegramArrayOutput `pulumi:"telegrams"`
	// A contact point that sends notifications to Threema.
	Threemas ContactPointThreemaArrayOutput `pulumi:"threemas"`
	// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
	Victorops ContactPointVictoropArrayOutput `pulumi:"victorops"`
	// A contact point that sends notifications to Cisco Webex.
	Webexes ContactPointWebexArrayOutput `pulumi:"webexes"`
	// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	Webhooks ContactPointWebhookArrayOutput `pulumi:"webhooks"`
	// A contact point that sends notifications to WeCom.
	Wecoms ContactPointWecomArrayOutput `pulumi:"wecoms"`
}

// NewContactPoint registers a new resource with the given unique name, arguments, and options.
func NewContactPoint(ctx *pulumi.Context,
	name string, args *ContactPointArgs, opts ...pulumi.ResourceOption) (*ContactPoint, error) {
	if args == nil {
		args = &ContactPointArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContactPoint
	err := ctx.RegisterResource("grafana:index/contactPoint:ContactPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactPoint gets an existing ContactPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactPointState, opts ...pulumi.ResourceOption) (*ContactPoint, error) {
	var resource ContactPoint
	err := ctx.ReadResource("grafana:index/contactPoint:ContactPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactPoint resources.
type contactPointState struct {
	// A contact point that sends notifications to other Alertmanager instances.
	Alertmanagers []ContactPointAlertmanager `pulumi:"alertmanagers"`
	// A contact point that sends notifications to DingDing.
	Dingdings []ContactPointDingding `pulumi:"dingdings"`
	// Allow modifying the contact point from other sources than Terraform or the Grafana API.
	DisableProvenance *bool `pulumi:"disableProvenance"`
	// A contact point that sends notifications as Discord messages
	Discords []ContactPointDiscord `pulumi:"discords"`
	// A contact point that sends notifications to an email address.
	Emails []ContactPointEmail `pulumi:"emails"`
	// A contact point that sends notifications to Google Chat.
	Googlechats []ContactPointGooglechat `pulumi:"googlechats"`
	// A contact point that publishes notifications to Apache Kafka topics.
	Kafkas []ContactPointKafka `pulumi:"kafkas"`
	// A contact point that sends notifications to LINE.me.
	Lines []ContactPointLine `pulumi:"lines"`
	// Name of the responder. Must be specified if username and id are empty.
	Name *string `pulumi:"name"`
	// A contact point that sends notifications to Grafana On-Call.
	Oncalls []ContactPointOncall `pulumi:"oncalls"`
	// A contact point that sends notifications to OpsGenie.
	Opsgenies []ContactPointOpsgeny `pulumi:"opsgenies"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// A contact point that sends notifications to PagerDuty.
	Pagerduties []ContactPointPagerduty `pulumi:"pagerduties"`
	// A contact point that sends notifications to Pushover.
	Pushovers []ContactPointPushover `pulumi:"pushovers"`
	// A contact point that sends notifications to SensuGo.
	Sensugos []ContactPointSensugo `pulumi:"sensugos"`
	// A contact point that sends notifications to Slack.
	Slacks []ContactPointSlack `pulumi:"slacks"`
	// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
	Sns []ContactPointSn `pulumi:"sns"`
	// A contact point that sends notifications to Microsoft Teams.
	Teams []ContactPointTeam `pulumi:"teams"`
	// A contact point that sends notifications to Telegram.
	Telegrams []ContactPointTelegram `pulumi:"telegrams"`
	// A contact point that sends notifications to Threema.
	Threemas []ContactPointThreema `pulumi:"threemas"`
	// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
	Victorops []ContactPointVictorop `pulumi:"victorops"`
	// A contact point that sends notifications to Cisco Webex.
	Webexes []ContactPointWebex `pulumi:"webexes"`
	// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	Webhooks []ContactPointWebhook `pulumi:"webhooks"`
	// A contact point that sends notifications to WeCom.
	Wecoms []ContactPointWecom `pulumi:"wecoms"`
}

type ContactPointState struct {
	// A contact point that sends notifications to other Alertmanager instances.
	Alertmanagers ContactPointAlertmanagerArrayInput
	// A contact point that sends notifications to DingDing.
	Dingdings ContactPointDingdingArrayInput
	// Allow modifying the contact point from other sources than Terraform or the Grafana API.
	DisableProvenance pulumi.BoolPtrInput
	// A contact point that sends notifications as Discord messages
	Discords ContactPointDiscordArrayInput
	// A contact point that sends notifications to an email address.
	Emails ContactPointEmailArrayInput
	// A contact point that sends notifications to Google Chat.
	Googlechats ContactPointGooglechatArrayInput
	// A contact point that publishes notifications to Apache Kafka topics.
	Kafkas ContactPointKafkaArrayInput
	// A contact point that sends notifications to LINE.me.
	Lines ContactPointLineArrayInput
	// Name of the responder. Must be specified if username and id are empty.
	Name pulumi.StringPtrInput
	// A contact point that sends notifications to Grafana On-Call.
	Oncalls ContactPointOncallArrayInput
	// A contact point that sends notifications to OpsGenie.
	Opsgenies ContactPointOpsgenyArrayInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// A contact point that sends notifications to PagerDuty.
	Pagerduties ContactPointPagerdutyArrayInput
	// A contact point that sends notifications to Pushover.
	Pushovers ContactPointPushoverArrayInput
	// A contact point that sends notifications to SensuGo.
	Sensugos ContactPointSensugoArrayInput
	// A contact point that sends notifications to Slack.
	Slacks ContactPointSlackArrayInput
	// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
	Sns ContactPointSnArrayInput
	// A contact point that sends notifications to Microsoft Teams.
	Teams ContactPointTeamArrayInput
	// A contact point that sends notifications to Telegram.
	Telegrams ContactPointTelegramArrayInput
	// A contact point that sends notifications to Threema.
	Threemas ContactPointThreemaArrayInput
	// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
	Victorops ContactPointVictoropArrayInput
	// A contact point that sends notifications to Cisco Webex.
	Webexes ContactPointWebexArrayInput
	// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	Webhooks ContactPointWebhookArrayInput
	// A contact point that sends notifications to WeCom.
	Wecoms ContactPointWecomArrayInput
}

func (ContactPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactPointState)(nil)).Elem()
}

type contactPointArgs struct {
	// A contact point that sends notifications to other Alertmanager instances.
	Alertmanagers []ContactPointAlertmanager `pulumi:"alertmanagers"`
	// A contact point that sends notifications to DingDing.
	Dingdings []ContactPointDingding `pulumi:"dingdings"`
	// Allow modifying the contact point from other sources than Terraform or the Grafana API.
	DisableProvenance *bool `pulumi:"disableProvenance"`
	// A contact point that sends notifications as Discord messages
	Discords []ContactPointDiscord `pulumi:"discords"`
	// A contact point that sends notifications to an email address.
	Emails []ContactPointEmail `pulumi:"emails"`
	// A contact point that sends notifications to Google Chat.
	Googlechats []ContactPointGooglechat `pulumi:"googlechats"`
	// A contact point that publishes notifications to Apache Kafka topics.
	Kafkas []ContactPointKafka `pulumi:"kafkas"`
	// A contact point that sends notifications to LINE.me.
	Lines []ContactPointLine `pulumi:"lines"`
	// Name of the responder. Must be specified if username and id are empty.
	Name *string `pulumi:"name"`
	// A contact point that sends notifications to Grafana On-Call.
	Oncalls []ContactPointOncall `pulumi:"oncalls"`
	// A contact point that sends notifications to OpsGenie.
	Opsgenies []ContactPointOpsgeny `pulumi:"opsgenies"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// A contact point that sends notifications to PagerDuty.
	Pagerduties []ContactPointPagerduty `pulumi:"pagerduties"`
	// A contact point that sends notifications to Pushover.
	Pushovers []ContactPointPushover `pulumi:"pushovers"`
	// A contact point that sends notifications to SensuGo.
	Sensugos []ContactPointSensugo `pulumi:"sensugos"`
	// A contact point that sends notifications to Slack.
	Slacks []ContactPointSlack `pulumi:"slacks"`
	// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
	Sns []ContactPointSn `pulumi:"sns"`
	// A contact point that sends notifications to Microsoft Teams.
	Teams []ContactPointTeam `pulumi:"teams"`
	// A contact point that sends notifications to Telegram.
	Telegrams []ContactPointTelegram `pulumi:"telegrams"`
	// A contact point that sends notifications to Threema.
	Threemas []ContactPointThreema `pulumi:"threemas"`
	// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
	Victorops []ContactPointVictorop `pulumi:"victorops"`
	// A contact point that sends notifications to Cisco Webex.
	Webexes []ContactPointWebex `pulumi:"webexes"`
	// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	Webhooks []ContactPointWebhook `pulumi:"webhooks"`
	// A contact point that sends notifications to WeCom.
	Wecoms []ContactPointWecom `pulumi:"wecoms"`
}

// The set of arguments for constructing a ContactPoint resource.
type ContactPointArgs struct {
	// A contact point that sends notifications to other Alertmanager instances.
	Alertmanagers ContactPointAlertmanagerArrayInput
	// A contact point that sends notifications to DingDing.
	Dingdings ContactPointDingdingArrayInput
	// Allow modifying the contact point from other sources than Terraform or the Grafana API.
	DisableProvenance pulumi.BoolPtrInput
	// A contact point that sends notifications as Discord messages
	Discords ContactPointDiscordArrayInput
	// A contact point that sends notifications to an email address.
	Emails ContactPointEmailArrayInput
	// A contact point that sends notifications to Google Chat.
	Googlechats ContactPointGooglechatArrayInput
	// A contact point that publishes notifications to Apache Kafka topics.
	Kafkas ContactPointKafkaArrayInput
	// A contact point that sends notifications to LINE.me.
	Lines ContactPointLineArrayInput
	// Name of the responder. Must be specified if username and id are empty.
	Name pulumi.StringPtrInput
	// A contact point that sends notifications to Grafana On-Call.
	Oncalls ContactPointOncallArrayInput
	// A contact point that sends notifications to OpsGenie.
	Opsgenies ContactPointOpsgenyArrayInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// A contact point that sends notifications to PagerDuty.
	Pagerduties ContactPointPagerdutyArrayInput
	// A contact point that sends notifications to Pushover.
	Pushovers ContactPointPushoverArrayInput
	// A contact point that sends notifications to SensuGo.
	Sensugos ContactPointSensugoArrayInput
	// A contact point that sends notifications to Slack.
	Slacks ContactPointSlackArrayInput
	// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
	Sns ContactPointSnArrayInput
	// A contact point that sends notifications to Microsoft Teams.
	Teams ContactPointTeamArrayInput
	// A contact point that sends notifications to Telegram.
	Telegrams ContactPointTelegramArrayInput
	// A contact point that sends notifications to Threema.
	Threemas ContactPointThreemaArrayInput
	// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
	Victorops ContactPointVictoropArrayInput
	// A contact point that sends notifications to Cisco Webex.
	Webexes ContactPointWebexArrayInput
	// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	Webhooks ContactPointWebhookArrayInput
	// A contact point that sends notifications to WeCom.
	Wecoms ContactPointWecomArrayInput
}

func (ContactPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactPointArgs)(nil)).Elem()
}

type ContactPointInput interface {
	pulumi.Input

	ToContactPointOutput() ContactPointOutput
	ToContactPointOutputWithContext(ctx context.Context) ContactPointOutput
}

func (*ContactPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactPoint)(nil)).Elem()
}

func (i *ContactPoint) ToContactPointOutput() ContactPointOutput {
	return i.ToContactPointOutputWithContext(context.Background())
}

func (i *ContactPoint) ToContactPointOutputWithContext(ctx context.Context) ContactPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactPointOutput)
}

func (i *ContactPoint) ToOutput(ctx context.Context) pulumix.Output[*ContactPoint] {
	return pulumix.Output[*ContactPoint]{
		OutputState: i.ToContactPointOutputWithContext(ctx).OutputState,
	}
}

// ContactPointArrayInput is an input type that accepts ContactPointArray and ContactPointArrayOutput values.
// You can construct a concrete instance of `ContactPointArrayInput` via:
//
//	ContactPointArray{ ContactPointArgs{...} }
type ContactPointArrayInput interface {
	pulumi.Input

	ToContactPointArrayOutput() ContactPointArrayOutput
	ToContactPointArrayOutputWithContext(context.Context) ContactPointArrayOutput
}

type ContactPointArray []ContactPointInput

func (ContactPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactPoint)(nil)).Elem()
}

func (i ContactPointArray) ToContactPointArrayOutput() ContactPointArrayOutput {
	return i.ToContactPointArrayOutputWithContext(context.Background())
}

func (i ContactPointArray) ToContactPointArrayOutputWithContext(ctx context.Context) ContactPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactPointArrayOutput)
}

func (i ContactPointArray) ToOutput(ctx context.Context) pulumix.Output[[]*ContactPoint] {
	return pulumix.Output[[]*ContactPoint]{
		OutputState: i.ToContactPointArrayOutputWithContext(ctx).OutputState,
	}
}

// ContactPointMapInput is an input type that accepts ContactPointMap and ContactPointMapOutput values.
// You can construct a concrete instance of `ContactPointMapInput` via:
//
//	ContactPointMap{ "key": ContactPointArgs{...} }
type ContactPointMapInput interface {
	pulumi.Input

	ToContactPointMapOutput() ContactPointMapOutput
	ToContactPointMapOutputWithContext(context.Context) ContactPointMapOutput
}

type ContactPointMap map[string]ContactPointInput

func (ContactPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactPoint)(nil)).Elem()
}

func (i ContactPointMap) ToContactPointMapOutput() ContactPointMapOutput {
	return i.ToContactPointMapOutputWithContext(context.Background())
}

func (i ContactPointMap) ToContactPointMapOutputWithContext(ctx context.Context) ContactPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactPointMapOutput)
}

func (i ContactPointMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContactPoint] {
	return pulumix.Output[map[string]*ContactPoint]{
		OutputState: i.ToContactPointMapOutputWithContext(ctx).OutputState,
	}
}

type ContactPointOutput struct{ *pulumi.OutputState }

func (ContactPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactPoint)(nil)).Elem()
}

func (o ContactPointOutput) ToContactPointOutput() ContactPointOutput {
	return o
}

func (o ContactPointOutput) ToContactPointOutputWithContext(ctx context.Context) ContactPointOutput {
	return o
}

func (o ContactPointOutput) ToOutput(ctx context.Context) pulumix.Output[*ContactPoint] {
	return pulumix.Output[*ContactPoint]{
		OutputState: o.OutputState,
	}
}

// A contact point that sends notifications to other Alertmanager instances.
func (o ContactPointOutput) Alertmanagers() ContactPointAlertmanagerArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointAlertmanagerArrayOutput { return v.Alertmanagers }).(ContactPointAlertmanagerArrayOutput)
}

// A contact point that sends notifications to DingDing.
func (o ContactPointOutput) Dingdings() ContactPointDingdingArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointDingdingArrayOutput { return v.Dingdings }).(ContactPointDingdingArrayOutput)
}

// Allow modifying the contact point from other sources than Terraform or the Grafana API.
func (o ContactPointOutput) DisableProvenance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContactPoint) pulumi.BoolPtrOutput { return v.DisableProvenance }).(pulumi.BoolPtrOutput)
}

// A contact point that sends notifications as Discord messages
func (o ContactPointOutput) Discords() ContactPointDiscordArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointDiscordArrayOutput { return v.Discords }).(ContactPointDiscordArrayOutput)
}

// A contact point that sends notifications to an email address.
func (o ContactPointOutput) Emails() ContactPointEmailArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointEmailArrayOutput { return v.Emails }).(ContactPointEmailArrayOutput)
}

// A contact point that sends notifications to Google Chat.
func (o ContactPointOutput) Googlechats() ContactPointGooglechatArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointGooglechatArrayOutput { return v.Googlechats }).(ContactPointGooglechatArrayOutput)
}

// A contact point that publishes notifications to Apache Kafka topics.
func (o ContactPointOutput) Kafkas() ContactPointKafkaArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointKafkaArrayOutput { return v.Kafkas }).(ContactPointKafkaArrayOutput)
}

// A contact point that sends notifications to LINE.me.
func (o ContactPointOutput) Lines() ContactPointLineArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointLineArrayOutput { return v.Lines }).(ContactPointLineArrayOutput)
}

// Name of the responder. Must be specified if username and id are empty.
func (o ContactPointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactPoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A contact point that sends notifications to Grafana On-Call.
func (o ContactPointOutput) Oncalls() ContactPointOncallArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointOncallArrayOutput { return v.Oncalls }).(ContactPointOncallArrayOutput)
}

// A contact point that sends notifications to OpsGenie.
func (o ContactPointOutput) Opsgenies() ContactPointOpsgenyArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointOpsgenyArrayOutput { return v.Opsgenies }).(ContactPointOpsgenyArrayOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o ContactPointOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactPoint) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// A contact point that sends notifications to PagerDuty.
func (o ContactPointOutput) Pagerduties() ContactPointPagerdutyArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointPagerdutyArrayOutput { return v.Pagerduties }).(ContactPointPagerdutyArrayOutput)
}

// A contact point that sends notifications to Pushover.
func (o ContactPointOutput) Pushovers() ContactPointPushoverArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointPushoverArrayOutput { return v.Pushovers }).(ContactPointPushoverArrayOutput)
}

// A contact point that sends notifications to SensuGo.
func (o ContactPointOutput) Sensugos() ContactPointSensugoArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointSensugoArrayOutput { return v.Sensugos }).(ContactPointSensugoArrayOutput)
}

// A contact point that sends notifications to Slack.
func (o ContactPointOutput) Slacks() ContactPointSlackArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointSlackArrayOutput { return v.Slacks }).(ContactPointSlackArrayOutput)
}

// A contact point that sends notifications to Amazon SNS. Requires Amazon Managed Grafana.
func (o ContactPointOutput) Sns() ContactPointSnArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointSnArrayOutput { return v.Sns }).(ContactPointSnArrayOutput)
}

// A contact point that sends notifications to Microsoft Teams.
func (o ContactPointOutput) Teams() ContactPointTeamArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointTeamArrayOutput { return v.Teams }).(ContactPointTeamArrayOutput)
}

// A contact point that sends notifications to Telegram.
func (o ContactPointOutput) Telegrams() ContactPointTelegramArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointTelegramArrayOutput { return v.Telegrams }).(ContactPointTelegramArrayOutput)
}

// A contact point that sends notifications to Threema.
func (o ContactPointOutput) Threemas() ContactPointThreemaArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointThreemaArrayOutput { return v.Threemas }).(ContactPointThreemaArrayOutput)
}

// A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
func (o ContactPointOutput) Victorops() ContactPointVictoropArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointVictoropArrayOutput { return v.Victorops }).(ContactPointVictoropArrayOutput)
}

// A contact point that sends notifications to Cisco Webex.
func (o ContactPointOutput) Webexes() ContactPointWebexArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointWebexArrayOutput { return v.Webexes }).(ContactPointWebexArrayOutput)
}

// A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
func (o ContactPointOutput) Webhooks() ContactPointWebhookArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointWebhookArrayOutput { return v.Webhooks }).(ContactPointWebhookArrayOutput)
}

// A contact point that sends notifications to WeCom.
func (o ContactPointOutput) Wecoms() ContactPointWecomArrayOutput {
	return o.ApplyT(func(v *ContactPoint) ContactPointWecomArrayOutput { return v.Wecoms }).(ContactPointWecomArrayOutput)
}

type ContactPointArrayOutput struct{ *pulumi.OutputState }

func (ContactPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactPoint)(nil)).Elem()
}

func (o ContactPointArrayOutput) ToContactPointArrayOutput() ContactPointArrayOutput {
	return o
}

func (o ContactPointArrayOutput) ToContactPointArrayOutputWithContext(ctx context.Context) ContactPointArrayOutput {
	return o
}

func (o ContactPointArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ContactPoint] {
	return pulumix.Output[[]*ContactPoint]{
		OutputState: o.OutputState,
	}
}

func (o ContactPointArrayOutput) Index(i pulumi.IntInput) ContactPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContactPoint {
		return vs[0].([]*ContactPoint)[vs[1].(int)]
	}).(ContactPointOutput)
}

type ContactPointMapOutput struct{ *pulumi.OutputState }

func (ContactPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactPoint)(nil)).Elem()
}

func (o ContactPointMapOutput) ToContactPointMapOutput() ContactPointMapOutput {
	return o
}

func (o ContactPointMapOutput) ToContactPointMapOutputWithContext(ctx context.Context) ContactPointMapOutput {
	return o
}

func (o ContactPointMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContactPoint] {
	return pulumix.Output[map[string]*ContactPoint]{
		OutputState: o.OutputState,
	}
}

func (o ContactPointMapOutput) MapIndex(k pulumi.StringInput) ContactPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContactPoint {
		return vs[0].(map[string]*ContactPoint)[vs[1].(string)]
	}).(ContactPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactPointInput)(nil)).Elem(), &ContactPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactPointArrayInput)(nil)).Elem(), ContactPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactPointMapInput)(nil)).Elem(), ContactPointMap{})
	pulumi.RegisterOutputType(ContactPointOutput{})
	pulumi.RegisterOutputType(ContactPointArrayOutput{})
	pulumi.RegisterOutputType(ContactPointMapOutput{})
}
