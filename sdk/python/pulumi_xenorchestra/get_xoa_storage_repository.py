# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = [
    'GetXoaStorageRepositoryResult',
    'AwaitableGetXoaStorageRepositoryResult',
    'get_xoa_storage_repository',
    'get_xoa_storage_repository_output',
]

@pulumi.output_type
class GetXoaStorageRepositoryResult:
    """
    A collection of values returned by getXoaStorageRepository.
    """
    def __init__(__self__, container=None, id=None, name_label=None, physical_usage=None, pool_id=None, size=None, sr_type=None, tags=None, usage=None, uuid=None):
        if container and not isinstance(container, str):
            raise TypeError("Expected argument 'container' to be a str")
        pulumi.set(__self__, "container", container)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_label and not isinstance(name_label, str):
            raise TypeError("Expected argument 'name_label' to be a str")
        pulumi.set(__self__, "name_label", name_label)
        if physical_usage and not isinstance(physical_usage, float):
            raise TypeError("Expected argument 'physical_usage' to be a float")
        pulumi.set(__self__, "physical_usage", physical_usage)
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        pulumi.set(__self__, "pool_id", pool_id)
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        pulumi.set(__self__, "size", size)
        if sr_type and not isinstance(sr_type, str):
            raise TypeError("Expected argument 'sr_type' to be a str")
        pulumi.set(__self__, "sr_type", sr_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if usage and not isinstance(usage, float):
            raise TypeError("Expected argument 'usage' to be a float")
        pulumi.set(__self__, "usage", usage)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def container(self) -> builtins.str:
        """
        The storage container.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameLabel")
    def name_label(self) -> builtins.str:
        """
        The name of the storage repository to look up
        """
        return pulumi.get(self, "name_label")

    @property
    @pulumi.getter(name="physicalUsage")
    def physical_usage(self) -> builtins.float:
        """
        The physical storage size.
        """
        return pulumi.get(self, "physical_usage")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[builtins.str]:
        """
        The Id of the pool the storage repository exists on.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def size(self) -> builtins.float:
        """
        The storage size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="srType")
    def sr_type(self) -> builtins.str:
        """
        The type of storage repository (lvm, udev, iso, user, etc).
        """
        return pulumi.get(self, "sr_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[builtins.str]]:
        """
        The tags (labels) applied to the given entity.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def usage(self) -> builtins.float:
        """
        The current usage for this storage repository.
        """
        return pulumi.get(self, "usage")

    @property
    @pulumi.getter
    def uuid(self) -> builtins.str:
        """
        uuid of the storage repository. This is equivalent to the id.
        """
        return pulumi.get(self, "uuid")


class AwaitableGetXoaStorageRepositoryResult(GetXoaStorageRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetXoaStorageRepositoryResult(
            container=self.container,
            id=self.id,
            name_label=self.name_label,
            physical_usage=self.physical_usage,
            pool_id=self.pool_id,
            size=self.size,
            sr_type=self.sr_type,
            tags=self.tags,
            usage=self.usage,
            uuid=self.uuid)


def get_xoa_storage_repository(name_label: Optional[builtins.str] = None,
                               pool_id: Optional[builtins.str] = None,
                               tags: Optional[Sequence[builtins.str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetXoaStorageRepositoryResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_xenorchestra as xenorchestra

    local_storage = xenorchestra.get_xoa_storage_repository(name_label="Your storage repository label")
    demo_vm = xenorchestra.Vm("demo-vm", disks=[{
        "sr_id": local_storage.id,
        "name_label": "Ubuntu Bionic Beaver 18.04_imavo",
        "size": 32212254720,
    }])
    ```


    :param builtins.str name_label: The name of the storage repository to look up
    :param builtins.str pool_id: The Id of the pool the storage repository exists on.
    :param Sequence[builtins.str] tags: The tags (labels) applied to the given entity.
    """
    __args__ = dict()
    __args__['nameLabel'] = name_label
    __args__['poolId'] = pool_id
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('xenorchestra:index/getXoaStorageRepository:getXoaStorageRepository', __args__, opts=opts, typ=GetXoaStorageRepositoryResult).value

    return AwaitableGetXoaStorageRepositoryResult(
        container=pulumi.get(__ret__, 'container'),
        id=pulumi.get(__ret__, 'id'),
        name_label=pulumi.get(__ret__, 'name_label'),
        physical_usage=pulumi.get(__ret__, 'physical_usage'),
        pool_id=pulumi.get(__ret__, 'pool_id'),
        size=pulumi.get(__ret__, 'size'),
        sr_type=pulumi.get(__ret__, 'sr_type'),
        tags=pulumi.get(__ret__, 'tags'),
        usage=pulumi.get(__ret__, 'usage'),
        uuid=pulumi.get(__ret__, 'uuid'))
def get_xoa_storage_repository_output(name_label: Optional[pulumi.Input[builtins.str]] = None,
                                      pool_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                      tags: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetXoaStorageRepositoryResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_xenorchestra as xenorchestra

    local_storage = xenorchestra.get_xoa_storage_repository(name_label="Your storage repository label")
    demo_vm = xenorchestra.Vm("demo-vm", disks=[{
        "sr_id": local_storage.id,
        "name_label": "Ubuntu Bionic Beaver 18.04_imavo",
        "size": 32212254720,
    }])
    ```


    :param builtins.str name_label: The name of the storage repository to look up
    :param builtins.str pool_id: The Id of the pool the storage repository exists on.
    :param Sequence[builtins.str] tags: The tags (labels) applied to the given entity.
    """
    __args__ = dict()
    __args__['nameLabel'] = name_label
    __args__['poolId'] = pool_id
    __args__['tags'] = tags
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('xenorchestra:index/getXoaStorageRepository:getXoaStorageRepository', __args__, opts=opts, typ=GetXoaStorageRepositoryResult)
    return __ret__.apply(lambda __response__: GetXoaStorageRepositoryResult(
        container=pulumi.get(__response__, 'container'),
        id=pulumi.get(__response__, 'id'),
        name_label=pulumi.get(__response__, 'name_label'),
        physical_usage=pulumi.get(__response__, 'physical_usage'),
        pool_id=pulumi.get(__response__, 'pool_id'),
        size=pulumi.get(__response__, 'size'),
        sr_type=pulumi.get(__response__, 'sr_type'),
        tags=pulumi.get(__response__, 'tags'),
        usage=pulumi.get(__response__, 'usage'),
        uuid=pulumi.get(__response__, 'uuid')))
