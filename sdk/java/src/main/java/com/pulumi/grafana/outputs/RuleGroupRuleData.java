// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.RuleGroupRuleDataRelativeTimeRange;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuleGroupRuleData {
    private String datasourceUid;
    private String model;
    private @Nullable String queryType;
    private String refId;
    private RuleGroupRuleDataRelativeTimeRange relativeTimeRange;

    private RuleGroupRuleData() {}
    public String datasourceUid() {
        return this.datasourceUid;
    }
    public String model() {
        return this.model;
    }
    public Optional<String> queryType() {
        return Optional.ofNullable(this.queryType);
    }
    public String refId() {
        return this.refId;
    }
    public RuleGroupRuleDataRelativeTimeRange relativeTimeRange() {
        return this.relativeTimeRange;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleGroupRuleData defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String datasourceUid;
        private String model;
        private @Nullable String queryType;
        private String refId;
        private RuleGroupRuleDataRelativeTimeRange relativeTimeRange;
        public Builder() {}
        public Builder(RuleGroupRuleData defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datasourceUid = defaults.datasourceUid;
    	      this.model = defaults.model;
    	      this.queryType = defaults.queryType;
    	      this.refId = defaults.refId;
    	      this.relativeTimeRange = defaults.relativeTimeRange;
        }

        @CustomType.Setter
        public Builder datasourceUid(String datasourceUid) {
            this.datasourceUid = Objects.requireNonNull(datasourceUid);
            return this;
        }
        @CustomType.Setter
        public Builder model(String model) {
            this.model = Objects.requireNonNull(model);
            return this;
        }
        @CustomType.Setter
        public Builder queryType(@Nullable String queryType) {
            this.queryType = queryType;
            return this;
        }
        @CustomType.Setter
        public Builder refId(String refId) {
            this.refId = Objects.requireNonNull(refId);
            return this;
        }
        @CustomType.Setter
        public Builder relativeTimeRange(RuleGroupRuleDataRelativeTimeRange relativeTimeRange) {
            this.relativeTimeRange = Objects.requireNonNull(relativeTimeRange);
            return this;
        }
        public RuleGroupRuleData build() {
            final var o = new RuleGroupRuleData();
            o.datasourceUid = datasourceUid;
            o.model = model;
            o.queryType = queryType;
            o.refId = refId;
            o.relativeTimeRange = relativeTimeRange;
            return o;
        }
    }
}
