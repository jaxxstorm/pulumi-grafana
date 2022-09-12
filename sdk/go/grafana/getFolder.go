// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/dashboard-folders/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/folder/)
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
//			test, err := grafana.NewFolder(ctx, "test", &grafana.FolderArgs{
//				Title: pulumi.String("test-folder"),
//				Uid:   pulumi.String("test-ds-folder-uid"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = grafana.LookupFolderOutput(ctx, GetFolderOutputArgs{
//				Title: test.Title,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFolderResult
	err := ctx.Invoke("grafana:index/getFolder:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type LookupFolderArgs struct {
	// The name of the Grafana folder.
	Title string `pulumi:"title"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	// The numerical ID of the Grafana folder.
	Id int `pulumi:"id"`
	// The name of the Grafana folder.
	Title string `pulumi:"title"`
	// The uid of the Grafana folder.
	Uid string `pulumi:"uid"`
	// The full URL of the folder.
	Url string `pulumi:"url"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderResult, error) {
			args := v.(LookupFolderArgs)
			r, err := LookupFolder(ctx, &args, opts...)
			var s LookupFolderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderResultOutput)
}

// A collection of arguments for invoking getFolder.
type LookupFolderOutputArgs struct {
	// The name of the Grafana folder.
	Title pulumi.StringInput `pulumi:"title"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

// A collection of values returned by getFolder.
type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// The numerical ID of the Grafana folder.
func (o LookupFolderResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFolderResult) int { return v.Id }).(pulumi.IntOutput)
}

// The name of the Grafana folder.
func (o LookupFolderResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Title }).(pulumi.StringOutput)
}

// The uid of the Grafana folder.
func (o LookupFolderResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The full URL of the folder.
func (o LookupFolderResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
