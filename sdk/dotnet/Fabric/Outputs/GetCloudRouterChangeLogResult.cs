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
    public sealed class GetCloudRouterChangeLogResult
    {
        public readonly string CreatedBy;
        public readonly string CreatedByEmail;
        public readonly string CreatedByFullName;
        public readonly string CreatedDateTime;
        public readonly string DeletedBy;
        public readonly string DeletedByEmail;
        public readonly string DeletedByFullName;
        public readonly string DeletedDateTime;
        public readonly string UpdatedBy;
        public readonly string UpdatedByEmail;
        public readonly string UpdatedByFullName;
        public readonly string UpdatedDateTime;

        [OutputConstructor]
        private GetCloudRouterChangeLogResult(
            string createdBy,

            string createdByEmail,

            string createdByFullName,

            string createdDateTime,

            string deletedBy,

            string deletedByEmail,

            string deletedByFullName,

            string deletedDateTime,

            string updatedBy,

            string updatedByEmail,

            string updatedByFullName,

            string updatedDateTime)
        {
            CreatedBy = createdBy;
            CreatedByEmail = createdByEmail;
            CreatedByFullName = createdByFullName;
            CreatedDateTime = createdDateTime;
            DeletedBy = deletedBy;
            DeletedByEmail = deletedByEmail;
            DeletedByFullName = deletedByFullName;
            DeletedDateTime = deletedDateTime;
            UpdatedBy = updatedBy;
            UpdatedByEmail = updatedByEmail;
            UpdatedByFullName = updatedByFullName;
            UpdatedDateTime = updatedDateTime;
        }
    }
}