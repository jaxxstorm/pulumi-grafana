# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOnCallSlackChannelResult',
    'AwaitableGetOnCallSlackChannelResult',
    'get_on_call_slack_channel',
    'get_on_call_slack_channel_output',
]

@pulumi.output_type
class GetOnCallSlackChannelResult:
    """
    A collection of values returned by getOnCallSlackChannel.
    """
    def __init__(__self__, id=None, name=None, slack_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slack_id and not isinstance(slack_id, str):
            raise TypeError("Expected argument 'slack_id' to be a str")
        pulumi.set(__self__, "slack_id", slack_id)

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
        The Slack channel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="slackId")
    def slack_id(self) -> str:
        """
        The Slack ID of the channel.
        """
        return pulumi.get(self, "slack_id")


class AwaitableGetOnCallSlackChannelResult(GetOnCallSlackChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOnCallSlackChannelResult(
            id=self.id,
            name=self.name,
            slack_id=self.slack_id)


def get_on_call_slack_channel(name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOnCallSlackChannelResult:
    """
    * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/slack_channels/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_slack_channel = grafana.get_on_call_slack_channel(name="example_slack_channel")
    ```


    :param str name: The Slack channel name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOnCallSlackChannel:getOnCallSlackChannel', __args__, opts=opts, typ=GetOnCallSlackChannelResult).value

    return AwaitableGetOnCallSlackChannelResult(
        id=__ret__.id,
        name=__ret__.name,
        slack_id=__ret__.slack_id)


@_utilities.lift_output_func(get_on_call_slack_channel)
def get_on_call_slack_channel_output(name: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOnCallSlackChannelResult]:
    """
    * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/slack_channels/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_slack_channel = grafana.get_on_call_slack_channel(name="example_slack_channel")
    ```


    :param str name: The Slack channel name.
    """
    ...
