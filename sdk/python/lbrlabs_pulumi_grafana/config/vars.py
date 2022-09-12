# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('grafana')


class _ExportableConfig(types.ModuleType):
    @property
    def auth(self) -> Optional[str]:
        """
        API token or basic auth `username:password`. May alternatively be set via the `GRAFANA_AUTH` environment variable.
        """
        return __config__.get('auth') or _utilities.get_env('GRAFANA_AUTH')

    @property
    def ca_cert(self) -> Optional[str]:
        """
        Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the
        `GRAFANA_CA_CERT` environment variable.
        """
        return __config__.get('caCert') or _utilities.get_env('GRAFANA_CA_CERT')

    @property
    def cloud_api_key(self) -> Optional[str]:
        """
        API key for Grafana Cloud. May alternatively be set via the `GRAFANA_CLOUD_API_KEY` environment variable.
        """
        return __config__.get('cloudApiKey') or _utilities.get_env('GRAFANA_CLOUD_API_KEY')

    @property
    def cloud_api_url(self) -> Optional[str]:
        """
        Grafana Cloud's API URL. May alternatively be set via the `GRAFANA_CLOUD_API_URL` environment variable.
        """
        return __config__.get('cloudApiUrl') or _utilities.get_env('GRAFANA_CLOUD_API_URL')

    @property
    def http_headers(self) -> Optional[str]:
        """
        Optional. HTTP headers mapping keys to values used for accessing the Grafana API. May alternatively be set via the
        `GRAFANA_HTTP_HEADERS` environment variable in JSON format.
        """
        return __config__.get('httpHeaders')

    @property
    def insecure_skip_verify(self) -> Optional[bool]:
        """
        Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
        """
        return __config__.get_bool('insecureSkipVerify') or _utilities.get_env_bool('GRAFANA_INSECURE_SKIP_VERIFY')

    @property
    def oncall_access_token(self) -> Optional[str]:
        """
        A Grafana OnCall access token. May alternatively be set via the `GRAFANA_ONCALL_ACCESS_TOKEN` environment variable.
        """
        return __config__.get('oncallAccessToken') or _utilities.get_env('GRAFANA_ONCALL_ACCESS_TOKEN')

    @property
    def oncall_url(self) -> Optional[str]:
        """
        An Grafana OnCall backend address. May alternatively be set via the `GRAFANA_ONCALL_URL` environment variable.
        """
        return __config__.get('oncallUrl') or _utilities.get_env('GRAFANA_ONCALL_URL')

    @property
    def org_id(self) -> Optional[int]:
        """
        The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment
        variable.
        """
        return __config__.get_int('orgId') or _utilities.get_env_int('GRAFANA_ORG_ID')

    @property
    def retries(self) -> Optional[int]:
        """
        The amount of retries to use for Grafana API calls. May alternatively be set via the `GRAFANA_RETRIES` environment
        variable.
        """
        return __config__.get_int('retries') or _utilities.get_env_int('GRAFANA_RETRIES')

    @property
    def sm_access_token(self) -> Optional[str]:
        """
        A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
        """
        return __config__.get('smAccessToken') or _utilities.get_env('GRAFANA_SM_ACCESS_TOKEN')

    @property
    def sm_url(self) -> Optional[str]:
        """
        Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable. The
        correct value for each service region is cited in the [Synthetic Monitoring
        documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/#probe-api-server-url). Note
        the `sm_url` value is optional, but it must correspond with the value specified as the `region_slug` in the
        `grafana_cloud_stack` resource. Also note that when a Terraform configuration contains multiple provider instances
        managing SM resources associated with the same Grafana stack, specifying an explicit `sm_url` set to the same value for
        each provider ensures all providers interact with the same SM API.
        """
        return __config__.get('smUrl') or _utilities.get_env('GRAFANA_SM_URL')

    @property
    def store_dashboard_sha256(self) -> Optional[bool]:
        """
        Set to true if you want to save only the sha256sum instead of complete dashboard model JSON in the tfstate.
        """
        return __config__.get_bool('storeDashboardSha256') or _utilities.get_env_bool('GRAFANA_STORE_DASHBOARD_SHA256')

    @property
    def tls_cert(self) -> Optional[str]:
        """
        Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the
        `GRAFANA_TLS_CERT` environment variable.
        """
        return __config__.get('tlsCert') or _utilities.get_env('GRAFANA_TLS_CERT')

    @property
    def tls_key(self) -> Optional[str]:
        """
        Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY`
        environment variable.
        """
        return __config__.get('tlsKey') or _utilities.get_env('GRAFANA_TLS_KEY')

    @property
    def url(self) -> Optional[str]:
        """
        The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
        """
        return __config__.get('url') or _utilities.get_env('GRAFANA_URL')

