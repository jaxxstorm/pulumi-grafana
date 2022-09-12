// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.grafana.inputs.RuleGroupRuleDataArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleGroupRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleGroupRuleArgs Empty = new RuleGroupRuleArgs();

    /**
     * Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    /**
     * @return Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
     * 
     */
    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * The `ref_id` of the query node in the `data` field to use as the alert condition.
     * 
     */
    @Import(name="condition", required=true)
    private Output<String> condition;

    /**
     * @return The `ref_id` of the query node in the `data` field to use as the alert condition.
     * 
     */
    public Output<String> condition() {
        return this.condition;
    }

    /**
     * A sequence of stages that describe the contents of the rule.
     * 
     */
    @Import(name="datas", required=true)
    private Output<List<RuleGroupRuleDataArgs>> datas;

    /**
     * @return A sequence of stages that describe the contents of the rule.
     * 
     */
    public Output<List<RuleGroupRuleDataArgs>> datas() {
        return this.datas;
    }

    /**
     * Describes what state to enter when the rule&#39;s query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
     * 
     */
    @Import(name="execErrState")
    private @Nullable Output<String> execErrState;

    /**
     * @return Describes what state to enter when the rule&#39;s query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
     * 
     */
    public Optional<Output<String>> execErrState() {
        return Optional.ofNullable(this.execErrState);
    }

    /**
     * The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
     * 
     */
    @Import(name="for")
    private @Nullable Output<String> for_;

    /**
     * @return The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
     * 
     */
    public Optional<Output<String>> for_() {
        return Optional.ofNullable(this.for_);
    }

    /**
     * Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The name of the alert rule.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the alert rule.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Describes what state to enter when the rule&#39;s query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
     * 
     */
    @Import(name="noDataState")
    private @Nullable Output<String> noDataState;

    /**
     * @return Describes what state to enter when the rule&#39;s query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
     * 
     */
    public Optional<Output<String>> noDataState() {
        return Optional.ofNullable(this.noDataState);
    }

    /**
     * The unique identifier of the alert rule.
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return The unique identifier of the alert rule.
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    private RuleGroupRuleArgs() {}

    private RuleGroupRuleArgs(RuleGroupRuleArgs $) {
        this.annotations = $.annotations;
        this.condition = $.condition;
        this.datas = $.datas;
        this.execErrState = $.execErrState;
        this.for_ = $.for_;
        this.labels = $.labels;
        this.name = $.name;
        this.noDataState = $.noDataState;
        this.uid = $.uid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleGroupRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleGroupRuleArgs $;

        public Builder() {
            $ = new RuleGroupRuleArgs();
        }

        public Builder(RuleGroupRuleArgs defaults) {
            $ = new RuleGroupRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param annotations Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        /**
         * @param annotations Key-value pairs of metadata to attach to the alert rule that may add user-defined context, but cannot be used for matching, grouping, or routing. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
        }

        /**
         * @param condition The `ref_id` of the query node in the `data` field to use as the alert condition.
         * 
         * @return builder
         * 
         */
        public Builder condition(Output<String> condition) {
            $.condition = condition;
            return this;
        }

        /**
         * @param condition The `ref_id` of the query node in the `data` field to use as the alert condition.
         * 
         * @return builder
         * 
         */
        public Builder condition(String condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param datas A sequence of stages that describe the contents of the rule.
         * 
         * @return builder
         * 
         */
        public Builder datas(Output<List<RuleGroupRuleDataArgs>> datas) {
            $.datas = datas;
            return this;
        }

        /**
         * @param datas A sequence of stages that describe the contents of the rule.
         * 
         * @return builder
         * 
         */
        public Builder datas(List<RuleGroupRuleDataArgs> datas) {
            return datas(Output.of(datas));
        }

        /**
         * @param datas A sequence of stages that describe the contents of the rule.
         * 
         * @return builder
         * 
         */
        public Builder datas(RuleGroupRuleDataArgs... datas) {
            return datas(List.of(datas));
        }

        /**
         * @param execErrState Describes what state to enter when the rule&#39;s query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
         * 
         * @return builder
         * 
         */
        public Builder execErrState(@Nullable Output<String> execErrState) {
            $.execErrState = execErrState;
            return this;
        }

        /**
         * @param execErrState Describes what state to enter when the rule&#39;s query is invalid and the rule cannot be executed. Options are OK, Error, and Alerting. Defaults to `Alerting`.
         * 
         * @return builder
         * 
         */
        public Builder execErrState(String execErrState) {
            return execErrState(Output.of(execErrState));
        }

        /**
         * @param for_ The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder for_(@Nullable Output<String> for_) {
            $.for_ = for_;
            return this;
        }

        /**
         * @param for_ The amount of time for which the rule must be breached for the rule to be considered to be Firing. Before this time has elapsed, the rule is only considered to be Pending. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder for_(String for_) {
            return for_(Output.of(for_));
        }

        /**
         * @param labels Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Key-value pairs to attach to the alert rule that can be used in matching, grouping, and routing. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param name The name of the alert rule.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the alert rule.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param noDataState Describes what state to enter when the rule&#39;s query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
         * 
         * @return builder
         * 
         */
        public Builder noDataState(@Nullable Output<String> noDataState) {
            $.noDataState = noDataState;
            return this;
        }

        /**
         * @param noDataState Describes what state to enter when the rule&#39;s query returns No Data. Options are OK, NoData, and Alerting. Defaults to `NoData`.
         * 
         * @return builder
         * 
         */
        public Builder noDataState(String noDataState) {
            return noDataState(Output.of(noDataState));
        }

        /**
         * @param uid The unique identifier of the alert rule.
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid The unique identifier of the alert rule.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        public RuleGroupRuleArgs build() {
            $.condition = Objects.requireNonNull($.condition, "expected parameter 'condition' to be non-null");
            $.datas = Objects.requireNonNull($.datas, "expected parameter 'datas' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
