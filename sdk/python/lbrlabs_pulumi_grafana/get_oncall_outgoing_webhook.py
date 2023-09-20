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
    'GetOncallOutgoingWebhookResult',
    'AwaitableGetOncallOutgoingWebhookResult',
    'get_oncall_outgoing_webhook',
    'get_oncall_outgoing_webhook_output',
]

@pulumi.output_type
class GetOncallOutgoingWebhookResult:
    """
    A collection of values returned by getOncallOutgoingWebhook.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

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
        The outgoing webhook name.
        """
        return pulumi.get(self, "name")


class AwaitableGetOncallOutgoingWebhookResult(GetOncallOutgoingWebhookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOncallOutgoingWebhookResult(
            id=self.id,
            name=self.name)


def get_oncall_outgoing_webhook(name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOncallOutgoingWebhookResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_outgoing_webhook = grafana.get_oncall_outgoing_webhook(name="example_outgoing_webhook")
    ```


    :param str name: The outgoing webhook name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOncallOutgoingWebhook:getOncallOutgoingWebhook', __args__, opts=opts, typ=GetOncallOutgoingWebhookResult).value

    return AwaitableGetOncallOutgoingWebhookResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_oncall_outgoing_webhook)
def get_oncall_outgoing_webhook_output(name: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOncallOutgoingWebhookResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_outgoing_webhook = grafana.get_oncall_outgoing_webhook(name="example_outgoing_webhook")
    ```


    :param str name: The outgoing webhook name.
    """
    ...
