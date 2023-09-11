// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class SLOQueryRatio {
    private @Nullable List<String> groupByLabels;
    private String successMetric;
    private String totalMetric;

    private SLOQueryRatio() {}
    public List<String> groupByLabels() {
        return this.groupByLabels == null ? List.of() : this.groupByLabels;
    }
    public String successMetric() {
        return this.successMetric;
    }
    public String totalMetric() {
        return this.totalMetric;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SLOQueryRatio defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> groupByLabels;
        private String successMetric;
        private String totalMetric;
        public Builder() {}
        public Builder(SLOQueryRatio defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupByLabels = defaults.groupByLabels;
    	      this.successMetric = defaults.successMetric;
    	      this.totalMetric = defaults.totalMetric;
        }

        @CustomType.Setter
        public Builder groupByLabels(@Nullable List<String> groupByLabels) {
            this.groupByLabels = groupByLabels;
            return this;
        }
        public Builder groupByLabels(String... groupByLabels) {
            return groupByLabels(List.of(groupByLabels));
        }
        @CustomType.Setter
        public Builder successMetric(String successMetric) {
            this.successMetric = Objects.requireNonNull(successMetric);
            return this;
        }
        @CustomType.Setter
        public Builder totalMetric(String totalMetric) {
            this.totalMetric = Objects.requireNonNull(totalMetric);
            return this;
        }
        public SLOQueryRatio build() {
            final var o = new SLOQueryRatio();
            o.groupByLabels = groupByLabels;
            o.successMetric = successMetric;
            o.totalMetric = totalMetric;
            return o;
        }
    }
}
