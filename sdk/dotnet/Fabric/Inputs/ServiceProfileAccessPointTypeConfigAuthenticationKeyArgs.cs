// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Fabric.Inputs
{

    public sealed class ServiceProfileAccessPointTypeConfigAuthenticationKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Label
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Required
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        public ServiceProfileAccessPointTypeConfigAuthenticationKeyArgs()
        {
        }
        public static new ServiceProfileAccessPointTypeConfigAuthenticationKeyArgs Empty => new ServiceProfileAccessPointTypeConfigAuthenticationKeyArgs();
    }
}
