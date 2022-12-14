// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSyntheticMonitoringProbesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSyntheticMonitoringProbesArgs Empty = new GetSyntheticMonitoringProbesArgs();

    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     * 
     */
    @Import(name="filterDeprecated")
    private @Nullable Output<Boolean> filterDeprecated;

    /**
     * @return If true, only probes that are not deprecated will be returned. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> filterDeprecated() {
        return Optional.ofNullable(this.filterDeprecated);
    }

    private GetSyntheticMonitoringProbesArgs() {}

    private GetSyntheticMonitoringProbesArgs(GetSyntheticMonitoringProbesArgs $) {
        this.filterDeprecated = $.filterDeprecated;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSyntheticMonitoringProbesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSyntheticMonitoringProbesArgs $;

        public Builder() {
            $ = new GetSyntheticMonitoringProbesArgs();
        }

        public Builder(GetSyntheticMonitoringProbesArgs defaults) {
            $ = new GetSyntheticMonitoringProbesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filterDeprecated If true, only probes that are not deprecated will be returned. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder filterDeprecated(@Nullable Output<Boolean> filterDeprecated) {
            $.filterDeprecated = filterDeprecated;
            return this;
        }

        /**
         * @param filterDeprecated If true, only probes that are not deprecated will be returned. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder filterDeprecated(Boolean filterDeprecated) {
            return filterDeprecated(Output.of(filterDeprecated));
        }

        public GetSyntheticMonitoringProbesArgs build() {
            return $;
        }
    }

}
