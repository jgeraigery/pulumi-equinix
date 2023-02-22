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
        [Input("accounts")]
        private InputList<Inputs.ConnectionASideAccessPointAccountGetArgs>? _accounts;

        /// <summary>
        /// Customer account information that is associated with this connection
        /// </summary>
        public InputList<Inputs.ConnectionASideAccessPointAccountGetArgs> Accounts
        {
            get => _accounts ?? (_accounts = new InputList<Inputs.ConnectionASideAccessPointAccountGetArgs>());
            set => _accounts = value;
        }

        [Input("authenticationKey")]
        public Input<string>? AuthenticationKey { get; set; }

        [Input("gateways")]
        private InputList<Inputs.ConnectionASideAccessPointGatewayGetArgs>? _gateways;
        public InputList<Inputs.ConnectionASideAccessPointGatewayGetArgs> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<Inputs.ConnectionASideAccessPointGatewayGetArgs>());
            set => _gateways = value;
        }

        [Input("interfaces")]
        private InputList<Inputs.ConnectionASideAccessPointInterfaceGetArgs>? _interfaces;
        public InputList<Inputs.ConnectionASideAccessPointInterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ConnectionASideAccessPointInterfaceGetArgs>());
            set => _interfaces = value;
        }

        [Input("linkProtocols")]
        public InputUnion<string, Pulumi.Equinix.Fabric.AccessPointLinkProtocolType>? LinkProtocols { get; set; }

        [Input("locations")]
        private InputList<Inputs.ConnectionASideAccessPointLocationGetArgs>? _locations;
        public InputList<Inputs.ConnectionASideAccessPointLocationGetArgs> Locations
        {
            get => _locations ?? (_locations = new InputList<Inputs.ConnectionASideAccessPointLocationGetArgs>());
            set => _locations = value;
        }

        [Input("peeringType")]
        public InputUnion<string, Pulumi.Equinix.Fabric.AccessPointPeeringType>? PeeringType { get; set; }

        [Input("ports")]
        private InputList<Inputs.ConnectionASideAccessPointPortGetArgs>? _ports;
        public InputList<Inputs.ConnectionASideAccessPointPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ConnectionASideAccessPointPortGetArgs>());
            set => _ports = value;
        }

        [Input("profiles")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileType>? Profiles { get; set; }

        [Input("providerConnectionId")]
        public Input<string>? ProviderConnectionId { get; set; }

        [Input("routingProtocols")]
        private InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs>? _routingProtocols;
        public InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs> RoutingProtocols
        {
            get => _routingProtocols ?? (_routingProtocols = new InputList<Inputs.ConnectionASideAccessPointRoutingProtocolGetArgs>());
            set => _routingProtocols = value;
        }

        [Input("sellerRegion")]
        public Input<string>? SellerRegion { get; set; }

        /// <summary>
        /// Defines the connection type like VG*VC, EVPL*VC, EPL*VC, EC*VC, GW*VC, ACCESS*EPL_VC
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.Equinix.Fabric.AccessPointType>? Type { get; set; }

        [Input("virtualDevices")]
        private InputList<Inputs.ConnectionASideAccessPointVirtualDeviceGetArgs>? _virtualDevices;
        public InputList<Inputs.ConnectionASideAccessPointVirtualDeviceGetArgs> VirtualDevices
        {
            get => _virtualDevices ?? (_virtualDevices = new InputList<Inputs.ConnectionASideAccessPointVirtualDeviceGetArgs>());
            set => _virtualDevices = value;
        }

        public ConnectionASideAccessPointGetArgs()
        {
        }
        public static new ConnectionASideAccessPointGetArgs Empty => new ConnectionASideAccessPointGetArgs();
    }
}
