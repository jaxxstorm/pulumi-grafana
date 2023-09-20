# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
    'get_organization_output',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, admins=None, editors=None, id=None, name=None, viewers=None):
        if admins and not isinstance(admins, list):
            raise TypeError("Expected argument 'admins' to be a list")
        pulumi.set(__self__, "admins", admins)
        if editors and not isinstance(editors, list):
            raise TypeError("Expected argument 'editors' to be a list")
        pulumi.set(__self__, "editors", editors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if viewers and not isinstance(viewers, list):
            raise TypeError("Expected argument 'viewers' to be a list")
        pulumi.set(__self__, "viewers", viewers)

    @property
    @pulumi.getter
    def admins(self) -> Sequence[str]:
        """
        A list of email addresses corresponding to users given admin access to the organization.
        """
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter
    def editors(self) -> Sequence[str]:
        """
        A list of email addresses corresponding to users given editor access to the organization.
        """
        return pulumi.get(self, "editors")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def viewers(self) -> Sequence[str]:
        """
        A list of email addresses corresponding to users given viewer access to the organization.
        """
        return pulumi.get(self, "viewers")


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            admins=self.admins,
            editors=self.editors,
            id=self.id,
            name=self.name,
            viewers=self.viewers)


def get_organization(name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_grafana as grafana
    import pulumi_grafana as grafana

    test = grafana.Organization("test",
        admin_user="admin",
        create_users=True,
        viewers=[
            "viewer-01@example.com",
            "viewer-02@example.com",
        ])
    from_name = grafana.get_organization_output(name=test.name)
    ```


    :param str name: The name of the Organization.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        admins=pulumi.get(__ret__, 'admins'),
        editors=pulumi.get(__ret__, 'editors'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        viewers=pulumi.get(__ret__, 'viewers'))


@_utilities.lift_output_func(get_organization)
def get_organization_output(name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)

    ## Example Usage

    ```python
    import pulumi
    import lbrlabs_pulumi_grafana as grafana
    import pulumi_grafana as grafana

    test = grafana.Organization("test",
        admin_user="admin",
        create_users=True,
        viewers=[
            "viewer-01@example.com",
            "viewer-02@example.com",
        ])
    from_name = grafana.get_organization_output(name=test.name)
    ```


    :param str name: The name of the Organization.
    """
    ...
