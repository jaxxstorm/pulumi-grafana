// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReportTimeRange {
    /**
     * @return Start of the time range.
     * 
     */
    private @Nullable String from;
    /**
     * @return End of the time range.
     * 
     */
    private @Nullable String to;

    private ReportTimeRange() {}
    /**
     * @return Start of the time range.
     * 
     */
    public Optional<String> from() {
        return Optional.ofNullable(this.from);
    }
    /**
     * @return End of the time range.
     * 
     */
    public Optional<String> to() {
        return Optional.ofNullable(this.to);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReportTimeRange defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String from;
        private @Nullable String to;
        public Builder() {}
        public Builder(ReportTimeRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.from = defaults.from;
    	      this.to = defaults.to;
        }

        @CustomType.Setter
        public Builder from(@Nullable String from) {
            this.from = from;
            return this;
        }
        @CustomType.Setter
        public Builder to(@Nullable String to) {
            this.to = to;
            return this;
        }
        public ReportTimeRange build() {
            final var o = new ReportTimeRange();
            o.from = from;
            o.to = to;
            return o;
        }
    }
}
