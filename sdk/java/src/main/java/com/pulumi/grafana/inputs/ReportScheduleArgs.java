// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReportScheduleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReportScheduleArgs Empty = new ReportScheduleArgs();

    /**
     * Custom interval of the report.
     * **Note:** This field is only available when frequency is set to `custom`.
     * 
     */
    @Import(name="customInterval")
    private @Nullable Output<String> customInterval;

    /**
     * @return Custom interval of the report.
     * **Note:** This field is only available when frequency is set to `custom`.
     * 
     */
    public Optional<Output<String>> customInterval() {
        return Optional.ofNullable(this.customInterval);
    }

    /**
     * End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana.
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<String> endTime;

    /**
     * @return End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana.
     * 
     */
    public Optional<Output<String>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * Frequency of the report. One of `never`, `once`, `hourly`, `daily`, `weekly`, `monthly` or `custom`.
     * 
     */
    @Import(name="frequency", required=true)
    private Output<String> frequency;

    /**
     * @return Frequency of the report. One of `never`, `once`, `hourly`, `daily`, `weekly`, `monthly` or `custom`.
     * 
     */
    public Output<String> frequency() {
        return this.frequency;
    }

    /**
     * Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana.
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<String> startTime;

    /**
     * @return Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana.
     * 
     */
    public Optional<Output<String>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * Whether to send the report only on work days. Defaults to `false`.
     * 
     */
    @Import(name="workdaysOnly")
    private @Nullable Output<Boolean> workdaysOnly;

    /**
     * @return Whether to send the report only on work days. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> workdaysOnly() {
        return Optional.ofNullable(this.workdaysOnly);
    }

    private ReportScheduleArgs() {}

    private ReportScheduleArgs(ReportScheduleArgs $) {
        this.customInterval = $.customInterval;
        this.endTime = $.endTime;
        this.frequency = $.frequency;
        this.startTime = $.startTime;
        this.workdaysOnly = $.workdaysOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReportScheduleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReportScheduleArgs $;

        public Builder() {
            $ = new ReportScheduleArgs();
        }

        public Builder(ReportScheduleArgs defaults) {
            $ = new ReportScheduleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customInterval Custom interval of the report.
         * **Note:** This field is only available when frequency is set to `custom`.
         * 
         * @return builder
         * 
         */
        public Builder customInterval(@Nullable Output<String> customInterval) {
            $.customInterval = customInterval;
            return this;
        }

        /**
         * @param customInterval Custom interval of the report.
         * **Note:** This field is only available when frequency is set to `custom`.
         * 
         * @return builder
         * 
         */
        public Builder customInterval(String customInterval) {
            return customInterval(Output.of(customInterval));
        }

        /**
         * @param endTime End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<String> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder endTime(String endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param frequency Frequency of the report. One of `never`, `once`, `hourly`, `daily`, `weekly`, `monthly` or `custom`.
         * 
         * @return builder
         * 
         */
        public Builder frequency(Output<String> frequency) {
            $.frequency = frequency;
            return this;
        }

        /**
         * @param frequency Frequency of the report. One of `never`, `once`, `hourly`, `daily`, `weekly`, `monthly` or `custom`.
         * 
         * @return builder
         * 
         */
        public Builder frequency(String frequency) {
            return frequency(Output.of(frequency));
        }

        /**
         * @param startTime Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<String> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder startTime(String startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param workdaysOnly Whether to send the report only on work days. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder workdaysOnly(@Nullable Output<Boolean> workdaysOnly) {
            $.workdaysOnly = workdaysOnly;
            return this;
        }

        /**
         * @param workdaysOnly Whether to send the report only on work days. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder workdaysOnly(Boolean workdaysOnly) {
            return workdaysOnly(Output.of(workdaysOnly));
        }

        public ReportScheduleArgs build() {
            $.frequency = Objects.requireNonNull($.frequency, "expected parameter 'frequency' to be non-null");
            return $;
        }
    }

}
