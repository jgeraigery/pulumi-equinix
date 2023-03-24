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
    /// Provides a resource to create and manage blocks of reserved IP addresses in a project.
    /// 
    /// When a user provisions first device in a facility, Equinix Metal API automatically allocates IPv6/56 and private IPv4/25 blocks.
    /// The new device then gets IPv6 and private IPv4 addresses from those block. It also gets a public IPv4/31 address.
    /// Every new device in the project and facility will automatically get IPv6 and private IPv4 addresses from these pre-allocated blocks.
    /// The IPv6 and private IPv4 blocks can't be created, only imported. With this resource, it's possible to create either public IPv4 blocks or global IPv4 blocks.
    /// 
    /// Public blocks are allocated in a facility. Addresses from public blocks can only be assigned to devices in the facility. Public blocks can have mask from /24 (256 addresses) to /32 (1 address). If you create public block with this resource, you must fill the facility argument.
    /// 
    /// Addresses from global blocks can be assigned in any facility. Global blocks can have mask from /30 (4 addresses), to /32 (1 address). If you create global block with this resource, you must specify type = "global_ipv4" and you must omit the facility argument.
    /// 
    /// Once IP block is allocated or imported, an address from it can be assigned to device with the `equinix.metal.IpAttachment` resource.
    /// 
    /// &gt; VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
    /// 
    /// ## Example Usage
    /// 
    /// Allocate reserved IP blocks:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility for myproject
    ///     var twoElasticAddresses = new Equinix.Metal.ReservedIpBlock("twoElasticAddresses", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         Facility = "sv15",
    ///         Quantity = 2,
    ///     });
    /// 
    ///     // Allocate 1 floating IP in Sillicon Valley (sv) metro
    ///     var testReservedIpBlock = new Equinix.Metal.ReservedIpBlock("testReservedIpBlock", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         Type = "public_ipv4",
    ///         Metro = "sv",
    ///         Quantity = 1,
    ///     });
    /// 
    ///     // Allocate 1 global floating IP, which can be assigned to device in any facility
    ///     var testMetal_reservedIpBlockReservedIpBlock = new Equinix.Metal.ReservedIpBlock("testMetal/reservedIpBlockReservedIpBlock", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         Type = "global_ipv4",
    ///         Quantity = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Allocate a block and run a device with public IPv4 from the block
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility
    ///     var example = new Equinix.Metal.ReservedIpBlock("example", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         Facility = "sv15",
    ///         Quantity = 2,
    ///     });
    /// 
    ///     // Run a device with both public IPv4 from the block assigned
    ///     var nodes = new Equinix.Metal.Device("nodes", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         Facilities = new[]
    ///         {
    ///             "sv15",
    ///         },
    ///         Plan = "c3.small.x86",
    ///         OperatingSystem = "ubuntu_20_04",
    ///         Hostname = "test",
    ///         BillingCycle = "hourly",
    ///         IpAddresses = new[]
    ///         {
    ///             new Equinix.Metal.Inputs.DeviceIpAddressArgs
    ///             {
    ///                 Type = "public_ipv4",
    ///                 Cidr = 31,
    ///                 ReservationIds = new[]
    ///                 {
    ///                     example.Id,
    ///                 },
    ///             },
    ///             new Equinix.Metal.Inputs.DeviceIpAddressArgs
    ///             {
    ///                 Type = "private_ipv4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing IP reservation ID
    /// 
    /// ```sh
    ///  $ pulumi import equinix:metal/reservedIpBlock:ReservedIpBlock equinix_metal_reserved_ip_block {existing_ip_reservation_id}
    /// ```
    /// </summary>
    [EquinixResourceType("equinix:metal/reservedIpBlock:ReservedIpBlock")]
    public partial class ReservedIpBlock : global::Pulumi.CustomResource
    {
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Address family as integer. One of `4` or `6`.
        /// </summary>
        [Output("addressFamily")]
        public Output<int> AddressFamily { get; private set; } = null!;

        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
        /// </summary>
        [Output("cidr")]
        public Output<int> Cidr { get; private set; } = null!;

        /// <summary>
        /// Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
        /// </summary>
        [Output("cidrNotation")]
        public Output<string> CidrNotation { get; private set; } = null!;

        /// <summary>
        /// Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
        /// be helpful for self-managed IPAM. The object must be valid JSON.
        /// </summary>
        [Output("customData")]
        public Output<string?> CustomData { get; private set; } = null!;

        /// <summary>
        /// Arbitrary description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Facility where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
        /// </summary>
        [Output("facility")]
        public Output<string?> Facility { get; private set; } = null!;

        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Boolean flag whether addresses from a block are global (i.e. can be assigned in any
        /// facility).
        /// </summary>
        [Output("global")]
        public Output<bool> Global { get; private set; } = null!;

        [Output("manageable")]
        public Output<bool> Manageable { get; private set; } = null!;

        [Output("management")]
        public Output<bool> Management { get; private set; } = null!;

        /// <summary>
        /// Metro where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
        /// </summary>
        [Output("metro")]
        public Output<string?> Metro { get; private set; } = null!;

        /// <summary>
        /// Mask in decimal notation, e.g. `255.255.255.0`.
        /// </summary>
        [Output("netmask")]
        public Output<string> Netmask { get; private set; } = null!;

        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The metal project ID where to allocate the address block.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Boolean flag whether addresses from a block are public.
        /// </summary>
        [Output("public")]
        public Output<bool> Public { get; private set; } = null!;

        /// <summary>
        /// The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
        /// </summary>
        [Output("quantity")]
        public Output<int> Quantity { get; private set; } = null!;

        /// <summary>
        /// String list of tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
        /// compatibility.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
        /// </summary>
        [Output("vrfId")]
        public Output<string?> VrfId { get; private set; } = null!;

        /// <summary>
        /// Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
        /// </summary>
        [Output("waitForState")]
        public Output<string?> WaitForState { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedIpBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedIpBlock(string name, ReservedIpBlockArgs args, CustomResourceOptions? options = null)
            : base("equinix:metal/reservedIpBlock:ReservedIpBlock", name, args ?? new ReservedIpBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedIpBlock(string name, Input<string> id, ReservedIpBlockState? state = null, CustomResourceOptions? options = null)
            : base("equinix:metal/reservedIpBlock:ReservedIpBlock", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/equinix/pulumi-equinix/releases/download/0.0.1-alpha.1679651896+b37a673a.dirty",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReservedIpBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedIpBlock Get(string name, Input<string> id, ReservedIpBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new ReservedIpBlock(name, id, state, options);
        }
    }

    public sealed class ReservedIpBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
        /// </summary>
        [Input("cidr")]
        public Input<int>? Cidr { get; set; }

        /// <summary>
        /// Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
        /// be helpful for self-managed IPAM. The object must be valid JSON.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// Arbitrary description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Facility where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
        /// </summary>
        [Input("facility")]
        public InputUnion<string, Pulumi.Equinix.Metal.Facility>? Facility { get; set; }

        /// <summary>
        /// Metro where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The metal project ID where to allocate the address block.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
        /// </summary>
        [Input("quantity")]
        public Input<int>? Quantity { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// String list of tags.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
        /// compatibility.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.Equinix.Metal.IpBlockType>? Type { get; set; }

        /// <summary>
        /// Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
        /// </summary>
        [Input("vrfId")]
        public Input<string>? VrfId { get; set; }

        /// <summary>
        /// Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
        /// </summary>
        [Input("waitForState")]
        public Input<string>? WaitForState { get; set; }

        public ReservedIpBlockArgs()
        {
        }
        public static new ReservedIpBlockArgs Empty => new ReservedIpBlockArgs();
    }

    public sealed class ReservedIpBlockState : global::Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Address family as integer. One of `4` or `6`.
        /// </summary>
        [Input("addressFamily")]
        public Input<int>? AddressFamily { get; set; }

        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
        /// </summary>
        [Input("cidr")]
        public Input<int>? Cidr { get; set; }

        /// <summary>
        /// Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
        /// </summary>
        [Input("cidrNotation")]
        public Input<string>? CidrNotation { get; set; }

        /// <summary>
        /// Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
        /// be helpful for self-managed IPAM. The object must be valid JSON.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        /// <summary>
        /// Arbitrary description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Facility where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
        /// </summary>
        [Input("facility")]
        public InputUnion<string, Pulumi.Equinix.Metal.Facility>? Facility { get; set; }

        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Boolean flag whether addresses from a block are global (i.e. can be assigned in any
        /// facility).
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        [Input("manageable")]
        public Input<bool>? Manageable { get; set; }

        [Input("management")]
        public Input<bool>? Management { get; set; }

        /// <summary>
        /// Metro where to allocate the public IP address block, makes sense only
        /// if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Mask in decimal notation, e.g. `255.255.255.0`.
        /// </summary>
        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        /// <summary>
        /// Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The metal project ID where to allocate the address block.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Boolean flag whether addresses from a block are public.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
        /// </summary>
        [Input("quantity")]
        public Input<int>? Quantity { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// String list of tags.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
        /// compatibility.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.Equinix.Metal.IpBlockType>? Type { get; set; }

        /// <summary>
        /// Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
        /// </summary>
        [Input("vrfId")]
        public Input<string>? VrfId { get; set; }

        /// <summary>
        /// Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
        /// </summary>
        [Input("waitForState")]
        public Input<string>? WaitForState { get; set; }

        public ReservedIpBlockState()
        {
        }
        public static new ReservedIpBlockState Empty => new ReservedIpBlockState();
    }
}
