// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCloudStackResult {
    /**
     * @return Name of the Alertmanager instance configured for this stack.
     * 
     */
    private String alertmanagerName;
    /**
     * @return Status of the Alertmanager instance configured for this stack.
     * 
     */
    private String alertmanagerStatus;
    /**
     * @return Base URL of the Alertmanager instance configured for this stack.
     * 
     */
    private String alertmanagerUrl;
    /**
     * @return User ID of the Alertmanager instance configured for this stack.
     * 
     */
    private Integer alertmanagerUserId;
    /**
     * @return Description of stack.
     * 
     */
    private String description;
    private String graphiteName;
    private String graphiteStatus;
    private String graphiteUrl;
    private Integer graphiteUserId;
    /**
     * @return The stack id assigned to this stack by Grafana.
     * 
     */
    private String id;
    private String logsName;
    private String logsStatus;
    private String logsUrl;
    private Integer logsUserId;
    /**
     * @return Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
     * 
     */
    private String name;
    /**
     * @return Organization id to assign to this stack.
     * 
     */
    private Integer orgId;
    /**
     * @return Organization name to assign to this stack.
     * 
     */
    private String orgName;
    /**
     * @return Organization slug to assign to this stack.
     * 
     */
    private String orgSlug;
    /**
     * @return Prometheus name for this instance.
     * 
     */
    private String prometheusName;
    /**
     * @return Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
     * 
     */
    private String prometheusRemoteEndpoint;
    /**
     * @return Use this URL to send prometheus metrics to Grafana cloud
     * 
     */
    private String prometheusRemoteWriteEndpoint;
    /**
     * @return Prometheus status for this instance.
     * 
     */
    private String prometheusStatus;
    /**
     * @return Prometheus url for this instance.
     * 
     */
    private String prometheusUrl;
    /**
     * @return Prometheus user ID. Used for e.g. remote_write.
     * 
     */
    private Integer prometheusUserId;
    /**
     * @return The region this stack is deployed to.
     * 
     */
    private String regionSlug;
    /**
     * @return Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net&#34;.
     * 
     */
    private String slug;
    /**
     * @return Status of the stack.
     * 
     */
    private String status;
    private String tracesName;
    private String tracesStatus;
    private String tracesUrl;
    private Integer tracesUserId;
    /**
     * @return Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     * 
     */
    private String url;

    private GetCloudStackResult() {}
    /**
     * @return Name of the Alertmanager instance configured for this stack.
     * 
     */
    public String alertmanagerName() {
        return this.alertmanagerName;
    }
    /**
     * @return Status of the Alertmanager instance configured for this stack.
     * 
     */
    public String alertmanagerStatus() {
        return this.alertmanagerStatus;
    }
    /**
     * @return Base URL of the Alertmanager instance configured for this stack.
     * 
     */
    public String alertmanagerUrl() {
        return this.alertmanagerUrl;
    }
    /**
     * @return User ID of the Alertmanager instance configured for this stack.
     * 
     */
    public Integer alertmanagerUserId() {
        return this.alertmanagerUserId;
    }
    /**
     * @return Description of stack.
     * 
     */
    public String description() {
        return this.description;
    }
    public String graphiteName() {
        return this.graphiteName;
    }
    public String graphiteStatus() {
        return this.graphiteStatus;
    }
    public String graphiteUrl() {
        return this.graphiteUrl;
    }
    public Integer graphiteUserId() {
        return this.graphiteUserId;
    }
    /**
     * @return The stack id assigned to this stack by Grafana.
     * 
     */
    public String id() {
        return this.id;
    }
    public String logsName() {
        return this.logsName;
    }
    public String logsStatus() {
        return this.logsStatus;
    }
    public String logsUrl() {
        return this.logsUrl;
    }
    public Integer logsUserId() {
        return this.logsUserId;
    }
    /**
     * @return Name of stack. Conventionally matches the url of the instance (e.g. “\n\n.grafana.net”).
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Organization id to assign to this stack.
     * 
     */
    public Integer orgId() {
        return this.orgId;
    }
    /**
     * @return Organization name to assign to this stack.
     * 
     */
    public String orgName() {
        return this.orgName;
    }
    /**
     * @return Organization slug to assign to this stack.
     * 
     */
    public String orgSlug() {
        return this.orgSlug;
    }
    /**
     * @return Prometheus name for this instance.
     * 
     */
    public String prometheusName() {
        return this.prometheusName;
    }
    /**
     * @return Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
     * 
     */
    public String prometheusRemoteEndpoint() {
        return this.prometheusRemoteEndpoint;
    }
    /**
     * @return Use this URL to send prometheus metrics to Grafana cloud
     * 
     */
    public String prometheusRemoteWriteEndpoint() {
        return this.prometheusRemoteWriteEndpoint;
    }
    /**
     * @return Prometheus status for this instance.
     * 
     */
    public String prometheusStatus() {
        return this.prometheusStatus;
    }
    /**
     * @return Prometheus url for this instance.
     * 
     */
    public String prometheusUrl() {
        return this.prometheusUrl;
    }
    /**
     * @return Prometheus user ID. Used for e.g. remote_write.
     * 
     */
    public Integer prometheusUserId() {
        return this.prometheusUserId;
    }
    /**
     * @return The region this stack is deployed to.
     * 
     */
    public String regionSlug() {
        return this.regionSlug;
    }
    /**
     * @return Subdomain that the Grafana instance will be available at (i.e. setting slug to “\n\n” will make the instance
     * available at “https://\n\n.grafana.net&#34;.
     * 
     */
    public String slug() {
        return this.slug;
    }
    /**
     * @return Status of the stack.
     * 
     */
    public String status() {
        return this.status;
    }
    public String tracesName() {
        return this.tracesName;
    }
    public String tracesStatus() {
        return this.tracesStatus;
    }
    public String tracesUrl() {
        return this.tracesUrl;
    }
    public Integer tracesUserId() {
        return this.tracesUserId;
    }
    /**
     * @return Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudStackResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String alertmanagerName;
        private String alertmanagerStatus;
        private String alertmanagerUrl;
        private Integer alertmanagerUserId;
        private String description;
        private String graphiteName;
        private String graphiteStatus;
        private String graphiteUrl;
        private Integer graphiteUserId;
        private String id;
        private String logsName;
        private String logsStatus;
        private String logsUrl;
        private Integer logsUserId;
        private String name;
        private Integer orgId;
        private String orgName;
        private String orgSlug;
        private String prometheusName;
        private String prometheusRemoteEndpoint;
        private String prometheusRemoteWriteEndpoint;
        private String prometheusStatus;
        private String prometheusUrl;
        private Integer prometheusUserId;
        private String regionSlug;
        private String slug;
        private String status;
        private String tracesName;
        private String tracesStatus;
        private String tracesUrl;
        private Integer tracesUserId;
        private String url;
        public Builder() {}
        public Builder(GetCloudStackResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alertmanagerName = defaults.alertmanagerName;
    	      this.alertmanagerStatus = defaults.alertmanagerStatus;
    	      this.alertmanagerUrl = defaults.alertmanagerUrl;
    	      this.alertmanagerUserId = defaults.alertmanagerUserId;
    	      this.description = defaults.description;
    	      this.graphiteName = defaults.graphiteName;
    	      this.graphiteStatus = defaults.graphiteStatus;
    	      this.graphiteUrl = defaults.graphiteUrl;
    	      this.graphiteUserId = defaults.graphiteUserId;
    	      this.id = defaults.id;
    	      this.logsName = defaults.logsName;
    	      this.logsStatus = defaults.logsStatus;
    	      this.logsUrl = defaults.logsUrl;
    	      this.logsUserId = defaults.logsUserId;
    	      this.name = defaults.name;
    	      this.orgId = defaults.orgId;
    	      this.orgName = defaults.orgName;
    	      this.orgSlug = defaults.orgSlug;
    	      this.prometheusName = defaults.prometheusName;
    	      this.prometheusRemoteEndpoint = defaults.prometheusRemoteEndpoint;
    	      this.prometheusRemoteWriteEndpoint = defaults.prometheusRemoteWriteEndpoint;
    	      this.prometheusStatus = defaults.prometheusStatus;
    	      this.prometheusUrl = defaults.prometheusUrl;
    	      this.prometheusUserId = defaults.prometheusUserId;
    	      this.regionSlug = defaults.regionSlug;
    	      this.slug = defaults.slug;
    	      this.status = defaults.status;
    	      this.tracesName = defaults.tracesName;
    	      this.tracesStatus = defaults.tracesStatus;
    	      this.tracesUrl = defaults.tracesUrl;
    	      this.tracesUserId = defaults.tracesUserId;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder alertmanagerName(String alertmanagerName) {
            this.alertmanagerName = Objects.requireNonNull(alertmanagerName);
            return this;
        }
        @CustomType.Setter
        public Builder alertmanagerStatus(String alertmanagerStatus) {
            this.alertmanagerStatus = Objects.requireNonNull(alertmanagerStatus);
            return this;
        }
        @CustomType.Setter
        public Builder alertmanagerUrl(String alertmanagerUrl) {
            this.alertmanagerUrl = Objects.requireNonNull(alertmanagerUrl);
            return this;
        }
        @CustomType.Setter
        public Builder alertmanagerUserId(Integer alertmanagerUserId) {
            this.alertmanagerUserId = Objects.requireNonNull(alertmanagerUserId);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder graphiteName(String graphiteName) {
            this.graphiteName = Objects.requireNonNull(graphiteName);
            return this;
        }
        @CustomType.Setter
        public Builder graphiteStatus(String graphiteStatus) {
            this.graphiteStatus = Objects.requireNonNull(graphiteStatus);
            return this;
        }
        @CustomType.Setter
        public Builder graphiteUrl(String graphiteUrl) {
            this.graphiteUrl = Objects.requireNonNull(graphiteUrl);
            return this;
        }
        @CustomType.Setter
        public Builder graphiteUserId(Integer graphiteUserId) {
            this.graphiteUserId = Objects.requireNonNull(graphiteUserId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logsName(String logsName) {
            this.logsName = Objects.requireNonNull(logsName);
            return this;
        }
        @CustomType.Setter
        public Builder logsStatus(String logsStatus) {
            this.logsStatus = Objects.requireNonNull(logsStatus);
            return this;
        }
        @CustomType.Setter
        public Builder logsUrl(String logsUrl) {
            this.logsUrl = Objects.requireNonNull(logsUrl);
            return this;
        }
        @CustomType.Setter
        public Builder logsUserId(Integer logsUserId) {
            this.logsUserId = Objects.requireNonNull(logsUserId);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder orgId(Integer orgId) {
            this.orgId = Objects.requireNonNull(orgId);
            return this;
        }
        @CustomType.Setter
        public Builder orgName(String orgName) {
            this.orgName = Objects.requireNonNull(orgName);
            return this;
        }
        @CustomType.Setter
        public Builder orgSlug(String orgSlug) {
            this.orgSlug = Objects.requireNonNull(orgSlug);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusName(String prometheusName) {
            this.prometheusName = Objects.requireNonNull(prometheusName);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusRemoteEndpoint(String prometheusRemoteEndpoint) {
            this.prometheusRemoteEndpoint = Objects.requireNonNull(prometheusRemoteEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusRemoteWriteEndpoint(String prometheusRemoteWriteEndpoint) {
            this.prometheusRemoteWriteEndpoint = Objects.requireNonNull(prometheusRemoteWriteEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusStatus(String prometheusStatus) {
            this.prometheusStatus = Objects.requireNonNull(prometheusStatus);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusUrl(String prometheusUrl) {
            this.prometheusUrl = Objects.requireNonNull(prometheusUrl);
            return this;
        }
        @CustomType.Setter
        public Builder prometheusUserId(Integer prometheusUserId) {
            this.prometheusUserId = Objects.requireNonNull(prometheusUserId);
            return this;
        }
        @CustomType.Setter
        public Builder regionSlug(String regionSlug) {
            this.regionSlug = Objects.requireNonNull(regionSlug);
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tracesName(String tracesName) {
            this.tracesName = Objects.requireNonNull(tracesName);
            return this;
        }
        @CustomType.Setter
        public Builder tracesStatus(String tracesStatus) {
            this.tracesStatus = Objects.requireNonNull(tracesStatus);
            return this;
        }
        @CustomType.Setter
        public Builder tracesUrl(String tracesUrl) {
            this.tracesUrl = Objects.requireNonNull(tracesUrl);
            return this;
        }
        @CustomType.Setter
        public Builder tracesUserId(Integer tracesUserId) {
            this.tracesUserId = Objects.requireNonNull(tracesUserId);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public GetCloudStackResult build() {
            final var o = new GetCloudStackResult();
            o.alertmanagerName = alertmanagerName;
            o.alertmanagerStatus = alertmanagerStatus;
            o.alertmanagerUrl = alertmanagerUrl;
            o.alertmanagerUserId = alertmanagerUserId;
            o.description = description;
            o.graphiteName = graphiteName;
            o.graphiteStatus = graphiteStatus;
            o.graphiteUrl = graphiteUrl;
            o.graphiteUserId = graphiteUserId;
            o.id = id;
            o.logsName = logsName;
            o.logsStatus = logsStatus;
            o.logsUrl = logsUrl;
            o.logsUserId = logsUserId;
            o.name = name;
            o.orgId = orgId;
            o.orgName = orgName;
            o.orgSlug = orgSlug;
            o.prometheusName = prometheusName;
            o.prometheusRemoteEndpoint = prometheusRemoteEndpoint;
            o.prometheusRemoteWriteEndpoint = prometheusRemoteWriteEndpoint;
            o.prometheusStatus = prometheusStatus;
            o.prometheusUrl = prometheusUrl;
            o.prometheusUserId = prometheusUserId;
            o.regionSlug = regionSlug;
            o.slug = slug;
            o.status = status;
            o.tracesName = tracesName;
            o.tracesStatus = tracesStatus;
            o.tracesUrl = tracesUrl;
            o.tracesUserId = tracesUserId;
            o.url = url;
            return o;
        }
    }
}
