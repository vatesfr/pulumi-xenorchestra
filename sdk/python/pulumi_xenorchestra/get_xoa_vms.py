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
from . import outputs

__all__ = [
    'GetXoaVmsResult',
    'AwaitableGetXoaVmsResult',
    'get_xoa_vms',
    'get_xoa_vms_output',
]

@pulumi.output_type
class GetXoaVmsResult:
    """
    A collection of values returned by getXoaVms.
    """
    def __init__(__self__, host=None, id=None, pool_id=None, power_state=None, vms=None):
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        pulumi.set(__self__, "pool_id", pool_id)
        if power_state and not isinstance(power_state, str):
            raise TypeError("Expected argument 'power_state' to be a str")
        pulumi.set(__self__, "power_state", power_state)
        if vms and not isinstance(vms, list):
            raise TypeError("Expected argument 'vms' to be a list")
        pulumi.set(__self__, "vms", vms)

    @property
    @pulumi.getter
    def host(self) -> Optional[builtins.str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> builtins.str:
        """
        The ID of the pool the VM belongs to.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> Optional[builtins.str]:
        """
        The power state of the vms. (Running, Halted)
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter
    def vms(self) -> Sequence['outputs.GetXoaVmsVmResult']:
        """
        A list of information for all vms found in this pool.
        """
        return pulumi.get(self, "vms")


class AwaitableGetXoaVmsResult(GetXoaVmsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetXoaVmsResult(
            host=self.host,
            id=self.id,
            pool_id=self.pool_id,
            power_state=self.power_state,
            vms=self.vms)


def get_xoa_vms(host: Optional[builtins.str] = None,
                pool_id: Optional[builtins.str] = None,
                power_state: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetXoaVmsResult:
    """
    Use this data source to filter Xenorchestra VMs by certain criteria (pool_id, power_state or host) for use in other resources.


    :param builtins.str pool_id: The ID of the pool the VM belongs to.
    :param builtins.str power_state: The power state of the vms. (Running, Halted)
    """
    __args__ = dict()
    __args__['host'] = host
    __args__['poolId'] = pool_id
    __args__['powerState'] = power_state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('xenorchestra:index/getXoaVms:getXoaVms', __args__, opts=opts, typ=GetXoaVmsResult).value

    return AwaitableGetXoaVmsResult(
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        pool_id=pulumi.get(__ret__, 'pool_id'),
        power_state=pulumi.get(__ret__, 'power_state'),
        vms=pulumi.get(__ret__, 'vms'))
def get_xoa_vms_output(host: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       pool_id: Optional[pulumi.Input[builtins.str]] = None,
                       power_state: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetXoaVmsResult]:
    """
    Use this data source to filter Xenorchestra VMs by certain criteria (pool_id, power_state or host) for use in other resources.


    :param builtins.str pool_id: The ID of the pool the VM belongs to.
    :param builtins.str power_state: The power state of the vms. (Running, Halted)
    """
    __args__ = dict()
    __args__['host'] = host
    __args__['poolId'] = pool_id
    __args__['powerState'] = power_state
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('xenorchestra:index/getXoaVms:getXoaVms', __args__, opts=opts, typ=GetXoaVmsResult)
    return __ret__.apply(lambda __response__: GetXoaVmsResult(
        host=pulumi.get(__response__, 'host'),
        id=pulumi.get(__response__, 'id'),
        pool_id=pulumi.get(__response__, 'pool_id'),
        power_state=pulumi.get(__response__, 'power_state'),
        vms=pulumi.get(__response__, 'vms')))
