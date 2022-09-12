// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsDnsValidateAnswerRrs;
import com.pulumi.grafana.outputs.SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SyntheticMonitoringCheckSettingsDns {
    private @Nullable String ipVersion;
    private @Nullable Integer port;
    private @Nullable String protocol;
    private @Nullable String recordType;
    private @Nullable String server;
    private @Nullable String sourceIpAddress;
    private @Nullable List<String> validRCodes;
    private @Nullable List<SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr> validateAdditionalRrs;
    private @Nullable SyntheticMonitoringCheckSettingsDnsValidateAnswerRrs validateAnswerRrs;
    private @Nullable SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs validateAuthorityRrs;

    private SyntheticMonitoringCheckSettingsDns() {}
    public Optional<String> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }
    public Optional<String> recordType() {
        return Optional.ofNullable(this.recordType);
    }
    public Optional<String> server() {
        return Optional.ofNullable(this.server);
    }
    public Optional<String> sourceIpAddress() {
        return Optional.ofNullable(this.sourceIpAddress);
    }
    public List<String> validRCodes() {
        return this.validRCodes == null ? List.of() : this.validRCodes;
    }
    public List<SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr> validateAdditionalRrs() {
        return this.validateAdditionalRrs == null ? List.of() : this.validateAdditionalRrs;
    }
    public Optional<SyntheticMonitoringCheckSettingsDnsValidateAnswerRrs> validateAnswerRrs() {
        return Optional.ofNullable(this.validateAnswerRrs);
    }
    public Optional<SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs> validateAuthorityRrs() {
        return Optional.ofNullable(this.validateAuthorityRrs);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SyntheticMonitoringCheckSettingsDns defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipVersion;
        private @Nullable Integer port;
        private @Nullable String protocol;
        private @Nullable String recordType;
        private @Nullable String server;
        private @Nullable String sourceIpAddress;
        private @Nullable List<String> validRCodes;
        private @Nullable List<SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr> validateAdditionalRrs;
        private @Nullable SyntheticMonitoringCheckSettingsDnsValidateAnswerRrs validateAnswerRrs;
        private @Nullable SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs validateAuthorityRrs;
        public Builder() {}
        public Builder(SyntheticMonitoringCheckSettingsDns defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipVersion = defaults.ipVersion;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
    	      this.recordType = defaults.recordType;
    	      this.server = defaults.server;
    	      this.sourceIpAddress = defaults.sourceIpAddress;
    	      this.validRCodes = defaults.validRCodes;
    	      this.validateAdditionalRrs = defaults.validateAdditionalRrs;
    	      this.validateAnswerRrs = defaults.validateAnswerRrs;
    	      this.validateAuthorityRrs = defaults.validateAuthorityRrs;
        }

        @CustomType.Setter
        public Builder ipVersion(@Nullable String ipVersion) {
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder recordType(@Nullable String recordType) {
            this.recordType = recordType;
            return this;
        }
        @CustomType.Setter
        public Builder server(@Nullable String server) {
            this.server = server;
            return this;
        }
        @CustomType.Setter
        public Builder sourceIpAddress(@Nullable String sourceIpAddress) {
            this.sourceIpAddress = sourceIpAddress;
            return this;
        }
        @CustomType.Setter
        public Builder validRCodes(@Nullable List<String> validRCodes) {
            this.validRCodes = validRCodes;
            return this;
        }
        public Builder validRCodes(String... validRCodes) {
            return validRCodes(List.of(validRCodes));
        }
        @CustomType.Setter
        public Builder validateAdditionalRrs(@Nullable List<SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr> validateAdditionalRrs) {
            this.validateAdditionalRrs = validateAdditionalRrs;
            return this;
        }
        public Builder validateAdditionalRrs(SyntheticMonitoringCheckSettingsDnsValidateAdditionalRr... validateAdditionalRrs) {
            return validateAdditionalRrs(List.of(validateAdditionalRrs));
        }
        @CustomType.Setter
        public Builder validateAnswerRrs(@Nullable SyntheticMonitoringCheckSettingsDnsValidateAnswerRrs validateAnswerRrs) {
            this.validateAnswerRrs = validateAnswerRrs;
            return this;
        }
        @CustomType.Setter
        public Builder validateAuthorityRrs(@Nullable SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrs validateAuthorityRrs) {
            this.validateAuthorityRrs = validateAuthorityRrs;
            return this;
        }
        public SyntheticMonitoringCheckSettingsDns build() {
            final var o = new SyntheticMonitoringCheckSettingsDns();
            o.ipVersion = ipVersion;
            o.port = port;
            o.protocol = protocol;
            o.recordType = recordType;
            o.server = server;
            o.sourceIpAddress = sourceIpAddress;
            o.validRCodes = validRCodes;
            o.validateAdditionalRrs = validateAdditionalRrs;
            o.validateAnswerRrs = validateAnswerRrs;
            o.validateAuthorityRrs = validateAuthorityRrs;
            return o;
        }
    }
}
