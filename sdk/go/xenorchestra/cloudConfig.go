// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xenorchestra

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/vatesfr/pulumi-xenorchestra/sdk/v2/go/xenorchestra/internal"
)

// Creates a Xen Orchestra cloud config resource.
//
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
//			demo, err := xenorchestra.NewCloudConfig(ctx, "demo", &xenorchestra.CloudConfigArgs{
//				Name: pulumi.String("cloud config name"),
//				Template: pulumi.String(`#cloud-config
//
// runcmd:
//   - [ ls, -l, / ]
//   - [ sh, -xc, "echo $(date) ': hello world!'" ]
//   - [ sh, -c, echo "=========hello world'=========" ]
//   - ls -l /root
//
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			_, err = xenorchestra.NewVm(ctx, "bar", &xenorchestra.VmArgs{
//				CloudConfig: demo.Template,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CloudConfig struct {
	pulumi.CustomResourceState

	// The name of the cloud config.
	Name pulumi.StringOutput `pulumi:"name"`
	// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewCloudConfig registers a new resource with the given unique name, arguments, and options.
func NewCloudConfig(ctx *pulumi.Context,
	name string, args *CloudConfigArgs, opts ...pulumi.ResourceOption) (*CloudConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudConfig
	err := ctx.RegisterResource("xenorchestra:index/cloudConfig:CloudConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudConfig gets an existing CloudConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudConfigState, opts ...pulumi.ResourceOption) (*CloudConfig, error) {
	var resource CloudConfig
	err := ctx.ReadResource("xenorchestra:index/cloudConfig:CloudConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudConfig resources.
type cloudConfigState struct {
	// The name of the cloud config.
	Name *string `pulumi:"name"`
	// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	Template *string `pulumi:"template"`
}

type CloudConfigState struct {
	// The name of the cloud config.
	Name pulumi.StringPtrInput
	// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	Template pulumi.StringPtrInput
}

func (CloudConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConfigState)(nil)).Elem()
}

type cloudConfigArgs struct {
	// The name of the cloud config.
	Name *string `pulumi:"name"`
	// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	Template string `pulumi:"template"`
}

// The set of arguments for constructing a CloudConfig resource.
type CloudConfigArgs struct {
	// The name of the cloud config.
	Name pulumi.StringPtrInput
	// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	Template pulumi.StringInput
}

func (CloudConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudConfigArgs)(nil)).Elem()
}

type CloudConfigInput interface {
	pulumi.Input

	ToCloudConfigOutput() CloudConfigOutput
	ToCloudConfigOutputWithContext(ctx context.Context) CloudConfigOutput
}

func (*CloudConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConfig)(nil)).Elem()
}

func (i *CloudConfig) ToCloudConfigOutput() CloudConfigOutput {
	return i.ToCloudConfigOutputWithContext(context.Background())
}

func (i *CloudConfig) ToCloudConfigOutputWithContext(ctx context.Context) CloudConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudConfigOutput)
}

// CloudConfigArrayInput is an input type that accepts CloudConfigArray and CloudConfigArrayOutput values.
// You can construct a concrete instance of `CloudConfigArrayInput` via:
//
//	CloudConfigArray{ CloudConfigArgs{...} }
type CloudConfigArrayInput interface {
	pulumi.Input

	ToCloudConfigArrayOutput() CloudConfigArrayOutput
	ToCloudConfigArrayOutputWithContext(context.Context) CloudConfigArrayOutput
}

type CloudConfigArray []CloudConfigInput

func (CloudConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudConfig)(nil)).Elem()
}

func (i CloudConfigArray) ToCloudConfigArrayOutput() CloudConfigArrayOutput {
	return i.ToCloudConfigArrayOutputWithContext(context.Background())
}

func (i CloudConfigArray) ToCloudConfigArrayOutputWithContext(ctx context.Context) CloudConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudConfigArrayOutput)
}

// CloudConfigMapInput is an input type that accepts CloudConfigMap and CloudConfigMapOutput values.
// You can construct a concrete instance of `CloudConfigMapInput` via:
//
//	CloudConfigMap{ "key": CloudConfigArgs{...} }
type CloudConfigMapInput interface {
	pulumi.Input

	ToCloudConfigMapOutput() CloudConfigMapOutput
	ToCloudConfigMapOutputWithContext(context.Context) CloudConfigMapOutput
}

type CloudConfigMap map[string]CloudConfigInput

func (CloudConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudConfig)(nil)).Elem()
}

func (i CloudConfigMap) ToCloudConfigMapOutput() CloudConfigMapOutput {
	return i.ToCloudConfigMapOutputWithContext(context.Background())
}

func (i CloudConfigMap) ToCloudConfigMapOutputWithContext(ctx context.Context) CloudConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudConfigMapOutput)
}

type CloudConfigOutput struct{ *pulumi.OutputState }

func (CloudConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudConfig)(nil)).Elem()
}

func (o CloudConfigOutput) ToCloudConfigOutput() CloudConfigOutput {
	return o
}

func (o CloudConfigOutput) ToCloudConfigOutputWithContext(ctx context.Context) CloudConfigOutput {
	return o
}

// The name of the cloud config.
func (o CloudConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
func (o CloudConfigOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudConfig) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

type CloudConfigArrayOutput struct{ *pulumi.OutputState }

func (CloudConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudConfig)(nil)).Elem()
}

func (o CloudConfigArrayOutput) ToCloudConfigArrayOutput() CloudConfigArrayOutput {
	return o
}

func (o CloudConfigArrayOutput) ToCloudConfigArrayOutputWithContext(ctx context.Context) CloudConfigArrayOutput {
	return o
}

func (o CloudConfigArrayOutput) Index(i pulumi.IntInput) CloudConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudConfig {
		return vs[0].([]*CloudConfig)[vs[1].(int)]
	}).(CloudConfigOutput)
}

type CloudConfigMapOutput struct{ *pulumi.OutputState }

func (CloudConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudConfig)(nil)).Elem()
}

func (o CloudConfigMapOutput) ToCloudConfigMapOutput() CloudConfigMapOutput {
	return o
}

func (o CloudConfigMapOutput) ToCloudConfigMapOutputWithContext(ctx context.Context) CloudConfigMapOutput {
	return o
}

func (o CloudConfigMapOutput) MapIndex(k pulumi.StringInput) CloudConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudConfig {
		return vs[0].(map[string]*CloudConfig)[vs[1].(string)]
	}).(CloudConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudConfigInput)(nil)).Elem(), &CloudConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudConfigArrayInput)(nil)).Elem(), CloudConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudConfigMapInput)(nil)).Elem(), CloudConfigMap{})
	pulumi.RegisterOutputType(CloudConfigOutput{})
	pulumi.RegisterOutputType(CloudConfigArrayOutput{})
	pulumi.RegisterOutputType(CloudConfigMapOutput{})
}
