// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOrganizationResult {
    /**
     * @return A list of email addresses corresponding to users given admin access to the organization.
     * 
     */
    private List<String> admins;
    /**
     * @return A list of email addresses corresponding to users given editor access to the organization.
     * 
     */
    private List<String> editors;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the Organization.
     * 
     */
    private String name;
    /**
     * @return A list of email addresses corresponding to users given viewer access to the organization.
     * 
     */
    private List<String> viewers;

    private GetOrganizationResult() {}
    /**
     * @return A list of email addresses corresponding to users given admin access to the organization.
     * 
     */
    public List<String> admins() {
        return this.admins;
    }
    /**
     * @return A list of email addresses corresponding to users given editor access to the organization.
     * 
     */
    public List<String> editors() {
        return this.editors;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Organization.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return A list of email addresses corresponding to users given viewer access to the organization.
     * 
     */
    public List<String> viewers() {
        return this.viewers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> admins;
        private List<String> editors;
        private String id;
        private String name;
        private List<String> viewers;
        public Builder() {}
        public Builder(GetOrganizationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.admins = defaults.admins;
    	      this.editors = defaults.editors;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.viewers = defaults.viewers;
        }

        @CustomType.Setter
        public Builder admins(List<String> admins) {
            this.admins = Objects.requireNonNull(admins);
            return this;
        }
        public Builder admins(String... admins) {
            return admins(List.of(admins));
        }
        @CustomType.Setter
        public Builder editors(List<String> editors) {
            this.editors = Objects.requireNonNull(editors);
            return this;
        }
        public Builder editors(String... editors) {
            return editors(List.of(editors));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder viewers(List<String> viewers) {
            this.viewers = Objects.requireNonNull(viewers);
            return this;
        }
        public Builder viewers(String... viewers) {
            return viewers(List.of(viewers));
        }
        public GetOrganizationResult build() {
            final var o = new GetOrganizationResult();
            o.admins = admins;
            o.editors = editors;
            o.id = id;
            o.name = name;
            o.viewers = viewers;
            return o;
        }
    }
}
