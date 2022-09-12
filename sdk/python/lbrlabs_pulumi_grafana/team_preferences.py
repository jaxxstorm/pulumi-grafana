# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TeamPreferencesArgs', 'TeamPreferences']

@pulumi.input_type
class TeamPreferencesArgs:
    def __init__(__self__, *,
                 team_id: pulumi.Input[int],
                 home_dashboard_id: Optional[pulumi.Input[int]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TeamPreferences resource.
        :param pulumi.Input[int] team_id: The numeric team ID.
        :param pulumi.Input[int] home_dashboard_id: The numeric ID of the dashboard to display when a team member logs in.
        :param pulumi.Input[str] theme: The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        :param pulumi.Input[str] timezone: The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        pulumi.set(__self__, "team_id", team_id)
        if home_dashboard_id is not None:
            pulumi.set(__self__, "home_dashboard_id", home_dashboard_id)
        if theme is not None:
            pulumi.set(__self__, "theme", theme)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Input[int]:
        """
        The numeric team ID.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="homeDashboardId")
    def home_dashboard_id(self) -> Optional[pulumi.Input[int]]:
        """
        The numeric ID of the dashboard to display when a team member logs in.
        """
        return pulumi.get(self, "home_dashboard_id")

    @home_dashboard_id.setter
    def home_dashboard_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "home_dashboard_id", value)

    @property
    @pulumi.getter
    def theme(self) -> Optional[pulumi.Input[str]]:
        """
        The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        """
        return pulumi.get(self, "theme")

    @theme.setter
    def theme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


@pulumi.input_type
class _TeamPreferencesState:
    def __init__(__self__, *,
                 home_dashboard_id: Optional[pulumi.Input[int]] = None,
                 team_id: Optional[pulumi.Input[int]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TeamPreferences resources.
        :param pulumi.Input[int] home_dashboard_id: The numeric ID of the dashboard to display when a team member logs in.
        :param pulumi.Input[int] team_id: The numeric team ID.
        :param pulumi.Input[str] theme: The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        :param pulumi.Input[str] timezone: The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        if home_dashboard_id is not None:
            pulumi.set(__self__, "home_dashboard_id", home_dashboard_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if theme is not None:
            pulumi.set(__self__, "theme", theme)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter(name="homeDashboardId")
    def home_dashboard_id(self) -> Optional[pulumi.Input[int]]:
        """
        The numeric ID of the dashboard to display when a team member logs in.
        """
        return pulumi.get(self, "home_dashboard_id")

    @home_dashboard_id.setter
    def home_dashboard_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "home_dashboard_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[int]]:
        """
        The numeric team ID.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def theme(self) -> Optional[pulumi.Input[str]]:
        """
        The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        """
        return pulumi.get(self, "theme")

    @theme.setter
    def theme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


class TeamPreferences(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 home_dashboard_id: Optional[pulumi.Input[int]] = None,
                 team_id: Optional[pulumi.Input[int]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/preferences/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/team/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_grafana as grafana

        metrics = grafana.Dashboard("metrics", config_json=(lambda path: open(path).read())("grafana-dashboard.json"))
        team = grafana.Team("team")
        team_preferences = grafana.TeamPreferences("teamPreferences",
            team_id=team.id,
            theme="dark",
            timezone="browser",
            home_dashboard_id=metrics.dashboard_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] home_dashboard_id: The numeric ID of the dashboard to display when a team member logs in.
        :param pulumi.Input[int] team_id: The numeric team ID.
        :param pulumi.Input[str] theme: The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        :param pulumi.Input[str] timezone: The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamPreferencesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/administration/preferences/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/team/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_grafana as grafana

        metrics = grafana.Dashboard("metrics", config_json=(lambda path: open(path).read())("grafana-dashboard.json"))
        team = grafana.Team("team")
        team_preferences = grafana.TeamPreferences("teamPreferences",
            team_id=team.id,
            theme="dark",
            timezone="browser",
            home_dashboard_id=metrics.dashboard_id)
        ```

        :param str resource_name: The name of the resource.
        :param TeamPreferencesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamPreferencesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 home_dashboard_id: Optional[pulumi.Input[int]] = None,
                 team_id: Optional[pulumi.Input[int]] = None,
                 theme: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamPreferencesArgs.__new__(TeamPreferencesArgs)

            __props__.__dict__["home_dashboard_id"] = home_dashboard_id
            if team_id is None and not opts.urn:
                raise TypeError("Missing required property 'team_id'")
            __props__.__dict__["team_id"] = team_id
            __props__.__dict__["theme"] = theme
            __props__.__dict__["timezone"] = timezone
        super(TeamPreferences, __self__).__init__(
            'grafana:index/teamPreferences:TeamPreferences',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            home_dashboard_id: Optional[pulumi.Input[int]] = None,
            team_id: Optional[pulumi.Input[int]] = None,
            theme: Optional[pulumi.Input[str]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'TeamPreferences':
        """
        Get an existing TeamPreferences resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] home_dashboard_id: The numeric ID of the dashboard to display when a team member logs in.
        :param pulumi.Input[int] team_id: The numeric team ID.
        :param pulumi.Input[str] theme: The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        :param pulumi.Input[str] timezone: The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamPreferencesState.__new__(_TeamPreferencesState)

        __props__.__dict__["home_dashboard_id"] = home_dashboard_id
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["theme"] = theme
        __props__.__dict__["timezone"] = timezone
        return TeamPreferences(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="homeDashboardId")
    def home_dashboard_id(self) -> pulumi.Output[Optional[int]]:
        """
        The numeric ID of the dashboard to display when a team member logs in.
        """
        return pulumi.get(self, "home_dashboard_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[int]:
        """
        The numeric team ID.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def theme(self) -> pulumi.Output[Optional[str]]:
        """
        The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
        """
        return pulumi.get(self, "theme")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.
        """
        return pulumi.get(self, "timezone")

