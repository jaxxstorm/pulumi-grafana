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

// Use the `teamSync` attribute of the `Team` resource instead.
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
//			_, err := grafana.NewTeamExternalGroup(ctx, "test-team-group", &grafana.TeamExternalGroupArgs{
//				Groups: pulumi.StringArray{
//					pulumi.String("test-group-1"),
//					pulumi.String("test-group-2"),
//				},
//				TeamId: pulumi.String("1"),
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
//	$ pulumi import grafana:index/teamExternalGroup:TeamExternalGroup main {{team_id}}
//
// ```
type TeamExternalGroup struct {
	pulumi.CustomResourceState

	// The team external groups list
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The Team ID
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeamExternalGroup registers a new resource with the given unique name, arguments, and options.
func NewTeamExternalGroup(ctx *pulumi.Context,
	name string, args *TeamExternalGroupArgs, opts ...pulumi.ResourceOption) (*TeamExternalGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TeamExternalGroup
	err := ctx.RegisterResource("grafana:index/teamExternalGroup:TeamExternalGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamExternalGroup gets an existing TeamExternalGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamExternalGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamExternalGroupState, opts ...pulumi.ResourceOption) (*TeamExternalGroup, error) {
	var resource TeamExternalGroup
	err := ctx.ReadResource("grafana:index/teamExternalGroup:TeamExternalGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamExternalGroup resources.
type teamExternalGroupState struct {
	// The team external groups list
	Groups []string `pulumi:"groups"`
	// The Team ID
	TeamId *string `pulumi:"teamId"`
}

type TeamExternalGroupState struct {
	// The team external groups list
	Groups pulumi.StringArrayInput
	// The Team ID
	TeamId pulumi.StringPtrInput
}

func (TeamExternalGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamExternalGroupState)(nil)).Elem()
}

type teamExternalGroupArgs struct {
	// The team external groups list
	Groups []string `pulumi:"groups"`
	// The Team ID
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamExternalGroup resource.
type TeamExternalGroupArgs struct {
	// The team external groups list
	Groups pulumi.StringArrayInput
	// The Team ID
	TeamId pulumi.StringInput
}

func (TeamExternalGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamExternalGroupArgs)(nil)).Elem()
}

type TeamExternalGroupInput interface {
	pulumi.Input

	ToTeamExternalGroupOutput() TeamExternalGroupOutput
	ToTeamExternalGroupOutputWithContext(ctx context.Context) TeamExternalGroupOutput
}

func (*TeamExternalGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamExternalGroup)(nil)).Elem()
}

func (i *TeamExternalGroup) ToTeamExternalGroupOutput() TeamExternalGroupOutput {
	return i.ToTeamExternalGroupOutputWithContext(context.Background())
}

func (i *TeamExternalGroup) ToTeamExternalGroupOutputWithContext(ctx context.Context) TeamExternalGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamExternalGroupOutput)
}

func (i *TeamExternalGroup) ToOutput(ctx context.Context) pulumix.Output[*TeamExternalGroup] {
	return pulumix.Output[*TeamExternalGroup]{
		OutputState: i.ToTeamExternalGroupOutputWithContext(ctx).OutputState,
	}
}

// TeamExternalGroupArrayInput is an input type that accepts TeamExternalGroupArray and TeamExternalGroupArrayOutput values.
// You can construct a concrete instance of `TeamExternalGroupArrayInput` via:
//
//	TeamExternalGroupArray{ TeamExternalGroupArgs{...} }
type TeamExternalGroupArrayInput interface {
	pulumi.Input

	ToTeamExternalGroupArrayOutput() TeamExternalGroupArrayOutput
	ToTeamExternalGroupArrayOutputWithContext(context.Context) TeamExternalGroupArrayOutput
}

type TeamExternalGroupArray []TeamExternalGroupInput

func (TeamExternalGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamExternalGroup)(nil)).Elem()
}

func (i TeamExternalGroupArray) ToTeamExternalGroupArrayOutput() TeamExternalGroupArrayOutput {
	return i.ToTeamExternalGroupArrayOutputWithContext(context.Background())
}

func (i TeamExternalGroupArray) ToTeamExternalGroupArrayOutputWithContext(ctx context.Context) TeamExternalGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamExternalGroupArrayOutput)
}

func (i TeamExternalGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*TeamExternalGroup] {
	return pulumix.Output[[]*TeamExternalGroup]{
		OutputState: i.ToTeamExternalGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// TeamExternalGroupMapInput is an input type that accepts TeamExternalGroupMap and TeamExternalGroupMapOutput values.
// You can construct a concrete instance of `TeamExternalGroupMapInput` via:
//
//	TeamExternalGroupMap{ "key": TeamExternalGroupArgs{...} }
type TeamExternalGroupMapInput interface {
	pulumi.Input

	ToTeamExternalGroupMapOutput() TeamExternalGroupMapOutput
	ToTeamExternalGroupMapOutputWithContext(context.Context) TeamExternalGroupMapOutput
}

type TeamExternalGroupMap map[string]TeamExternalGroupInput

func (TeamExternalGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamExternalGroup)(nil)).Elem()
}

func (i TeamExternalGroupMap) ToTeamExternalGroupMapOutput() TeamExternalGroupMapOutput {
	return i.ToTeamExternalGroupMapOutputWithContext(context.Background())
}

func (i TeamExternalGroupMap) ToTeamExternalGroupMapOutputWithContext(ctx context.Context) TeamExternalGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamExternalGroupMapOutput)
}

func (i TeamExternalGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TeamExternalGroup] {
	return pulumix.Output[map[string]*TeamExternalGroup]{
		OutputState: i.ToTeamExternalGroupMapOutputWithContext(ctx).OutputState,
	}
}

type TeamExternalGroupOutput struct{ *pulumi.OutputState }

func (TeamExternalGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamExternalGroup)(nil)).Elem()
}

func (o TeamExternalGroupOutput) ToTeamExternalGroupOutput() TeamExternalGroupOutput {
	return o
}

func (o TeamExternalGroupOutput) ToTeamExternalGroupOutputWithContext(ctx context.Context) TeamExternalGroupOutput {
	return o
}

func (o TeamExternalGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*TeamExternalGroup] {
	return pulumix.Output[*TeamExternalGroup]{
		OutputState: o.OutputState,
	}
}

// The team external groups list
func (o TeamExternalGroupOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TeamExternalGroup) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// The Team ID
func (o TeamExternalGroupOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamExternalGroup) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamExternalGroupArrayOutput struct{ *pulumi.OutputState }

func (TeamExternalGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamExternalGroup)(nil)).Elem()
}

func (o TeamExternalGroupArrayOutput) ToTeamExternalGroupArrayOutput() TeamExternalGroupArrayOutput {
	return o
}

func (o TeamExternalGroupArrayOutput) ToTeamExternalGroupArrayOutputWithContext(ctx context.Context) TeamExternalGroupArrayOutput {
	return o
}

func (o TeamExternalGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TeamExternalGroup] {
	return pulumix.Output[[]*TeamExternalGroup]{
		OutputState: o.OutputState,
	}
}

func (o TeamExternalGroupArrayOutput) Index(i pulumi.IntInput) TeamExternalGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamExternalGroup {
		return vs[0].([]*TeamExternalGroup)[vs[1].(int)]
	}).(TeamExternalGroupOutput)
}

type TeamExternalGroupMapOutput struct{ *pulumi.OutputState }

func (TeamExternalGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamExternalGroup)(nil)).Elem()
}

func (o TeamExternalGroupMapOutput) ToTeamExternalGroupMapOutput() TeamExternalGroupMapOutput {
	return o
}

func (o TeamExternalGroupMapOutput) ToTeamExternalGroupMapOutputWithContext(ctx context.Context) TeamExternalGroupMapOutput {
	return o
}

func (o TeamExternalGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TeamExternalGroup] {
	return pulumix.Output[map[string]*TeamExternalGroup]{
		OutputState: o.OutputState,
	}
}

func (o TeamExternalGroupMapOutput) MapIndex(k pulumi.StringInput) TeamExternalGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamExternalGroup {
		return vs[0].(map[string]*TeamExternalGroup)[vs[1].(string)]
	}).(TeamExternalGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamExternalGroupInput)(nil)).Elem(), &TeamExternalGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamExternalGroupArrayInput)(nil)).Elem(), TeamExternalGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamExternalGroupMapInput)(nil)).Elem(), TeamExternalGroupMap{})
	pulumi.RegisterOutputType(TeamExternalGroupOutput{})
	pulumi.RegisterOutputType(TeamExternalGroupArrayOutput{})
	pulumi.RegisterOutputType(TeamExternalGroupMapOutput{})
}
