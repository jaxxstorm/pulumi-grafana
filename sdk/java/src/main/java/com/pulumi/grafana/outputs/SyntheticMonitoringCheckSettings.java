// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsDns;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsHttp;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsPing;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsTcp;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsTraceroute;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettings {
    /**
     * @return Settings for DNS check. The target must be a valid hostname (or IP address for `PTR` records).
     * 
     */
    private @Nullable SyntheticMonitoringCheckSettingsDns dns;
    /**
     * @return Settings for HTTP check. The target must be a URL (http or https).
     * 
     */
    private @Nullable SyntheticMonitoringCheckSettingsHttp http;
    /**
     * @return Settings for ping (ICMP) check. The target must be a valid hostname or IP address.
     * 
     */
    private @Nullable SyntheticMonitoringCheckSettingsPing ping;
    /**
     * @return Settings for TCP check. The target must be of the form `&lt;host&gt;:&lt;port&gt;`, where the host portion must be a valid hostname or IP address.
     * 
     */
    private @Nullable SyntheticMonitoringCheckSettingsTcp tcp;
    /**
     * @return Settings for traceroute check. The target must be a valid hostname or IP address
     * 
     */
    private @Nullable SyntheticMonitoringCheckSettingsTraceroute traceroute;

    private SyntheticMonitoringCheckSettings() {}
    /**
     * @return Settings for DNS check. The target must be a valid hostname (or IP address for `PTR` records).
     * 
     */
    public Optional<SyntheticMonitoringCheckSettingsDns> dns() {
        return Optional.ofNullable(this.dns);
    }
    /**
     * @return Settings for HTTP check. The target must be a URL (http or https).
     * 
     */
    public Optional<SyntheticMonitoringCheckSettingsHttp> http() {
        return Optional.ofNullable(this.http);
    }
    /**
     * @return Settings for ping (ICMP) check. The target must be a valid hostname or IP address.
     * 
     */
    public Optional<SyntheticMonitoringCheckSettingsPing> ping() {
        return Optional.ofNullable(this.ping);
    }
    /**
     * @return Settings for TCP check. The target must be of the form `&lt;host&gt;:&lt;port&gt;`, where the host portion must be a valid hostname or IP address.
     * 
     */
    public Optional<SyntheticMonitoringCheckSettingsTcp> tcp() {
        return Optional.ofNullable(this.tcp);
    }
    /**
     * @return Settings for traceroute check. The target must be a valid hostname or IP address
     * 
     */
    public Optional<SyntheticMonitoringCheckSettingsTraceroute> traceroute() {
        return Optional.ofNullable(this.traceroute);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable SyntheticMonitoringCheckSettingsDns dns;
        private @Nullable SyntheticMonitoringCheckSettingsHttp http;
        private @Nullable SyntheticMonitoringCheckSettingsPing ping;
        private @Nullable SyntheticMonitoringCheckSettingsTcp tcp;
        private @Nullable SyntheticMonitoringCheckSettingsTraceroute traceroute;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dns = defaults.dns;
    	      this.http = defaults.http;
    	      this.ping = defaults.ping;
    	      this.tcp = defaults.tcp;
    	      this.traceroute = defaults.traceroute;
        }

        @CustomType.Setter
        public Builder dns(@Nullable SyntheticMonitoringCheckSettingsDns dns) {
            this.dns = dns;
            return this;
        }
        @CustomType.Setter
        public Builder http(@Nullable SyntheticMonitoringCheckSettingsHttp http) {
            this.http = http;
            return this;
        }
        @CustomType.Setter
        public Builder ping(@Nullable SyntheticMonitoringCheckSettingsPing ping) {
            this.ping = ping;
            return this;
        }
        @CustomType.Setter
        public Builder tcp(@Nullable SyntheticMonitoringCheckSettingsTcp tcp) {
            this.tcp = tcp;
            return this;
        }
        @CustomType.Setter
        public Builder traceroute(@Nullable SyntheticMonitoringCheckSettingsTraceroute traceroute) {
            this.traceroute = traceroute;
            return this;
        }
        public SyntheticMonitoringCheckSettings build() {
            final var o = new SyntheticMonitoringCheckSettings();
            o.dns = dns;
            o.http = http;
            o.ping = ping;
            o.tcp = tcp;
            o.traceroute = traceroute;
            return o;
        }
    }
}
