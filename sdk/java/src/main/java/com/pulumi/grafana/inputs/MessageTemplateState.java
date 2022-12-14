// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MessageTemplateState extends com.pulumi.resources.ResourceArgs {

    public static final MessageTemplateState Empty = new MessageTemplateState();

    /**
     * The name of the message template.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the message template.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The content of the message template.
     * 
     */
    @Import(name="template")
    private @Nullable Output<String> template;

    /**
     * @return The content of the message template.
     * 
     */
    public Optional<Output<String>> template() {
        return Optional.ofNullable(this.template);
    }

    private MessageTemplateState() {}

    private MessageTemplateState(MessageTemplateState $) {
        this.name = $.name;
        this.template = $.template;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MessageTemplateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MessageTemplateState $;

        public Builder() {
            $ = new MessageTemplateState();
        }

        public Builder(MessageTemplateState defaults) {
            $ = new MessageTemplateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the message template.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the message template.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param template The content of the message template.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<String> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template The content of the message template.
         * 
         * @return builder
         * 
         */
        public Builder template(String template) {
            return template(Output.of(template));
        }

        public MessageTemplateState build() {
            return $;
        }
    }

}
