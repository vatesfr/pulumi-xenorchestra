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

// ## Example Usage
type Vm struct {
	pulumi.CustomResourceState

	// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
	// be rescheduled.
	AffinityHost pulumi.StringPtrOutput `pulumi:"affinityHost"`
	// If the VM will automatically turn on. Defaults to `false`.
	AutoPoweron pulumi.BoolPtrOutput `pulumi:"autoPoweron"`
	// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
	// hard_shutdown, pause, shutdown, suspend, destroy. See:
	// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
	BlockedOperations pulumi.StringArrayOutput `pulumi:"blockedOperations"`
	// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
	// from `xe template-list`) and install the OS from the following ISO.
	Cdrom VmCdromPtrOutput `pulumi:"cdrom"`
	// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
	// perform a `full` clone, the VM template must not be a disk template.
	CloneType pulumi.StringPtrOutput `pulumi:"cloneType"`
	// The content of the cloud-init config to use. See the cloud init docs for more
	// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	CloudConfig pulumi.StringPtrOutput `pulumi:"cloudConfig"`
	// The content of the cloud-init network configuration for the VM (uses [version
	// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
	CloudNetworkConfig pulumi.StringPtrOutput `pulumi:"cloudNetworkConfig"`
	CoreOs             pulumi.BoolPtrOutput   `pulumi:"coreOs"`
	CpuCap             pulumi.IntPtrOutput    `pulumi:"cpuCap"`
	CpuWeight          pulumi.IntPtrOutput    `pulumi:"cpuWeight"`
	// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
	// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
	// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
	// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
	Cpus pulumi.IntOutput `pulumi:"cpus"`
	// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
	// powerState must be set to `Running`.
	DestroyCloudConfigVdiAfterBoot pulumi.BoolPtrOutput `pulumi:"destroyCloudConfigVdiAfterBoot"`
	// The disk the VM will have access to.
	Disks VmDiskArrayOutput `pulumi:"disks"`
	// Boolean parameter that allows a VM to use nested virtualization.
	ExpNestedHvm pulumi.BoolPtrOutput `pulumi:"expNestedHvm"`
	// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
	// Defaults to empty string
	HighAvailability pulumi.StringPtrOutput `pulumi:"highAvailability"`
	Host             pulumi.StringPtrOutput `pulumi:"host"`
	// The firmware to use for the VM. Possible values are `bios` and `uefi`.
	HvmBootFirmware pulumi.StringPtrOutput `pulumi:"hvmBootFirmware"`
	// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
	InstallationMethod pulumi.StringPtrOutput   `pulumi:"installationMethod"`
	Ipv4Addresses      pulumi.StringArrayOutput `pulumi:"ipv4Addresses"`
	// This is only accessible if guest-tools is installed in the VM and if `expectedIpCidr` is set on any network interfaces.
	// This will contain a list of the ipv6 addresses across all network interfaces in order.
	Ipv6Addresses pulumi.StringArrayOutput `pulumi:"ipv6Addresses"`
	// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
	// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
	// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
	// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
	// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
	MemoryMax pulumi.IntOutput `pulumi:"memoryMax"`
	// The description of the VM.
	NameDescription pulumi.StringPtrOutput `pulumi:"nameDescription"`
	// The name of the VM.
	NameLabel pulumi.StringOutput `pulumi:"nameLabel"`
	// The network for the VM.
	Networks VmNetworkArrayOutput `pulumi:"networks"`
	// The power state of the VM. This can be Running, Halted, Paused or Suspended.
	PowerState  pulumi.StringPtrOutput `pulumi:"powerState"`
	ResourceSet pulumi.StringPtrOutput `pulumi:"resourceSet"`
	// Number of seconds the VM should be delayed from starting.
	StartDelay pulumi.IntPtrOutput `pulumi:"startDelay"`
	// The tags (labels) applied to the given entity.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The ID of the VM template to create the new VM from.
	Template pulumi.StringOutput `pulumi:"template"`
	// The video adapter the VM should use. Possible values include std and cirrus.
	Vga pulumi.StringPtrOutput `pulumi:"vga"`
	// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
	Videoram pulumi.IntPtrOutput `pulumi:"videoram"`
	// The key value pairs to be populated in xenstore.
	Xenstore pulumi.StringMapOutput `pulumi:"xenstore"`
}

// NewVm registers a new resource with the given unique name, arguments, and options.
func NewVm(ctx *pulumi.Context,
	name string, args *VmArgs, opts ...pulumi.ResourceOption) (*Vm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cpus == nil {
		return nil, errors.New("invalid value for required argument 'Cpus'")
	}
	if args.Disks == nil {
		return nil, errors.New("invalid value for required argument 'Disks'")
	}
	if args.MemoryMax == nil {
		return nil, errors.New("invalid value for required argument 'MemoryMax'")
	}
	if args.NameLabel == nil {
		return nil, errors.New("invalid value for required argument 'NameLabel'")
	}
	if args.Networks == nil {
		return nil, errors.New("invalid value for required argument 'Networks'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vm
	err := ctx.RegisterResource("xenorchestra:index/vm:Vm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVm gets an existing Vm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmState, opts ...pulumi.ResourceOption) (*Vm, error) {
	var resource Vm
	err := ctx.ReadResource("xenorchestra:index/vm:Vm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vm resources.
type vmState struct {
	// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
	// be rescheduled.
	AffinityHost *string `pulumi:"affinityHost"`
	// If the VM will automatically turn on. Defaults to `false`.
	AutoPoweron *bool `pulumi:"autoPoweron"`
	// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
	// hard_shutdown, pause, shutdown, suspend, destroy. See:
	// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
	BlockedOperations []string `pulumi:"blockedOperations"`
	// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
	// from `xe template-list`) and install the OS from the following ISO.
	Cdrom *VmCdrom `pulumi:"cdrom"`
	// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
	// perform a `full` clone, the VM template must not be a disk template.
	CloneType *string `pulumi:"cloneType"`
	// The content of the cloud-init config to use. See the cloud init docs for more
	// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	CloudConfig *string `pulumi:"cloudConfig"`
	// The content of the cloud-init network configuration for the VM (uses [version
	// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
	CloudNetworkConfig *string `pulumi:"cloudNetworkConfig"`
	CoreOs             *bool   `pulumi:"coreOs"`
	CpuCap             *int    `pulumi:"cpuCap"`
	CpuWeight          *int    `pulumi:"cpuWeight"`
	// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
	// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
	// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
	// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
	Cpus *int `pulumi:"cpus"`
	// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
	// powerState must be set to `Running`.
	DestroyCloudConfigVdiAfterBoot *bool `pulumi:"destroyCloudConfigVdiAfterBoot"`
	// The disk the VM will have access to.
	Disks []VmDisk `pulumi:"disks"`
	// Boolean parameter that allows a VM to use nested virtualization.
	ExpNestedHvm *bool `pulumi:"expNestedHvm"`
	// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
	// Defaults to empty string
	HighAvailability *string `pulumi:"highAvailability"`
	Host             *string `pulumi:"host"`
	// The firmware to use for the VM. Possible values are `bios` and `uefi`.
	HvmBootFirmware *string `pulumi:"hvmBootFirmware"`
	// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
	InstallationMethod *string  `pulumi:"installationMethod"`
	Ipv4Addresses      []string `pulumi:"ipv4Addresses"`
	// This is only accessible if guest-tools is installed in the VM and if `expectedIpCidr` is set on any network interfaces.
	// This will contain a list of the ipv6 addresses across all network interfaces in order.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
	// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
	// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
	// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
	// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
	MemoryMax *int `pulumi:"memoryMax"`
	// The description of the VM.
	NameDescription *string `pulumi:"nameDescription"`
	// The name of the VM.
	NameLabel *string `pulumi:"nameLabel"`
	// The network for the VM.
	Networks []VmNetwork `pulumi:"networks"`
	// The power state of the VM. This can be Running, Halted, Paused or Suspended.
	PowerState  *string `pulumi:"powerState"`
	ResourceSet *string `pulumi:"resourceSet"`
	// Number of seconds the VM should be delayed from starting.
	StartDelay *int `pulumi:"startDelay"`
	// The tags (labels) applied to the given entity.
	Tags []string `pulumi:"tags"`
	// The ID of the VM template to create the new VM from.
	Template *string `pulumi:"template"`
	// The video adapter the VM should use. Possible values include std and cirrus.
	Vga *string `pulumi:"vga"`
	// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
	Videoram *int `pulumi:"videoram"`
	// The key value pairs to be populated in xenstore.
	Xenstore map[string]string `pulumi:"xenstore"`
}

type VmState struct {
	// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
	// be rescheduled.
	AffinityHost pulumi.StringPtrInput
	// If the VM will automatically turn on. Defaults to `false`.
	AutoPoweron pulumi.BoolPtrInput
	// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
	// hard_shutdown, pause, shutdown, suspend, destroy. See:
	// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
	BlockedOperations pulumi.StringArrayInput
	// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
	// from `xe template-list`) and install the OS from the following ISO.
	Cdrom VmCdromPtrInput
	// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
	// perform a `full` clone, the VM template must not be a disk template.
	CloneType pulumi.StringPtrInput
	// The content of the cloud-init config to use. See the cloud init docs for more
	// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	CloudConfig pulumi.StringPtrInput
	// The content of the cloud-init network configuration for the VM (uses [version
	// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
	CloudNetworkConfig pulumi.StringPtrInput
	CoreOs             pulumi.BoolPtrInput
	CpuCap             pulumi.IntPtrInput
	CpuWeight          pulumi.IntPtrInput
	// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
	// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
	// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
	// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
	Cpus pulumi.IntPtrInput
	// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
	// powerState must be set to `Running`.
	DestroyCloudConfigVdiAfterBoot pulumi.BoolPtrInput
	// The disk the VM will have access to.
	Disks VmDiskArrayInput
	// Boolean parameter that allows a VM to use nested virtualization.
	ExpNestedHvm pulumi.BoolPtrInput
	// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
	// Defaults to empty string
	HighAvailability pulumi.StringPtrInput
	Host             pulumi.StringPtrInput
	// The firmware to use for the VM. Possible values are `bios` and `uefi`.
	HvmBootFirmware pulumi.StringPtrInput
	// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
	InstallationMethod pulumi.StringPtrInput
	Ipv4Addresses      pulumi.StringArrayInput
	// This is only accessible if guest-tools is installed in the VM and if `expectedIpCidr` is set on any network interfaces.
	// This will contain a list of the ipv6 addresses across all network interfaces in order.
	Ipv6Addresses pulumi.StringArrayInput
	// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
	// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
	// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
	// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
	// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
	MemoryMax pulumi.IntPtrInput
	// The description of the VM.
	NameDescription pulumi.StringPtrInput
	// The name of the VM.
	NameLabel pulumi.StringPtrInput
	// The network for the VM.
	Networks VmNetworkArrayInput
	// The power state of the VM. This can be Running, Halted, Paused or Suspended.
	PowerState  pulumi.StringPtrInput
	ResourceSet pulumi.StringPtrInput
	// Number of seconds the VM should be delayed from starting.
	StartDelay pulumi.IntPtrInput
	// The tags (labels) applied to the given entity.
	Tags pulumi.StringArrayInput
	// The ID of the VM template to create the new VM from.
	Template pulumi.StringPtrInput
	// The video adapter the VM should use. Possible values include std and cirrus.
	Vga pulumi.StringPtrInput
	// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
	Videoram pulumi.IntPtrInput
	// The key value pairs to be populated in xenstore.
	Xenstore pulumi.StringMapInput
}

func (VmState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmState)(nil)).Elem()
}

type vmArgs struct {
	// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
	// be rescheduled.
	AffinityHost *string `pulumi:"affinityHost"`
	// If the VM will automatically turn on. Defaults to `false`.
	AutoPoweron *bool `pulumi:"autoPoweron"`
	// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
	// hard_shutdown, pause, shutdown, suspend, destroy. See:
	// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
	BlockedOperations []string `pulumi:"blockedOperations"`
	// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
	// from `xe template-list`) and install the OS from the following ISO.
	Cdrom *VmCdrom `pulumi:"cdrom"`
	// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
	// perform a `full` clone, the VM template must not be a disk template.
	CloneType *string `pulumi:"cloneType"`
	// The content of the cloud-init config to use. See the cloud init docs for more
	// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	CloudConfig *string `pulumi:"cloudConfig"`
	// The content of the cloud-init network configuration for the VM (uses [version
	// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
	CloudNetworkConfig *string `pulumi:"cloudNetworkConfig"`
	CoreOs             *bool   `pulumi:"coreOs"`
	CpuCap             *int    `pulumi:"cpuCap"`
	CpuWeight          *int    `pulumi:"cpuWeight"`
	// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
	// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
	// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
	// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
	Cpus int `pulumi:"cpus"`
	// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
	// powerState must be set to `Running`.
	DestroyCloudConfigVdiAfterBoot *bool `pulumi:"destroyCloudConfigVdiAfterBoot"`
	// The disk the VM will have access to.
	Disks []VmDisk `pulumi:"disks"`
	// Boolean parameter that allows a VM to use nested virtualization.
	ExpNestedHvm *bool `pulumi:"expNestedHvm"`
	// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
	// Defaults to empty string
	HighAvailability *string `pulumi:"highAvailability"`
	Host             *string `pulumi:"host"`
	// The firmware to use for the VM. Possible values are `bios` and `uefi`.
	HvmBootFirmware *string `pulumi:"hvmBootFirmware"`
	// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
	InstallationMethod *string `pulumi:"installationMethod"`
	// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
	// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
	// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
	// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
	// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
	MemoryMax int `pulumi:"memoryMax"`
	// The description of the VM.
	NameDescription *string `pulumi:"nameDescription"`
	// The name of the VM.
	NameLabel string `pulumi:"nameLabel"`
	// The network for the VM.
	Networks []VmNetwork `pulumi:"networks"`
	// The power state of the VM. This can be Running, Halted, Paused or Suspended.
	PowerState  *string `pulumi:"powerState"`
	ResourceSet *string `pulumi:"resourceSet"`
	// Number of seconds the VM should be delayed from starting.
	StartDelay *int `pulumi:"startDelay"`
	// The tags (labels) applied to the given entity.
	Tags []string `pulumi:"tags"`
	// The ID of the VM template to create the new VM from.
	Template string `pulumi:"template"`
	// The video adapter the VM should use. Possible values include std and cirrus.
	Vga *string `pulumi:"vga"`
	// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
	Videoram *int `pulumi:"videoram"`
	// The key value pairs to be populated in xenstore.
	Xenstore map[string]string `pulumi:"xenstore"`
}

// The set of arguments for constructing a Vm resource.
type VmArgs struct {
	// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
	// be rescheduled.
	AffinityHost pulumi.StringPtrInput
	// If the VM will automatically turn on. Defaults to `false`.
	AutoPoweron pulumi.BoolPtrInput
	// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
	// hard_shutdown, pause, shutdown, suspend, destroy. See:
	// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
	BlockedOperations pulumi.StringArrayInput
	// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
	// from `xe template-list`) and install the OS from the following ISO.
	Cdrom VmCdromPtrInput
	// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
	// perform a `full` clone, the VM template must not be a disk template.
	CloneType pulumi.StringPtrInput
	// The content of the cloud-init config to use. See the cloud init docs for more
	// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
	CloudConfig pulumi.StringPtrInput
	// The content of the cloud-init network configuration for the VM (uses [version
	// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
	CloudNetworkConfig pulumi.StringPtrInput
	CoreOs             pulumi.BoolPtrInput
	CpuCap             pulumi.IntPtrInput
	CpuWeight          pulumi.IntPtrInput
	// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
	// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
	// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
	// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
	Cpus pulumi.IntInput
	// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
	// powerState must be set to `Running`.
	DestroyCloudConfigVdiAfterBoot pulumi.BoolPtrInput
	// The disk the VM will have access to.
	Disks VmDiskArrayInput
	// Boolean parameter that allows a VM to use nested virtualization.
	ExpNestedHvm pulumi.BoolPtrInput
	// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
	// Defaults to empty string
	HighAvailability pulumi.StringPtrInput
	Host             pulumi.StringPtrInput
	// The firmware to use for the VM. Possible values are `bios` and `uefi`.
	HvmBootFirmware pulumi.StringPtrInput
	// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
	InstallationMethod pulumi.StringPtrInput
	// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
	// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
	// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
	// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
	// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
	MemoryMax pulumi.IntInput
	// The description of the VM.
	NameDescription pulumi.StringPtrInput
	// The name of the VM.
	NameLabel pulumi.StringInput
	// The network for the VM.
	Networks VmNetworkArrayInput
	// The power state of the VM. This can be Running, Halted, Paused or Suspended.
	PowerState  pulumi.StringPtrInput
	ResourceSet pulumi.StringPtrInput
	// Number of seconds the VM should be delayed from starting.
	StartDelay pulumi.IntPtrInput
	// The tags (labels) applied to the given entity.
	Tags pulumi.StringArrayInput
	// The ID of the VM template to create the new VM from.
	Template pulumi.StringInput
	// The video adapter the VM should use. Possible values include std and cirrus.
	Vga pulumi.StringPtrInput
	// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
	Videoram pulumi.IntPtrInput
	// The key value pairs to be populated in xenstore.
	Xenstore pulumi.StringMapInput
}

func (VmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmArgs)(nil)).Elem()
}

type VmInput interface {
	pulumi.Input

	ToVmOutput() VmOutput
	ToVmOutputWithContext(ctx context.Context) VmOutput
}

func (*Vm) ElementType() reflect.Type {
	return reflect.TypeOf((**Vm)(nil)).Elem()
}

func (i *Vm) ToVmOutput() VmOutput {
	return i.ToVmOutputWithContext(context.Background())
}

func (i *Vm) ToVmOutputWithContext(ctx context.Context) VmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmOutput)
}

// VmArrayInput is an input type that accepts VmArray and VmArrayOutput values.
// You can construct a concrete instance of `VmArrayInput` via:
//
//	VmArray{ VmArgs{...} }
type VmArrayInput interface {
	pulumi.Input

	ToVmArrayOutput() VmArrayOutput
	ToVmArrayOutputWithContext(context.Context) VmArrayOutput
}

type VmArray []VmInput

func (VmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vm)(nil)).Elem()
}

func (i VmArray) ToVmArrayOutput() VmArrayOutput {
	return i.ToVmArrayOutputWithContext(context.Background())
}

func (i VmArray) ToVmArrayOutputWithContext(ctx context.Context) VmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmArrayOutput)
}

// VmMapInput is an input type that accepts VmMap and VmMapOutput values.
// You can construct a concrete instance of `VmMapInput` via:
//
//	VmMap{ "key": VmArgs{...} }
type VmMapInput interface {
	pulumi.Input

	ToVmMapOutput() VmMapOutput
	ToVmMapOutputWithContext(context.Context) VmMapOutput
}

type VmMap map[string]VmInput

func (VmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vm)(nil)).Elem()
}

func (i VmMap) ToVmMapOutput() VmMapOutput {
	return i.ToVmMapOutputWithContext(context.Background())
}

func (i VmMap) ToVmMapOutputWithContext(ctx context.Context) VmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmMapOutput)
}

type VmOutput struct{ *pulumi.OutputState }

func (VmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vm)(nil)).Elem()
}

func (o VmOutput) ToVmOutput() VmOutput {
	return o
}

func (o VmOutput) ToVmOutputWithContext(ctx context.Context) VmOutput {
	return o
}

// The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to
// be rescheduled.
func (o VmOutput) AffinityHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.AffinityHost }).(pulumi.StringPtrOutput)
}

// If the VM will automatically turn on. Defaults to `false`.
func (o VmOutput) AutoPoweron() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.AutoPoweron }).(pulumi.BoolPtrOutput)
}

// List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot,
// hard_shutdown, pause, shutdown, suspend, destroy. See:
// https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
func (o VmOutput) BlockedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.BlockedOperations }).(pulumi.StringArrayOutput)
}

// The ISO that should be attached to VM. This allows you to create a VM from a diskless template (any templates available
// from `xe template-list`) and install the OS from the following ISO.
func (o VmOutput) Cdrom() VmCdromPtrOutput {
	return o.ApplyT(func(v *Vm) VmCdromPtrOutput { return v.Cdrom }).(VmCdromPtrOutput)
}

// The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to
// perform a `full` clone, the VM template must not be a disk template.
func (o VmOutput) CloneType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.CloneType }).(pulumi.StringPtrOutput)
}

// The content of the cloud-init config to use. See the cloud init docs for more
// [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
func (o VmOutput) CloudConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.CloudConfig }).(pulumi.StringPtrOutput)
}

// The content of the cloud-init network configuration for the VM (uses [version
// 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
func (o VmOutput) CloudNetworkConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.CloudNetworkConfig }).(pulumi.StringPtrOutput)
}

func (o VmOutput) CoreOs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.CoreOs }).(pulumi.BoolPtrOutput)
}

func (o VmOutput) CpuCap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntPtrOutput { return v.CpuCap }).(pulumi.IntPtrOutput)
}

func (o VmOutput) CpuWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntPtrOutput { return v.CpuWeight }).(pulumi.IntPtrOutput)
}

// The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is
// greater than the max CPU value. This can be determined with the following command: ```$ xo-cli xo.getAllObjects
// filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs' { "max": 4, "number": 2 } # Updating the VM
// to use 3 CPUs would happen without stopping/starting the VM # Updating the VM to use 5 CPUs would stop/start the VM```
func (o VmOutput) Cpus() pulumi.IntOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntOutput { return v.Cpus }).(pulumi.IntOutput)
}

// Determines whether the cloud config VDI should be deleted once the VM has booted. Defaults to `false`. If set to `true`,
// powerState must be set to `Running`.
func (o VmOutput) DestroyCloudConfigVdiAfterBoot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.DestroyCloudConfigVdiAfterBoot }).(pulumi.BoolPtrOutput)
}

// The disk the VM will have access to.
func (o VmOutput) Disks() VmDiskArrayOutput {
	return o.ApplyT(func(v *Vm) VmDiskArrayOutput { return v.Disks }).(VmDiskArrayOutput)
}

// Boolean parameter that allows a VM to use nested virtualization.
func (o VmOutput) ExpNestedHvm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.ExpNestedHvm }).(pulumi.BoolPtrOutput)
}

// The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure.
// Defaults to empty string
func (o VmOutput) HighAvailability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.HighAvailability }).(pulumi.StringPtrOutput)
}

func (o VmOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

// The firmware to use for the VM. Possible values are `bios` and `uefi`.
func (o VmOutput) HvmBootFirmware() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.HvmBootFirmware }).(pulumi.StringPtrOutput)
}

// This cannot be used with `cdrom`. Possible values are `network` which allows a VM to boot via PXE.
func (o VmOutput) InstallationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.InstallationMethod }).(pulumi.StringPtrOutput)
}

func (o VmOutput) Ipv4Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.Ipv4Addresses }).(pulumi.StringArrayOutput)
}

// This is only accessible if guest-tools is installed in the VM and if `expectedIpCidr` is set on any network interfaces.
// This will contain a list of the ipv6 addresses across all network interfaces in order.
func (o VmOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

// The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new
// value is greater than the dynamic memory max. This can be determined with the following command: ```$ xo-cli
// xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic' [ 2147483648, #
// memory dynamic min 4294967296 # memory dynamic max (4GB) ] # Updating the VM to use 3GB of memory would happen without
// stopping/starting the VM # Updating the VM to use 5GB of memory would stop/start the VM```
func (o VmOutput) MemoryMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntOutput { return v.MemoryMax }).(pulumi.IntOutput)
}

// The description of the VM.
func (o VmOutput) NameDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.NameDescription }).(pulumi.StringPtrOutput)
}

// The name of the VM.
func (o VmOutput) NameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.NameLabel }).(pulumi.StringOutput)
}

// The network for the VM.
func (o VmOutput) Networks() VmNetworkArrayOutput {
	return o.ApplyT(func(v *Vm) VmNetworkArrayOutput { return v.Networks }).(VmNetworkArrayOutput)
}

// The power state of the VM. This can be Running, Halted, Paused or Suspended.
func (o VmOutput) PowerState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.PowerState }).(pulumi.StringPtrOutput)
}

func (o VmOutput) ResourceSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.ResourceSet }).(pulumi.StringPtrOutput)
}

// Number of seconds the VM should be delayed from starting.
func (o VmOutput) StartDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntPtrOutput { return v.StartDelay }).(pulumi.IntPtrOutput)
}

// The tags (labels) applied to the given entity.
func (o VmOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The ID of the VM template to create the new VM from.
func (o VmOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

// The video adapter the VM should use. Possible values include std and cirrus.
func (o VmOutput) Vga() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.Vga }).(pulumi.StringPtrOutput)
}

// The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
func (o VmOutput) Videoram() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntPtrOutput { return v.Videoram }).(pulumi.IntPtrOutput)
}

// The key value pairs to be populated in xenstore.
func (o VmOutput) Xenstore() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringMapOutput { return v.Xenstore }).(pulumi.StringMapOutput)
}

type VmArrayOutput struct{ *pulumi.OutputState }

func (VmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vm)(nil)).Elem()
}

func (o VmArrayOutput) ToVmArrayOutput() VmArrayOutput {
	return o
}

func (o VmArrayOutput) ToVmArrayOutputWithContext(ctx context.Context) VmArrayOutput {
	return o
}

func (o VmArrayOutput) Index(i pulumi.IntInput) VmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vm {
		return vs[0].([]*Vm)[vs[1].(int)]
	}).(VmOutput)
}

type VmMapOutput struct{ *pulumi.OutputState }

func (VmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vm)(nil)).Elem()
}

func (o VmMapOutput) ToVmMapOutput() VmMapOutput {
	return o
}

func (o VmMapOutput) ToVmMapOutputWithContext(ctx context.Context) VmMapOutput {
	return o
}

func (o VmMapOutput) MapIndex(k pulumi.StringInput) VmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vm {
		return vs[0].(map[string]*Vm)[vs[1].(string)]
	}).(VmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmInput)(nil)).Elem(), &Vm{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmArrayInput)(nil)).Elem(), VmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmMapInput)(nil)).Elem(), VmMap{})
	pulumi.RegisterOutputType(VmOutput{})
	pulumi.RegisterOutputType(VmArrayOutput{})
	pulumi.RegisterOutputType(VmMapOutput{})
}
