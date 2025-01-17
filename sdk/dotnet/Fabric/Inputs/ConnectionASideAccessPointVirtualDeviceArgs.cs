// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class ConnectionASideAccessPointVirtualDeviceArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Interface type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Equinix-assigned interface identifier
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public ConnectionASideAccessPointVirtualDeviceArgs()
        {
        }
        public static new ConnectionASideAccessPointVirtualDeviceArgs Empty => new ConnectionASideAccessPointVirtualDeviceArgs();
    }
}
