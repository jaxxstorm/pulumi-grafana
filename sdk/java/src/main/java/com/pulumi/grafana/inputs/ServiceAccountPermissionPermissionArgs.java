// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceAccountPermissionPermissionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAccountPermissionPermissionArgs Empty = new ServiceAccountPermissionPermissionArgs();

    /**
     * Permission to associate with item. Must be `Edit` or `Admin`.
     * 
     */
    @Import(name="permission", required=true)
    private Output<String> permission;

    /**
     * @return Permission to associate with item. Must be `Edit` or `Admin`.
     * 
     */
    public Output<String> permission() {
        return this.permission;
    }

    /**
     * ID of the team to manage permissions for. Specify either this or `user_id`. Defaults to `0`.
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<String> teamId;

    /**
     * @return ID of the team to manage permissions for. Specify either this or `user_id`. Defaults to `0`.
     * 
     */
    public Optional<Output<String>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    /**
     * ID of the user to manage permissions for. Specify either this or `team_id`. Defaults to `0`.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return ID of the user to manage permissions for. Specify either this or `team_id`. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private ServiceAccountPermissionPermissionArgs() {}

    private ServiceAccountPermissionPermissionArgs(ServiceAccountPermissionPermissionArgs $) {
        this.permission = $.permission;
        this.teamId = $.teamId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAccountPermissionPermissionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAccountPermissionPermissionArgs $;

        public Builder() {
            $ = new ServiceAccountPermissionPermissionArgs();
        }

        public Builder(ServiceAccountPermissionPermissionArgs defaults) {
            $ = new ServiceAccountPermissionPermissionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permission Permission to associate with item. Must be `Edit` or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder permission(Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Permission to associate with item. Must be `Edit` or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param teamId ID of the team to manage permissions for. Specify either this or `user_id`. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId ID of the team to manage permissions for. Specify either this or `user_id`. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        /**
         * @param userId ID of the user to manage permissions for. Specify either this or `team_id`. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId ID of the user to manage permissions for. Specify either this or `team_id`. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public ServiceAccountPermissionPermissionArgs build() {
            $.permission = Objects.requireNonNull($.permission, "expected parameter 'permission' to be non-null");
            return $;
        }
    }

}
