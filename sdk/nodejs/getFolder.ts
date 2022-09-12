// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/dashboard-folders/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/folder/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const test = new grafana.Folder("test", {
 *     title: "test-folder",
 *     uid: "test-ds-folder-uid",
 * });
 * const fromTitle = grafana.getFolderOutput({
 *     title: test.title,
 * });
 * ```
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("grafana:index/getFolder:getFolder", {
        "title": args.title,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderArgs {
    /**
     * The name of the Grafana folder.
     */
    title: string;
}

/**
 * A collection of values returned by getFolder.
 */
export interface GetFolderResult {
    /**
     * The numerical ID of the Grafana folder.
     */
    readonly id: number;
    /**
     * The name of the Grafana folder.
     */
    readonly title: string;
    /**
     * The uid of the Grafana folder.
     */
    readonly uid: string;
    /**
     * The full URL of the folder.
     */
    readonly url: string;
}

export function getFolderOutput(args: GetFolderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFolderResult> {
    return pulumi.output(args).apply(a => getFolder(a, opts))
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderOutputArgs {
    /**
     * The name of the Grafana folder.
     */
    title: pulumi.Input<string>;
}
