// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class ConnectionASideAccessPointRoutingProtocolGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Routing protocol instance state
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

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

        public ConnectionASideAccessPointRoutingProtocolGetArgs()
        {
        }
        public static new ConnectionASideAccessPointRoutingProtocolGetArgs Empty => new ConnectionASideAccessPointRoutingProtocolGetArgs();
    }
}
