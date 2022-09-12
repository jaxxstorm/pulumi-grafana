// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetSyntheticMonitoringProbeResult {
    /**
     * @return The ID of the probe.
     * 
     */
    private String id;
    /**
     * @return Custom labels to be included with collected metrics and logs.
     * 
     */
    private Map<String,String> labels;
    /**
     * @return Latitude coordinates.
     * 
     */
    private Double latitude;
    /**
     * @return Longitude coordinates.
     * 
     */
    private Double longitude;
    /**
     * @return Name of the probe.
     * 
     */
    private String name;
    /**
     * @return Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`.
     * 
     */
    private Boolean public_;
    /**
     * @return Region of the probe.
     * 
     */
    private String region;
    /**
     * @return The tenant ID of the probe.
     * 
     */
    private Integer tenantId;

    private GetSyntheticMonitoringProbeResult() {}
    /**
     * @return The ID of the probe.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Custom labels to be included with collected metrics and logs.
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }
    /**
     * @return Latitude coordinates.
     * 
     */
    public Double latitude() {
        return this.latitude;
    }
    /**
     * @return Longitude coordinates.
     * 
     */
    public Double longitude() {
        return this.longitude;
    }
    /**
     * @return Name of the probe.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`.
     * 
     */
    public Boolean public_() {
        return this.public_;
    }
    /**
     * @return Region of the probe.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The tenant ID of the probe.
     * 
     */
    public Integer tenantId() {
        return this.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSyntheticMonitoringProbeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Map<String,String> labels;
        private Double latitude;
        private Double longitude;
        private String name;
        private Boolean public_;
        private String region;
        private Integer tenantId;
        public Builder() {}
        public Builder(GetSyntheticMonitoringProbeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.labels = defaults.labels;
    	      this.latitude = defaults.latitude;
    	      this.longitude = defaults.longitude;
    	      this.name = defaults.name;
    	      this.public_ = defaults.public_;
    	      this.region = defaults.region;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        @CustomType.Setter
        public Builder latitude(Double latitude) {
            this.latitude = Objects.requireNonNull(latitude);
            return this;
        }
        @CustomType.Setter
        public Builder longitude(Double longitude) {
            this.longitude = Objects.requireNonNull(longitude);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter("public")
        public Builder public_(Boolean public_) {
            this.public_ = Objects.requireNonNull(public_);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder tenantId(Integer tenantId) {
            this.tenantId = Objects.requireNonNull(tenantId);
            return this;
        }
        public GetSyntheticMonitoringProbeResult build() {
            final var o = new GetSyntheticMonitoringProbeResult();
            o.id = id;
            o.labels = labels;
            o.latitude = latitude;
            o.longitude = longitude;
            o.name = name;
            o.public_ = public_;
            o.region = region;
            o.tenantId = tenantId;
            return o;
        }
    }
}
