// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class RoutingProtocolOperationErrorAdditionalInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("reason")]
        public Input<string>? Reason { get; set; }

        public RoutingProtocolOperationErrorAdditionalInfoArgs()
        {
        }
        public static new RoutingProtocolOperationErrorAdditionalInfoArgs Empty => new RoutingProtocolOperationErrorAdditionalInfoArgs();
    }
}
