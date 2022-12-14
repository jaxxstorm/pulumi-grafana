// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudPluginInstallationState extends com.pulumi.resources.ResourceArgs {

    public static final CloudPluginInstallationState Empty = new CloudPluginInstallationState();

    /**
     * Slug of the plugin to be installed.
     * 
     */
    @Import(name="slug")
    private @Nullable Output<String> slug;

    /**
     * @return Slug of the plugin to be installed.
     * 
     */
    public Optional<Output<String>> slug() {
        return Optional.ofNullable(this.slug);
    }

    /**
     * The stack id to which the plugin should be installed.
     * 
     */
    @Import(name="stackSlug")
    private @Nullable Output<String> stackSlug;

    /**
     * @return The stack id to which the plugin should be installed.
     * 
     */
    public Optional<Output<String>> stackSlug() {
        return Optional.ofNullable(this.stackSlug);
    }

    /**
     * Version of the plugin to be installed.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the plugin to be installed.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private CloudPluginInstallationState() {}

    private CloudPluginInstallationState(CloudPluginInstallationState $) {
        this.slug = $.slug;
        this.stackSlug = $.stackSlug;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudPluginInstallationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudPluginInstallationState $;

        public Builder() {
            $ = new CloudPluginInstallationState();
        }

        public Builder(CloudPluginInstallationState defaults) {
            $ = new CloudPluginInstallationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param slug Slug of the plugin to be installed.
         * 
         * @return builder
         * 
         */
        public Builder slug(@Nullable Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug Slug of the plugin to be installed.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
        }

        /**
         * @param stackSlug The stack id to which the plugin should be installed.
         * 
         * @return builder
         * 
         */
        public Builder stackSlug(@Nullable Output<String> stackSlug) {
            $.stackSlug = stackSlug;
            return this;
        }

        /**
         * @param stackSlug The stack id to which the plugin should be installed.
         * 
         * @return builder
         * 
         */
        public Builder stackSlug(String stackSlug) {
            return stackSlug(Output.of(stackSlug));
        }

        /**
         * @param version Version of the plugin to be installed.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the plugin to be installed.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public CloudPluginInstallationState build() {
            return $;
        }
    }

}
