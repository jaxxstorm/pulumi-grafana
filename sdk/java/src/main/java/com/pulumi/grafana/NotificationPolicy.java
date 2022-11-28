// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.NotificationPolicyArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.NotificationPolicyState;
import com.pulumi.grafana.outputs.NotificationPolicyPolicy;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Sets the global notification policy for Grafana. Note that this resource manages the entire notification policy tree, and will overwrite any existing policies.
 * 
 * * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/notifications/)
 * * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#notification-policies)
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
 * import com.pulumi.grafana.ContactPoint;
 * import com.pulumi.grafana.ContactPointArgs;
 * import com.pulumi.grafana.inputs.ContactPointEmailArgs;
 * import com.pulumi.grafana.MuteTiming;
 * import com.pulumi.grafana.MuteTimingArgs;
 * import com.pulumi.grafana.inputs.MuteTimingIntervalArgs;
 * import com.pulumi.grafana.NotificationPolicy;
 * import com.pulumi.grafana.NotificationPolicyArgs;
 * import com.pulumi.grafana.inputs.NotificationPolicyPolicyArgs;
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
 *         var aContactPoint = new ContactPoint(&#34;aContactPoint&#34;, ContactPointArgs.builder()        
 *             .emails(ContactPointEmailArgs.builder()
 *                 .addresses(                
 *                     &#34;one@company.org&#34;,
 *                     &#34;two@company.org&#34;)
 *                 .message(&#34;{{ len .Alerts.Firing }} firing.&#34;)
 *                 .build())
 *             .build());
 * 
 *         var aMuteTiming = new MuteTiming(&#34;aMuteTiming&#34;, MuteTimingArgs.builder()        
 *             .intervals(MuteTimingIntervalArgs.builder()
 *                 .weekdays(&#34;monday&#34;)
 *                 .build())
 *             .build());
 * 
 *         var myNotificationPolicy = new NotificationPolicy(&#34;myNotificationPolicy&#34;, NotificationPolicyArgs.builder()        
 *             .groupBies(&#34;...&#34;)
 *             .contactPoint(aContactPoint.name())
 *             .groupWait(&#34;45s&#34;)
 *             .groupInterval(&#34;6m&#34;)
 *             .repeatInterval(&#34;3h&#34;)
 *             .policies(            
 *                 NotificationPolicyPolicyArgs.builder()
 *                     .matchers(NotificationPolicyPolicyMatcherArgs.builder()
 *                         .label(&#34;mylabel&#34;)
 *                         .match(&#34;=&#34;)
 *                         .value(&#34;myvalue&#34;)
 *                         .build())
 *                     .contactPoint(aContactPoint.name())
 *                     .groupBies(&#34;alertname&#34;)
 *                     .continue_(true)
 *                     .muteTimings(aMuteTiming.name())
 *                     .groupWait(&#34;45s&#34;)
 *                     .groupInterval(&#34;6m&#34;)
 *                     .repeatInterval(&#34;3h&#34;)
 *                     .policies(NotificationPolicyPolicyPolicyArgs.builder()
 *                         .matchers(NotificationPolicyPolicyPolicyMatcherArgs.builder()
 *                             .label(&#34;sublabel&#34;)
 *                             .match(&#34;=&#34;)
 *                             .value(&#34;subvalue&#34;)
 *                             .build())
 *                         .contactPoint(aContactPoint.name())
 *                         .groupBies(&#34;...&#34;)
 *                         .build())
 *                     .build(),
 *                 NotificationPolicyPolicyArgs.builder()
 *                     .matchers(NotificationPolicyPolicyMatcherArgs.builder()
 *                         .label(&#34;anotherlabel&#34;)
 *                         .match(&#34;=~&#34;)
 *                         .value(&#34;another value.*&#34;)
 *                         .build())
 *                     .contactPoint(aContactPoint.name())
 *                     .groupBies(&#34;...&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * # The policy is a singleton, so the ID is a constant &#34;policy&#34; value.
 * 
 * ```sh
 *  $ pulumi import grafana:index/notificationPolicy:NotificationPolicy notification_policy_name &#34;policy&#34;
 * ```
 * 
 */
@ResourceType(type="grafana:index/notificationPolicy:NotificationPolicy")
public class NotificationPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The default contact point to route all unmatched notifications to.
     * 
     */
    @Export(name="contactPoint", type=String.class, parameters={})
    private Output<String> contactPoint;

    /**
     * @return The default contact point to route all unmatched notifications to.
     * 
     */
    public Output<String> contactPoint() {
        return this.contactPoint;
    }
    /**
     * A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
     * 
     */
    @Export(name="groupBies", type=List.class, parameters={String.class})
    private Output<List<String>> groupBies;

    /**
     * @return A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
     * 
     */
    public Output<List<String>> groupBies() {
        return this.groupBies;
    }
    /**
     * Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    @Export(name="groupInterval", type=String.class, parameters={})
    private Output</* @Nullable */ String> groupInterval;

    /**
     * @return Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    public Output<Optional<String>> groupInterval() {
        return Codegen.optional(this.groupInterval);
    }
    /**
     * Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    @Export(name="groupWait", type=String.class, parameters={})
    private Output</* @Nullable */ String> groupWait;

    /**
     * @return Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    public Output<Optional<String>> groupWait() {
        return Codegen.optional(this.groupWait);
    }
    /**
     * Routing rules for specific label sets.
     * 
     */
    @Export(name="policies", type=List.class, parameters={NotificationPolicyPolicy.class})
    private Output</* @Nullable */ List<NotificationPolicyPolicy>> policies;

    /**
     * @return Routing rules for specific label sets.
     * 
     */
    public Output<Optional<List<NotificationPolicyPolicy>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    @Export(name="repeatInterval", type=String.class, parameters={})
    private Output</* @Nullable */ String> repeatInterval;

    /**
     * @return Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    public Output<Optional<String>> repeatInterval() {
        return Codegen.optional(this.repeatInterval);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NotificationPolicy(String name) {
        this(name, NotificationPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NotificationPolicy(String name, NotificationPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NotificationPolicy(String name, NotificationPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/notificationPolicy:NotificationPolicy", name, args == null ? NotificationPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NotificationPolicy(String name, Output<String> id, @Nullable NotificationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/notificationPolicy:NotificationPolicy", name, state, makeResourceOptions(options, id));
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
    public static NotificationPolicy get(String name, Output<String> id, @Nullable NotificationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NotificationPolicy(name, id, state, options);
    }
}
