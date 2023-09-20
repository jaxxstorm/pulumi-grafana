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

// Sets the global notification policy for Grafana.
//
// !> This resource manages the entire notification policy tree, and will overwrite any existing policies.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/manage-notifications/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/)
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
//			aContactPoint, err := grafana.NewContactPoint(ctx, "aContactPoint", &grafana.ContactPointArgs{
//				Emails: grafana.ContactPointEmailArray{
//					&grafana.ContactPointEmailArgs{
//						Addresses: pulumi.StringArray{
//							pulumi.String("one@company.org"),
//							pulumi.String("two@company.org"),
//						},
//						Message: pulumi.String("{{ len .Alerts.Firing }} firing."),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			aMuteTiming, err := grafana.NewMuteTiming(ctx, "aMuteTiming", &grafana.MuteTimingArgs{
//				Intervals: grafana.MuteTimingIntervalArray{
//					&grafana.MuteTimingIntervalArgs{
//						Weekdays: pulumi.StringArray{
//							pulumi.String("monday"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewNotificationPolicy(ctx, "myNotificationPolicy", &grafana.NotificationPolicyArgs{
//				GroupBies: pulumi.StringArray{
//					pulumi.String("..."),
//				},
//				ContactPoint:   aContactPoint.Name,
//				GroupWait:      pulumi.String("45s"),
//				GroupInterval:  pulumi.String("6m"),
//				RepeatInterval: pulumi.String("3h"),
//				Policies: grafana.NotificationPolicyPolicyArray{
//					&grafana.NotificationPolicyPolicyArgs{
//						Matchers: grafana.NotificationPolicyPolicyMatcherArray{
//							&grafana.NotificationPolicyPolicyMatcherArgs{
//								Label: pulumi.String("mylabel"),
//								Match: pulumi.String("="),
//								Value: pulumi.String("myvalue"),
//							},
//						},
//						ContactPoint: aContactPoint.Name,
//						GroupBies: pulumi.StringArray{
//							pulumi.String("alertname"),
//						},
//						Continue: pulumi.Bool(true),
//						MuteTimings: pulumi.StringArray{
//							aMuteTiming.Name,
//						},
//						GroupWait:      pulumi.String("45s"),
//						GroupInterval:  pulumi.String("6m"),
//						RepeatInterval: pulumi.String("3h"),
//						Policies: grafana.NotificationPolicyPolicyPolicyArray{
//							&grafana.NotificationPolicyPolicyPolicyArgs{
//								Matchers: grafana.NotificationPolicyPolicyPolicyMatcherArray{
//									&grafana.NotificationPolicyPolicyPolicyMatcherArgs{
//										Label: pulumi.String("sublabel"),
//										Match: pulumi.String("="),
//										Value: pulumi.String("subvalue"),
//									},
//								},
//								ContactPoint: aContactPoint.Name,
//								GroupBies: pulumi.StringArray{
//									pulumi.String("..."),
//								},
//							},
//						},
//					},
//					&grafana.NotificationPolicyPolicyArgs{
//						Matchers: grafana.NotificationPolicyPolicyMatcherArray{
//							&grafana.NotificationPolicyPolicyMatcherArgs{
//								Label: pulumi.String("anotherlabel"),
//								Match: pulumi.String("=~"),
//								Value: pulumi.String("another value.*"),
//							},
//						},
//						ContactPoint: aContactPoint.Name,
//						GroupBies: pulumi.StringArray{
//							pulumi.String("..."),
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
// The policy is a singleton, so the ID is a constant "policy" value.
//
// ```sh
//
//	$ pulumi import grafana:index/notificationPolicy:NotificationPolicy notification_policy_name "policy"
//
// ```
type NotificationPolicy struct {
	pulumi.CustomResourceState

	// The default contact point to route all unmatched notifications to.
	ContactPoint pulumix.Output[string] `pulumi:"contactPoint"`
	// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
	GroupBies pulumix.ArrayOutput[string] `pulumi:"groupBies"`
	// Minimum time interval between two notifications for the same group. Default is 5 minutes.
	GroupInterval pulumix.Output[*string] `pulumi:"groupInterval"`
	// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
	GroupWait pulumix.Output[*string] `pulumi:"groupWait"`
	// Routing rules for specific label sets.
	Policies pulumix.GArrayOutput[NotificationPolicyPolicy, NotificationPolicyPolicyOutput] `pulumi:"policies"`
	// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
	RepeatInterval pulumix.Output[*string] `pulumi:"repeatInterval"`
}

// NewNotificationPolicy registers a new resource with the given unique name, arguments, and options.
func NewNotificationPolicy(ctx *pulumi.Context,
	name string, args *NotificationPolicyArgs, opts ...pulumi.ResourceOption) (*NotificationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactPoint == nil {
		return nil, errors.New("invalid value for required argument 'ContactPoint'")
	}
	if args.GroupBies == nil {
		return nil, errors.New("invalid value for required argument 'GroupBies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NotificationPolicy
	err := ctx.RegisterResource("grafana:index/notificationPolicy:NotificationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationPolicy gets an existing NotificationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationPolicyState, opts ...pulumi.ResourceOption) (*NotificationPolicy, error) {
	var resource NotificationPolicy
	err := ctx.ReadResource("grafana:index/notificationPolicy:NotificationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationPolicy resources.
type notificationPolicyState struct {
	// The default contact point to route all unmatched notifications to.
	ContactPoint *string `pulumi:"contactPoint"`
	// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
	GroupBies []string `pulumi:"groupBies"`
	// Minimum time interval between two notifications for the same group. Default is 5 minutes.
	GroupInterval *string `pulumi:"groupInterval"`
	// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
	GroupWait *string `pulumi:"groupWait"`
	// Routing rules for specific label sets.
	Policies []NotificationPolicyPolicy `pulumi:"policies"`
	// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
	RepeatInterval *string `pulumi:"repeatInterval"`
}

type NotificationPolicyState struct {
	// The default contact point to route all unmatched notifications to.
	ContactPoint pulumix.Input[*string]
	// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
	GroupBies pulumix.Input[[]string]
	// Minimum time interval between two notifications for the same group. Default is 5 minutes.
	GroupInterval pulumix.Input[*string]
	// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
	GroupWait pulumix.Input[*string]
	// Routing rules for specific label sets.
	Policies pulumix.Input[[]*NotificationPolicyPolicyArgs]
	// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
	RepeatInterval pulumix.Input[*string]
}

func (NotificationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationPolicyState)(nil)).Elem()
}

type notificationPolicyArgs struct {
	// The default contact point to route all unmatched notifications to.
	ContactPoint string `pulumi:"contactPoint"`
	// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
	GroupBies []string `pulumi:"groupBies"`
	// Minimum time interval between two notifications for the same group. Default is 5 minutes.
	GroupInterval *string `pulumi:"groupInterval"`
	// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
	GroupWait *string `pulumi:"groupWait"`
	// Routing rules for specific label sets.
	Policies []NotificationPolicyPolicy `pulumi:"policies"`
	// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
	RepeatInterval *string `pulumi:"repeatInterval"`
}

// The set of arguments for constructing a NotificationPolicy resource.
type NotificationPolicyArgs struct {
	// The default contact point to route all unmatched notifications to.
	ContactPoint pulumix.Input[string]
	// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
	GroupBies pulumix.Input[[]string]
	// Minimum time interval between two notifications for the same group. Default is 5 minutes.
	GroupInterval pulumix.Input[*string]
	// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
	GroupWait pulumix.Input[*string]
	// Routing rules for specific label sets.
	Policies pulumix.Input[[]*NotificationPolicyPolicyArgs]
	// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
	RepeatInterval pulumix.Input[*string]
}

func (NotificationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationPolicyArgs)(nil)).Elem()
}

type NotificationPolicyOutput struct{ *pulumi.OutputState }

func (NotificationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPolicy)(nil)).Elem()
}

func (o NotificationPolicyOutput) ToNotificationPolicyOutput() NotificationPolicyOutput {
	return o
}

func (o NotificationPolicyOutput) ToNotificationPolicyOutputWithContext(ctx context.Context) NotificationPolicyOutput {
	return o
}

func (o NotificationPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[NotificationPolicy] {
	return pulumix.Output[NotificationPolicy]{
		OutputState: o.OutputState,
	}
}

// The default contact point to route all unmatched notifications to.
func (o NotificationPolicyOutput) ContactPoint() pulumix.Output[string] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.Output[string] { return v.ContactPoint })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
func (o NotificationPolicyOutput) GroupBies() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.ArrayOutput[string] { return v.GroupBies })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

// Minimum time interval between two notifications for the same group. Default is 5 minutes.
func (o NotificationPolicyOutput) GroupInterval() pulumix.Output[*string] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.Output[*string] { return v.GroupInterval })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
func (o NotificationPolicyOutput) GroupWait() pulumix.Output[*string] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.Output[*string] { return v.GroupWait })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Routing rules for specific label sets.
func (o NotificationPolicyOutput) Policies() pulumix.GArrayOutput[NotificationPolicyPolicy, NotificationPolicyPolicyOutput] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.GArrayOutput[NotificationPolicyPolicy, NotificationPolicyPolicyOutput] {
		return v.Policies
	})
	unwrapped := pulumix.Flatten[[]NotificationPolicyPolicy, pulumix.GArrayOutput[NotificationPolicyPolicy, NotificationPolicyPolicyOutput]](value)
	return pulumix.GArrayOutput[NotificationPolicyPolicy, NotificationPolicyPolicyOutput]{OutputState: unwrapped.OutputState}
}

// Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
func (o NotificationPolicyOutput) RepeatInterval() pulumix.Output[*string] {
	value := pulumix.Apply[NotificationPolicy](o, func(v NotificationPolicy) pulumix.Output[*string] { return v.RepeatInterval })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

func init() {
	pulumi.RegisterOutputType(NotificationPolicyOutput{})
}
