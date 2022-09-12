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


public final class DashboardPermissionPermissionArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardPermissionPermissionArgs Empty = new DashboardPermissionPermissionArgs();

    /**
     * Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
     * 
     */
    @Import(name="permission", required=true)
    private Output<String> permission;

    /**
     * @return Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
     * 
     */
    public Output<String> permission() {
        return this.permission;
    }

    /**
     * Manage permissions for `Viewer` or `Editor` roles.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Manage permissions for `Viewer` or `Editor` roles.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * ID of the team to manage permissions for. Defaults to `0`.
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<Integer> teamId;

    /**
     * @return ID of the team to manage permissions for. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    /**
     * ID of the user to manage permissions for. Defaults to `0`.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return ID of the user to manage permissions for. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private DashboardPermissionPermissionArgs() {}

    private DashboardPermissionPermissionArgs(DashboardPermissionPermissionArgs $) {
        this.permission = $.permission;
        this.role = $.role;
        this.teamId = $.teamId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardPermissionPermissionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardPermissionPermissionArgs $;

        public Builder() {
            $ = new DashboardPermissionPermissionArgs();
        }

        public Builder(DashboardPermissionPermissionArgs defaults) {
            $ = new DashboardPermissionPermissionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permission Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder permission(Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param role Manage permissions for `Viewer` or `Editor` roles.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Manage permissions for `Viewer` or `Editor` roles.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param teamId ID of the team to manage permissions for. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<Integer> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId ID of the team to manage permissions for. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder teamId(Integer teamId) {
            return teamId(Output.of(teamId));
        }

        /**
         * @param userId ID of the user to manage permissions for. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId ID of the user to manage permissions for. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public DashboardPermissionPermissionArgs build() {
            $.permission = Objects.requireNonNull($.permission, "expected parameter 'permission' to be non-null");
            return $;
        }
    }

}
