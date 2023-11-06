// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Outputs
{

    [OutputType]
    public sealed class RoutingProtocolChange
    {
        public readonly string? Href;
        public readonly string? Type;
        public readonly string? Uuid;

        [OutputConstructor]
        private RoutingProtocolChange(
            string? href,

            string? type,

            string? uuid)
        {
            Href = href;
            Type = type;
            Uuid = uuid;
        }
    }
}