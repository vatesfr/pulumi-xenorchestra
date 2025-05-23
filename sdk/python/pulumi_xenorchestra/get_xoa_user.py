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
    'GetXoaUserResult',
    'AwaitableGetXoaUserResult',
    'get_xoa_user',
    'get_xoa_user_output',
]

@pulumi.output_type
class GetXoaUserResult:
    """
    A collection of values returned by getXoaUser.
    """
    def __init__(__self__, id=None, search_in_session=None, username=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if search_in_session and not isinstance(search_in_session, bool):
            raise TypeError("Expected argument 'search_in_session' to be a bool")
        pulumi.set(__self__, "search_in_session", search_in_session)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="searchInSession")
    def search_in_session(self) -> Optional[builtins.bool]:
        """
        A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
        """
        return pulumi.get(self, "search_in_session")

    @property
    @pulumi.getter
    def username(self) -> builtins.str:
        """
        The username of the XO user.
        """
        return pulumi.get(self, "username")


class AwaitableGetXoaUserResult(GetXoaUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetXoaUserResult(
            id=self.id,
            search_in_session=self.search_in_session,
            username=self.username)


def get_xoa_user(search_in_session: Optional[builtins.bool] = None,
                 username: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetXoaUserResult:
    """
    Provides information about a Xen Orchestra user. If the Xen Orchestra user account you are using is not an admin, see the `search_in_session` parameter.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_xenorchestra as xenorchestra

    user = xenorchestra.get_xoa_user(username="my-username")
    ```


    :param builtins.bool search_in_session: A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
    :param builtins.str username: The username of the XO user.
    """
    __args__ = dict()
    __args__['searchInSession'] = search_in_session
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('xenorchestra:index/getXoaUser:getXoaUser', __args__, opts=opts, typ=GetXoaUserResult).value

    return AwaitableGetXoaUserResult(
        id=pulumi.get(__ret__, 'id'),
        search_in_session=pulumi.get(__ret__, 'search_in_session'),
        username=pulumi.get(__ret__, 'username'))
def get_xoa_user_output(search_in_session: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                        username: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetXoaUserResult]:
    """
    Provides information about a Xen Orchestra user. If the Xen Orchestra user account you are using is not an admin, see the `search_in_session` parameter.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_xenorchestra as xenorchestra

    user = xenorchestra.get_xoa_user(username="my-username")
    ```


    :param builtins.bool search_in_session: A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
    :param builtins.str username: The username of the XO user.
    """
    __args__ = dict()
    __args__['searchInSession'] = search_in_session
    __args__['username'] = username
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('xenorchestra:index/getXoaUser:getXoaUser', __args__, opts=opts, typ=GetXoaUserResult)
    return __ret__.apply(lambda __response__: GetXoaUserResult(
        id=pulumi.get(__response__, 'id'),
        search_in_session=pulumi.get(__response__, 'search_in_session'),
        username=pulumi.get(__response__, 'username')))
