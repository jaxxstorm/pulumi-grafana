// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const test = new grafana.Team("test", {
 *     email: "test-team-email@test.com",
 *     preferences: {
 *         theme: "dark",
 *         timezone: "utc",
 *     },
 * });
 * const fromName = grafana.getTeamOutput({
 *     name: test.name,
 * });
 * ```
 */
export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getTeam:getTeam", {
        "name": args.name,
        "orgId": args.orgId,
        "readTeamSync": args.readTeamSync,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    name: string;
    orgId?: string;
    readTeamSync?: boolean;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    readonly email: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly members: string[];
    readonly name: string;
    readonly orgId?: string;
    readonly preferences: outputs.GetTeamPreference[];
    readonly readTeamSync?: boolean;
    readonly teamId: number;
    readonly teamSyncs: outputs.GetTeamTeamSync[];
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const test = new grafana.Team("test", {
 *     email: "test-team-email@test.com",
 *     preferences: {
 *         theme: "dark",
 *         timezone: "utc",
 *     },
 * });
 * const fromName = grafana.getTeamOutput({
 *     name: test.name,
 * });
 * ```
 */
export function getTeamOutput(args: GetTeamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTeamResult> {
    return pulumi.output(args).apply((a: any) => getTeam(a, opts))
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamOutputArgs {
    name: pulumi.Input<string>;
    orgId?: pulumi.Input<string>;
    readTeamSync?: pulumi.Input<boolean>;
}
