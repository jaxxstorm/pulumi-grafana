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

// **Note:** This resource is available only with Grafana Enterprise 8.+.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/access_control/)
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
//			test, err := grafana.NewRole(ctx, "test", &grafana.RoleArgs{
//				Description: pulumi.String("test-role description"),
//				Uid:         pulumi.String("test-ds-role-uid"),
//				Version:     pulumi.Int(1),
//				Global:      pulumi.Bool(true),
//				Permissions: grafana.RolePermissionArray{
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:add"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:write"),
//						Scope:  pulumi.String("users:*"),
//					},
//					&grafana.RolePermissionArgs{
//						Action: pulumi.String("org.users:read"),
//						Scope:  pulumi.String("users:*"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = grafana.LookupRoleOutput(ctx, grafana.GetRoleOutputArgs{
//				Name: test.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleResult
	err := ctx.Invoke("grafana:index/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	// Name of the role
	Name string `pulumi:"name"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	// Description of the role.
	Description string `pulumi:"description"`
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName string `pulumi:"displayName"`
	// Boolean to state whether the role is available across all organizations or not.
	Global bool `pulumi:"global"`
	// Group of the role. Available with Grafana 8.5+.
	Group string `pulumi:"group"`
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+.
	Hidden bool `pulumi:"hidden"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the role
	Name string `pulumi:"name"`
	// Specific set of actions granted by the role.
	Permissions []GetRolePermission `pulumi:"permissions"`
	// Unique identifier of the role. Used for assignments.
	Uid string `pulumi:"uid"`
	// Version of the role. A role is updated only on version increase.
	Version int `pulumi:"version"`
}

func LookupRoleOutput(ctx *pulumi.Context, args LookupRoleOutputArgs, opts ...pulumi.InvokeOption) LookupRoleResultOutput {
	outputResult := pulumix.ApplyErr[*LookupRoleArgs](args.ToOutput(), func(plainArgs *LookupRoleArgs) (*LookupRoleResult, error) {
		return LookupRole(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[LookupRoleResultOutput, *LookupRoleResult](outputResult)
}

// A collection of arguments for invoking getRole.
type LookupRoleOutputArgs struct {
	// Name of the role
	Name pulumix.Input[string] `pulumi:"name"`
}

func (args LookupRoleOutputArgs) ToOutput() pulumix.Output[*LookupRoleArgs] {
	allArgs := pulumix.All(
		args.Name.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *LookupRoleArgs {
		return &LookupRoleArgs{
			Name: resolvedArgs[0].(string),
		}
	})
}

type LookupRoleResultOutput struct{ *pulumi.OutputState }

func (LookupRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleResult)(nil)).Elem()
}

func (o LookupRoleResultOutput) ToOutput(context.Context) pulumix.Output[*LookupRoleResult] {
	return pulumix.Output[*LookupRoleResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRoleResultOutput) Description() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.Description })
}

func (o LookupRoleResultOutput) DisplayName() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.DisplayName })
}

func (o LookupRoleResultOutput) Global() pulumix.Output[bool] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) bool { return v.Global })
}

func (o LookupRoleResultOutput) Group() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.Group })
}

func (o LookupRoleResultOutput) Hidden() pulumix.Output[bool] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) bool { return v.Hidden })
}

func (o LookupRoleResultOutput) Id() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.Id })
}

func (o LookupRoleResultOutput) Name() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.Name })
}

func (o LookupRoleResultOutput) Permissions() pulumix.GArrayOutput[GetRolePermission, GetRolePermissionOutput] {
	value := pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) []GetRolePermission { return v.Permissions })
	return pulumix.GArrayOutput[GetRolePermission, GetRolePermissionOutput]{
		OutputState: value.OutputState,
	}
}

func (o LookupRoleResultOutput) Uid() pulumix.Output[string] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) string { return v.Uid })
}

func (o LookupRoleResultOutput) Version() pulumix.Output[int] {
	return pulumix.Apply[*LookupRoleResult](o, func(v *LookupRoleResult) int { return v.Version })
}
