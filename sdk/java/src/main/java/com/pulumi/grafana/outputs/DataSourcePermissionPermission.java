// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DataSourcePermissionPermission {
    /**
     * @return Permission to associate with item. Must be `Query`.
     * 
     */
    private String permission;
    /**
     * @return ID of the team to manage permissions for. Defaults to `0`.
     * 
     */
    private @Nullable Integer teamId;
    /**
     * @return ID of the user to manage permissions for. Defaults to `0`.
     * 
     */
    private @Nullable Integer userId;

    private DataSourcePermissionPermission() {}
    /**
     * @return Permission to associate with item. Must be `Query`.
     * 
     */
    public String permission() {
        return this.permission;
    }
    /**
     * @return ID of the team to manage permissions for. Defaults to `0`.
     * 
     */
    public Optional<Integer> teamId() {
        return Optional.ofNullable(this.teamId);
    }
    /**
     * @return ID of the user to manage permissions for. Defaults to `0`.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataSourcePermissionPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String permission;
        private @Nullable Integer teamId;
        private @Nullable Integer userId;
        public Builder() {}
        public Builder(DataSourcePermissionPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.permission = defaults.permission;
    	      this.teamId = defaults.teamId;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder permission(String permission) {
            this.permission = Objects.requireNonNull(permission);
            return this;
        }
        @CustomType.Setter
        public Builder teamId(@Nullable Integer teamId) {
            this.teamId = teamId;
            return this;
        }
        @CustomType.Setter
        public Builder userId(@Nullable Integer userId) {
            this.userId = userId;
            return this;
        }
        public DataSourcePermissionPermission build() {
            final var o = new DataSourcePermissionPermission();
            o.permission = permission;
            o.teamId = teamId;
            o.userId = userId;
            return o;
        }
    }
}
