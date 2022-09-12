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


public final class DashboardArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardArgs Empty = new DashboardArgs();

    /**
     * The complete dashboard model JSON.
     * 
     */
    @Import(name="configJson", required=true)
    private Output<String> configJson;

    /**
     * @return The complete dashboard model JSON.
     * 
     */
    public Output<String> configJson() {
        return this.configJson;
    }

    /**
     * The id of the folder to save the dashboard in. This attribute is a string to reflect the type of the folder&#39;s id.
     * 
     */
    @Import(name="folder")
    private @Nullable Output<String> folder;

    /**
     * @return The id of the folder to save the dashboard in. This attribute is a string to reflect the type of the folder&#39;s id.
     * 
     */
    public Optional<Output<String>> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * Set a commit message for the version history.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return Set a commit message for the version history.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
     * 
     */
    @Import(name="overwrite")
    private @Nullable Output<Boolean> overwrite;

    /**
     * @return Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
     * 
     */
    public Optional<Output<Boolean>> overwrite() {
        return Optional.ofNullable(this.overwrite);
    }

    private DashboardArgs() {}

    private DashboardArgs(DashboardArgs $) {
        this.configJson = $.configJson;
        this.folder = $.folder;
        this.message = $.message;
        this.overwrite = $.overwrite;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardArgs $;

        public Builder() {
            $ = new DashboardArgs();
        }

        public Builder(DashboardArgs defaults) {
            $ = new DashboardArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configJson The complete dashboard model JSON.
         * 
         * @return builder
         * 
         */
        public Builder configJson(Output<String> configJson) {
            $.configJson = configJson;
            return this;
        }

        /**
         * @param configJson The complete dashboard model JSON.
         * 
         * @return builder
         * 
         */
        public Builder configJson(String configJson) {
            return configJson(Output.of(configJson));
        }

        /**
         * @param folder The id of the folder to save the dashboard in. This attribute is a string to reflect the type of the folder&#39;s id.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable Output<String> folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param folder The id of the folder to save the dashboard in. This attribute is a string to reflect the type of the folder&#39;s id.
         * 
         * @return builder
         * 
         */
        public Builder folder(String folder) {
            return folder(Output.of(folder));
        }

        /**
         * @param message Set a commit message for the version history.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message Set a commit message for the version history.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param overwrite Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
         * 
         * @return builder
         * 
         */
        public Builder overwrite(@Nullable Output<Boolean> overwrite) {
            $.overwrite = overwrite;
            return this;
        }

        /**
         * @param overwrite Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
         * 
         * @return builder
         * 
         */
        public Builder overwrite(Boolean overwrite) {
            return overwrite(Output.of(overwrite));
        }

        public DashboardArgs build() {
            $.configJson = Objects.requireNonNull($.configJson, "expected parameter 'configJson' to be non-null");
            return $;
        }
    }

}
