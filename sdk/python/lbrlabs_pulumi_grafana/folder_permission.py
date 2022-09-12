# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FolderPermissionArgs', 'FolderPermission']

@pulumi.input_type
class FolderPermissionArgs:
    def __init__(__self__, *,
                 folder_uid: pulumi.Input[str],
                 permissions: pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]):
        """
        The set of arguments for constructing a FolderPermission resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        pulumi.set(__self__, "folder_uid", folder_uid)
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Input[str]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder_uid", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]):
        pulumi.set(self, "permissions", value)


@pulumi.input_type
class _FolderPermissionState:
    def __init__(__self__, *,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]] = None):
        """
        Input properties used for looking up and filtering FolderPermission resources.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        if folder_uid is not None:
            pulumi.set(__self__, "folder_uid", folder_uid)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

    @folder_uid.setter
    def folder_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_uid", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


class FolderPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionPermissionArgs']]]]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/folder_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_grafana as grafana

        team = grafana.Team("team")
        user = grafana.User("user", email="user.name@example.com")
        collection = grafana.Folder("collection", title="Folder Title")
        collection_permission = grafana.FolderPermission("collectionPermission",
            folder_uid=collection.uid,
            permissions=[
                grafana.FolderPermissionPermissionArgs(
                    role="Editor",
                    permission="Edit",
                ),
                grafana.FolderPermissionPermissionArgs(
                    team_id=team.id,
                    permission="View",
                ),
                grafana.FolderPermissionPermissionArgs(
                    user_id=user.id,
                    permission="Admin",
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/folder_permissions/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_grafana as grafana

        team = grafana.Team("team")
        user = grafana.User("user", email="user.name@example.com")
        collection = grafana.Folder("collection", title="Folder Title")
        collection_permission = grafana.FolderPermission("collectionPermission",
            folder_uid=collection.uid,
            permissions=[
                grafana.FolderPermissionPermissionArgs(
                    role="Editor",
                    permission="Edit",
                ),
                grafana.FolderPermissionPermissionArgs(
                    team_id=team.id,
                    permission="View",
                ),
                grafana.FolderPermissionPermissionArgs(
                    user_id=user.id,
                    permission="Admin",
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param FolderPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_uid: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionPermissionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderPermissionArgs.__new__(FolderPermissionArgs)

            if folder_uid is None and not opts.urn:
                raise TypeError("Missing required property 'folder_uid'")
            __props__.__dict__["folder_uid"] = folder_uid
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
        super(FolderPermission, __self__).__init__(
            'grafana:index/folderPermission:FolderPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            folder_uid: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionPermissionArgs']]]]] = None) -> 'FolderPermission':
        """
        Get an existing FolderPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_uid: The UID of the folder.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionPermissionArgs']]]] permissions: The permission items to add/update. Items that are omitted from the list will be removed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderPermissionState.__new__(_FolderPermissionState)

        __props__.__dict__["folder_uid"] = folder_uid
        __props__.__dict__["permissions"] = permissions
        return FolderPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="folderUid")
    def folder_uid(self) -> pulumi.Output[str]:
        """
        The UID of the folder.
        """
        return pulumi.get(self, "folder_uid")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence['outputs.FolderPermissionPermission']]:
        """
        The permission items to add/update. Items that are omitted from the list will be removed.
        """
        return pulumi.get(self, "permissions")

