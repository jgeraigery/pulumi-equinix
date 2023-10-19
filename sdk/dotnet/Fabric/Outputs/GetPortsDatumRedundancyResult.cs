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
    public sealed class GetPortsDatumRedundancyResult
    {
        public readonly bool Enabled;
        public readonly string Group;
        public readonly string Priority;

        [OutputConstructor]
        private GetPortsDatumRedundancyResult(
            bool enabled,

            string group,

            string priority)
        {
            Enabled = enabled;
            Group = group;
            Priority = priority;
        }
    }
}
