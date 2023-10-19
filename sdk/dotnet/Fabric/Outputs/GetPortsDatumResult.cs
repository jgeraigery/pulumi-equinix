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
    public sealed class GetPortsDatumResult
    {
        public readonly ImmutableArray<Outputs.GetPortsDatumAccountResult> Accounts;
        public readonly int AvailableBandwidth;
        public readonly int Bandwidth;
        public readonly ImmutableArray<Outputs.GetPortsDatumChangeLogResult> ChangeLogs;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetPortsDatumDeviceResult> Devices;
        public readonly ImmutableArray<Outputs.GetPortsDatumEncapsulationResult> Encapsulations;
        public readonly string Href;
        public readonly bool LagEnabled;
        public readonly ImmutableArray<Outputs.GetPortsDatumLocationResult> Locations;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetPortsDatumOperationResult> Operations;
        public readonly ImmutableArray<Outputs.GetPortsDatumRedundancyResult> Redundancies;
        public readonly string ServiceType;
        public readonly string State;
        public readonly string Type;
        public readonly int UsedBandwidth;
        public readonly string? Uuid;

        [OutputConstructor]
        private GetPortsDatumResult(
            ImmutableArray<Outputs.GetPortsDatumAccountResult> accounts,

            int availableBandwidth,

            int bandwidth,

            ImmutableArray<Outputs.GetPortsDatumChangeLogResult> changeLogs,

            string description,

            ImmutableArray<Outputs.GetPortsDatumDeviceResult> devices,

            ImmutableArray<Outputs.GetPortsDatumEncapsulationResult> encapsulations,

            string href,

            bool lagEnabled,

            ImmutableArray<Outputs.GetPortsDatumLocationResult> locations,

            string name,

            ImmutableArray<Outputs.GetPortsDatumOperationResult> operations,

            ImmutableArray<Outputs.GetPortsDatumRedundancyResult> redundancies,

            string serviceType,

            string state,

            string type,

            int usedBandwidth,

            string? uuid)
        {
            Accounts = accounts;
            AvailableBandwidth = availableBandwidth;
            Bandwidth = bandwidth;
            ChangeLogs = changeLogs;
            Description = description;
            Devices = devices;
            Encapsulations = encapsulations;
            Href = href;
            LagEnabled = lagEnabled;
            Locations = locations;
            Name = name;
            Operations = operations;
            Redundancies = redundancies;
            ServiceType = serviceType;
            State = state;
            Type = type;
            UsedBandwidth = usedBandwidth;
            Uuid = uuid;
        }
    }
}
