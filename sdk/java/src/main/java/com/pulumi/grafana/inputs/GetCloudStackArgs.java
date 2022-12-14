// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCloudStackArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudStackArgs Empty = new GetCloudStackArgs();

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

    private GetCloudStackArgs() {}

    private GetCloudStackArgs(GetCloudStackArgs $) {
        this.slug = $.slug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudStackArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudStackArgs $;

        public Builder() {
            $ = new GetCloudStackArgs();
        }

        public Builder(GetCloudStackArgs defaults) {
            $ = new GetCloudStackArgs(Objects.requireNonNull(defaults));
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

        public GetCloudStackArgs build() {
            $.slug = Objects.requireNonNull($.slug, "expected parameter 'slug' to be non-null");
            return $;
        }
    }

}
