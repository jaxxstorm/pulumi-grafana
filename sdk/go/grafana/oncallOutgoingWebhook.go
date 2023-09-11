// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.NewOncallOutgoingWebhook(ctx, "test-acc-outgoingWebhook", &grafana.OncallOutgoingWebhookArgs{
//				Url: pulumi.String("https://example.com/"),
//			}, pulumi.Provider(grafana.Oncall))
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
//	$ pulumi import grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook outgoing_webhook_name {{outgoing_webhook_id}}
//
// ```
type OncallOutgoingWebhook struct {
	pulumi.CustomResourceState

	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrOutput `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// Forwards whole payload of the alert to the webhook's url as POST data.
	ForwardWholePayload pulumi.BoolPtrOutput `pulumi:"forwardWholePayload"`
	// The name of the outgoing webhook.
	Name pulumi.StringOutput `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The webhook URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// The auth data of the webhook. Used for Basic authentication.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewOncallOutgoingWebhook registers a new resource with the given unique name, arguments, and options.
func NewOncallOutgoingWebhook(ctx *pulumi.Context,
	name string, args *OncallOutgoingWebhookArgs, opts ...pulumi.ResourceOption) (*OncallOutgoingWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.AuthorizationHeader != nil {
		args.AuthorizationHeader = pulumi.ToSecret(args.AuthorizationHeader).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authorizationHeader",
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OncallOutgoingWebhook
	err := ctx.RegisterResource("grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOncallOutgoingWebhook gets an existing OncallOutgoingWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOncallOutgoingWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OncallOutgoingWebhookState, opts ...pulumi.ResourceOption) (*OncallOutgoingWebhook, error) {
	var resource OncallOutgoingWebhook
	err := ctx.ReadResource("grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OncallOutgoingWebhook resources.
type oncallOutgoingWebhookState struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader *string `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data *string `pulumi:"data"`
	// Forwards whole payload of the alert to the webhook's url as POST data.
	ForwardWholePayload *bool `pulumi:"forwardWholePayload"`
	// The name of the outgoing webhook.
	Name *string `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password *string `pulumi:"password"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
	TeamId *string `pulumi:"teamId"`
	// The webhook URL.
	Url *string `pulumi:"url"`
	// The auth data of the webhook. Used for Basic authentication.
	User *string `pulumi:"user"`
}

type OncallOutgoingWebhookState struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrInput
	// The data of the webhook.
	Data pulumi.StringPtrInput
	// Forwards whole payload of the alert to the webhook's url as POST data.
	ForwardWholePayload pulumi.BoolPtrInput
	// The name of the outgoing webhook.
	Name pulumi.StringPtrInput
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrInput
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
	TeamId pulumi.StringPtrInput
	// The webhook URL.
	Url pulumi.StringPtrInput
	// The auth data of the webhook. Used for Basic authentication.
	User pulumi.StringPtrInput
}

func (OncallOutgoingWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallOutgoingWebhookState)(nil)).Elem()
}

type oncallOutgoingWebhookArgs struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader *string `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data *string `pulumi:"data"`
	// Forwards whole payload of the alert to the webhook's url as POST data.
	ForwardWholePayload *bool `pulumi:"forwardWholePayload"`
	// The name of the outgoing webhook.
	Name *string `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password *string `pulumi:"password"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
	TeamId *string `pulumi:"teamId"`
	// The webhook URL.
	Url string `pulumi:"url"`
	// The auth data of the webhook. Used for Basic authentication.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a OncallOutgoingWebhook resource.
type OncallOutgoingWebhookArgs struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrInput
	// The data of the webhook.
	Data pulumi.StringPtrInput
	// Forwards whole payload of the alert to the webhook's url as POST data.
	ForwardWholePayload pulumi.BoolPtrInput
	// The name of the outgoing webhook.
	Name pulumi.StringPtrInput
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrInput
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
	TeamId pulumi.StringPtrInput
	// The webhook URL.
	Url pulumi.StringInput
	// The auth data of the webhook. Used for Basic authentication.
	User pulumi.StringPtrInput
}

func (OncallOutgoingWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallOutgoingWebhookArgs)(nil)).Elem()
}

type OncallOutgoingWebhookInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput
	ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput
}

func (*OncallOutgoingWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallOutgoingWebhook)(nil)).Elem()
}

func (i *OncallOutgoingWebhook) ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput {
	return i.ToOncallOutgoingWebhookOutputWithContext(context.Background())
}

func (i *OncallOutgoingWebhook) ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookOutput)
}

func (i *OncallOutgoingWebhook) ToOutput(ctx context.Context) pulumix.Output[*OncallOutgoingWebhook] {
	return pulumix.Output[*OncallOutgoingWebhook]{
		OutputState: i.ToOncallOutgoingWebhookOutputWithContext(ctx).OutputState,
	}
}

// OncallOutgoingWebhookArrayInput is an input type that accepts OncallOutgoingWebhookArray and OncallOutgoingWebhookArrayOutput values.
// You can construct a concrete instance of `OncallOutgoingWebhookArrayInput` via:
//
//	OncallOutgoingWebhookArray{ OncallOutgoingWebhookArgs{...} }
type OncallOutgoingWebhookArrayInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput
	ToOncallOutgoingWebhookArrayOutputWithContext(context.Context) OncallOutgoingWebhookArrayOutput
}

type OncallOutgoingWebhookArray []OncallOutgoingWebhookInput

func (OncallOutgoingWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallOutgoingWebhook)(nil)).Elem()
}

func (i OncallOutgoingWebhookArray) ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput {
	return i.ToOncallOutgoingWebhookArrayOutputWithContext(context.Background())
}

func (i OncallOutgoingWebhookArray) ToOncallOutgoingWebhookArrayOutputWithContext(ctx context.Context) OncallOutgoingWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookArrayOutput)
}

func (i OncallOutgoingWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*OncallOutgoingWebhook] {
	return pulumix.Output[[]*OncallOutgoingWebhook]{
		OutputState: i.ToOncallOutgoingWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// OncallOutgoingWebhookMapInput is an input type that accepts OncallOutgoingWebhookMap and OncallOutgoingWebhookMapOutput values.
// You can construct a concrete instance of `OncallOutgoingWebhookMapInput` via:
//
//	OncallOutgoingWebhookMap{ "key": OncallOutgoingWebhookArgs{...} }
type OncallOutgoingWebhookMapInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput
	ToOncallOutgoingWebhookMapOutputWithContext(context.Context) OncallOutgoingWebhookMapOutput
}

type OncallOutgoingWebhookMap map[string]OncallOutgoingWebhookInput

func (OncallOutgoingWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallOutgoingWebhook)(nil)).Elem()
}

func (i OncallOutgoingWebhookMap) ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput {
	return i.ToOncallOutgoingWebhookMapOutputWithContext(context.Background())
}

func (i OncallOutgoingWebhookMap) ToOncallOutgoingWebhookMapOutputWithContext(ctx context.Context) OncallOutgoingWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookMapOutput)
}

func (i OncallOutgoingWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OncallOutgoingWebhook] {
	return pulumix.Output[map[string]*OncallOutgoingWebhook]{
		OutputState: i.ToOncallOutgoingWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type OncallOutgoingWebhookOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookOutput) ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput {
	return o
}

func (o OncallOutgoingWebhookOutput) ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput {
	return o
}

func (o OncallOutgoingWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*OncallOutgoingWebhook] {
	return pulumix.Output[*OncallOutgoingWebhook]{
		OutputState: o.OutputState,
	}
}

// The auth data of the webhook. Used in Authorization header instead of user/password auth.
func (o OncallOutgoingWebhookOutput) AuthorizationHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.AuthorizationHeader }).(pulumi.StringPtrOutput)
}

// The data of the webhook.
func (o OncallOutgoingWebhookOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.Data }).(pulumi.StringPtrOutput)
}

// Forwards whole payload of the alert to the webhook's url as POST data.
func (o OncallOutgoingWebhookOutput) ForwardWholePayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.BoolPtrOutput { return v.ForwardWholePayload }).(pulumi.BoolPtrOutput)
}

// The name of the outgoing webhook.
func (o OncallOutgoingWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The auth data of the webhook. Used for Basic authentication
func (o OncallOutgoingWebhookOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `getOncallTeam` datasource.
func (o OncallOutgoingWebhookOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

// The webhook URL.
func (o OncallOutgoingWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The auth data of the webhook. Used for Basic authentication.
func (o OncallOutgoingWebhookOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type OncallOutgoingWebhookArrayOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookArrayOutput) ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput {
	return o
}

func (o OncallOutgoingWebhookArrayOutput) ToOncallOutgoingWebhookArrayOutputWithContext(ctx context.Context) OncallOutgoingWebhookArrayOutput {
	return o
}

func (o OncallOutgoingWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OncallOutgoingWebhook] {
	return pulumix.Output[[]*OncallOutgoingWebhook]{
		OutputState: o.OutputState,
	}
}

func (o OncallOutgoingWebhookArrayOutput) Index(i pulumi.IntInput) OncallOutgoingWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OncallOutgoingWebhook {
		return vs[0].([]*OncallOutgoingWebhook)[vs[1].(int)]
	}).(OncallOutgoingWebhookOutput)
}

type OncallOutgoingWebhookMapOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookMapOutput) ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput {
	return o
}

func (o OncallOutgoingWebhookMapOutput) ToOncallOutgoingWebhookMapOutputWithContext(ctx context.Context) OncallOutgoingWebhookMapOutput {
	return o
}

func (o OncallOutgoingWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OncallOutgoingWebhook] {
	return pulumix.Output[map[string]*OncallOutgoingWebhook]{
		OutputState: o.OutputState,
	}
}

func (o OncallOutgoingWebhookMapOutput) MapIndex(k pulumi.StringInput) OncallOutgoingWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OncallOutgoingWebhook {
		return vs[0].(map[string]*OncallOutgoingWebhook)[vs[1].(string)]
	}).(OncallOutgoingWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookInput)(nil)).Elem(), &OncallOutgoingWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookArrayInput)(nil)).Elem(), OncallOutgoingWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookMapInput)(nil)).Elem(), OncallOutgoingWebhookMap{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookOutput{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookArrayOutput{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookMapOutput{})
}
