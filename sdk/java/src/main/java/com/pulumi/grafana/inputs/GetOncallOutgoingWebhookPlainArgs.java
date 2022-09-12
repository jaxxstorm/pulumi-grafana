// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetOncallOutgoingWebhookPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOncallOutgoingWebhookPlainArgs Empty = new GetOncallOutgoingWebhookPlainArgs();

    /**
     * The outgoing webhook name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The outgoing webhook name.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetOncallOutgoingWebhookPlainArgs() {}

    private GetOncallOutgoingWebhookPlainArgs(GetOncallOutgoingWebhookPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOncallOutgoingWebhookPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOncallOutgoingWebhookPlainArgs $;

        public Builder() {
            $ = new GetOncallOutgoingWebhookPlainArgs();
        }

        public Builder(GetOncallOutgoingWebhookPlainArgs defaults) {
            $ = new GetOncallOutgoingWebhookPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The outgoing webhook name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetOncallOutgoingWebhookPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
