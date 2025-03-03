# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'ResourceSetLimit',
    'VmCdrom',
    'VmDisk',
    'VmNetwork',
    'GetXoaHostsHostResult',
    'GetXoaVmsVmResult',
    'GetXoaVmsVmDiskResult',
    'GetXoaVmsVmNetworkResult',
]

@pulumi.output_type
class ResourceSetLimit(dict):
    def __init__(__self__, *,
                 quantity: int,
                 type: str):
        """
        :param int quantity: The numerical limit for the given type.
        :param str type: The type of resource set limit. Must be cpus, memory or disk.
        """
        pulumi.set(__self__, "quantity", quantity)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def quantity(self) -> int:
        """
        The numerical limit for the given type.
        """
        return pulumi.get(self, "quantity")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of resource set limit. Must be cpus, memory or disk.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class VmCdrom(dict):
    def __init__(__self__, *,
                 id: str):
        """
        :param str id: The ID of the ISO (VDI) to attach to the VM. This can be easily provided by using the `vdi` data source.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the ISO (VDI) to attach to the VM. This can be easily provided by using the `vdi` data source.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class VmDisk(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nameLabel":
            suggest = "name_label"
        elif key == "srId":
            suggest = "sr_id"
        elif key == "nameDescription":
            suggest = "name_description"
        elif key == "vbdId":
            suggest = "vbd_id"
        elif key == "vdiId":
            suggest = "vdi_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VmDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VmDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VmDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name_label: str,
                 size: int,
                 sr_id: str,
                 attached: Optional[bool] = None,
                 name_description: Optional[str] = None,
                 position: Optional[str] = None,
                 vbd_id: Optional[str] = None,
                 vdi_id: Optional[str] = None):
        """
        :param str name_label: The name for the disk
        :param int size: The size in bytes for the disk.
        :param str sr_id: The storage repository ID to use.
        :param bool attached: Whether the device should be attached to the VM.
        :param str name_description: The description for the disk
        :param str position: Indicates the order of the block device.
        """
        pulumi.set(__self__, "name_label", name_label)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "sr_id", sr_id)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if name_description is not None:
            pulumi.set(__self__, "name_description", name_description)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if vbd_id is not None:
            pulumi.set(__self__, "vbd_id", vbd_id)
        if vdi_id is not None:
            pulumi.set(__self__, "vdi_id", vdi_id)

    @property
    @pulumi.getter(name="nameLabel")
    def name_label(self) -> str:
        """
        The name for the disk
        """
        return pulumi.get(self, "name_label")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size in bytes for the disk.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="srId")
    def sr_id(self) -> str:
        """
        The storage repository ID to use.
        """
        return pulumi.get(self, "sr_id")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        """
        Whether the device should be attached to the VM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="nameDescription")
    def name_description(self) -> Optional[str]:
        """
        The description for the disk
        """
        return pulumi.get(self, "name_description")

    @property
    @pulumi.getter
    def position(self) -> Optional[str]:
        """
        Indicates the order of the block device.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="vbdId")
    def vbd_id(self) -> Optional[str]:
        return pulumi.get(self, "vbd_id")

    @property
    @pulumi.getter(name="vdiId")
    def vdi_id(self) -> Optional[str]:
        return pulumi.get(self, "vdi_id")


@pulumi.output_type
class VmNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkId":
            suggest = "network_id"
        elif key == "expectedIpCidr":
            suggest = "expected_ip_cidr"
        elif key == "ipv4Addresses":
            suggest = "ipv4_addresses"
        elif key == "ipv6Addresses":
            suggest = "ipv6_addresses"
        elif key == "macAddress":
            suggest = "mac_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VmNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VmNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VmNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_id: str,
                 attached: Optional[bool] = None,
                 device: Optional[str] = None,
                 expected_ip_cidr: Optional[str] = None,
                 ipv4_addresses: Optional[Sequence[str]] = None,
                 ipv6_addresses: Optional[Sequence[str]] = None,
                 mac_address: Optional[str] = None):
        """
        :param str network_id: The ID of the network the VM will be on.
        :param bool attached: Whether the device should be attached to the VM.
        """
        pulumi.set(__self__, "network_id", network_id)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if expected_ip_cidr is not None:
            pulumi.set(__self__, "expected_ip_cidr", expected_ip_cidr)
        if ipv4_addresses is not None:
            pulumi.set(__self__, "ipv4_addresses", ipv4_addresses)
        if ipv6_addresses is not None:
            pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        The ID of the network the VM will be on.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        """
        Whether the device should be attached to the VM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="expectedIpCidr")
    def expected_ip_cidr(self) -> Optional[str]:
        return pulumi.get(self, "expected_ip_cidr")

    @property
    @pulumi.getter(name="ipv4Addresses")
    def ipv4_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ipv4_addresses")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        return pulumi.get(self, "mac_address")


@pulumi.output_type
class GetXoaHostsHostResult(dict):
    def __init__(__self__, *,
                 cpus: Mapping[str, int],
                 id: str,
                 memory: int,
                 memory_usage: int,
                 name_label: str,
                 pool_id: str,
                 tags: Optional[Sequence[str]] = None):
        """
        :param Mapping[str, int] cpus: CPU information about the host. The 'cores' key will contain the number of cpu cores and the 'sockets' key will contain the number of sockets.
        :param int memory: The memory size of the host.
        :param int memory_usage: The memory usage of the host.
        :param str name_label: The name label of the host.
        :param str pool_id: Id of the pool that the host belongs to.
        :param Sequence[str] tags: The tags (labels) applied to the given entity.
        """
        pulumi.set(__self__, "cpus", cpus)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "memory_usage", memory_usage)
        pulumi.set(__self__, "name_label", name_label)
        pulumi.set(__self__, "pool_id", pool_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def cpus(self) -> Mapping[str, int]:
        """
        CPU information about the host. The 'cores' key will contain the number of cpu cores and the 'sockets' key will contain the number of sockets.
        """
        return pulumi.get(self, "cpus")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        The memory size of the host.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="memoryUsage")
    def memory_usage(self) -> int:
        """
        The memory usage of the host.
        """
        return pulumi.get(self, "memory_usage")

    @property
    @pulumi.getter(name="nameLabel")
    def name_label(self) -> str:
        """
        The name label of the host.
        """
        return pulumi.get(self, "name_label")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> str:
        """
        Id of the pool that the host belongs to.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The tags (labels) applied to the given entity.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class GetXoaVmsVmResult(dict):
    def __init__(__self__, *,
                 cpus: int,
                 disks: Sequence['outputs.GetXoaVmsVmDiskResult'],
                 id: str,
                 ipv4_addresses: Sequence[str],
                 ipv6_addresses: Sequence[str],
                 memory_max: int,
                 name_label: str,
                 networks: Sequence['outputs.GetXoaVmsVmNetworkResult'],
                 template: str,
                 affinity_host: Optional[str] = None,
                 auto_poweron: Optional[bool] = None,
                 blocked_operations: Optional[Sequence[str]] = None,
                 clone_type: Optional[str] = None,
                 cloud_config: Optional[str] = None,
                 cloud_network_config: Optional[str] = None,
                 core_os: Optional[bool] = None,
                 cpu_cap: Optional[int] = None,
                 cpu_weight: Optional[int] = None,
                 exp_nested_hvm: Optional[bool] = None,
                 high_availability: Optional[str] = None,
                 host: Optional[str] = None,
                 hvm_boot_firmware: Optional[str] = None,
                 name_description: Optional[str] = None,
                 power_state: Optional[str] = None,
                 resource_set: Optional[str] = None,
                 start_delay: Optional[int] = None,
                 tags: Optional[Sequence[str]] = None,
                 vga: Optional[str] = None,
                 videoram: Optional[int] = None,
                 xenstore: Optional[Mapping[str, str]] = None):
        """
        :param int cpus: The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is greater than the max CPU value. This can be determined with the following command:
               ```
               
               $ xo-cli xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs'
               {
                 "max": 4,
                 "number": 2
               }
               
               # Updating the VM to use 3 CPUs would happen without stopping/starting the VM
               # Updating the VM to use 5 CPUs would stop/start the VM
               ```
        :param Sequence['GetXoaVmsVmDiskArgs'] disks: The disk the VM will have access to.
        :param Sequence[str] ipv6_addresses: This is only accessible if guest-tools is installed in the VM and if `expected_ip_cidr` is set on any network interfaces. This will contain a list of the ipv6 addresses across all network interfaces in order.
        :param int memory_max: The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new value is greater than the dynamic memory max. This can be determined with the following command:
               ```
               
               
               $ xo-cli xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic'
               [
                 2147483648, # memory dynamic min
                 4294967296  # memory dynamic max (4GB)
               ]
               # Updating the VM to use 3GB of memory would happen without stopping/starting the VM
               # Updating the VM to use 5GB of memory would stop/start the VM
               ```
        :param str name_label: The name of the VM.
        :param Sequence['GetXoaVmsVmNetworkArgs'] networks: The network for the VM.
        :param str template: The ID of the VM template to create the new VM from.
        :param str affinity_host: The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to be rescheduled.
        :param bool auto_poweron: If the VM will automatically turn on. Defaults to `false`.
        :param Sequence[str] blocked_operations: List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot, hard_shutdown, pause, shutdown, suspend, destroy. See: https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
        :param str clone_type: The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to perform a `full` clone, the VM template must not be a disk template.
        :param str cloud_config: The content of the cloud-init config to use. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
        :param str cloud_network_config: The content of the cloud-init network configuration for the VM (uses [version 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
        :param bool exp_nested_hvm: Boolean parameter that allows a VM to use nested virtualization.
        :param str high_availability: The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure. Defaults to empty string
        :param str hvm_boot_firmware: The firmware to use for the VM. Possible values are `bios` and `uefi`.
        :param str name_description: The description of the VM.
        :param str power_state: The power state of the VM. This can be Running, Halted, Paused or Suspended.
        :param int start_delay: Number of seconds the VM should be delayed from starting.
        :param Sequence[str] tags: The tags (labels) applied to the given entity.
        :param str vga: The video adapter the VM should use. Possible values include std and cirrus.
        :param int videoram: The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
        :param Mapping[str, str] xenstore: The key value pairs to be populated in xenstore.
        """
        pulumi.set(__self__, "cpus", cpus)
        pulumi.set(__self__, "disks", disks)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ipv4_addresses", ipv4_addresses)
        pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        pulumi.set(__self__, "memory_max", memory_max)
        pulumi.set(__self__, "name_label", name_label)
        pulumi.set(__self__, "networks", networks)
        pulumi.set(__self__, "template", template)
        if affinity_host is not None:
            pulumi.set(__self__, "affinity_host", affinity_host)
        if auto_poweron is not None:
            pulumi.set(__self__, "auto_poweron", auto_poweron)
        if blocked_operations is not None:
            pulumi.set(__self__, "blocked_operations", blocked_operations)
        if clone_type is not None:
            pulumi.set(__self__, "clone_type", clone_type)
        if cloud_config is not None:
            pulumi.set(__self__, "cloud_config", cloud_config)
        if cloud_network_config is not None:
            pulumi.set(__self__, "cloud_network_config", cloud_network_config)
        if core_os is not None:
            pulumi.set(__self__, "core_os", core_os)
        if cpu_cap is not None:
            pulumi.set(__self__, "cpu_cap", cpu_cap)
        if cpu_weight is not None:
            pulumi.set(__self__, "cpu_weight", cpu_weight)
        if exp_nested_hvm is not None:
            pulumi.set(__self__, "exp_nested_hvm", exp_nested_hvm)
        if high_availability is not None:
            pulumi.set(__self__, "high_availability", high_availability)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if hvm_boot_firmware is not None:
            pulumi.set(__self__, "hvm_boot_firmware", hvm_boot_firmware)
        if name_description is not None:
            pulumi.set(__self__, "name_description", name_description)
        if power_state is not None:
            pulumi.set(__self__, "power_state", power_state)
        if resource_set is not None:
            pulumi.set(__self__, "resource_set", resource_set)
        if start_delay is not None:
            pulumi.set(__self__, "start_delay", start_delay)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vga is not None:
            pulumi.set(__self__, "vga", vga)
        if videoram is not None:
            pulumi.set(__self__, "videoram", videoram)
        if xenstore is not None:
            pulumi.set(__self__, "xenstore", xenstore)

    @property
    @pulumi.getter
    def cpus(self) -> int:
        """
        The number of CPUs the VM will have. Updates to this field will cause a stop and start of the VM if the new CPU value is greater than the max CPU value. This can be determined with the following command:
        ```

        $ xo-cli xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].CPUs'
        {
          "max": 4,
          "number": 2
        }

        # Updating the VM to use 3 CPUs would happen without stopping/starting the VM
        # Updating the VM to use 5 CPUs would stop/start the VM
        ```
        """
        return pulumi.get(self, "cpus")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.GetXoaVmsVmDiskResult']:
        """
        The disk the VM will have access to.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipv4Addresses")
    def ipv4_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ipv4_addresses")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Sequence[str]:
        """
        This is only accessible if guest-tools is installed in the VM and if `expected_ip_cidr` is set on any network interfaces. This will contain a list of the ipv6 addresses across all network interfaces in order.
        """
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="memoryMax")
    def memory_max(self) -> int:
        """
        The amount of memory in bytes the VM will have. Updates to this field will case a stop and start of the VM if the new value is greater than the dynamic memory max. This can be determined with the following command:
        ```


        $ xo-cli xo.getAllObjects filter='json:{"id": "cf7b5d7d-3cd5-6b7c-5025-5c935c8cd0b8"}' | jq '.[].memory.dynamic'
        [
          2147483648, # memory dynamic min
          4294967296  # memory dynamic max (4GB)
        ]
        # Updating the VM to use 3GB of memory would happen without stopping/starting the VM
        # Updating the VM to use 5GB of memory would stop/start the VM
        ```
        """
        return pulumi.get(self, "memory_max")

    @property
    @pulumi.getter(name="nameLabel")
    def name_label(self) -> str:
        """
        The name of the VM.
        """
        return pulumi.get(self, "name_label")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetXoaVmsVmNetworkResult']:
        """
        The network for the VM.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def template(self) -> str:
        """
        The ID of the VM template to create the new VM from.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="affinityHost")
    def affinity_host(self) -> Optional[str]:
        """
        The preferred host you would like the VM to run on. If changed on an existing VM it will require a reboot for the VM to be rescheduled.
        """
        return pulumi.get(self, "affinity_host")

    @property
    @pulumi.getter(name="autoPoweron")
    def auto_poweron(self) -> Optional[bool]:
        """
        If the VM will automatically turn on. Defaults to `false`.
        """
        return pulumi.get(self, "auto_poweron")

    @property
    @pulumi.getter(name="blockedOperations")
    def blocked_operations(self) -> Optional[Sequence[str]]:
        """
        List of operations on a VM that are not permitted. Examples include: clean_reboot, clean_shutdown, hard_reboot, hard_shutdown, pause, shutdown, suspend, destroy. See: https://xapi-project.github.io/xen-api/classes/vm.html#enum_vm_operations
        """
        return pulumi.get(self, "blocked_operations")

    @property
    @pulumi.getter(name="cloneType")
    def clone_type(self) -> Optional[str]:
        """
        The type of clone to perform for the VM. Possible values include `fast` or `full` and defaults to `fast`. In order to perform a `full` clone, the VM template must not be a disk template.
        """
        return pulumi.get(self, "clone_type")

    @property
    @pulumi.getter(name="cloudConfig")
    def cloud_config(self) -> Optional[str]:
        """
        The content of the cloud-init config to use. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
        """
        return pulumi.get(self, "cloud_config")

    @property
    @pulumi.getter(name="cloudNetworkConfig")
    def cloud_network_config(self) -> Optional[str]:
        """
        The content of the cloud-init network configuration for the VM (uses [version 1](https://cloudinit.readthedocs.io/en/latest/topics/network-config-format-v1.html))
        """
        return pulumi.get(self, "cloud_network_config")

    @property
    @pulumi.getter(name="coreOs")
    def core_os(self) -> Optional[bool]:
        return pulumi.get(self, "core_os")

    @property
    @pulumi.getter(name="cpuCap")
    def cpu_cap(self) -> Optional[int]:
        return pulumi.get(self, "cpu_cap")

    @property
    @pulumi.getter(name="cpuWeight")
    def cpu_weight(self) -> Optional[int]:
        return pulumi.get(self, "cpu_weight")

    @property
    @pulumi.getter(name="expNestedHvm")
    def exp_nested_hvm(self) -> Optional[bool]:
        """
        Boolean parameter that allows a VM to use nested virtualization.
        """
        return pulumi.get(self, "exp_nested_hvm")

    @property
    @pulumi.getter(name="highAvailability")
    def high_availability(self) -> Optional[str]:
        """
        The restart priority for the VM. Possible values are `best-effort`, `restart` and empty string (no restarts on failure. Defaults to empty string
        """
        return pulumi.get(self, "high_availability")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="hvmBootFirmware")
    def hvm_boot_firmware(self) -> Optional[str]:
        """
        The firmware to use for the VM. Possible values are `bios` and `uefi`.
        """
        return pulumi.get(self, "hvm_boot_firmware")

    @property
    @pulumi.getter(name="nameDescription")
    def name_description(self) -> Optional[str]:
        """
        The description of the VM.
        """
        return pulumi.get(self, "name_description")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> Optional[str]:
        """
        The power state of the VM. This can be Running, Halted, Paused or Suspended.
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter(name="resourceSet")
    def resource_set(self) -> Optional[str]:
        return pulumi.get(self, "resource_set")

    @property
    @pulumi.getter(name="startDelay")
    def start_delay(self) -> Optional[int]:
        """
        Number of seconds the VM should be delayed from starting.
        """
        return pulumi.get(self, "start_delay")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The tags (labels) applied to the given entity.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vga(self) -> Optional[str]:
        """
        The video adapter the VM should use. Possible values include std and cirrus.
        """
        return pulumi.get(self, "vga")

    @property
    @pulumi.getter
    def videoram(self) -> Optional[int]:
        """
        The videoram option the VM should use. Possible values include 1, 2, 4, 8, 16
        """
        return pulumi.get(self, "videoram")

    @property
    @pulumi.getter
    def xenstore(self) -> Optional[Mapping[str, str]]:
        """
        The key value pairs to be populated in xenstore.
        """
        return pulumi.get(self, "xenstore")


@pulumi.output_type
class GetXoaVmsVmDiskResult(dict):
    def __init__(__self__, *,
                 name_label: str,
                 position: str,
                 size: int,
                 sr_id: str,
                 vbd_id: str,
                 vdi_id: str,
                 attached: Optional[bool] = None,
                 name_description: Optional[str] = None):
        """
        :param str name_label: The name for the disk
        :param str position: Indicates the order of the block device.
        :param int size: The size in bytes for the disk.
        :param str sr_id: The storage repository ID to use.
        :param bool attached: Whether the device should be attached to the VM.
        :param str name_description: The description for the disk
        """
        pulumi.set(__self__, "name_label", name_label)
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "sr_id", sr_id)
        pulumi.set(__self__, "vbd_id", vbd_id)
        pulumi.set(__self__, "vdi_id", vdi_id)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if name_description is not None:
            pulumi.set(__self__, "name_description", name_description)

    @property
    @pulumi.getter(name="nameLabel")
    def name_label(self) -> str:
        """
        The name for the disk
        """
        return pulumi.get(self, "name_label")

    @property
    @pulumi.getter
    def position(self) -> str:
        """
        Indicates the order of the block device.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size in bytes for the disk.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="srId")
    def sr_id(self) -> str:
        """
        The storage repository ID to use.
        """
        return pulumi.get(self, "sr_id")

    @property
    @pulumi.getter(name="vbdId")
    def vbd_id(self) -> str:
        return pulumi.get(self, "vbd_id")

    @property
    @pulumi.getter(name="vdiId")
    def vdi_id(self) -> str:
        return pulumi.get(self, "vdi_id")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        """
        Whether the device should be attached to the VM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="nameDescription")
    def name_description(self) -> Optional[str]:
        """
        The description for the disk
        """
        return pulumi.get(self, "name_description")


@pulumi.output_type
class GetXoaVmsVmNetworkResult(dict):
    def __init__(__self__, *,
                 device: str,
                 ipv4_addresses: Sequence[str],
                 ipv6_addresses: Sequence[str],
                 mac_address: str,
                 network_id: str,
                 attached: Optional[bool] = None,
                 expected_ip_cidr: Optional[str] = None):
        """
        :param str network_id: The ID of the network the VM will be on.
        :param bool attached: Whether the device should be attached to the VM.
        """
        pulumi.set(__self__, "device", device)
        pulumi.set(__self__, "ipv4_addresses", ipv4_addresses)
        pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        pulumi.set(__self__, "mac_address", mac_address)
        pulumi.set(__self__, "network_id", network_id)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if expected_ip_cidr is not None:
            pulumi.set(__self__, "expected_ip_cidr", expected_ip_cidr)

    @property
    @pulumi.getter
    def device(self) -> str:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="ipv4Addresses")
    def ipv4_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ipv4_addresses")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        The ID of the network the VM will be on.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        """
        Whether the device should be attached to the VM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="expectedIpCidr")
    def expected_ip_cidr(self) -> Optional[str]:
        return pulumi.get(self, "expected_ip_cidr")


