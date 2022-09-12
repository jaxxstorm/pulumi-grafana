// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudStackArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudStackArgs Empty = new CloudStackArgs();

    /**
     * Description of stack.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of stack.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Region slug to assign to this stack.
     * Changing region will destroy the existing stack and create a new one in the desired region
     * 
     */
    @Import(name="regionSlug")
    private @Nullable Output<String> regionSlug;

    /**
     * @return Region slug to assign to this stack.
     * Changing region will destroy the existing stack and create a new one in the desired region
     * 
     */
    public Optional<Output<String>> regionSlug() {
        return Optional.ofNullable(this.regionSlug);
    }

    /**
     * Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net&#34;.
     * 
     */
    @Import(name="slug", required=true)
    private Output<String> slug;

    /**
     * @return Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net&#34;.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }

    /**
     * Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
     * 
     */
    @Import(name="waitForReadiness")
    private @Nullable Output<Boolean> waitForReadiness;

    /**
     * @return Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> waitForReadiness() {
        return Optional.ofNullable(this.waitForReadiness);
    }

    /**
     * How long to wait for readiness (if enabled). Defaults to `5m0s`.
     * 
     */
    @Import(name="waitForReadinessTimeout")
    private @Nullable Output<String> waitForReadinessTimeout;

    /**
     * @return How long to wait for readiness (if enabled). Defaults to `5m0s`.
     * 
     */
    public Optional<Output<String>> waitForReadinessTimeout() {
        return Optional.ofNullable(this.waitForReadinessTimeout);
    }

    private CloudStackArgs() {}

    private CloudStackArgs(CloudStackArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.regionSlug = $.regionSlug;
        this.slug = $.slug;
        this.url = $.url;
        this.waitForReadiness = $.waitForReadiness;
        this.waitForReadinessTimeout = $.waitForReadinessTimeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudStackArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudStackArgs $;

        public Builder() {
            $ = new CloudStackArgs();
        }

        public Builder(CloudStackArgs defaults) {
            $ = new CloudStackArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of stack.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of stack.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param regionSlug Region slug to assign to this stack.
         * Changing region will destroy the existing stack and create a new one in the desired region
         * 
         * @return builder
         * 
         */
        public Builder regionSlug(@Nullable Output<String> regionSlug) {
            $.regionSlug = regionSlug;
            return this;
        }

        /**
         * @param regionSlug Region slug to assign to this stack.
         * Changing region will destroy the existing stack and create a new one in the desired region
         * 
         * @return builder
         * 
         */
        public Builder regionSlug(String regionSlug) {
            return regionSlug(Output.of(regionSlug));
        }

        /**
         * @param slug Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
         * available at “https://\n\n.grafana.net&#34;.
         * 
         * @return builder
         * 
         */
        public Builder slug(Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
         * available at “https://\n\n.grafana.net&#34;.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
        }

        /**
         * @param url Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param waitForReadiness Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder waitForReadiness(@Nullable Output<Boolean> waitForReadiness) {
            $.waitForReadiness = waitForReadiness;
            return this;
        }

        /**
         * @param waitForReadiness Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder waitForReadiness(Boolean waitForReadiness) {
            return waitForReadiness(Output.of(waitForReadiness));
        }

        /**
         * @param waitForReadinessTimeout How long to wait for readiness (if enabled). Defaults to `5m0s`.
         * 
         * @return builder
         * 
         */
        public Builder waitForReadinessTimeout(@Nullable Output<String> waitForReadinessTimeout) {
            $.waitForReadinessTimeout = waitForReadinessTimeout;
            return this;
        }

        /**
         * @param waitForReadinessTimeout How long to wait for readiness (if enabled). Defaults to `5m0s`.
         * 
         * @return builder
         * 
         */
        public Builder waitForReadinessTimeout(String waitForReadinessTimeout) {
            return waitForReadinessTimeout(Output.of(waitForReadinessTimeout));
        }

        public CloudStackArgs build() {
            $.slug = Objects.requireNonNull($.slug, "expected parameter 'slug' to be non-null");
            return $;
        }
    }

}
