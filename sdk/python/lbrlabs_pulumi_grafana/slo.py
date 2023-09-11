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

__all__ = ['SLOArgs', 'SLO']

@pulumi.input_type
class SLOArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 objectives: pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]],
                 queries: pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]],
                 alertings: Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SLO resource.
        :param pulumi.Input[str] description: Description is a free-text field that can provide more context to an SLO.
        :param pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]] objectives: Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        :param pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]] queries: Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        :param pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]] alertings: Configures the alerting rules that will be generated for each
               			time window associated with the SLO. Grafana SLOs can generate
               			alerts when the short-term error budget burn is very high, the
               			long-term error budget burn rate is high, or when the remaining
               			error budget is below a certain threshold. Annotations and Labels support templating.
        :param pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]] labels: Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        :param pulumi.Input[str] name: Name should be a short description of your indicator. Consider names like "API Availability"
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "objectives", objectives)
        pulumi.set(__self__, "queries", queries)
        if alertings is not None:
            pulumi.set(__self__, "alertings", alertings)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description is a free-text field that can provide more context to an SLO.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def objectives(self) -> pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]]:
        """
        Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        """
        return pulumi.get(self, "objectives")

    @objectives.setter
    def objectives(self, value: pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]]):
        pulumi.set(self, "objectives", value)

    @property
    @pulumi.getter
    def queries(self) -> pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]]:
        """
        Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        return pulumi.get(self, "queries")

    @queries.setter
    def queries(self, value: pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]]):
        pulumi.set(self, "queries", value)

    @property
    @pulumi.getter
    def alertings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]]:
        """
        Configures the alerting rules that will be generated for each
        			time window associated with the SLO. Grafana SLOs can generate
        			alerts when the short-term error budget burn is very high, the
        			long-term error budget burn rate is high, or when the remaining
        			error budget is below a certain threshold. Annotations and Labels support templating.
        """
        return pulumi.get(self, "alertings")

    @alertings.setter
    def alertings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]]):
        pulumi.set(self, "alertings", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]]:
        """
        Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name should be a short description of your indicator. Consider names like "API Availability"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SLOState:
    def __init__(__self__, *,
                 alertings: Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 objectives: Optional[pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]]] = None,
                 queries: Optional[pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]]] = None):
        """
        Input properties used for looking up and filtering SLO resources.
        :param pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]] alertings: Configures the alerting rules that will be generated for each
               			time window associated with the SLO. Grafana SLOs can generate
               			alerts when the short-term error budget burn is very high, the
               			long-term error budget burn rate is high, or when the remaining
               			error budget is below a certain threshold. Annotations and Labels support templating.
        :param pulumi.Input[str] description: Description is a free-text field that can provide more context to an SLO.
        :param pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]] labels: Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        :param pulumi.Input[str] name: Name should be a short description of your indicator. Consider names like "API Availability"
        :param pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]] objectives: Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        :param pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]] queries: Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        if alertings is not None:
            pulumi.set(__self__, "alertings", alertings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if objectives is not None:
            pulumi.set(__self__, "objectives", objectives)
        if queries is not None:
            pulumi.set(__self__, "queries", queries)

    @property
    @pulumi.getter
    def alertings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]]:
        """
        Configures the alerting rules that will be generated for each
        			time window associated with the SLO. Grafana SLOs can generate
        			alerts when the short-term error budget burn is very high, the
        			long-term error budget burn rate is high, or when the remaining
        			error budget is below a certain threshold. Annotations and Labels support templating.
        """
        return pulumi.get(self, "alertings")

    @alertings.setter
    def alertings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOAlertingArgs']]]]):
        pulumi.set(self, "alertings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description is a free-text field that can provide more context to an SLO.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]]:
        """
        Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name should be a short description of your indicator. Consider names like "API Availability"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def objectives(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]]]:
        """
        Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        """
        return pulumi.get(self, "objectives")

    @objectives.setter
    def objectives(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOObjectiveArgs']]]]):
        pulumi.set(self, "objectives", value)

    @property
    @pulumi.getter
    def queries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]]]:
        """
        Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        return pulumi.get(self, "queries")

    @queries.setter
    def queries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SLOQueryArgs']]]]):
        pulumi.set(self, "queries", value)


class SLO(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alertings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOAlertingArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOLabelArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 objectives: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOObjectiveArgs']]]]] = None,
                 queries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOQueryArgs']]]]] = None,
                 __props__=None):
        """
        Resource manages Grafana SLOs.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/slo/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/slo/api/)
        * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOAlertingArgs']]]] alertings: Configures the alerting rules that will be generated for each
               			time window associated with the SLO. Grafana SLOs can generate
               			alerts when the short-term error budget burn is very high, the
               			long-term error budget burn rate is high, or when the remaining
               			error budget is below a certain threshold. Annotations and Labels support templating.
        :param pulumi.Input[str] description: Description is a free-text field that can provide more context to an SLO.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOLabelArgs']]]] labels: Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        :param pulumi.Input[str] name: Name should be a short description of your indicator. Consider names like "API Availability"
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOObjectiveArgs']]]] objectives: Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOQueryArgs']]]] queries: Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SLOArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource manages Grafana SLOs.

        * [Official documentation](https://grafana.com/docs/grafana-cloud/slo/)
        * [API documentation](https://grafana.com/docs/grafana-cloud/slo/api/)
        * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param SLOArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SLOArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alertings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOAlertingArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOLabelArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 objectives: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOObjectiveArgs']]]]] = None,
                 queries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOQueryArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SLOArgs.__new__(SLOArgs)

            __props__.__dict__["alertings"] = alertings
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if objectives is None and not opts.urn:
                raise TypeError("Missing required property 'objectives'")
            __props__.__dict__["objectives"] = objectives
            if queries is None and not opts.urn:
                raise TypeError("Missing required property 'queries'")
            __props__.__dict__["queries"] = queries
        super(SLO, __self__).__init__(
            'grafana:index/sLO:SLO',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alertings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOAlertingArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOLabelArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            objectives: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOObjectiveArgs']]]]] = None,
            queries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOQueryArgs']]]]] = None) -> 'SLO':
        """
        Get an existing SLO resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOAlertingArgs']]]] alertings: Configures the alerting rules that will be generated for each
               			time window associated with the SLO. Grafana SLOs can generate
               			alerts when the short-term error budget burn is very high, the
               			long-term error budget burn rate is high, or when the remaining
               			error budget is below a certain threshold. Annotations and Labels support templating.
        :param pulumi.Input[str] description: Description is a free-text field that can provide more context to an SLO.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOLabelArgs']]]] labels: Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        :param pulumi.Input[str] name: Name should be a short description of your indicator. Consider names like "API Availability"
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOObjectiveArgs']]]] objectives: Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SLOQueryArgs']]]] queries: Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SLOState.__new__(_SLOState)

        __props__.__dict__["alertings"] = alertings
        __props__.__dict__["description"] = description
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["objectives"] = objectives
        __props__.__dict__["queries"] = queries
        return SLO(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alertings(self) -> pulumi.Output[Optional[Sequence['outputs.SLOAlerting']]]:
        """
        Configures the alerting rules that will be generated for each
        			time window associated with the SLO. Grafana SLOs can generate
        			alerts when the short-term error budget burn is very high, the
        			long-term error budget burn rate is high, or when the remaining
        			error budget is below a certain threshold. Annotations and Labels support templating.
        """
        return pulumi.get(self, "alertings")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description is a free-text field that can provide more context to an SLO.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence['outputs.SLOLabel']]]:
        """
        Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name should be a short description of your indicator. Consider names like "API Availability"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def objectives(self) -> pulumi.Output[Sequence['outputs.SLOObjective']]:
        """
        Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
        """
        return pulumi.get(self, "objectives")

    @property
    @pulumi.getter
    def queries(self) -> pulumi.Output[Sequence['outputs.SLOQuery']]:
        """
        Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
        """
        return pulumi.get(self, "queries")

