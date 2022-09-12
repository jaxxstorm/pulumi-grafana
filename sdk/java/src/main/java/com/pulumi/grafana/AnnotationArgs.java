// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AnnotationArgs extends com.pulumi.resources.ResourceArgs {

    public static final AnnotationArgs Empty = new AnnotationArgs();

    /**
     * The ID of the dashboard on which to create the annotation.
     * 
     */
    @Import(name="dashboardId")
    private @Nullable Output<Integer> dashboardId;

    /**
     * @return The ID of the dashboard on which to create the annotation.
     * 
     */
    public Optional<Output<Integer>> dashboardId() {
        return Optional.ofNullable(this.dashboardId);
    }

    /**
     * The ID of the dashboard panel on which to create the annotation.
     * 
     */
    @Import(name="panelId")
    private @Nullable Output<Integer> panelId;

    /**
     * @return The ID of the dashboard panel on which to create the annotation.
     * 
     */
    public Optional<Output<Integer>> panelId() {
        return Optional.ofNullable(this.panelId);
    }

    /**
     * The tags to associate with the annotation.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags to associate with the annotation.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The text to associate with the annotation.
     * 
     */
    @Import(name="text", required=true)
    private Output<String> text;

    /**
     * @return The text to associate with the annotation.
     * 
     */
    public Output<String> text() {
        return this.text;
    }

    /**
     * The RFC 3339-formatted time string indicating the annotation&#39;s time.
     * 
     */
    @Import(name="time")
    private @Nullable Output<String> time;

    /**
     * @return The RFC 3339-formatted time string indicating the annotation&#39;s time.
     * 
     */
    public Optional<Output<String>> time() {
        return Optional.ofNullable(this.time);
    }

    /**
     * The RFC 3339-formatted time string indicating the annotation&#39;s end time.
     * 
     */
    @Import(name="timeEnd")
    private @Nullable Output<String> timeEnd;

    /**
     * @return The RFC 3339-formatted time string indicating the annotation&#39;s end time.
     * 
     */
    public Optional<Output<String>> timeEnd() {
        return Optional.ofNullable(this.timeEnd);
    }

    private AnnotationArgs() {}

    private AnnotationArgs(AnnotationArgs $) {
        this.dashboardId = $.dashboardId;
        this.panelId = $.panelId;
        this.tags = $.tags;
        this.text = $.text;
        this.time = $.time;
        this.timeEnd = $.timeEnd;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AnnotationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AnnotationArgs $;

        public Builder() {
            $ = new AnnotationArgs();
        }

        public Builder(AnnotationArgs defaults) {
            $ = new AnnotationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dashboardId The ID of the dashboard on which to create the annotation.
         * 
         * @return builder
         * 
         */
        public Builder dashboardId(@Nullable Output<Integer> dashboardId) {
            $.dashboardId = dashboardId;
            return this;
        }

        /**
         * @param dashboardId The ID of the dashboard on which to create the annotation.
         * 
         * @return builder
         * 
         */
        public Builder dashboardId(Integer dashboardId) {
            return dashboardId(Output.of(dashboardId));
        }

        /**
         * @param panelId The ID of the dashboard panel on which to create the annotation.
         * 
         * @return builder
         * 
         */
        public Builder panelId(@Nullable Output<Integer> panelId) {
            $.panelId = panelId;
            return this;
        }

        /**
         * @param panelId The ID of the dashboard panel on which to create the annotation.
         * 
         * @return builder
         * 
         */
        public Builder panelId(Integer panelId) {
            return panelId(Output.of(panelId));
        }

        /**
         * @param tags The tags to associate with the annotation.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags to associate with the annotation.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags to associate with the annotation.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param text The text to associate with the annotation.
         * 
         * @return builder
         * 
         */
        public Builder text(Output<String> text) {
            $.text = text;
            return this;
        }

        /**
         * @param text The text to associate with the annotation.
         * 
         * @return builder
         * 
         */
        public Builder text(String text) {
            return text(Output.of(text));
        }

        /**
         * @param time The RFC 3339-formatted time string indicating the annotation&#39;s time.
         * 
         * @return builder
         * 
         */
        public Builder time(@Nullable Output<String> time) {
            $.time = time;
            return this;
        }

        /**
         * @param time The RFC 3339-formatted time string indicating the annotation&#39;s time.
         * 
         * @return builder
         * 
         */
        public Builder time(String time) {
            return time(Output.of(time));
        }

        /**
         * @param timeEnd The RFC 3339-formatted time string indicating the annotation&#39;s end time.
         * 
         * @return builder
         * 
         */
        public Builder timeEnd(@Nullable Output<String> timeEnd) {
            $.timeEnd = timeEnd;
            return this;
        }

        /**
         * @param timeEnd The RFC 3339-formatted time string indicating the annotation&#39;s end time.
         * 
         * @return builder
         * 
         */
        public Builder timeEnd(String timeEnd) {
            return timeEnd(Output.of(timeEnd));
        }

        public AnnotationArgs build() {
            $.text = Objects.requireNonNull($.text, "expected parameter 'text' to be non-null");
            return $;
        }
    }

}
