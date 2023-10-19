// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a [connection resource](https://metal.equinix.com/developers/docs/networking/fabric/)
 *
 * > Equinix Metal connection with with Service Token A-side / Z-side (service_token_type) is not generally available and may not be enabled yet for your organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@equinix-labs/pulumi-equinix";
 *
 * const example = equinix.metal.getInterconnection({
 *     connectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
 * });
 * ```
 */
export function getInterconnection(args: GetInterconnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetInterconnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("equinix:metal/getInterconnection:getInterconnection", {
        "connectionId": args.connectionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInterconnection.
 */
export interface GetInterconnectionArgs {
    /**
     * ID of the connection resource.
     */
    connectionId: string;
}

/**
 * A collection of values returned by getInterconnection.
 */
export interface GetInterconnectionResult {
    readonly connectionId: string;
    /**
     * The preferred email used for communication and notifications about the Equinix Fabric interconnection.
     */
    readonly contactEmail: string;
    /**
     * Description of the connection resource.
     */
    readonly description: string;
    /**
     * (**Deprecated**) Slug of a facility to which the connection belongs. Use metro instead; read the facility to metro migration guide
     *
     * @deprecated Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
     */
    readonly facility: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Slug of a metro to which the connection belongs.
     */
    readonly metro: string;
    /**
     * Mode for connections in IBX facilities with the dedicated type - standard or tunnel.
     */
    readonly mode: string;
    /**
     * Port name.
     */
    readonly name: string;
    /**
     * ID of the organization where the connection is scoped to.
     */
    readonly organizationId: string;
    /**
     * List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
     */
    readonly ports: outputs.metal.GetInterconnectionPort[];
    /**
     * ID of project to which the connection belongs.
     */
    readonly projectId: string;
    /**
     * Connection redundancy, reduntant or primary.
     */
    readonly redundancy: string;
    /**
     * Type of service token, aSide or z_side. One available in shared connection.
     */
    readonly serviceTokenType: string;
    /**
     * List of connection service tokens with attributes
     */
    readonly serviceTokens: outputs.metal.GetInterconnectionServiceToken[];
    /**
     * Port speed in bits per second.
     */
    readonly speed: string;
    /**
     * Port status.
     */
    readonly status: string;
    /**
     * String list of tags.
     */
    readonly tags: string[];
    /**
     * (Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the equinixEcxL2Connection resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). If your organization already has connection service tokens enabled, use `serviceTokens` instead.
     *
     * @deprecated If your organization already has connection service tokens enabled, use `service_tokens` instead
     */
    readonly token: string;
    /**
     * Token type, `aSide` or `zSide`.
     */
    readonly type: string;
    /**
     * Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection.
     */
    readonly vlans: number[];
}
/**
 * Use this data source to retrieve a [connection resource](https://metal.equinix.com/developers/docs/networking/fabric/)
 *
 * > Equinix Metal connection with with Service Token A-side / Z-side (service_token_type) is not generally available and may not be enabled yet for your organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@equinix-labs/pulumi-equinix";
 *
 * const example = equinix.metal.getInterconnection({
 *     connectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
 * });
 * ```
 */
export function getInterconnectionOutput(args: GetInterconnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInterconnectionResult> {
    return pulumi.output(args).apply((a: any) => getInterconnection(a, opts))
}

/**
 * A collection of arguments for invoking getInterconnection.
 */
export interface GetInterconnectionOutputArgs {
    /**
     * ID of the connection resource.
     */
    connectionId: pulumi.Input<string>;
}
