// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApiKeyState extends com.pulumi.resources.ResourceArgs {

    public static final ApiKeyState Empty = new ApiKeyState();

    /**
     * If set, the API key will be created for the given Cloud stack. This can be used to bootstrap a management API key for a new stack. **Note**: This requires a cloud token to be configured.
     * 
     */
    @Import(name="cloudStackSlug")
    private @Nullable Output<String> cloudStackSlug;

    /**
     * @return If set, the API key will be created for the given Cloud stack. This can be used to bootstrap a management API key for a new stack. **Note**: This requires a cloud token to be configured.
     * 
     */
    public Optional<Output<String>> cloudStackSlug() {
        return Optional.ofNullable(this.cloudStackSlug);
    }

    @Import(name="expiration")
    private @Nullable Output<String> expiration;

    public Optional<Output<String>> expiration() {
        return Optional.ofNullable(this.expiration);
    }

    @Import(name="key")
    private @Nullable Output<String> key;

    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="role")
    private @Nullable Output<String> role;

    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="secondsToLive")
    private @Nullable Output<Integer> secondsToLive;

    public Optional<Output<Integer>> secondsToLive() {
        return Optional.ofNullable(this.secondsToLive);
    }

    private ApiKeyState() {}

    private ApiKeyState(ApiKeyState $) {
        this.cloudStackSlug = $.cloudStackSlug;
        this.expiration = $.expiration;
        this.key = $.key;
        this.name = $.name;
        this.role = $.role;
        this.secondsToLive = $.secondsToLive;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApiKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApiKeyState $;

        public Builder() {
            $ = new ApiKeyState();
        }

        public Builder(ApiKeyState defaults) {
            $ = new ApiKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cloudStackSlug If set, the API key will be created for the given Cloud stack. This can be used to bootstrap a management API key for a new stack. **Note**: This requires a cloud token to be configured.
         * 
         * @return builder
         * 
         */
        public Builder cloudStackSlug(@Nullable Output<String> cloudStackSlug) {
            $.cloudStackSlug = cloudStackSlug;
            return this;
        }

        /**
         * @param cloudStackSlug If set, the API key will be created for the given Cloud stack. This can be used to bootstrap a management API key for a new stack. **Note**: This requires a cloud token to be configured.
         * 
         * @return builder
         * 
         */
        public Builder cloudStackSlug(String cloudStackSlug) {
            return cloudStackSlug(Output.of(cloudStackSlug));
        }

        public Builder expiration(@Nullable Output<String> expiration) {
            $.expiration = expiration;
            return this;
        }

        public Builder expiration(String expiration) {
            return expiration(Output.of(expiration));
        }

        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder secondsToLive(@Nullable Output<Integer> secondsToLive) {
            $.secondsToLive = secondsToLive;
            return this;
        }

        public Builder secondsToLive(Integer secondsToLive) {
            return secondsToLive(Output.of(secondsToLive));
        }

        public ApiKeyState build() {
            return $;
        }
    }

}
