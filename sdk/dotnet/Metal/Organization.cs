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
    /// Provides a resource to manage organization resource in Equinix Metal.
    /// 
    /// ## Example Usage
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var orgResource = new Equinix.Metal.Organization("org", new()
    ///     {
    ///         Name = "Foo Organization",
    ///         Address = new Equinix.Metal.Inputs.OrganizationAddressArgs
    ///         {
    ///             Address = "org street",
    ///             City = "london",
    ///             Country = "GB",
    ///             ZipCode = "12345",
    ///         },
    ///         Description = "An organization",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["org"] = orgResource.Id,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing organization ID: &lt;break&gt;&lt;break&gt;```sh&lt;break&gt; $ pulumi import equinix:metal/organization:Organization equinix_metal_organization {existing_organization_id} &lt;break&gt;```&lt;break&gt;&lt;break&gt;
    /// </summary>
    [EquinixResourceType("equinix:metal/organization:Organization")]
    public partial class Organization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An object that has the address information. See Address
        /// below for more details.
        /// </summary>
        [Output("address")]
        public Output<Outputs.OrganizationAddress> Address { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the organization was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Description string.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Output("logo")]
        public Output<string?> Logo { get; private set; } = null!;

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Output("twitter")]
        public Output<string?> Twitter { get; private set; } = null!;

        /// <summary>
        /// The timestamp for the last time the organization was updated.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;

        /// <summary>
        /// Website link.
        /// </summary>
        [Output("website")]
        public Output<string?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs args, CustomResourceOptions? options = null)
            : base("equinix:metal/organization:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
            : base("equinix:metal/organization:Organization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, state, options);
        }
    }

    public sealed class OrganizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object that has the address information. See Address
        /// below for more details.
        /// </summary>
        [Input("address", required: true)]
        public Input<Inputs.OrganizationAddressArgs> Address { get; set; } = null!;

        /// <summary>
        /// Description string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        /// <summary>
        /// Website link.
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public OrganizationArgs()
        {
        }
        public static new OrganizationArgs Empty => new OrganizationArgs();
    }

    public sealed class OrganizationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object that has the address information. See Address
        /// below for more details.
        /// </summary>
        [Input("address")]
        public Input<Inputs.OrganizationAddressGetArgs>? Address { get; set; }

        /// <summary>
        /// The timestamp for when the organization was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Description string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        /// <summary>
        /// The timestamp for the last time the organization was updated.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Website link.
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public OrganizationState()
        {
        }
        public static new OrganizationState Empty => new OrganizationState();
    }
}
