// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
 *
 * > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@equinix-labs/pulumi-equinix";
 *
 * // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
 * const testVlan = new equinix.metal.Vlan("testVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testGateway = equinix.metal.getGateway({
 *     gatewayId: local.gateway_id,
 * });
 * ```
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("equinix:metal/getGateway:getGateway", {
        "gatewayId": args.gatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGateway.
 */
export interface GetGatewayArgs {
    /**
     * UUID of the metal gateway resource to retrieve.
     */
    gatewayId: string;
}

/**
 * A collection of values returned by getGateway.
 */
export interface GetGatewayResult {
    readonly gatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * UUID of IP reservation block bound to the gateway.
     */
    readonly ipReservationId: string;
    /**
     * Size of the private IPv4 subnet bound to this metal gateway. One of
     * `8`, `16`, `32`, `64`, `128`.
     */
    readonly privateIpv4SubnetSize: number;
    /**
     * UUID of the project where the gateway is scoped to.
     */
    readonly projectId: string;
    /**
     * Status of the gateway resource.
     */
    readonly state: string;
    /**
     * UUID of the VLAN where the gateway is scoped to.
     */
    readonly vlanId: string;
    /**
     * UUID of the VRF associated with the IP Reservation.
     */
    readonly vrfId: string;
}
/**
 * Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
 *
 * > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@equinix-labs/pulumi-equinix";
 *
 * // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
 * const testVlan = new equinix.metal.Vlan("testVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testGateway = equinix.metal.getGateway({
 *     gatewayId: local.gateway_id,
 * });
 * ```
 */
export function getGatewayOutput(args: GetGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayResult> {
    return pulumi.output(args).apply((a: any) => getGateway(a, opts))
}

/**
 * A collection of arguments for invoking getGateway.
 */
export interface GetGatewayOutputArgs {
    /**
     * UUID of the metal gateway resource to retrieve.
     */
    gatewayId: pulumi.Input<string>;
}
