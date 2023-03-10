// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource `equinix.networkedge.NetworkFile` allows creation and management of Equinix Network Edge files.
 *
 * ## Import
 *
 * This resource can be imported using an existing ID
 *
 * ```sh
 *  $ pulumi import equinix:networkedge/networkFile:NetworkFile example {existing_id}
 * ```
 *
 *  The `content`, `self_managed` and `byol` fields can not be imported.
 */
export class NetworkFile extends pulumi.CustomResource {
    /**
     * Get an existing NetworkFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkFileState, opts?: pulumi.CustomResourceOptions): NetworkFile {
        return new NetworkFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix:networkedge/networkFile:NetworkFile';

    /**
     * Returns true if the given object is an instance of NetworkFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkFile.__pulumiType;
    }

    /**
     * Boolean value that determines device licensing mode, i.e.,
     * `bring your own license` or `subscription`.
     */
    public readonly byol!: pulumi.Output<boolean>;
    /**
     * Uploaded file content, expected to be a UTF-8 encoded string.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Device type code
     */
    public readonly deviceTypeCode!: pulumi.Output<string>;
    /**
     * File name.
     */
    public readonly fileName!: pulumi.Output<string>;
    /**
     * File upload location metro code. It should match the device location metro code.
     */
    public readonly metroCode!: pulumi.Output<string>;
    /**
     * File process type (LICENSE or CLOUD_INIT).
     */
    public readonly processType!: pulumi.Output<string>;
    /**
     * Boolean value that determines device management mode, i.e.,
     * `self-managed` or `Equinix-managed`.
     */
    public readonly selfManaged!: pulumi.Output<boolean>;
    /**
     * File upload status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Unique identifier of file resource.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a NetworkFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkFileArgs | NetworkFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkFileState | undefined;
            resourceInputs["byol"] = state ? state.byol : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["deviceTypeCode"] = state ? state.deviceTypeCode : undefined;
            resourceInputs["fileName"] = state ? state.fileName : undefined;
            resourceInputs["metroCode"] = state ? state.metroCode : undefined;
            resourceInputs["processType"] = state ? state.processType : undefined;
            resourceInputs["selfManaged"] = state ? state.selfManaged : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as NetworkFileArgs | undefined;
            if ((!args || args.byol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'byol'");
            }
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.deviceTypeCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceTypeCode'");
            }
            if ((!args || args.fileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileName'");
            }
            if ((!args || args.metroCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metroCode'");
            }
            if ((!args || args.processType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'processType'");
            }
            if ((!args || args.selfManaged === undefined) && !opts.urn) {
                throw new Error("Missing required property 'selfManaged'");
            }
            resourceInputs["byol"] = args ? args.byol : undefined;
            resourceInputs["content"] = args?.content ? pulumi.secret(args.content) : undefined;
            resourceInputs["deviceTypeCode"] = args ? args.deviceTypeCode : undefined;
            resourceInputs["fileName"] = args ? args.fileName : undefined;
            resourceInputs["metroCode"] = args ? args.metroCode : undefined;
            resourceInputs["processType"] = args ? args.processType : undefined;
            resourceInputs["selfManaged"] = args ? args.selfManaged : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["content"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(NetworkFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkFile resources.
 */
export interface NetworkFileState {
    /**
     * Boolean value that determines device licensing mode, i.e.,
     * `bring your own license` or `subscription`.
     */
    byol?: pulumi.Input<boolean>;
    /**
     * Uploaded file content, expected to be a UTF-8 encoded string.
     */
    content?: pulumi.Input<string>;
    /**
     * Device type code
     */
    deviceTypeCode?: pulumi.Input<string>;
    /**
     * File name.
     */
    fileName?: pulumi.Input<string>;
    /**
     * File upload location metro code. It should match the device location metro code.
     */
    metroCode?: pulumi.Input<string | enums.Metro>;
    /**
     * File process type (LICENSE or CLOUD_INIT).
     */
    processType?: pulumi.Input<string | enums.networkedge.FileType>;
    /**
     * Boolean value that determines device management mode, i.e.,
     * `self-managed` or `Equinix-managed`.
     */
    selfManaged?: pulumi.Input<boolean>;
    /**
     * File upload status.
     */
    status?: pulumi.Input<string>;
    /**
     * Unique identifier of file resource.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkFile resource.
 */
export interface NetworkFileArgs {
    /**
     * Boolean value that determines device licensing mode, i.e.,
     * `bring your own license` or `subscription`.
     */
    byol: pulumi.Input<boolean>;
    /**
     * Uploaded file content, expected to be a UTF-8 encoded string.
     */
    content: pulumi.Input<string>;
    /**
     * Device type code
     */
    deviceTypeCode: pulumi.Input<string>;
    /**
     * File name.
     */
    fileName: pulumi.Input<string>;
    /**
     * File upload location metro code. It should match the device location metro code.
     */
    metroCode: pulumi.Input<string | enums.Metro>;
    /**
     * File process type (LICENSE or CLOUD_INIT).
     */
    processType: pulumi.Input<string | enums.networkedge.FileType>;
    /**
     * Boolean value that determines device management mode, i.e.,
     * `self-managed` or `Equinix-managed`.
     */
    selfManaged: pulumi.Input<boolean>;
}
