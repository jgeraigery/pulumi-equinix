// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Metal.Inputs
{

    public sealed class InterconnectionServiceTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("maxAllowedSpeed")]
        public Input<string>? MaxAllowedSpeed { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Connection type - dedicated or shared.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public InterconnectionServiceTokenArgs()
        {
        }
        public static new InterconnectionServiceTokenArgs Empty => new InterconnectionServiceTokenArgs();
    }
}
