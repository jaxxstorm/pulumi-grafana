// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.grafana.inputs.ServiceAccountPermissionPermissionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceAccountPermissionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAccountPermissionArgs Empty = new ServiceAccountPermissionArgs();

    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return The Organization ID. If not set, the Org ID defined in the provider block will be used.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     * 
     */
    @Import(name="permissions", required=true)
    private Output<List<ServiceAccountPermissionPermissionArgs>> permissions;

    /**
     * @return The permission items to add/update. Items that are omitted from the list will be removed.
     * 
     */
    public Output<List<ServiceAccountPermissionPermissionArgs>> permissions() {
        return this.permissions;
    }

    /**
     * The id of the service account.
     * 
     */
    @Import(name="serviceAccountId", required=true)
    private Output<String> serviceAccountId;

    /**
     * @return The id of the service account.
     * 
     */
    public Output<String> serviceAccountId() {
        return this.serviceAccountId;
    }

    private ServiceAccountPermissionArgs() {}

    private ServiceAccountPermissionArgs(ServiceAccountPermissionArgs $) {
        this.orgId = $.orgId;
        this.permissions = $.permissions;
        this.serviceAccountId = $.serviceAccountId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAccountPermissionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAccountPermissionArgs $;

        public Builder() {
            $ = new ServiceAccountPermissionArgs();
        }

        public Builder(ServiceAccountPermissionArgs defaults) {
            $ = new ServiceAccountPermissionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param orgId The Organization ID. If not set, the Org ID defined in the provider block will be used.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId The Organization ID. If not set, the Org ID defined in the provider block will be used.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param permissions The permission items to add/update. Items that are omitted from the list will be removed.
         * 
         * @return builder
         * 
         */
        public Builder permissions(Output<List<ServiceAccountPermissionPermissionArgs>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions The permission items to add/update. Items that are omitted from the list will be removed.
         * 
         * @return builder
         * 
         */
        public Builder permissions(List<ServiceAccountPermissionPermissionArgs> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param permissions The permission items to add/update. Items that are omitted from the list will be removed.
         * 
         * @return builder
         * 
         */
        public Builder permissions(ServiceAccountPermissionPermissionArgs... permissions) {
            return permissions(List.of(permissions));
        }

        /**
         * @param serviceAccountId The id of the service account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(Output<String> serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        /**
         * @param serviceAccountId The id of the service account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            return serviceAccountId(Output.of(serviceAccountId));
        }

        public ServiceAccountPermissionArgs build() {
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            $.serviceAccountId = Objects.requireNonNull($.serviceAccountId, "expected parameter 'serviceAccountId' to be non-null");
            return $;
        }
    }

}
