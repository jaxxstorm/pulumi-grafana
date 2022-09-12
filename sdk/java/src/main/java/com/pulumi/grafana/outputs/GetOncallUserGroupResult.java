// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOncallUserGroupResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String slackHandle;
    private String slackId;

    private GetOncallUserGroupResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String slackHandle() {
        return this.slackHandle;
    }
    public String slackId() {
        return this.slackId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOncallUserGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String slackHandle;
        private String slackId;
        public Builder() {}
        public Builder(GetOncallUserGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.slackHandle = defaults.slackHandle;
    	      this.slackId = defaults.slackId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder slackHandle(String slackHandle) {
            this.slackHandle = Objects.requireNonNull(slackHandle);
            return this;
        }
        @CustomType.Setter
        public Builder slackId(String slackId) {
            this.slackId = Objects.requireNonNull(slackId);
            return this;
        }
        public GetOncallUserGroupResult build() {
            final var o = new GetOncallUserGroupResult();
            o.id = id;
            o.slackHandle = slackHandle;
            o.slackId = slackId;
            return o;
        }
    }
}
