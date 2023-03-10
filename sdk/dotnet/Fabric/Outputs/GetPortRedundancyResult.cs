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
    public sealed class GetPortRedundancyResult
    {
        public readonly bool Enabled;
        public readonly int Group;
        public readonly string Priority;

        [OutputConstructor]
        private GetPortRedundancyResult(
            bool enabled,

            int group,

            string priority)
        {
            Enabled = enabled;
            Group = group;
            Priority = priority;
        }
    }
}
