// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Metal
{
    /// <summary>
    /// Use this resource to associate VLAN with a Dedicated Port from
    /// [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/#associating-a-vlan-with-a-dedicated-port).
    /// 
    /// &gt; VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
    /// 
    /// ## Example Usage
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var projectId = config.Require("projectId");
    ///     var connectionId = config.Require("connectionId");
    ///     var vlanId = config.Require("vlanId");
    ///     var portId = Equinix.Metal.GetInterconnection.Invoke(new()
    ///     {
    ///         ConnectionId = connectionId,
    ///     }).Apply(invoke =&gt; invoke.Ports[0]?.Id);
    /// 
    ///     var vc = new Equinix.Metal.VirtualCircuit("vc", new()
    ///     {
    ///         ConnectionId = connectionId,
    ///         ProjectId = projectId,
    ///         PortId = portId,
    ///         VlanId = vlanId,
    ///         NniVlan = 1056,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["vcStatus"] = vc.Status,
    ///         ["vcVnid"] = vc.Vnid,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing Virtual Circuit ID: &lt;break&gt;&lt;break&gt;```sh&lt;break&gt; $ pulumi import equinix:metal/virtualCircuit:VirtualCircuit equinix_metal_virtual_circuit {existing_id} &lt;break&gt;```&lt;break&gt;&lt;break&gt;
    /// </summary>
    [EquinixResourceType("equinix:metal/virtualCircuit:VirtualCircuit")]
    public partial class VirtualCircuit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to.
        /// </summary>
        [Output("connectionId")]
        public Output<string> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
        /// </summary>
        [Output("customerIp")]
        public Output<string?> CustomerIp { get; private set; } = null!;

        /// <summary>
        /// Description for the Virtual Circuit resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The password that can be set for the VRF BGP peer
        /// </summary>
        [Output("md5")]
        public Output<string?> Md5 { get; private set; } = null!;

        /// <summary>
        /// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
        /// </summary>
        [Output("metalIp")]
        public Output<string?> MetalIp { get; private set; } = null!;

        /// <summary>
        /// Name of the Virtual Circuit resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID.
        /// </summary>
        [Output("nniVlan")]
        public Output<int?> NniVlan { get; private set; } = null!;

        /// <summary>
        /// NNI VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
        /// </summary>
        [Output("nniVnid")]
        public Output<int> NniVnid { get; private set; } = null!;

        /// <summary>
        /// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
        /// </summary>
        [Output("peerAsn")]
        public Output<int?> PeerAsn { get; private set; } = null!;

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to.
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// UUID of the Project where the VC is scoped to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Speed of the Virtual Circuit resource.
        /// </summary>
        [Output("speed")]
        public Output<string> Speed { get; private set; } = null!;

        /// <summary>
        /// Status of the virtal circuit.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A subnet from one of the IP
        /// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
        /// * For a /31 block, it will only have two IP addresses, which will be used for
        /// the metal_ip and customer_ip.
        /// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
        /// </summary>
        [Output("subnet")]
        public Output<string?> Subnet { get; private set; } = null!;

        /// <summary>
        /// Tags for the Virtual Circuit resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// UUID of the VLAN to associate.
        /// </summary>
        [Output("vlanId")]
        public Output<string?> VlanId { get; private set; } = null!;

        /// <summary>
        /// VNID VLAN parameter, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
        /// </summary>
        [Output("vnid")]
        public Output<int> Vnid { get; private set; } = null!;

        /// <summary>
        /// UUID of the VRF to associate.
        /// </summary>
        [Output("vrfId")]
        public Output<string?> VrfId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualCircuit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualCircuit(string name, VirtualCircuitArgs args, CustomResourceOptions? options = null)
            : base("equinix:metal/virtualCircuit:VirtualCircuit", name, args ?? new VirtualCircuitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualCircuit(string name, Input<string> id, VirtualCircuitState? state = null, CustomResourceOptions? options = null)
            : base("equinix:metal/virtualCircuit:VirtualCircuit", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/equinix",
                AdditionalSecretOutputs =
                {
                    "md5",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualCircuit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualCircuit Get(string name, Input<string> id, VirtualCircuitState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualCircuit(name, id, state, options);
        }
    }

    public sealed class VirtualCircuitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to.
        /// </summary>
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        /// <summary>
        /// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
        /// </summary>
        [Input("customerIp")]
        public Input<string>? CustomerIp { get; set; }

        /// <summary>
        /// Description for the Virtual Circuit resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("md5")]
        private Input<string>? _md5;

        /// <summary>
        /// The password that can be set for the VRF BGP peer
        /// </summary>
        public Input<string>? Md5
        {
            get => _md5;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _md5 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
        /// </summary>
        [Input("metalIp")]
        public Input<string>? MetalIp { get; set; }

        /// <summary>
        /// Name of the Virtual Circuit resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID.
        /// </summary>
        [Input("nniVlan")]
        public Input<int>? NniVlan { get; set; }

        /// <summary>
        /// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// UUID of the Project where the VC is scoped to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Speed of the Virtual Circuit resource.
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// A subnet from one of the IP
        /// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
        /// * For a /31 block, it will only have two IP addresses, which will be used for
        /// the metal_ip and customer_ip.
        /// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for the Virtual Circuit resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// UUID of the VLAN to associate.
        /// </summary>
        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        /// <summary>
        /// UUID of the VRF to associate.
        /// </summary>
        [Input("vrfId")]
        public Input<string>? VrfId { get; set; }

        public VirtualCircuitArgs()
        {
        }
        public static new VirtualCircuitArgs Empty => new VirtualCircuitArgs();
    }

    public sealed class VirtualCircuitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to.
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
        /// </summary>
        [Input("customerIp")]
        public Input<string>? CustomerIp { get; set; }

        /// <summary>
        /// Description for the Virtual Circuit resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("md5")]
        private Input<string>? _md5;

        /// <summary>
        /// The password that can be set for the VRF BGP peer
        /// </summary>
        public Input<string>? Md5
        {
            get => _md5;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _md5 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
        /// </summary>
        [Input("metalIp")]
        public Input<string>? MetalIp { get; set; }

        /// <summary>
        /// Name of the Virtual Circuit resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID.
        /// </summary>
        [Input("nniVlan")]
        public Input<int>? NniVlan { get; set; }

        /// <summary>
        /// NNI VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
        /// </summary>
        [Input("nniVnid")]
        public Input<int>? NniVnid { get; set; }

        /// <summary>
        /// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// UUID of the Project where the VC is scoped to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Speed of the Virtual Circuit resource.
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Status of the virtal circuit.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A subnet from one of the IP
        /// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
        /// * For a /31 block, it will only have two IP addresses, which will be used for
        /// the metal_ip and customer_ip.
        /// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for the Virtual Circuit resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// UUID of the VLAN to associate.
        /// </summary>
        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        /// <summary>
        /// VNID VLAN parameter, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
        /// </summary>
        [Input("vnid")]
        public Input<int>? Vnid { get; set; }

        /// <summary>
        /// UUID of the VRF to associate.
        /// </summary>
        [Input("vrfId")]
        public Input<string>? VrfId { get; set; }

        public VirtualCircuitState()
        {
        }
        public static new VirtualCircuitState Empty => new VirtualCircuitState();
    }
}
