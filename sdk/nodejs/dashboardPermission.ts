// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/dashboard_permissions/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 *
 * const team = new grafana.Team("team", {});
 * const user = new grafana.User("user", {email: "user.name@example.com"});
 * const metrics = new grafana.Dashboard("metrics", {configJson: fs.readFileSync("grafana-dashboard.json")});
 * const collectionPermission = new grafana.DashboardPermission("collectionPermission", {
 *     dashboardId: metrics.dashboardId,
 *     permissions: [
 *         {
 *             role: "Editor",
 *             permission: "Edit",
 *         },
 *         {
 *             teamId: team.id,
 *             permission: "View",
 *         },
 *         {
 *             userId: user.id,
 *             permission: "Admin",
 *         },
 *     ],
 * });
 * ```
 */
export class DashboardPermission extends pulumi.CustomResource {
    /**
     * Get an existing DashboardPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardPermissionState, opts?: pulumi.CustomResourceOptions): DashboardPermission {
        return new DashboardPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/dashboardPermission:DashboardPermission';

    /**
     * Returns true if the given object is an instance of DashboardPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DashboardPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DashboardPermission.__pulumiType;
    }

    /**
     * ID of the dashboard to apply permissions to.
     */
    public readonly dashboardId!: pulumi.Output<number>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    public readonly permissions!: pulumi.Output<outputs.DashboardPermissionPermission[]>;

    /**
     * Create a DashboardPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardPermissionArgs | DashboardPermissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardPermissionState | undefined;
            resourceInputs["dashboardId"] = state ? state.dashboardId : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
        } else {
            const args = argsOrState as DashboardPermissionArgs | undefined;
            if ((!args || args.dashboardId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dashboardId'");
            }
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            resourceInputs["dashboardId"] = args ? args.dashboardId : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DashboardPermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DashboardPermission resources.
 */
export interface DashboardPermissionState {
    /**
     * ID of the dashboard to apply permissions to.
     */
    dashboardId?: pulumi.Input<number>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.DashboardPermissionPermission>[]>;
}

/**
 * The set of arguments for constructing a DashboardPermission resource.
 */
export interface DashboardPermissionArgs {
    /**
     * ID of the dashboard to apply permissions to.
     */
    dashboardId: pulumi.Input<number>;
    /**
     * The permission items to add/update. Items that are omitted from the list will be removed.
     */
    permissions: pulumi.Input<pulumi.Input<inputs.DashboardPermissionPermission>[]>;
}
