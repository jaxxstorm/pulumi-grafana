// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetOncallActionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOncallActionPlainArgs Empty = new GetOncallActionPlainArgs();

    /**
     * The action name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The action name.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetOncallActionPlainArgs() {}

    private GetOncallActionPlainArgs(GetOncallActionPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOncallActionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOncallActionPlainArgs $;

        public Builder() {
            $ = new GetOncallActionPlainArgs();
        }

        public Builder(GetOncallActionPlainArgs defaults) {
            $ = new GetOncallActionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The action name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetOncallActionPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
