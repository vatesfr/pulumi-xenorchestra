// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xenorchestra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/vatesfr/pulumi-xenorchestra/sdk/v2/go/xenorchestra/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/vatesfr/pulumi-xenorchestra/sdk/v2/go/xenorchestra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := xenorchestra.GetXoaCloudConfig(ctx, &xenorchestra.GetXoaCloudConfigArgs{
//				Name: "Name of cloud config",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetXoaCloudConfig(ctx *pulumi.Context, args *GetXoaCloudConfigArgs, opts ...pulumi.InvokeOption) (*GetXoaCloudConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetXoaCloudConfigResult
	err := ctx.Invoke("xenorchestra:index/getXoaCloudConfig:getXoaCloudConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getXoaCloudConfig.
type GetXoaCloudConfigArgs struct {
	// The name of the cloud config you want to look up.
	Name string `pulumi:"name"`
}

// A collection of values returned by getXoaCloudConfig.
type GetXoaCloudConfigResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the cloud config you want to look up.
	Name string `pulumi:"name"`
	// The contents of the cloud-config.
	Template string `pulumi:"template"`
}

func GetXoaCloudConfigOutput(ctx *pulumi.Context, args GetXoaCloudConfigOutputArgs, opts ...pulumi.InvokeOption) GetXoaCloudConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetXoaCloudConfigResultOutput, error) {
			args := v.(GetXoaCloudConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("xenorchestra:index/getXoaCloudConfig:getXoaCloudConfig", args, GetXoaCloudConfigResultOutput{}, options).(GetXoaCloudConfigResultOutput), nil
		}).(GetXoaCloudConfigResultOutput)
}

// A collection of arguments for invoking getXoaCloudConfig.
type GetXoaCloudConfigOutputArgs struct {
	// The name of the cloud config you want to look up.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetXoaCloudConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetXoaCloudConfigArgs)(nil)).Elem()
}

// A collection of values returned by getXoaCloudConfig.
type GetXoaCloudConfigResultOutput struct{ *pulumi.OutputState }

func (GetXoaCloudConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetXoaCloudConfigResult)(nil)).Elem()
}

func (o GetXoaCloudConfigResultOutput) ToGetXoaCloudConfigResultOutput() GetXoaCloudConfigResultOutput {
	return o
}

func (o GetXoaCloudConfigResultOutput) ToGetXoaCloudConfigResultOutputWithContext(ctx context.Context) GetXoaCloudConfigResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetXoaCloudConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetXoaCloudConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the cloud config you want to look up.
func (o GetXoaCloudConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetXoaCloudConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The contents of the cloud-config.
func (o GetXoaCloudConfigResultOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v GetXoaCloudConfigResult) string { return v.Template }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetXoaCloudConfigResultOutput{})
}
