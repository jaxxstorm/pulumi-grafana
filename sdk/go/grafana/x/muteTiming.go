// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages Grafana Alerting mute timings.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/manage-notifications/mute-timings/)
// * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#mute-timings)
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
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.NewMuteTiming(ctx, "myMuteTiming", &grafana.MuteTimingArgs{
//				Intervals: grafana.MuteTimingIntervalArray{
//					&grafana.MuteTimingIntervalArgs{
//						DaysOfMonths: pulumi.StringArray{
//							pulumi.String("1:7"),
//							pulumi.String("-1"),
//						},
//						Months: pulumi.StringArray{
//							pulumi.String("1:3"),
//							pulumi.String("december"),
//						},
//						Times: grafana.MuteTimingIntervalTimeArray{
//							&grafana.MuteTimingIntervalTimeArgs{
//								End:   pulumi.String("14:17"),
//								Start: pulumi.String("04:56"),
//							},
//						},
//						Weekdays: pulumi.StringArray{
//							pulumi.String("monday"),
//							pulumi.String("tuesday:thursday"),
//						},
//						Years: pulumi.StringArray{
//							pulumi.String("2030"),
//							pulumi.String("2025:2026"),
//						},
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
//	$ pulumi import grafana:index/muteTiming:MuteTiming mute_timing_name {{mute_timing_name}}
//
// ```
type MuteTiming struct {
	pulumi.CustomResourceState

	// The time intervals at which to mute notifications.
	Intervals pulumix.GArrayOutput[MuteTimingInterval, MuteTimingIntervalOutput] `pulumi:"intervals"`
	// The name of the mute timing.
	Name pulumix.Output[string] `pulumi:"name"`
}

// NewMuteTiming registers a new resource with the given unique name, arguments, and options.
func NewMuteTiming(ctx *pulumi.Context,
	name string, args *MuteTimingArgs, opts ...pulumi.ResourceOption) (*MuteTiming, error) {
	if args == nil {
		args = &MuteTimingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MuteTiming
	err := ctx.RegisterResource("grafana:index/muteTiming:MuteTiming", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMuteTiming gets an existing MuteTiming resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMuteTiming(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MuteTimingState, opts ...pulumi.ResourceOption) (*MuteTiming, error) {
	var resource MuteTiming
	err := ctx.ReadResource("grafana:index/muteTiming:MuteTiming", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MuteTiming resources.
type muteTimingState struct {
	// The time intervals at which to mute notifications.
	Intervals []MuteTimingInterval `pulumi:"intervals"`
	// The name of the mute timing.
	Name *string `pulumi:"name"`
}

type MuteTimingState struct {
	// The time intervals at which to mute notifications.
	Intervals pulumix.Input[[]*MuteTimingIntervalArgs]
	// The name of the mute timing.
	Name pulumix.Input[*string]
}

func (MuteTimingState) ElementType() reflect.Type {
	return reflect.TypeOf((*muteTimingState)(nil)).Elem()
}

type muteTimingArgs struct {
	// The time intervals at which to mute notifications.
	Intervals []MuteTimingInterval `pulumi:"intervals"`
	// The name of the mute timing.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a MuteTiming resource.
type MuteTimingArgs struct {
	// The time intervals at which to mute notifications.
	Intervals pulumix.Input[[]*MuteTimingIntervalArgs]
	// The name of the mute timing.
	Name pulumix.Input[*string]
}

func (MuteTimingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*muteTimingArgs)(nil)).Elem()
}

type MuteTimingOutput struct{ *pulumi.OutputState }

func (MuteTimingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MuteTiming)(nil)).Elem()
}

func (o MuteTimingOutput) ToMuteTimingOutput() MuteTimingOutput {
	return o
}

func (o MuteTimingOutput) ToMuteTimingOutputWithContext(ctx context.Context) MuteTimingOutput {
	return o
}

func (o MuteTimingOutput) ToOutput(ctx context.Context) pulumix.Output[MuteTiming] {
	return pulumix.Output[MuteTiming]{
		OutputState: o.OutputState,
	}
}

// The time intervals at which to mute notifications.
func (o MuteTimingOutput) Intervals() pulumix.GArrayOutput[MuteTimingInterval, MuteTimingIntervalOutput] {
	value := pulumix.Apply[MuteTiming](o, func(v MuteTiming) pulumix.GArrayOutput[MuteTimingInterval, MuteTimingIntervalOutput] {
		return v.Intervals
	})
	unwrapped := pulumix.Flatten[[]MuteTimingInterval, pulumix.GArrayOutput[MuteTimingInterval, MuteTimingIntervalOutput]](value)
	return pulumix.GArrayOutput[MuteTimingInterval, MuteTimingIntervalOutput]{OutputState: unwrapped.OutputState}
}

// The name of the mute timing.
func (o MuteTimingOutput) Name() pulumix.Output[string] {
	value := pulumix.Apply[MuteTiming](o, func(v MuteTiming) pulumix.Output[string] { return v.Name })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(MuteTimingOutput{})
}
