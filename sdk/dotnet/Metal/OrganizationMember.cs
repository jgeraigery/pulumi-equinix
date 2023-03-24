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
    /// Manage the membership of existing and new invitees within an Equinix Metal organization and its projects.
    /// 
    /// ## Example Usage
    /// 
    /// Add a member to an organization to collaborate on given projects:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Equinix.Metal.OrganizationMember("member", new()
    ///     {
    ///         Invitee = "member@example.com",
    ///         Roles = new[]
    ///         {
    ///             "limited_collaborator",
    ///         },
    ///         ProjectsIds = new[]
    ///         {
    ///             @var.Project_id,
    ///         },
    ///         OrganizationId = @var.Organization_id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Add a member to an organization as an organization administrator:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var owner = new Equinix.Metal.OrganizationMember("owner", new()
    ///     {
    ///         Invitee = "admin@example.com",
    ///         Roles = new[]
    ///         {
    ///             "owner",
    ///         },
    ///         ProjectsIds = new[] {},
    ///         OrganizationId = @var.Organization_id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the `invitee` and `organization_id` as colon separated arguments
    /// 
    /// ```sh
    ///  $ pulumi import equinix:metal/organizationMember:OrganizationMember resource_name {invitee}:{organization_id}
    /// ```
    /// </summary>
    [EquinixResourceType("equinix:metal/organizationMember:OrganizationMember")]
    public partial class OrganizationMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When the invitation was created (only known in the invitation stage)
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The user_id of the user that sent the invitation (only known in the invitation stage)
        /// </summary>
        [Output("invitedBy")]
        public Output<string> InvitedBy { get; private set; } = null!;

        /// <summary>
        /// The email address of the user to invite
        /// </summary>
        [Output("invitee")]
        public Output<string> Invitee { get; private set; } = null!;

        /// <summary>
        /// A message to include in the emailed invitation.
        /// </summary>
        [Output("message")]
        public Output<string?> Message { get; private set; } = null!;

        /// <summary>
        /// The nonce for the invitation (only known in the invitation stage)
        /// </summary>
        [Output("nonce")]
        public Output<string> Nonce { get; private set; } = null!;

        /// <summary>
        /// The organization to invite the user to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
        /// </summary>
        [Output("projectsIds")]
        public Output<ImmutableArray<string>> ProjectsIds { get; private set; } = null!;

        /// <summary>
        /// Organization roles (admin, collaborator, limited_collaborator, billing)
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// When the invitation was updated (only known in the invitation stage)
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationMember(string name, OrganizationMemberArgs args, CustomResourceOptions? options = null)
            : base("equinix:metal/organizationMember:OrganizationMember", name, args ?? new OrganizationMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationMember(string name, Input<string> id, OrganizationMemberState? state = null, CustomResourceOptions? options = null)
            : base("equinix:metal/organizationMember:OrganizationMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationMember Get(string name, Input<string> id, OrganizationMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationMember(name, id, state, options);
        }
    }

    public sealed class OrganizationMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email address of the user to invite
        /// </summary>
        [Input("invitee", required: true)]
        public Input<string> Invitee { get; set; } = null!;

        /// <summary>
        /// A message to include in the emailed invitation.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The organization to invite the user to
        /// </summary>
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("projectsIds", required: true)]
        private InputList<string>? _projectsIds;

        /// <summary>
        /// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
        /// </summary>
        public InputList<string> ProjectsIds
        {
            get => _projectsIds ?? (_projectsIds = new InputList<string>());
            set => _projectsIds = value;
        }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// Organization roles (admin, collaborator, limited_collaborator, billing)
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public OrganizationMemberArgs()
        {
        }
        public static new OrganizationMemberArgs Empty => new OrganizationMemberArgs();
    }

    public sealed class OrganizationMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the invitation was created (only known in the invitation stage)
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The user_id of the user that sent the invitation (only known in the invitation stage)
        /// </summary>
        [Input("invitedBy")]
        public Input<string>? InvitedBy { get; set; }

        /// <summary>
        /// The email address of the user to invite
        /// </summary>
        [Input("invitee")]
        public Input<string>? Invitee { get; set; }

        /// <summary>
        /// A message to include in the emailed invitation.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The nonce for the invitation (only known in the invitation stage)
        /// </summary>
        [Input("nonce")]
        public Input<string>? Nonce { get; set; }

        /// <summary>
        /// The organization to invite the user to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("projectsIds")]
        private InputList<string>? _projectsIds;

        /// <summary>
        /// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
        /// </summary>
        public InputList<string> ProjectsIds
        {
            get => _projectsIds ?? (_projectsIds = new InputList<string>());
            set => _projectsIds = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Organization roles (admin, collaborator, limited_collaborator, billing)
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// When the invitation was updated (only known in the invitation stage)
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public OrganizationMemberState()
        {
        }
        public static new OrganizationMemberState Empty => new OrganizationMemberState();
    }
}
