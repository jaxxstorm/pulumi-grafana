// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Data source for retrieving a single library panel by name or uid.
 */
export function getLibraryPanel(args?: GetLibraryPanelArgs, opts?: pulumi.InvokeOptions): Promise<GetLibraryPanelResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("grafana:index/getLibraryPanel:getLibraryPanel", {
        "name": args.name,
        "uid": args.uid,
    }, opts);
}

/**
 * A collection of arguments for invoking getLibraryPanel.
 */
export interface GetLibraryPanelArgs {
    /**
     * Name of the library panel.
     */
    name?: string;
    /**
     * The unique identifier (UID) of the library panel.
     */
    uid?: string;
}

/**
 * A collection of values returned by getLibraryPanel.
 */
export interface GetLibraryPanelResult {
    /**
     * Timestamp when the library panel was created.
     */
    readonly created: string;
    /**
     * Numerical IDs of Grafana dashboards containing the library panel.
     */
    readonly dashboardIds: number[];
    /**
     * Description of the library panel.
     */
    readonly description: string;
    /**
     * ID of the folder where the library panel is stored.
     */
    readonly folderId: number;
    /**
     * Name of the folder containing the library panel.
     */
    readonly folderName: string;
    /**
     * Unique ID (UID) of the folder containing the library panel.
     */
    readonly folderUid: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The JSON model for the library panel.
     */
    readonly modelJson: string;
    /**
     * Name of the library panel.
     */
    readonly name?: string;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    readonly orgId: number;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    readonly panelId: number;
    /**
     * Type of the library panel (eg. text).
     */
    readonly type: string;
    /**
     * The unique identifier (UID) of the library panel.
     */
    readonly uid?: string;
    /**
     * Timestamp when the library panel was last modified.
     */
    readonly updated: string;
    /**
     * Version of the library panel.
     */
    readonly version: number;
}

export function getLibraryPanelOutput(args?: GetLibraryPanelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLibraryPanelResult> {
    return pulumi.output(args).apply(a => getLibraryPanel(a, opts))
}

/**
 * A collection of arguments for invoking getLibraryPanel.
 */
export interface GetLibraryPanelOutputArgs {
    /**
     * Name of the library panel.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique identifier (UID) of the library panel.
     */
    uid?: pulumi.Input<string>;
}
