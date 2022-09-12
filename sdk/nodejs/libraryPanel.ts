// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Grafana library panels.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/panels/panel-library/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/http_api/library_element/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 *
 * const test = new grafana.LibraryPanel("test", {modelJson: JSON.stringify({
 *     title: "updated name",
 *     id: 12,
 *     version: 35,
 * })});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import grafana:index/libraryPanel:LibraryPanel panel_name {{library_panel_slug}}
 * ```
 */
export class LibraryPanel extends pulumi.CustomResource {
    /**
     * Get an existing LibraryPanel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LibraryPanelState, opts?: pulumi.CustomResourceOptions): LibraryPanel {
        return new LibraryPanel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/libraryPanel:LibraryPanel';

    /**
     * Returns true if the given object is an instance of LibraryPanel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LibraryPanel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LibraryPanel.__pulumiType;
    }

    /**
     * Timestamp when the library panel was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Numerical IDs of Grafana dashboards containing the library panel.
     */
    public /*out*/ readonly dashboardIds!: pulumi.Output<number[]>;
    /**
     * Description of the library panel.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * ID of the folder where the library panel is stored.
     */
    public readonly folderId!: pulumi.Output<number | undefined>;
    /**
     * Name of the folder containing the library panel.
     */
    public /*out*/ readonly folderName!: pulumi.Output<string>;
    /**
     * Unique ID (UID) of the folder containing the library panel.
     */
    public /*out*/ readonly folderUid!: pulumi.Output<string>;
    /**
     * The JSON model for the library panel.
     */
    public readonly modelJson!: pulumi.Output<string>;
    /**
     * Name of the library panel.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    public /*out*/ readonly orgId!: pulumi.Output<number>;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    public /*out*/ readonly panelId!: pulumi.Output<number>;
    /**
     * Type of the library panel (eg. text).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * Timestamp when the library panel was last modified.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;
    /**
     * Version of the library panel.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a LibraryPanel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LibraryPanelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LibraryPanelArgs | LibraryPanelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LibraryPanelState | undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["dashboardIds"] = state ? state.dashboardIds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["folderName"] = state ? state.folderName : undefined;
            resourceInputs["folderUid"] = state ? state.folderUid : undefined;
            resourceInputs["modelJson"] = state ? state.modelJson : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["panelId"] = state ? state.panelId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updated"] = state ? state.updated : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as LibraryPanelArgs | undefined;
            if ((!args || args.modelJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelJson'");
            }
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["modelJson"] = args ? args.modelJson : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uid"] = args ? args.uid : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["dashboardIds"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["folderName"] = undefined /*out*/;
            resourceInputs["folderUid"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["panelId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LibraryPanel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LibraryPanel resources.
 */
export interface LibraryPanelState {
    /**
     * Timestamp when the library panel was created.
     */
    created?: pulumi.Input<string>;
    /**
     * Numerical IDs of Grafana dashboards containing the library panel.
     */
    dashboardIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Description of the library panel.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder where the library panel is stored.
     */
    folderId?: pulumi.Input<number>;
    /**
     * Name of the folder containing the library panel.
     */
    folderName?: pulumi.Input<string>;
    /**
     * Unique ID (UID) of the folder containing the library panel.
     */
    folderUid?: pulumi.Input<string>;
    /**
     * The JSON model for the library panel.
     */
    modelJson?: pulumi.Input<string>;
    /**
     * Name of the library panel.
     */
    name?: pulumi.Input<string>;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    orgId?: pulumi.Input<number>;
    /**
     * The numeric ID of the library panel computed by Grafana.
     */
    panelId?: pulumi.Input<number>;
    /**
     * Type of the library panel (eg. text).
     */
    type?: pulumi.Input<string>;
    /**
     * The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
     */
    uid?: pulumi.Input<string>;
    /**
     * Timestamp when the library panel was last modified.
     */
    updated?: pulumi.Input<string>;
    /**
     * Version of the library panel.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LibraryPanel resource.
 */
export interface LibraryPanelArgs {
    /**
     * ID of the folder where the library panel is stored.
     */
    folderId?: pulumi.Input<number>;
    /**
     * The JSON model for the library panel.
     */
    modelJson: pulumi.Input<string>;
    /**
     * Name of the library panel.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
     */
    uid?: pulumi.Input<string>;
}
