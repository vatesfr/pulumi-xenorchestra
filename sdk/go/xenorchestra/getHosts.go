// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xenorchestra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/vatesfr/pulumi-xenorchestra/sdk/go/xenorchestra/internal"
)

// Use this data source to filter Xenorchestra hosts by certain criteria (name_label, tags) for use in other resources.
//
// Deprecated: xenorchestra.index/gethosts.getHosts has been deprecated in favor of xenorchestra.index/getxoahosts.getXoaHosts
func GetHosts(ctx *pulumi.Context, args *GetHostsArgs, opts ...pulumi.InvokeOption) (*GetHostsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHostsResult
	err := ctx.Invoke("xenorchestra:index/getHosts:getHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHosts.
type GetHostsArgs struct {
	// The pool id used to filter the resulting hosts by.
	PoolId string `pulumi:"poolId"`
	// The host field to sort the results by (id and nameLabel are supported).
	SortBy *string `pulumi:"sortBy"`
	// Valid options are `asc` or `desc` and sort order is applied to `sortBy` argument.
	SortOrder *string `pulumi:"sortOrder"`
	// The tags (labels) applied to the given entity.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getHosts.
type GetHostsResult struct {
	// The resulting hosts after applying the argument filtering.
	Hosts []GetHostsHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The primary host of the pool.
	Master string `pulumi:"master"`
	// The pool id used to filter the resulting hosts by.
	PoolId string `pulumi:"poolId"`
	// The host field to sort the results by (id and nameLabel are supported).
	SortBy *string `pulumi:"sortBy"`
	// Valid options are `asc` or `desc` and sort order is applied to `sortBy` argument.
	SortOrder *string `pulumi:"sortOrder"`
	// The tags (labels) applied to the given entity.
	Tags []string `pulumi:"tags"`
}

func GetHostsOutput(ctx *pulumi.Context, args GetHostsOutputArgs, opts ...pulumi.InvokeOption) GetHostsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetHostsResultOutput, error) {
			args := v.(GetHostsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("xenorchestra:index/getHosts:getHosts", args, GetHostsResultOutput{}, options).(GetHostsResultOutput), nil
		}).(GetHostsResultOutput)
}

// A collection of arguments for invoking getHosts.
type GetHostsOutputArgs struct {
	// The pool id used to filter the resulting hosts by.
	PoolId pulumi.StringInput `pulumi:"poolId"`
	// The host field to sort the results by (id and nameLabel are supported).
	SortBy pulumi.StringPtrInput `pulumi:"sortBy"`
	// Valid options are `asc` or `desc` and sort order is applied to `sortBy` argument.
	SortOrder pulumi.StringPtrInput `pulumi:"sortOrder"`
	// The tags (labels) applied to the given entity.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (GetHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostsArgs)(nil)).Elem()
}

// A collection of values returned by getHosts.
type GetHostsResultOutput struct{ *pulumi.OutputState }

func (GetHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostsResult)(nil)).Elem()
}

func (o GetHostsResultOutput) ToGetHostsResultOutput() GetHostsResultOutput {
	return o
}

func (o GetHostsResultOutput) ToGetHostsResultOutputWithContext(ctx context.Context) GetHostsResultOutput {
	return o
}

// The resulting hosts after applying the argument filtering.
func (o GetHostsResultOutput) Hosts() GetHostsHostArrayOutput {
	return o.ApplyT(func(v GetHostsResult) []GetHostsHost { return v.Hosts }).(GetHostsHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The primary host of the pool.
func (o GetHostsResultOutput) Master() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostsResult) string { return v.Master }).(pulumi.StringOutput)
}

// The pool id used to filter the resulting hosts by.
func (o GetHostsResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostsResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// The host field to sort the results by (id and nameLabel are supported).
func (o GetHostsResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostsResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

// Valid options are `asc` or `desc` and sort order is applied to `sortBy` argument.
func (o GetHostsResultOutput) SortOrder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostsResult) *string { return v.SortOrder }).(pulumi.StringPtrOutput)
}

// The tags (labels) applied to the given entity.
func (o GetHostsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHostsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostsResultOutput{})
}
