// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class RoutingProtocolBfdGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bidirectional Forwarding Detection enablement
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Interval range between the received BFD control packets
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        public RoutingProtocolBfdGetArgs()
        {
        }
        public static new RoutingProtocolBfdGetArgs Empty => new RoutingProtocolBfdGetArgs();
    }
}
