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
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.GetOncallSlackChannel(ctx, &GetOncallSlackChannelArgs{
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
func GetOncallSlackChannel(ctx *pulumi.Context, args *GetOncallSlackChannelArgs, opts ...pulumi.InvokeOption) (*GetOncallSlackChannelResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOncallSlackChannelResult
	err := ctx.Invoke("grafana:index/getOncallSlackChannel:getOncallSlackChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallSlackChannel.
type GetOncallSlackChannelArgs struct {
	// The Slack channel name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOncallSlackChannel.
type GetOncallSlackChannelResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Slack channel name.
	Name string `pulumi:"name"`
	// The Slack ID of the channel.
	SlackId string `pulumi:"slackId"`
}

func GetOncallSlackChannelOutput(ctx *pulumi.Context, args GetOncallSlackChannelOutputArgs, opts ...pulumi.InvokeOption) GetOncallSlackChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOncallSlackChannelResult, error) {
			args := v.(GetOncallSlackChannelArgs)
			r, err := GetOncallSlackChannel(ctx, &args, opts...)
			var s GetOncallSlackChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOncallSlackChannelResultOutput)
}

// A collection of arguments for invoking getOncallSlackChannel.
type GetOncallSlackChannelOutputArgs struct {
	// The Slack channel name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOncallSlackChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallSlackChannelArgs)(nil)).Elem()
}

// A collection of values returned by getOncallSlackChannel.
type GetOncallSlackChannelResultOutput struct{ *pulumi.OutputState }

func (GetOncallSlackChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallSlackChannelResult)(nil)).Elem()
}

func (o GetOncallSlackChannelResultOutput) ToGetOncallSlackChannelResultOutput() GetOncallSlackChannelResultOutput {
	return o
}

func (o GetOncallSlackChannelResultOutput) ToGetOncallSlackChannelResultOutputWithContext(ctx context.Context) GetOncallSlackChannelResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOncallSlackChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallSlackChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Slack channel name.
func (o GetOncallSlackChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallSlackChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Slack ID of the channel.
func (o GetOncallSlackChannelResultOutput) SlackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallSlackChannelResult) string { return v.SlackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOncallSlackChannelResultOutput{})
}
