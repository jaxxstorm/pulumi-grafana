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
public final class OncallIntegrationDefaultRouteTelegram {
    private @Nullable Boolean enabled;
    /**
     * @return The ID of this resource.
     * 
     */
    private @Nullable String id;

    private OncallIntegrationDefaultRouteTelegram() {}
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OncallIntegrationDefaultRouteTelegram defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable String id;
        public Builder() {}
        public Builder(OncallIntegrationDefaultRouteTelegram defaults) {
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
        public OncallIntegrationDefaultRouteTelegram build() {
            final var o = new OncallIntegrationDefaultRouteTelegram();
            o.enabled = enabled;
            o.id = id;
            return o;
        }
    }
}
