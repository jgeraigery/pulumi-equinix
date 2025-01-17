// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class ConnectionASideAccessPointPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique Resource Identifier
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// Port name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("redundancies")]
        private InputList<Inputs.ConnectionASideAccessPointPortRedundancyArgs>? _redundancies;

        /// <summary>
        /// Redundancy Information
        /// </summary>
        public InputList<Inputs.ConnectionASideAccessPointPortRedundancyArgs> Redundancies
        {
            get => _redundancies ?? (_redundancies = new InputList<Inputs.ConnectionASideAccessPointPortRedundancyArgs>());
            set => _redundancies = value;
        }

        /// <summary>
        /// Equinix-assigned interface identifier
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public ConnectionASideAccessPointPortArgs()
        {
        }
        public static new ConnectionASideAccessPointPortArgs Empty => new ConnectionASideAccessPointPortArgs();
    }
}
