// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class RoutingProtocolChangeLogGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        [Input("createdByEmail")]
        public Input<string>? CreatedByEmail { get; set; }

        [Input("createdByFullName")]
        public Input<string>? CreatedByFullName { get; set; }

        [Input("createdDateTime")]
        public Input<string>? CreatedDateTime { get; set; }

        [Input("deletedBy")]
        public Input<string>? DeletedBy { get; set; }

        [Input("deletedByEmail")]
        public Input<string>? DeletedByEmail { get; set; }

        [Input("deletedByFullName")]
        public Input<string>? DeletedByFullName { get; set; }

        [Input("deletedDateTime")]
        public Input<string>? DeletedDateTime { get; set; }

        [Input("updatedBy")]
        public Input<string>? UpdatedBy { get; set; }

        [Input("updatedByEmail")]
        public Input<string>? UpdatedByEmail { get; set; }

        [Input("updatedByFullName")]
        public Input<string>? UpdatedByFullName { get; set; }

        [Input("updatedDateTime")]
        public Input<string>? UpdatedDateTime { get; set; }

        public RoutingProtocolChangeLogGetArgs()
        {
        }
        public static new RoutingProtocolChangeLogGetArgs Empty => new RoutingProtocolChangeLogGetArgs();
    }
}