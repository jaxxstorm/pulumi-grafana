// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.MachineLearningJobArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.MachineLearningJobState;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A job defines the queries and model parameters for a machine learning task.
 * 
 */
@ResourceType(type="grafana:index/machineLearningJob:MachineLearningJob")
public class MachineLearningJob extends com.pulumi.resources.CustomResource {
    /**
     * The id of the datasource to query.
     * 
     */
    @Export(name="datasourceId", type=Integer.class, parameters={})
    private Output<Integer> datasourceId;

    /**
     * @return The id of the datasource to query.
     * 
     */
    public Output<Integer> datasourceId() {
        return this.datasourceId;
    }
    /**
     * The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     * 
     */
    @Export(name="datasourceType", type=String.class, parameters={})
    private Output<String> datasourceType;

    /**
     * @return The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
     * 
     */
    public Output<String> datasourceType() {
        return this.datasourceType;
    }
    /**
     * A description of the job.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the job.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/ for the full list of available hyperparameters. Defaults to `map[]`.
     * 
     */
    @Export(name="hyperParams", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> hyperParams;

    /**
     * @return The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/ for the full list of available hyperparameters. Defaults to `map[]`.
     * 
     */
    public Output<Optional<Map<String,Object>>> hyperParams() {
        return Codegen.optional(this.hyperParams);
    }
    /**
     * The data interval in seconds to train the data on. Defaults to `300`.
     * 
     */
    @Export(name="interval", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> interval;

    /**
     * @return The data interval in seconds to train the data on. Defaults to `300`.
     * 
     */
    public Output<Optional<Integer>> interval() {
        return Codegen.optional(this.interval);
    }
    /**
     * The metric used to query the job results.
     * 
     */
    @Export(name="metric", type=String.class, parameters={})
    private Output<String> metric;

    /**
     * @return The metric used to query the job results.
     * 
     */
    public Output<String> metric() {
        return this.metric;
    }
    /**
     * The name of the job.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the job.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * An object representing the query params to query Grafana with.
     * 
     */
    @Export(name="queryParams", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> queryParams;

    /**
     * @return An object representing the query params to query Grafana with.
     * 
     */
    public Output<Map<String,Object>> queryParams() {
        return this.queryParams;
    }
    /**
     * The data interval in seconds to train the data on. Defaults to `7776000`.
     * 
     */
    @Export(name="trainingWindow", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> trainingWindow;

    /**
     * @return The data interval in seconds to train the data on. Defaults to `7776000`.
     * 
     */
    public Output<Optional<Integer>> trainingWindow() {
        return Codegen.optional(this.trainingWindow);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MachineLearningJob(String name) {
        this(name, MachineLearningJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MachineLearningJob(String name, MachineLearningJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MachineLearningJob(String name, MachineLearningJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/machineLearningJob:MachineLearningJob", name, args == null ? MachineLearningJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MachineLearningJob(String name, Output<String> id, @Nullable MachineLearningJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/machineLearningJob:MachineLearningJob", name, state, makeResourceOptions(options, id));
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
    public static MachineLearningJob get(String name, Output<String> id, @Nullable MachineLearningJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MachineLearningJob(name, id, state, options);
    }
}
