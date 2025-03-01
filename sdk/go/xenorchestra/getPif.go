// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xenorchestra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/vatesfr/pulumi-xenorchestra/sdk/go/xenorchestra/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/vatesfr/pulumi-xenorchestra/sdk/go/xenorchestra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			eth0, err := xenorchestra.GetXoaPif(ctx, &xenorchestra.GetXoaPifArgs{
//				Device: "eth0",
//				Vlan:   -1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = xenorchestra.NewVm(ctx, "demo-vm", &xenorchestra.VmArgs{
//				Networks: xenorchestra.VmNetworkArray{
//					&xenorchestra.VmNetworkArgs{
//						NetworkId: pulumi.String(eth0.Network),
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
// Deprecated: xenorchestra.index/getpif.getPif has been deprecated in favor of xenorchestra.index/getxoapif.getXoaPif
func GetPif(ctx *pulumi.Context, args *GetPifArgs, opts ...pulumi.InvokeOption) (*GetPifResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPifResult
	err := ctx.Invoke("xenorchestra:index/getPif:getPif", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPif.
type GetPifArgs struct {
	// The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
	Device string `pulumi:"device"`
	// The ID of the host that the PIF belongs to.
	HostId *string `pulumi:"hostId"`
	// The VLAN the PIF belongs to.
	Vlan int `pulumi:"vlan"`
}

// A collection of values returned by getPif.
type GetPifResult struct {
	// If the PIF is attached to the network.
	Attached bool `pulumi:"attached"`
	// The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
	Device string `pulumi:"device"`
	// The host the PIF is associated with.
	Host string `pulumi:"host"`
	// The ID of the host that the PIF belongs to.
	HostId string `pulumi:"hostId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The network the PIF is associated with.
	Network string `pulumi:"network"`
	// The pool the PIF is associated with.
	PoolId string `pulumi:"poolId"`
	// The uuid of the PIF.
	Uuid string `pulumi:"uuid"`
	// The VLAN the PIF belongs to.
	Vlan int `pulumi:"vlan"`
}

func GetPifOutput(ctx *pulumi.Context, args GetPifOutputArgs, opts ...pulumi.InvokeOption) GetPifResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetPifResultOutput, error) {
			args := v.(GetPifArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("xenorchestra:index/getPif:getPif", args, GetPifResultOutput{}, options).(GetPifResultOutput), nil
		}).(GetPifResultOutput)
}

// A collection of arguments for invoking getPif.
type GetPifOutputArgs struct {
	// The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
	Device pulumi.StringInput `pulumi:"device"`
	// The ID of the host that the PIF belongs to.
	HostId pulumi.StringPtrInput `pulumi:"hostId"`
	// The VLAN the PIF belongs to.
	Vlan pulumi.IntInput `pulumi:"vlan"`
}

func (GetPifOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPifArgs)(nil)).Elem()
}

// A collection of values returned by getPif.
type GetPifResultOutput struct{ *pulumi.OutputState }

func (GetPifResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPifResult)(nil)).Elem()
}

func (o GetPifResultOutput) ToGetPifResultOutput() GetPifResultOutput {
	return o
}

func (o GetPifResultOutput) ToGetPifResultOutputWithContext(ctx context.Context) GetPifResultOutput {
	return o
}

// If the PIF is attached to the network.
func (o GetPifResultOutput) Attached() pulumi.BoolOutput {
	return o.ApplyT(func(v GetPifResult) bool { return v.Attached }).(pulumi.BoolOutput)
}

// The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
func (o GetPifResultOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.Device }).(pulumi.StringOutput)
}

// The host the PIF is associated with.
func (o GetPifResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.Host }).(pulumi.StringOutput)
}

// The ID of the host that the PIF belongs to.
func (o GetPifResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.HostId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPifResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.Id }).(pulumi.StringOutput)
}

// The network the PIF is associated with.
func (o GetPifResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.Network }).(pulumi.StringOutput)
}

// The pool the PIF is associated with.
func (o GetPifResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// The uuid of the PIF.
func (o GetPifResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetPifResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The VLAN the PIF belongs to.
func (o GetPifResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v GetPifResult) int { return v.Vlan }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPifResultOutput{})
}
