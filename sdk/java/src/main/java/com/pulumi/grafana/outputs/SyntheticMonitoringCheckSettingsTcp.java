// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsTcpQueryResponse;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsTcpTlsConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettingsTcp {
    private @Nullable String ipVersion;
    private @Nullable List<SyntheticMonitoringCheckSettingsTcpQueryResponse> queryResponses;
    private @Nullable String sourceIpAddress;
    private @Nullable Boolean tls;
    private @Nullable SyntheticMonitoringCheckSettingsTcpTlsConfig tlsConfig;

    private SyntheticMonitoringCheckSettingsTcp() {}
    public Optional<String> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }
    public List<SyntheticMonitoringCheckSettingsTcpQueryResponse> queryResponses() {
        return this.queryResponses == null ? List.of() : this.queryResponses;
    }
    public Optional<String> sourceIpAddress() {
        return Optional.ofNullable(this.sourceIpAddress);
    }
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    public Optional<SyntheticMonitoringCheckSettingsTcpTlsConfig> tlsConfig() {
        return Optional.ofNullable(this.tlsConfig);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettingsTcp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipVersion;
        private @Nullable List<SyntheticMonitoringCheckSettingsTcpQueryResponse> queryResponses;
        private @Nullable String sourceIpAddress;
        private @Nullable Boolean tls;
        private @Nullable SyntheticMonitoringCheckSettingsTcpTlsConfig tlsConfig;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettingsTcp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipVersion = defaults.ipVersion;
    	      this.queryResponses = defaults.queryResponses;
    	      this.sourceIpAddress = defaults.sourceIpAddress;
    	      this.tls = defaults.tls;
    	      this.tlsConfig = defaults.tlsConfig;
        }

        @CustomType.Setter
        public Builder ipVersion(@Nullable String ipVersion) {
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder queryResponses(@Nullable List<SyntheticMonitoringCheckSettingsTcpQueryResponse> queryResponses) {
            this.queryResponses = queryResponses;
            return this;
        }
        public Builder queryResponses(SyntheticMonitoringCheckSettingsTcpQueryResponse... queryResponses) {
            return queryResponses(List.of(queryResponses));
        }
        @CustomType.Setter
        public Builder sourceIpAddress(@Nullable String sourceIpAddress) {
            this.sourceIpAddress = sourceIpAddress;
            return this;
        }
        @CustomType.Setter
        public Builder tls(@Nullable Boolean tls) {
            this.tls = tls;
            return this;
        }
        @CustomType.Setter
        public Builder tlsConfig(@Nullable SyntheticMonitoringCheckSettingsTcpTlsConfig tlsConfig) {
            this.tlsConfig = tlsConfig;
            return this;
        }
        public SyntheticMonitoringCheckSettingsTcp build() {
            final var o = new SyntheticMonitoringCheckSettingsTcp();
            o.ipVersion = ipVersion;
            o.queryResponses = queryResponses;
            o.sourceIpAddress = sourceIpAddress;
            o.tls = tls;
            o.tlsConfig = tlsConfig;
            return o;
        }
    }
}
