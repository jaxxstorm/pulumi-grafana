// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Grafana.Outputs
{

    [OutputType]
    public sealed class RuleGroupRule
    {
        /// <summary>
        /// Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The `ref_id` of the query node in the `data` field to use as the alert condition.
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// A sequence of stages that describe the contents of the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupRuleData> Datas;
        /// <summary>
        /// Describes what state to enter when the rule's query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
        /// </summary>
        public readonly string? ExecErrState;
        /// <summary>
        /// The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
        /// </summary>
        public readonly string? For;
        /// <summary>
        /// Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes what state to enter when the rule's query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
        /// </summary>
        public readonly string? NoDataState;
        /// <summary>
        /// The unique identifier of the alert rule.
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private RuleGroupRule(
            ImmutableDictionary<string, string>? annotations,

            string condition,

            ImmutableArray<Outputs.RuleGroupRuleData> datas,

            string? execErrState,

            string? @for,

            ImmutableDictionary<string, string>? labels,

            string name,

            string? noDataState,

            string? uid)
        {
            Annotations = annotations;
            Condition = condition;
            Datas = datas;
            ExecErrState = execErrState;
            For = @for;
            Labels = labels;
            Name = name;
            NoDataState = noDataState;
            Uid = uid;
        }
    }
}
