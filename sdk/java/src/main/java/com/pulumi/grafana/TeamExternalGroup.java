// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.TeamExternalGroupArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.TeamExternalGroupState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/enterprise/team-sync/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/external_group_sync/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.TeamExternalGroup;
 * import com.pulumi.grafana.TeamExternalGroupArgs;
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
 *         var test_team_group = new TeamExternalGroup(&#34;test-team-group&#34;, TeamExternalGroupArgs.builder()        
 *             .groups(            
 *                 &#34;test-group-1&#34;,
 *                 &#34;test-group-2&#34;)
 *             .teamId(1)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/teamExternalGroup:TeamExternalGroup main {{team_id}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/teamExternalGroup:TeamExternalGroup")
public class TeamExternalGroup extends com.pulumi.resources.CustomResource {
    /**
     * The team external groups list
     * 
     */
    @Export(name="groups", type=List.class, parameters={String.class})
    private Output<List<String>> groups;

    /**
     * @return The team external groups list
     * 
     */
    public Output<List<String>> groups() {
        return this.groups;
    }
    /**
     * The Team ID
     * 
     */
    @Export(name="teamId", type=Integer.class, parameters={})
    private Output<Integer> teamId;

    /**
     * @return The Team ID
     * 
     */
    public Output<Integer> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamExternalGroup(String name) {
        this(name, TeamExternalGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamExternalGroup(String name, TeamExternalGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamExternalGroup(String name, TeamExternalGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/teamExternalGroup:TeamExternalGroup", name, args == null ? TeamExternalGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamExternalGroup(String name, Output<String> id, @Nullable TeamExternalGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/teamExternalGroup:TeamExternalGroup", name, state, makeResourceOptions(options, id));
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
    public static TeamExternalGroup get(String name, Output<String> id, @Nullable TeamExternalGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamExternalGroup(name, id, state, options);
    }
}
