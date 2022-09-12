// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDashboardResult {
    /**
     * @return The complete dashboard model JSON.
     * 
     */
    private String configJson;
    /**
     * @return The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
     * 
     */
    private @Nullable Integer dashboardId;
    /**
     * @return The numerical ID of the folder where the Grafana dashboard is found.
     * 
     */
    private Integer folder;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Whether or not the Grafana dashboard is starred. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you’re interested in.
     * 
     */
    private Boolean isStarred;
    /**
     * @return URL slug of the dashboard (deprecated).
     * 
     */
    private String slug;
    /**
     * @return The title of the Grafana dashboard.
     * 
     */
    private String title;
    /**
     * @return The uid of the Grafana dashboard. Specify either this or `dashboard_id`. Defaults to ``.
     * 
     */
    private @Nullable String uid;
    /**
     * @return The full URL of the dashboard.
     * 
     */
    private String url;
    /**
     * @return The numerical version of the Grafana dashboard.
     * 
     */
    private Integer version;

    private GetDashboardResult() {}
    /**
     * @return The complete dashboard model JSON.
     * 
     */
    public String configJson() {
        return this.configJson;
    }
    /**
     * @return The numerical ID of the Grafana dashboard. Specify either this or `uid`. Defaults to `-1`.
     * 
     */
    public Optional<Integer> dashboardId() {
        return Optional.ofNullable(this.dashboardId);
    }
    /**
     * @return The numerical ID of the folder where the Grafana dashboard is found.
     * 
     */
    public Integer folder() {
        return this.folder;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether or not the Grafana dashboard is starred. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you’re interested in.
     * 
     */
    public Boolean isStarred() {
        return this.isStarred;
    }
    /**
     * @return URL slug of the dashboard (deprecated).
     * 
     */
    public String slug() {
        return this.slug;
    }
    /**
     * @return The title of the Grafana dashboard.
     * 
     */
    public String title() {
        return this.title;
    }
    /**
     * @return The uid of the Grafana dashboard. Specify either this or `dashboard_id`. Defaults to ``.
     * 
     */
    public Optional<String> uid() {
        return Optional.ofNullable(this.uid);
    }
    /**
     * @return The full URL of the dashboard.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return The numerical version of the Grafana dashboard.
     * 
     */
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDashboardResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String configJson;
        private @Nullable Integer dashboardId;
        private Integer folder;
        private String id;
        private Boolean isStarred;
        private String slug;
        private String title;
        private @Nullable String uid;
        private String url;
        private Integer version;
        public Builder() {}
        public Builder(GetDashboardResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configJson = defaults.configJson;
    	      this.dashboardId = defaults.dashboardId;
    	      this.folder = defaults.folder;
    	      this.id = defaults.id;
    	      this.isStarred = defaults.isStarred;
    	      this.slug = defaults.slug;
    	      this.title = defaults.title;
    	      this.uid = defaults.uid;
    	      this.url = defaults.url;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder configJson(String configJson) {
            this.configJson = Objects.requireNonNull(configJson);
            return this;
        }
        @CustomType.Setter
        public Builder dashboardId(@Nullable Integer dashboardId) {
            this.dashboardId = dashboardId;
            return this;
        }
        @CustomType.Setter
        public Builder folder(Integer folder) {
            this.folder = Objects.requireNonNull(folder);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isStarred(Boolean isStarred) {
            this.isStarred = Objects.requireNonNull(isStarred);
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        @CustomType.Setter
        public Builder uid(@Nullable String uid) {
            this.uid = uid;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetDashboardResult build() {
            final var o = new GetDashboardResult();
            o.configJson = configJson;
            o.dashboardId = dashboardId;
            o.folder = folder;
            o.id = id;
            o.isStarred = isStarred;
            o.slug = slug;
            o.title = title;
            o.uid = uid;
            o.url = url;
            o.version = version;
            return o;
        }
    }
}
