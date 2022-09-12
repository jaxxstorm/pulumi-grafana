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
public final class OncallRouteMsteams {
    /**
     * @return Enable notification in MS teams. Defaults to `true`.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
     * 
     */
    private @Nullable String id;

    private OncallRouteMsteams() {}
    /**
     * @return Enable notification in MS teams. Defaults to `true`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OncallRouteMsteams defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable String id;
        public Builder() {}
        public Builder(OncallRouteMsteams defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        public OncallRouteMsteams build() {
            final var o = new OncallRouteMsteams();
            o.enabled = enabled;
            o.id = id;
            return o;
        }
    }
}
