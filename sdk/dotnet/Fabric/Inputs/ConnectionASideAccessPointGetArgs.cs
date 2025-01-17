// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class ConnectionASideAccessPointGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account
        /// </summary>
        [Input("account")]
        public Input<Inputs.ConnectionASideAccessPointAccountGetArgs>? Account { get; set; }

        /// <summary>
        /// Authentication key for provider based connections
        /// </summary>
        [Input("authenticationKey")]
        public Input<string>? AuthenticationKey { get; set; }

        [Input("gateway")]
        public Input<Inputs.ConnectionASideAccessPointGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// Virtual device interface
        /// </summary>
        [Input("interface")]
        public Input<Inputs.ConnectionASideAccessPointInterfaceGetArgs>? Interface { get; set; }

        /// <summary>
        /// Connection link protocol
        /// </summary>
        [Input("linkProtocol")]
        public Input<Inputs.ConnectionASideAccessPointLinkProtocolGetArgs>? LinkProtocol { get; set; }

        /// <summary>
        /// Access point location
        /// </summary>
        [Input("location")]
        public Input<Inputs.ConnectionASideAccessPointLocationGetArgs>? Location { get; set; }

        /// <summary>
        /// Simplified Network
        /// </summary>
        [Input("network")]
        public Input<Inputs.ConnectionASideAccessPointNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// Peering Type- PRIVATE,MICROSOFT,PUBLIC, MANUAL
        /// </summary>
        [Input("peeringType")]
        public InputUnion<string, Pulumi.Equinix.Fabric.AccessPointPeeringType>? PeeringType { get; set; }

        /// <summary>
        /// Port access point information
        /// </summary>
        [Input("port")]
        public Input<Inputs.ConnectionASideAccessPointPortGetArgs>? Port { get; set; }

        /// <summary>
        /// Service Profile
        /// </summary>
        [Input("profile")]
        public Input<Inputs.ConnectionASideAccessPointProfileGetArgs>? Profile { get; set; }

        /// <summary>
        /// Provider assigned Connection Id
        /// </summary>
        [Input("providerConnectionId")]
        public Input<string>? ProviderConnectionId { get; set; }

        /// <summary>
        /// Cloud Router access point information that replaces `gateway` (refers to below for nested schema)
        /// </summary>
        [Input("router")]
        public Input<Inputs.ConnectionASideAccessPointRouterGetArgs>? Router { get; set; }

        [Input("routingProtocols")]
        private InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs>? _routingProtocols;

        /// <summary>
        /// Access point routing protocols configuration
        /// </summary>
        public InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs> RoutingProtocols
        {
            get => _routingProtocols ?? (_routingProtocols = new InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs>());
            set => _routingProtocols = value;
        }

        /// <summary>
        /// Access point seller region
        /// </summary>
        [Input("sellerRegion")]
        public Input<string>? SellerRegion { get; set; }

        /// <summary>
        /// Interface type
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.Equinix.Fabric.AccessPointType>? Type { get; set; }

        /// <summary>
        /// Virtual device
        /// </summary>
        [Input("virtualDevice")]
        public Input<Inputs.ConnectionASideAccessPointVirtualDeviceGetArgs>? VirtualDevice { get; set; }

        public ConnectionASideAccessPointGetArgs()
        {
        }
        public static new ConnectionASideAccessPointGetArgs Empty => new ConnectionASideAccessPointGetArgs();
    }
}
