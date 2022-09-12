// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.ContactPointArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.ContactPointState;
import com.pulumi.grafana.outputs.ContactPointAlertmanager;
import com.pulumi.grafana.outputs.ContactPointDingding;
import com.pulumi.grafana.outputs.ContactPointDiscord;
import com.pulumi.grafana.outputs.ContactPointEmail;
import com.pulumi.grafana.outputs.ContactPointGooglechat;
import com.pulumi.grafana.outputs.ContactPointKafka;
import com.pulumi.grafana.outputs.ContactPointOpsgeny;
import com.pulumi.grafana.outputs.ContactPointPagerduty;
import com.pulumi.grafana.outputs.ContactPointPushover;
import com.pulumi.grafana.outputs.ContactPointSensugo;
import com.pulumi.grafana.outputs.ContactPointSlack;
import com.pulumi.grafana.outputs.ContactPointTeam;
import com.pulumi.grafana.outputs.ContactPointTelegram;
import com.pulumi.grafana.outputs.ContactPointThreema;
import com.pulumi.grafana.outputs.ContactPointVictorop;
import com.pulumi.grafana.outputs.ContactPointWebhook;
import com.pulumi.grafana.outputs.ContactPointWecom;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Grafana Alerting contact points.
 * 
 * * [Official documentation](https://grafana.com/docs/grafana/next/alerting/contact-points)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#contact-points)
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
 *         var myContactPoint = new ContactPoint(&#34;myContactPoint&#34;, ContactPointArgs.builder()        
 *             .emails(ContactPointEmailArgs.builder()
 *                 .addresses(                
 *                     &#34;one@company.org&#34;,
 *                     &#34;two@company.org&#34;)
 *                 .disableResolveMessage(false)
 *                 .message(&#34;{{ len .Alerts.Firing }} firing.&#34;)
 *                 .singleEmail(true)
 *                 .subject(&#34;{{ template \&#34;default.title\&#34; .}}&#34;)
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
 *  $ pulumi import grafana:index/contactPoint:ContactPoint contact_point_name {{contact_point_name}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/contactPoint:ContactPoint")
public class ContactPoint extends com.pulumi.resources.CustomResource {
    /**
     * A contact point that sends notifications to other Alertmanager instances.
     * 
     */
    @Export(name="alertmanagers", type=List.class, parameters={ContactPointAlertmanager.class})
    private Output</* @Nullable */ List<ContactPointAlertmanager>> alertmanagers;

    /**
     * @return A contact point that sends notifications to other Alertmanager instances.
     * 
     */
    public Output<Optional<List<ContactPointAlertmanager>>> alertmanagers() {
        return Codegen.optional(this.alertmanagers);
    }
    /**
     * A contact point that sends notifications to DingDing.
     * 
     */
    @Export(name="dingdings", type=List.class, parameters={ContactPointDingding.class})
    private Output</* @Nullable */ List<ContactPointDingding>> dingdings;

    /**
     * @return A contact point that sends notifications to DingDing.
     * 
     */
    public Output<Optional<List<ContactPointDingding>>> dingdings() {
        return Codegen.optional(this.dingdings);
    }
    /**
     * A contact point that sends notifications as Discord messages
     * 
     */
    @Export(name="discords", type=List.class, parameters={ContactPointDiscord.class})
    private Output</* @Nullable */ List<ContactPointDiscord>> discords;

    /**
     * @return A contact point that sends notifications as Discord messages
     * 
     */
    public Output<Optional<List<ContactPointDiscord>>> discords() {
        return Codegen.optional(this.discords);
    }
    /**
     * A contact point that sends notifications to an email address.
     * 
     */
    @Export(name="emails", type=List.class, parameters={ContactPointEmail.class})
    private Output</* @Nullable */ List<ContactPointEmail>> emails;

    /**
     * @return A contact point that sends notifications to an email address.
     * 
     */
    public Output<Optional<List<ContactPointEmail>>> emails() {
        return Codegen.optional(this.emails);
    }
    /**
     * A contact point that sends notifications to Google Chat.
     * 
     */
    @Export(name="googlechats", type=List.class, parameters={ContactPointGooglechat.class})
    private Output</* @Nullable */ List<ContactPointGooglechat>> googlechats;

    /**
     * @return A contact point that sends notifications to Google Chat.
     * 
     */
    public Output<Optional<List<ContactPointGooglechat>>> googlechats() {
        return Codegen.optional(this.googlechats);
    }
    /**
     * A contact point that publishes notifications to Apache Kafka topics.
     * 
     */
    @Export(name="kafkas", type=List.class, parameters={ContactPointKafka.class})
    private Output</* @Nullable */ List<ContactPointKafka>> kafkas;

    /**
     * @return A contact point that publishes notifications to Apache Kafka topics.
     * 
     */
    public Output<Optional<List<ContactPointKafka>>> kafkas() {
        return Codegen.optional(this.kafkas);
    }
    /**
     * The name of the contact point.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the contact point.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A contact point that sends notifications to OpsGenie.
     * 
     */
    @Export(name="opsgenies", type=List.class, parameters={ContactPointOpsgeny.class})
    private Output</* @Nullable */ List<ContactPointOpsgeny>> opsgenies;

    /**
     * @return A contact point that sends notifications to OpsGenie.
     * 
     */
    public Output<Optional<List<ContactPointOpsgeny>>> opsgenies() {
        return Codegen.optional(this.opsgenies);
    }
    /**
     * A contact point that sends notifications to PagerDuty.
     * 
     */
    @Export(name="pagerduties", type=List.class, parameters={ContactPointPagerduty.class})
    private Output</* @Nullable */ List<ContactPointPagerduty>> pagerduties;

    /**
     * @return A contact point that sends notifications to PagerDuty.
     * 
     */
    public Output<Optional<List<ContactPointPagerduty>>> pagerduties() {
        return Codegen.optional(this.pagerduties);
    }
    /**
     * A contact point that sends notifications to Pushover.
     * 
     */
    @Export(name="pushovers", type=List.class, parameters={ContactPointPushover.class})
    private Output</* @Nullable */ List<ContactPointPushover>> pushovers;

    /**
     * @return A contact point that sends notifications to Pushover.
     * 
     */
    public Output<Optional<List<ContactPointPushover>>> pushovers() {
        return Codegen.optional(this.pushovers);
    }
    /**
     * A contact point that sends notifications to SensuGo.
     * 
     */
    @Export(name="sensugos", type=List.class, parameters={ContactPointSensugo.class})
    private Output</* @Nullable */ List<ContactPointSensugo>> sensugos;

    /**
     * @return A contact point that sends notifications to SensuGo.
     * 
     */
    public Output<Optional<List<ContactPointSensugo>>> sensugos() {
        return Codegen.optional(this.sensugos);
    }
    /**
     * A contact point that sends notifications to Slack.
     * 
     */
    @Export(name="slacks", type=List.class, parameters={ContactPointSlack.class})
    private Output</* @Nullable */ List<ContactPointSlack>> slacks;

    /**
     * @return A contact point that sends notifications to Slack.
     * 
     */
    public Output<Optional<List<ContactPointSlack>>> slacks() {
        return Codegen.optional(this.slacks);
    }
    /**
     * A contact point that sends notifications to Microsoft Teams.
     * 
     */
    @Export(name="teams", type=List.class, parameters={ContactPointTeam.class})
    private Output</* @Nullable */ List<ContactPointTeam>> teams;

    /**
     * @return A contact point that sends notifications to Microsoft Teams.
     * 
     */
    public Output<Optional<List<ContactPointTeam>>> teams() {
        return Codegen.optional(this.teams);
    }
    /**
     * A contact point that sends notifications to Telegram.
     * 
     */
    @Export(name="telegrams", type=List.class, parameters={ContactPointTelegram.class})
    private Output</* @Nullable */ List<ContactPointTelegram>> telegrams;

    /**
     * @return A contact point that sends notifications to Telegram.
     * 
     */
    public Output<Optional<List<ContactPointTelegram>>> telegrams() {
        return Codegen.optional(this.telegrams);
    }
    /**
     * A contact point that sends notifications to Threema.
     * 
     */
    @Export(name="threemas", type=List.class, parameters={ContactPointThreema.class})
    private Output</* @Nullable */ List<ContactPointThreema>> threemas;

    /**
     * @return A contact point that sends notifications to Threema.
     * 
     */
    public Output<Optional<List<ContactPointThreema>>> threemas() {
        return Codegen.optional(this.threemas);
    }
    /**
     * A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
     * 
     */
    @Export(name="victorops", type=List.class, parameters={ContactPointVictorop.class})
    private Output</* @Nullable */ List<ContactPointVictorop>> victorops;

    /**
     * @return A contact point that sends notifications to VictorOps (now known as Splunk OnCall).
     * 
     */
    public Output<Optional<List<ContactPointVictorop>>> victorops() {
        return Codegen.optional(this.victorops);
    }
    /**
     * A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
     * 
     */
    @Export(name="webhooks", type=List.class, parameters={ContactPointWebhook.class})
    private Output</* @Nullable */ List<ContactPointWebhook>> webhooks;

    /**
     * @return A contact point that sends notifications to an arbitrary webhook, using the Prometheus webhook format defined here: https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
     * 
     */
    public Output<Optional<List<ContactPointWebhook>>> webhooks() {
        return Codegen.optional(this.webhooks);
    }
    /**
     * A contact point that sends notifications to WeCom.
     * 
     */
    @Export(name="wecoms", type=List.class, parameters={ContactPointWecom.class})
    private Output</* @Nullable */ List<ContactPointWecom>> wecoms;

    /**
     * @return A contact point that sends notifications to WeCom.
     * 
     */
    public Output<Optional<List<ContactPointWecom>>> wecoms() {
        return Codegen.optional(this.wecoms);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContactPoint(String name) {
        this(name, ContactPointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContactPoint(String name, @Nullable ContactPointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContactPoint(String name, @Nullable ContactPointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/contactPoint:ContactPoint", name, args == null ? ContactPointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContactPoint(String name, Output<String> id, @Nullable ContactPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/contactPoint:ContactPoint", name, state, makeResourceOptions(options, id));
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
    public static ContactPoint get(String name, Output<String> id, @Nullable ContactPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContactPoint(name, id, state, options);
    }
}
