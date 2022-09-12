// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr {
    private @Nullable List<String> failIfMatchesRegexps;
    private @Nullable List<String> failIfNotMatchesRegexps;

    private SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr() {}
    public List<String> failIfMatchesRegexps() {
        return this.failIfMatchesRegexps == null ? List.of() : this.failIfMatchesRegexps;
    }
    public List<String> failIfNotMatchesRegexps() {
        return this.failIfNotMatchesRegexps == null ? List.of() : this.failIfNotMatchesRegexps;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> failIfMatchesRegexps;
        private @Nullable List<String> failIfNotMatchesRegexps;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failIfMatchesRegexps = defaults.failIfMatchesRegexps;
    	      this.failIfNotMatchesRegexps = defaults.failIfNotMatchesRegexps;
        }

        @CustomType.Setter
        public Builder failIfMatchesRegexps(@Nullable List<String> failIfMatchesRegexps) {
            this.failIfMatchesRegexps = failIfMatchesRegexps;
            return this;
        }
        public Builder failIfMatchesRegexps(String... failIfMatchesRegexps) {
            return failIfMatchesRegexps(List.of(failIfMatchesRegexps));
        }
        @CustomType.Setter
        public Builder failIfNotMatchesRegexps(@Nullable List<String> failIfNotMatchesRegexps) {
            this.failIfNotMatchesRegexps = failIfNotMatchesRegexps;
            return this;
        }
        public Builder failIfNotMatchesRegexps(String... failIfNotMatchesRegexps) {
            return failIfNotMatchesRegexps(List.of(failIfNotMatchesRegexps));
        }
        public SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr build() {
            final var o = new SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr();
            o.failIfMatchesRegexps = failIfMatchesRegexps;
            o.failIfNotMatchesRegexps = failIfNotMatchesRegexps;
            return o;
        }
    }
}
