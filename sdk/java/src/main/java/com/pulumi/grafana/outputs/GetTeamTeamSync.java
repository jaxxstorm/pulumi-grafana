// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetTeamTeamSync {
    private @Nullable List<String> groups;

    private GetTeamTeamSync() {}
    public List<String> groups() {
        return this.groups == null ? List.of() : this.groups;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamTeamSync defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> groups;
        public Builder() {}
        public Builder(GetTeamTeamSync defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<String> groups) {
            this.groups = groups;
            return this;
        }
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }
        public GetTeamTeamSync build() {
            final var o = new GetTeamTeamSync();
            o.groups = groups;
            return o;
        }
    }
}