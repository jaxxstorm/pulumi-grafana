// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp {
    private @Nullable Boolean allowMissing;
    private String header;
    private String regexp;

    private SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp() {}
    public Optional<Boolean> allowMissing() {
        return Optional.ofNullable(this.allowMissing);
    }
    public String header() {
        return this.header;
    }
    public String regexp() {
        return this.regexp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowMissing;
        private String header;
        private String regexp;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowMissing = defaults.allowMissing;
    	      this.header = defaults.header;
    	      this.regexp = defaults.regexp;
        }

        @CustomType.Setter
        public Builder allowMissing(@Nullable Boolean allowMissing) {
            this.allowMissing = allowMissing;
            return this;
        }
        @CustomType.Setter
        public Builder header(String header) {
            this.header = Objects.requireNonNull(header);
            return this;
        }
        @CustomType.Setter
        public Builder regexp(String regexp) {
            this.regexp = Objects.requireNonNull(regexp);
            return this;
        }
        public SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp build() {
            final var o = new SyntheticMonitoringCheckSettingsHttpFailIfHeaderNotMatchesRegexp();
            o.allowMissing = allowMissing;
            o.header = header;
            o.regexp = regexp;
            return o;
        }
    }
}
