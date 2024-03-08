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
    'GetDashboardResult',
    'AwaitableGetDashboardResult',
    'get_dashboard',
    'get_dashboard_output',
]

@pulumi.output_type
class GetDashboardResult:
    """
    A collection of values returned by getDashboard.
    """
    def __init__(__self__, config_json=None, dashboard_id=None, folder=None, id=None, is_starred=None, org_id=None, slug=None, title=None, uid=None, url=None, version=None):
        if config_json and not isinstance(config_json, str):
            raise TypeError("Expected argument 'config_json' to be a str")
        pulumi.set(__self__, "config_json", config_json)
        if dashboard_id and not isinstance(dashboard_id, int):
            raise TypeError("Expected argument 'dashboard_id' to be a int")
        pulumi.set(__self__, "dashboard_id", dashboard_id)
        if folder and not isinstance(folder, int):
            raise TypeError("Expected argument 'folder' to be a int")
        pulumi.set(__self__, "folder", folder)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_starred and not isinstance(is_starred, bool):
            raise TypeError("Expected argument 'is_starred' to be a bool")
        pulumi.set(__self__, "is_starred", is_starred)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="configJson")
    def config_json(self) -> str:
        """
        The complete dashboard model JSON.
        """
        return pulumi.get(self, "config_json")

    @property
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> Optional[int]:
        """
        The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
        """
        return pulumi.get(self, "dashboard_id")

    @property
    @pulumi.getter
    def folder(self) -> int:
        """
        The numerical ID of the folder where the Grafana dashboard is found.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isStarred")
    def is_starred(self) -> bool:
        """
        Whether or not the Grafana dashboard is starred. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you’re interested in.
        """
        return pulumi.get(self, "is_starred")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def slug(self) -> str:
        """
        URL slug of the dashboard (deprecated).
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the Grafana dashboard.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        The uid of the Grafana dashboard. Specify either this or `dashboard_id`. Defaults to ``.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The full URL of the dashboard.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> int:
        """
        The numerical version of the Grafana dashboard.
        """
        return pulumi.get(self, "version")


class AwaitableGetDashboardResult(GetDashboardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDashboardResult(
            config_json=self.config_json,
            dashboard_id=self.dashboard_id,
            folder=self.folder,
            id=self.id,
            is_starred=self.is_starred,
            org_id=self.org_id,
            slug=self.slug,
            title=self.title,
            uid=self.uid,
            url=self.url,
            version=self.version)


def get_dashboard(dashboard_id: Optional[int] = None,
                  org_id: Optional[str] = None,
                  uid: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDashboardResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
    * [Folder/Dashboard Search HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_dashboard_search/)
    * [Dashboard HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)

    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.Dashboard("test", config_json=json.dumps({
        "id": 12345,
        "uid": "test-ds-dashboard-uid",
        "title": "Production Overview",
        "tags": ["templated"],
        "timezone": "browser",
        "schemaVersion": 16,
        "version": 0,
        "refresh": "25s",
    }))
    from_id = grafana.get_dashboard_output(dashboard_id=test.dashboard_id)
    from_uid = grafana.get_dashboard(uid="test-ds-dashboard-uid")
    ```


    :param int dashboard_id: The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
    :param str org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
    :param str uid: The uid of the Grafana dashboard. Specify either this or `dashboard_id`. Defaults to ``.
    """
    __args__ = dict()
    __args__['dashboardId'] = dashboard_id
    __args__['orgId'] = org_id
    __args__['uid'] = uid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getDashboard:getDashboard', __args__, opts=opts, typ=GetDashboardResult).value

    return AwaitableGetDashboardResult(
        config_json=pulumi.get(__ret__, 'config_json'),
        dashboard_id=pulumi.get(__ret__, 'dashboard_id'),
        folder=pulumi.get(__ret__, 'folder'),
        id=pulumi.get(__ret__, 'id'),
        is_starred=pulumi.get(__ret__, 'is_starred'),
        org_id=pulumi.get(__ret__, 'org_id'),
        slug=pulumi.get(__ret__, 'slug'),
        title=pulumi.get(__ret__, 'title'),
        uid=pulumi.get(__ret__, 'uid'),
        url=pulumi.get(__ret__, 'url'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_dashboard)
def get_dashboard_output(dashboard_id: Optional[pulumi.Input[Optional[int]]] = None,
                         org_id: Optional[pulumi.Input[Optional[str]]] = None,
                         uid: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDashboardResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
    * [Folder/Dashboard Search HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder_dashboard_search/)
    * [Dashboard HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)

    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test = grafana.Dashboard("test", config_json=json.dumps({
        "id": 12345,
        "uid": "test-ds-dashboard-uid",
        "title": "Production Overview",
        "tags": ["templated"],
        "timezone": "browser",
        "schemaVersion": 16,
        "version": 0,
        "refresh": "25s",
    }))
    from_id = grafana.get_dashboard_output(dashboard_id=test.dashboard_id)
    from_uid = grafana.get_dashboard(uid="test-ds-dashboard-uid")
    ```


    :param int dashboard_id: The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
    :param str org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
    :param str uid: The uid of the Grafana dashboard. Specify either this or `dashboard_id`. Defaults to ``.
    """
    ...
