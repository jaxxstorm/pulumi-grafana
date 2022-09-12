// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContactPointDiscordArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContactPointDiscordArgs Empty = new ContactPointDiscordArgs();

    /**
     * The URL of a custom avatar image to use. Defaults to ``.
     * 
     */
    @Import(name="avatarUrl")
    private @Nullable Output<String> avatarUrl;

    /**
     * @return The URL of a custom avatar image to use. Defaults to ``.
     * 
     */
    public Optional<Output<String>> avatarUrl() {
        return Optional.ofNullable(this.avatarUrl);
    }

    /**
     * Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    @Import(name="disableResolveMessage")
    private @Nullable Output<Boolean> disableResolveMessage;

    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> disableResolveMessage() {
        return Optional.ofNullable(this.disableResolveMessage);
    }

    /**
     * The templated content of the message. Defaults to ``.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return The templated content of the message. Defaults to ``.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    @Import(name="settings")
    private @Nullable Output<Map<String,String>> settings;

    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    public Optional<Output<Map<String,String>>> settings() {
        return Optional.ofNullable(this.settings);
    }

    /**
     * The UID of the contact point.
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return The UID of the contact point.
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    /**
     * The discord webhook URL.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The discord webhook URL.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * Whether to use the bot account&#39;s plain username instead of &#34;Grafana.&#34; Defaults to `false`.
     * 
     */
    @Import(name="useDiscordUsername")
    private @Nullable Output<Boolean> useDiscordUsername;

    /**
     * @return Whether to use the bot account&#39;s plain username instead of &#34;Grafana.&#34; Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> useDiscordUsername() {
        return Optional.ofNullable(this.useDiscordUsername);
    }

    private ContactPointDiscordArgs() {}

    private ContactPointDiscordArgs(ContactPointDiscordArgs $) {
        this.avatarUrl = $.avatarUrl;
        this.disableResolveMessage = $.disableResolveMessage;
        this.message = $.message;
        this.settings = $.settings;
        this.uid = $.uid;
        this.url = $.url;
        this.useDiscordUsername = $.useDiscordUsername;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContactPointDiscordArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContactPointDiscordArgs $;

        public Builder() {
            $ = new ContactPointDiscordArgs();
        }

        public Builder(ContactPointDiscordArgs defaults) {
            $ = new ContactPointDiscordArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param avatarUrl The URL of a custom avatar image to use. Defaults to ``.
         * 
         * @return builder
         * 
         */
        public Builder avatarUrl(@Nullable Output<String> avatarUrl) {
            $.avatarUrl = avatarUrl;
            return this;
        }

        /**
         * @param avatarUrl The URL of a custom avatar image to use. Defaults to ``.
         * 
         * @return builder
         * 
         */
        public Builder avatarUrl(String avatarUrl) {
            return avatarUrl(Output.of(avatarUrl));
        }

        /**
         * @param disableResolveMessage Whether to disable sending resolve messages. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableResolveMessage(@Nullable Output<Boolean> disableResolveMessage) {
            $.disableResolveMessage = disableResolveMessage;
            return this;
        }

        /**
         * @param disableResolveMessage Whether to disable sending resolve messages. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableResolveMessage(Boolean disableResolveMessage) {
            return disableResolveMessage(Output.of(disableResolveMessage));
        }

        /**
         * @param message The templated content of the message. Defaults to ``.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message The templated content of the message. Defaults to ``.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param settings Additional custom properties to attach to the notifier. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder settings(@Nullable Output<Map<String,String>> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings Additional custom properties to attach to the notifier. Defaults to `map[]`.
         * 
         * @return builder
         * 
         */
        public Builder settings(Map<String,String> settings) {
            return settings(Output.of(settings));
        }

        /**
         * @param uid The UID of the contact point.
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid The UID of the contact point.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        /**
         * @param url The discord webhook URL.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The discord webhook URL.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param useDiscordUsername Whether to use the bot account&#39;s plain username instead of &#34;Grafana.&#34; Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder useDiscordUsername(@Nullable Output<Boolean> useDiscordUsername) {
            $.useDiscordUsername = useDiscordUsername;
            return this;
        }

        /**
         * @param useDiscordUsername Whether to use the bot account&#39;s plain username instead of &#34;Grafana.&#34; Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder useDiscordUsername(Boolean useDiscordUsername) {
            return useDiscordUsername(Output.of(useDiscordUsername));
        }

        public ContactPointDiscordArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
