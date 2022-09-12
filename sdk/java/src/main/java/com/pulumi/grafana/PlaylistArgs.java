// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.grafana.inputs.PlaylistItemArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PlaylistArgs extends com.pulumi.resources.ResourceArgs {

    public static final PlaylistArgs Empty = new PlaylistArgs();

    @Import(name="interval", required=true)
    private Output<String> interval;

    public Output<String> interval() {
        return this.interval;
    }

    @Import(name="items", required=true)
    private Output<List<PlaylistItemArgs>> items;

    public Output<List<PlaylistItemArgs>> items() {
        return this.items;
    }

    /**
     * The name of the playlist.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the playlist.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private PlaylistArgs() {}

    private PlaylistArgs(PlaylistArgs $) {
        this.interval = $.interval;
        this.items = $.items;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PlaylistArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PlaylistArgs $;

        public Builder() {
            $ = new PlaylistArgs();
        }

        public Builder(PlaylistArgs defaults) {
            $ = new PlaylistArgs(Objects.requireNonNull(defaults));
        }

        public Builder interval(Output<String> interval) {
            $.interval = interval;
            return this;
        }

        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        public Builder items(Output<List<PlaylistItemArgs>> items) {
            $.items = items;
            return this;
        }

        public Builder items(List<PlaylistItemArgs> items) {
            return items(Output.of(items));
        }

        public Builder items(PlaylistItemArgs... items) {
            return items(List.of(items));
        }

        /**
         * @param name The name of the playlist.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the playlist.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public PlaylistArgs build() {
            $.interval = Objects.requireNonNull($.interval, "expected parameter 'interval' to be non-null");
            $.items = Objects.requireNonNull($.items, "expected parameter 'items' to be non-null");
            return $;
        }
    }

}
