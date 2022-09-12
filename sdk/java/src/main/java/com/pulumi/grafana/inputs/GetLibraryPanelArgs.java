// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLibraryPanelArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLibraryPanelArgs Empty = new GetLibraryPanelArgs();

    /**
     * Name of the library panel.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the library panel.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The unique identifier (UID) of the library panel.
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return The unique identifier (UID) of the library panel.
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    private GetLibraryPanelArgs() {}

    private GetLibraryPanelArgs(GetLibraryPanelArgs $) {
        this.name = $.name;
        this.uid = $.uid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLibraryPanelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLibraryPanelArgs $;

        public Builder() {
            $ = new GetLibraryPanelArgs();
        }

        public Builder(GetLibraryPanelArgs defaults) {
            $ = new GetLibraryPanelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the library panel.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the library panel.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param uid The unique identifier (UID) of the library panel.
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid The unique identifier (UID) of the library panel.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        public GetLibraryPanelArgs build() {
            return $;
        }
    }

}
