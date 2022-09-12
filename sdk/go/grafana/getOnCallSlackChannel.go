// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/slack_channels/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.GetOnCallSlackChannel(ctx, &GetOnCallSlackChannelArgs{
//				Name: "example_slack_channel",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOnCallSlackChannel(ctx *pulumi.Context, args *GetOnCallSlackChannelArgs, opts ...pulumi.InvokeOption) (*GetOnCallSlackChannelResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOnCallSlackChannelResult
	err := ctx.Invoke("grafana:index/getOnCallSlackChannel:getOnCallSlackChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOnCallSlackChannel.
type GetOnCallSlackChannelArgs struct {
	// The Slack channel name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOnCallSlackChannel.
type GetOnCallSlackChannelResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Slack channel name.
	Name string `pulumi:"name"`
	// The Slack ID of the channel.
	SlackId string `pulumi:"slackId"`
}

func GetOnCallSlackChannelOutput(ctx *pulumi.Context, args GetOnCallSlackChannelOutputArgs, opts ...pulumi.InvokeOption) GetOnCallSlackChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOnCallSlackChannelResult, error) {
			args := v.(GetOnCallSlackChannelArgs)
			r, err := GetOnCallSlackChannel(ctx, &args, opts...)
			var s GetOnCallSlackChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOnCallSlackChannelResultOutput)
}

// A collection of arguments for invoking getOnCallSlackChannel.
type GetOnCallSlackChannelOutputArgs struct {
	// The Slack channel name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOnCallSlackChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnCallSlackChannelArgs)(nil)).Elem()
}

// A collection of values returned by getOnCallSlackChannel.
type GetOnCallSlackChannelResultOutput struct{ *pulumi.OutputState }

func (GetOnCallSlackChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnCallSlackChannelResult)(nil)).Elem()
}

func (o GetOnCallSlackChannelResultOutput) ToGetOnCallSlackChannelResultOutput() GetOnCallSlackChannelResultOutput {
	return o
}

func (o GetOnCallSlackChannelResultOutput) ToGetOnCallSlackChannelResultOutputWithContext(ctx context.Context) GetOnCallSlackChannelResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOnCallSlackChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOnCallSlackChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Slack channel name.
func (o GetOnCallSlackChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOnCallSlackChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Slack ID of the channel.
func (o GetOnCallSlackChannelResultOutput) SlackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOnCallSlackChannelResult) string { return v.SlackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOnCallSlackChannelResultOutput{})
}
