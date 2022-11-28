// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.MuteTimingArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.MuteTimingState;
import com.pulumi.grafana.outputs.MuteTimingInterval;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Grafana Alerting mute timings.
 * 
 * * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/notifications/mute-timings/)
 * * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#mute-timings)
 * 
 * This resource requires Grafana 9.1.0 or later.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.MuteTiming;
 * import com.pulumi.grafana.MuteTimingArgs;
 * import com.pulumi.grafana.inputs.MuteTimingIntervalArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var myMuteTiming = new MuteTiming(&#34;myMuteTiming&#34;, MuteTimingArgs.builder()        
 *             .intervals(MuteTimingIntervalArgs.builder()
 *                 .daysOfMonths(                
 *                     &#34;1:7&#34;,
 *                     &#34;-1&#34;)
 *                 .months(                
 *                     &#34;1:3&#34;,
 *                     &#34;december&#34;)
 *                 .times(MuteTimingIntervalTimeArgs.builder()
 *                     .end(&#34;14:17&#34;)
 *                     .start(&#34;04:56&#34;)
 *                     .build())
 *                 .weekdays(                
 *                     &#34;monday&#34;,
 *                     &#34;tuesday:thursday&#34;)
 *                 .years(                
 *                     &#34;2030&#34;,
 *                     &#34;2025:2026&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/muteTiming:MuteTiming mute_timing_name {{mute_timing_name}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/muteTiming:MuteTiming")
public class MuteTiming extends com.pulumi.resources.CustomResource {
    /**
     * The time intervals at which to mute notifications.
     * 
     */
    @Export(name="intervals", type=List.class, parameters={MuteTimingInterval.class})
    private Output</* @Nullable */ List<MuteTimingInterval>> intervals;

    /**
     * @return The time intervals at which to mute notifications.
     * 
     */
    public Output<Optional<List<MuteTimingInterval>>> intervals() {
        return Codegen.optional(this.intervals);
    }
    /**
     * The name of the mute timing.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the mute timing.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MuteTiming(String name) {
        this(name, MuteTimingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MuteTiming(String name, @Nullable MuteTimingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MuteTiming(String name, @Nullable MuteTimingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/muteTiming:MuteTiming", name, args == null ? MuteTimingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MuteTiming(String name, Output<String> id, @Nullable MuteTimingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/muteTiming:MuteTiming", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static MuteTiming get(String name, Output<String> id, @Nullable MuteTimingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MuteTiming(name, id, state, options);
    }
}
