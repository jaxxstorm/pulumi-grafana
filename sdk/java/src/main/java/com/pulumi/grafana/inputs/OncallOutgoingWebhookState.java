// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OncallOutgoingWebhookState extends com.pulumi.resources.ResourceArgs {

    public static final OncallOutgoingWebhookState Empty = new OncallOutgoingWebhookState();

    /**
     * The auth data of the webhook. Used in Authorization header instead of user/password auth.
     * 
     */
    @Import(name="authorizationHeader")
    private @Nullable Output<String> authorizationHeader;

    /**
     * @return The auth data of the webhook. Used in Authorization header instead of user/password auth.
     * 
     */
    public Optional<Output<String>> authorizationHeader() {
        return Optional.ofNullable(this.authorizationHeader);
    }

    /**
     * The data of the webhook.
     * 
     */
    @Import(name="data")
    private @Nullable Output<String> data;

    /**
     * @return The data of the webhook.
     * 
     */
    public Optional<Output<String>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * Forwards whole payload of the alert to the webhook&#39;s url as POST data.
     * 
     */
    @Import(name="forwardWholePayload")
    private @Nullable Output<Boolean> forwardWholePayload;

    /**
     * @return Forwards whole payload of the alert to the webhook&#39;s url as POST data.
     * 
     */
    public Optional<Output<Boolean>> forwardWholePayload() {
        return Optional.ofNullable(this.forwardWholePayload);
    }

    /**
     * The name of the outgoing webhook.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the outgoing webhook.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The auth data of the webhook. Used for Basic authentication
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The auth data of the webhook. Used for Basic authentication
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<String> teamId;

    /**
     * @return The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     * 
     */
    public Optional<Output<String>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    /**
     * The webhook URL.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The webhook URL.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * The auth data of the webhook. Used for Basic authentication.
     * 
     */
    @Import(name="user")
    private @Nullable Output<String> user;

    /**
     * @return The auth data of the webhook. Used for Basic authentication.
     * 
     */
    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    private OncallOutgoingWebhookState() {}

    private OncallOutgoingWebhookState(OncallOutgoingWebhookState $) {
        this.authorizationHeader = $.authorizationHeader;
        this.data = $.data;
        this.forwardWholePayload = $.forwardWholePayload;
        this.name = $.name;
        this.password = $.password;
        this.teamId = $.teamId;
        this.url = $.url;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OncallOutgoingWebhookState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OncallOutgoingWebhookState $;

        public Builder() {
            $ = new OncallOutgoingWebhookState();
        }

        public Builder(OncallOutgoingWebhookState defaults) {
            $ = new OncallOutgoingWebhookState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizationHeader The auth data of the webhook. Used in Authorization header instead of user/password auth.
         * 
         * @return builder
         * 
         */
        public Builder authorizationHeader(@Nullable Output<String> authorizationHeader) {
            $.authorizationHeader = authorizationHeader;
            return this;
        }

        /**
         * @param authorizationHeader The auth data of the webhook. Used in Authorization header instead of user/password auth.
         * 
         * @return builder
         * 
         */
        public Builder authorizationHeader(String authorizationHeader) {
            return authorizationHeader(Output.of(authorizationHeader));
        }

        /**
         * @param data The data of the webhook.
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<String> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data The data of the webhook.
         * 
         * @return builder
         * 
         */
        public Builder data(String data) {
            return data(Output.of(data));
        }

        /**
         * @param forwardWholePayload Forwards whole payload of the alert to the webhook&#39;s url as POST data.
         * 
         * @return builder
         * 
         */
        public Builder forwardWholePayload(@Nullable Output<Boolean> forwardWholePayload) {
            $.forwardWholePayload = forwardWholePayload;
            return this;
        }

        /**
         * @param forwardWholePayload Forwards whole payload of the alert to the webhook&#39;s url as POST data.
         * 
         * @return builder
         * 
         */
        public Builder forwardWholePayload(Boolean forwardWholePayload) {
            return forwardWholePayload(Output.of(forwardWholePayload));
        }

        /**
         * @param name The name of the outgoing webhook.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the outgoing webhook.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The auth data of the webhook. Used for Basic authentication
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The auth data of the webhook. Used for Basic authentication
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param teamId The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        /**
         * @param url The webhook URL.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The webhook URL.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param user The auth data of the webhook. Used for Basic authentication.
         * 
         * @return builder
         * 
         */
        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        /**
         * @param user The auth data of the webhook. Used for Basic authentication.
         * 
         * @return builder
         * 
         */
        public Builder user(String user) {
            return user(Output.of(user));
        }

        public OncallOutgoingWebhookState build() {
            return $;
        }
    }

}
