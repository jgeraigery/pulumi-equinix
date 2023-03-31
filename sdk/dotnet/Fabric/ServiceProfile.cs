// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric
{
    /// <summary>
    /// ## Example Usage
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var profile = new Equinix.Fabric.ServiceProfile("profile", new()
    ///     {
    ///         Name = "Example Cloud Provider",
    ///         Description = "50 to 500 Mbps Hosted Connection to Example Cloud",
    ///         Type = "L2_PROFILE",
    ///         AccessPointTypeConfigs = new[]
    ///         {
    ///             new Equinix.Fabric.Inputs.ServiceProfileAccessPointTypeConfigArgs
    ///             {
    ///                 Type = "COLO",
    ///                 SupportedBandwidths = new[]
    ///                 {
    ///                     50,
    ///                     100,
    ///                     200,
    ///                     500,
    ///                 },
    ///                 AllowRemoteConnections = true,
    ///                 AllowCustomBandwidth = false,
    ///                 AllowBandwidthAutoApproval = false,
    ///                 LinkProtocolConfig = new Equinix.Fabric.Inputs.ServiceProfileAccessPointTypeConfigLinkProtocolConfigArgs
    ///                 {
    ///                     EncapsulationStrategy = "CTAGED",
    ///                     ReuseVlanSTag = false,
    ///                     Encapsulation = "DOT1Q",
    ///                 },
    ///                 EnableAutoGenerateServiceKey = "false,",
    ///                 ConnectionRedundancyRequired = "false,",
    ///                 ApiConfig = new Equinix.Fabric.Inputs.ServiceProfileAccessPointTypeConfigApiConfigArgs
    ///                 {
    ///                     ApiAvailable = true,
    ///                     IntegrationId = "Example-Connect-01",
    ///                     BandwidthFromApi = false,
    ///                 },
    ///                 ConnectionLabel = "Virtual Circuit Name",
    ///                 AuthenticationKey = new Equinix.Fabric.Inputs.ServiceProfileAccessPointTypeConfigAuthenticationKeyArgs
    ///                 {
    ///                     Required = true,
    ///                     Label = "Example ACCOUNT ID",
    ///                 },
    ///             },
    ///         },
    ///         Account = new Equinix.Fabric.Inputs.ServiceProfileAccountArgs
    ///         {
    ///             OrganizationName = "Example Cloud",
    ///             GlobalOrganizationName = "Example Global",
    ///         },
    ///         Metros = null,
    ///         Visibility = "PUBLIC",
    ///         MarketingInfo = new Equinix.Fabric.Inputs.ServiceProfileMarketingInfoArgs
    ///         {
    ///             Promotion = true,
    ///         },
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["profileId"] = profile.Id,
    ///     };
    /// });
    /// ```
    /// </summary>
    [EquinixResourceType("equinix:fabric/serviceProfile:ServiceProfile")]
    public partial class ServiceProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access point config information
        /// </summary>
        [Output("accessPointTypeConfigs")]
        public Output<ImmutableArray<Outputs.ServiceProfileAccessPointTypeConfig>> AccessPointTypeConfigs { get; private set; } = null!;

        /// <summary>
        /// Account
        /// </summary>
        [Output("account")]
        public Output<Outputs.ServiceProfileAccount?> Account { get; private set; } = null!;

        /// <summary>
        /// Array of contact emails
        /// </summary>
        [Output("allowedEmails")]
        public Output<ImmutableArray<string>> AllowedEmails { get; private set; } = null!;

        /// <summary>
        /// Captures connection lifecycle change information
        /// </summary>
        [Output("changeLog")]
        public Output<Outputs.ServiceProfileChangeLog> ChangeLog { get; private set; } = null!;

        /// <summary>
        /// Custom Fields
        /// </summary>
        [Output("customFields")]
        public Output<ImmutableArray<Outputs.ServiceProfileCustomField>> CustomFields { get; private set; } = null!;

        /// <summary>
        /// User-provided service description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Service Profile URI response attribute
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// Marketing Info
        /// </summary>
        [Output("marketingInfo")]
        public Output<Outputs.ServiceProfileMarketingInfo?> MarketingInfo { get; private set; } = null!;

        /// <summary>
        /// Access point config information
        /// </summary>
        [Output("metros")]
        public Output<ImmutableArray<Outputs.ServiceProfileMetro>> Metros { get; private set; } = null!;

        /// <summary>
        /// Customer-assigned service profile name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Preferences for notifications on connection configuration or status changes
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableArray<Outputs.ServiceProfileNotification>> Notifications { get; private set; } = null!;

        /// <summary>
        /// Ports
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.ServiceProfilePort>> Ports { get; private set; } = null!;

        /// <summary>
        /// Project information
        /// </summary>
        [Output("project")]
        public Output<Outputs.ServiceProfileProject?> Project { get; private set; } = null!;

        /// <summary>
        /// Self Profile
        /// </summary>
        [Output("selfProfile")]
        public Output<bool?> SelfProfile { get; private set; } = null!;

        /// <summary>
        /// Service profile state - ACTIVE, PENDING_APPROVAL, DELETED, REJECTED
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Tags attached to the connection
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Service profile type - L2*PROFILE, L3*PROFILE, ECIA*PROFILE, ECMC*PROFILE
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Equinix assigned service profile identifier
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Virtual Devices
        /// </summary>
        [Output("virtualDevices")]
        public Output<ImmutableArray<Outputs.ServiceProfileVirtualDevice>> VirtualDevices { get; private set; } = null!;

        /// <summary>
        /// Service profile visibility - PUBLIC, PRIVATE
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceProfile(string name, ServiceProfileArgs args, CustomResourceOptions? options = null)
            : base("equinix:fabric/serviceProfile:ServiceProfile", name, args ?? new ServiceProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceProfile(string name, Input<string> id, ServiceProfileState? state = null, CustomResourceOptions? options = null)
            : base("equinix:fabric/serviceProfile:ServiceProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/equinix",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceProfile Get(string name, Input<string> id, ServiceProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceProfile(name, id, state, options);
        }
    }

    public sealed class ServiceProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessPointTypeConfigs")]
        private InputList<Inputs.ServiceProfileAccessPointTypeConfigArgs>? _accessPointTypeConfigs;

        /// <summary>
        /// Access point config information
        /// </summary>
        public InputList<Inputs.ServiceProfileAccessPointTypeConfigArgs> AccessPointTypeConfigs
        {
            get => _accessPointTypeConfigs ?? (_accessPointTypeConfigs = new InputList<Inputs.ServiceProfileAccessPointTypeConfigArgs>());
            set => _accessPointTypeConfigs = value;
        }

        /// <summary>
        /// Account
        /// </summary>
        [Input("account")]
        public Input<Inputs.ServiceProfileAccountArgs>? Account { get; set; }

        [Input("allowedEmails")]
        private InputList<string>? _allowedEmails;

        /// <summary>
        /// Array of contact emails
        /// </summary>
        public InputList<string> AllowedEmails
        {
            get => _allowedEmails ?? (_allowedEmails = new InputList<string>());
            set => _allowedEmails = value;
        }

        [Input("customFields")]
        private InputList<Inputs.ServiceProfileCustomFieldArgs>? _customFields;

        /// <summary>
        /// Custom Fields
        /// </summary>
        public InputList<Inputs.ServiceProfileCustomFieldArgs> CustomFields
        {
            get => _customFields ?? (_customFields = new InputList<Inputs.ServiceProfileCustomFieldArgs>());
            set => _customFields = value;
        }

        /// <summary>
        /// User-provided service description
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Marketing Info
        /// </summary>
        [Input("marketingInfo")]
        public Input<Inputs.ServiceProfileMarketingInfoArgs>? MarketingInfo { get; set; }

        [Input("metros")]
        private InputList<Inputs.ServiceProfileMetroArgs>? _metros;

        /// <summary>
        /// Access point config information
        /// </summary>
        public InputList<Inputs.ServiceProfileMetroArgs> Metros
        {
            get => _metros ?? (_metros = new InputList<Inputs.ServiceProfileMetroArgs>());
            set => _metros = value;
        }

        /// <summary>
        /// Customer-assigned service profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<Inputs.ServiceProfileNotificationArgs>? _notifications;

        /// <summary>
        /// Preferences for notifications on connection configuration or status changes
        /// </summary>
        public InputList<Inputs.ServiceProfileNotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.ServiceProfileNotificationArgs>());
            set => _notifications = value;
        }

        [Input("ports")]
        private InputList<Inputs.ServiceProfilePortArgs>? _ports;

        /// <summary>
        /// Ports
        /// </summary>
        public InputList<Inputs.ServiceProfilePortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServiceProfilePortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Project information
        /// </summary>
        [Input("project")]
        public Input<Inputs.ServiceProfileProjectArgs>? Project { get; set; }

        /// <summary>
        /// Self Profile
        /// </summary>
        [Input("selfProfile")]
        public Input<bool>? SelfProfile { get; set; }

        /// <summary>
        /// Service profile state - ACTIVE, PENDING_APPROVAL, DELETED, REJECTED
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileState>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags attached to the connection
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Service profile type - L2*PROFILE, L3*PROFILE, ECIA*PROFILE, ECMC*PROFILE
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileType> Type { get; set; } = null!;

        [Input("virtualDevices")]
        private InputList<Inputs.ServiceProfileVirtualDeviceArgs>? _virtualDevices;

        /// <summary>
        /// Virtual Devices
        /// </summary>
        public InputList<Inputs.ServiceProfileVirtualDeviceArgs> VirtualDevices
        {
            get => _virtualDevices ?? (_virtualDevices = new InputList<Inputs.ServiceProfileVirtualDeviceArgs>());
            set => _virtualDevices = value;
        }

        /// <summary>
        /// Service profile visibility - PUBLIC, PRIVATE
        /// </summary>
        [Input("visibility")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileVisibility>? Visibility { get; set; }

        public ServiceProfileArgs()
        {
        }
        public static new ServiceProfileArgs Empty => new ServiceProfileArgs();
    }

    public sealed class ServiceProfileState : global::Pulumi.ResourceArgs
    {
        [Input("accessPointTypeConfigs")]
        private InputList<Inputs.ServiceProfileAccessPointTypeConfigGetArgs>? _accessPointTypeConfigs;

        /// <summary>
        /// Access point config information
        /// </summary>
        public InputList<Inputs.ServiceProfileAccessPointTypeConfigGetArgs> AccessPointTypeConfigs
        {
            get => _accessPointTypeConfigs ?? (_accessPointTypeConfigs = new InputList<Inputs.ServiceProfileAccessPointTypeConfigGetArgs>());
            set => _accessPointTypeConfigs = value;
        }

        /// <summary>
        /// Account
        /// </summary>
        [Input("account")]
        public Input<Inputs.ServiceProfileAccountGetArgs>? Account { get; set; }

        [Input("allowedEmails")]
        private InputList<string>? _allowedEmails;

        /// <summary>
        /// Array of contact emails
        /// </summary>
        public InputList<string> AllowedEmails
        {
            get => _allowedEmails ?? (_allowedEmails = new InputList<string>());
            set => _allowedEmails = value;
        }

        /// <summary>
        /// Captures connection lifecycle change information
        /// </summary>
        [Input("changeLog")]
        public Input<Inputs.ServiceProfileChangeLogGetArgs>? ChangeLog { get; set; }

        [Input("customFields")]
        private InputList<Inputs.ServiceProfileCustomFieldGetArgs>? _customFields;

        /// <summary>
        /// Custom Fields
        /// </summary>
        public InputList<Inputs.ServiceProfileCustomFieldGetArgs> CustomFields
        {
            get => _customFields ?? (_customFields = new InputList<Inputs.ServiceProfileCustomFieldGetArgs>());
            set => _customFields = value;
        }

        /// <summary>
        /// User-provided service description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Service Profile URI response attribute
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// Marketing Info
        /// </summary>
        [Input("marketingInfo")]
        public Input<Inputs.ServiceProfileMarketingInfoGetArgs>? MarketingInfo { get; set; }

        [Input("metros")]
        private InputList<Inputs.ServiceProfileMetroGetArgs>? _metros;

        /// <summary>
        /// Access point config information
        /// </summary>
        public InputList<Inputs.ServiceProfileMetroGetArgs> Metros
        {
            get => _metros ?? (_metros = new InputList<Inputs.ServiceProfileMetroGetArgs>());
            set => _metros = value;
        }

        /// <summary>
        /// Customer-assigned service profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<Inputs.ServiceProfileNotificationGetArgs>? _notifications;

        /// <summary>
        /// Preferences for notifications on connection configuration or status changes
        /// </summary>
        public InputList<Inputs.ServiceProfileNotificationGetArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.ServiceProfileNotificationGetArgs>());
            set => _notifications = value;
        }

        [Input("ports")]
        private InputList<Inputs.ServiceProfilePortGetArgs>? _ports;

        /// <summary>
        /// Ports
        /// </summary>
        public InputList<Inputs.ServiceProfilePortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServiceProfilePortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Project information
        /// </summary>
        [Input("project")]
        public Input<Inputs.ServiceProfileProjectGetArgs>? Project { get; set; }

        /// <summary>
        /// Self Profile
        /// </summary>
        [Input("selfProfile")]
        public Input<bool>? SelfProfile { get; set; }

        /// <summary>
        /// Service profile state - ACTIVE, PENDING_APPROVAL, DELETED, REJECTED
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileState>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags attached to the connection
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Service profile type - L2*PROFILE, L3*PROFILE, ECIA*PROFILE, ECMC*PROFILE
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileType>? Type { get; set; }

        /// <summary>
        /// Equinix assigned service profile identifier
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        [Input("virtualDevices")]
        private InputList<Inputs.ServiceProfileVirtualDeviceGetArgs>? _virtualDevices;

        /// <summary>
        /// Virtual Devices
        /// </summary>
        public InputList<Inputs.ServiceProfileVirtualDeviceGetArgs> VirtualDevices
        {
            get => _virtualDevices ?? (_virtualDevices = new InputList<Inputs.ServiceProfileVirtualDeviceGetArgs>());
            set => _virtualDevices = value;
        }

        /// <summary>
        /// Service profile visibility - PUBLIC, PRIVATE
        /// </summary>
        [Input("visibility")]
        public InputUnion<string, Pulumi.Equinix.Fabric.ProfileVisibility>? Visibility { get; set; }

        public ServiceProfileState()
        {
        }
        public static new ServiceProfileState Empty => new ServiceProfileState();
    }
}
