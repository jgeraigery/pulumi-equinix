// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package metal

import (
	"context"
	"reflect"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The datasource can be used to fetch a single device.
//
// If you need to fetch a list of devices which meet filter criteria, you can use the metal.getDevices datasource.
//
// > **Note:** All arguments including the `rootPassword` and `userData` will be stored in
//
//	the raw state as plain-text.
//
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/equinix/pulumi-equinix/sdk/go/equinix/metal"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := metal.LookupDevice(ctx, &metal.LookupDeviceArgs{
//				ProjectId: pulumi.StringRef(local.Project_id),
//				Hostname:  pulumi.StringRef("mydevice"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("id", test.Id)
//			return nil
//		})
//	}
//
// ```
//
// ```go
// package main
//
// import (
//
//	"github.com/equinix/pulumi-equinix/sdk/go/equinix/metal"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := metal.LookupDevice(ctx, &metal.LookupDeviceArgs{
//				DeviceId: pulumi.StringRef("4c641195-25e5-4c3c-b2b7-4cd7a42c7b40"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ipv4", test.AccessPublicIpv4)
//			return nil
//		})
//	}
//
// ```
func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeviceResult
	err := ctx.Invoke("equinix:metal/getDevice:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDevice.
type LookupDeviceArgs struct {
	// Device ID.
	//
	// > **NOTE:** You should pass either `deviceId`, or both `projectId` and `hostname`.
	DeviceId *string `pulumi:"deviceId"`
	// The device name.
	Hostname *string `pulumi:"hostname"`
	// The id of the project in which the devices exists.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getDevice.
type LookupDeviceResult struct {
	// The ipv4 private IP assigned to the device.
	AccessPrivateIpv4 string `pulumi:"accessPrivateIpv4"`
	// The ipv4 management IP assigned to the device.
	AccessPublicIpv4 string `pulumi:"accessPublicIpv4"`
	// The ipv6 management IP assigned to the device.
	AccessPublicIpv6 string `pulumi:"accessPublicIpv6"`
	AlwaysPxe        bool   `pulumi:"alwaysPxe"`
	// The billing cycle of the device (monthly or hourly).
	BillingCycle string `pulumi:"billingCycle"`
	// Description string for the device.
	Description string `pulumi:"description"`
	DeviceId    string `pulumi:"deviceId"`
	// (**Deprecated**) The facility where the device is deployed. Use metro instead; read the facility to metro migration guide
	//
	// Deprecated: Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
	Facility string `pulumi:"facility"`
	// The id of hardware reservation which this device occupies.
	HardwareReservationId string `pulumi:"hardwareReservationId"`
	Hostname              string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IpxeScriptUrl string `pulumi:"ipxeScriptUrl"`
	// The metro where the device is deployed
	Metro string `pulumi:"metro"`
	// L2 network type of the device, one of `layer3`, `layer2-bonded`,
	// `layer2-individual`, `hybrid`.
	NetworkType string `pulumi:"networkType"`
	// The device's private and public IP (v4 and v6) network details. See
	// Network Attribute below for more details.
	Networks []GetDeviceNetwork `pulumi:"networks"`
	// The operating system running on the device.
	OperatingSystem string `pulumi:"operatingSystem"`
	// The hardware config of the device.
	Plan string `pulumi:"plan"`
	// List of ports assigned to the device. See Ports Attribute below for
	// more details.
	Ports     []GetDevicePort `pulumi:"ports"`
	ProjectId string          `pulumi:"projectId"`
	// Root password to the server (if still available).
	RootPassword string `pulumi:"rootPassword"`
	// The hostname to use for [Serial over SSH](https://deploy.equinix.com/developers/docs/metal/resilience-recovery/serial-over-ssh/) access to the device
	SosHostname string `pulumi:"sosHostname"`
	// List of IDs of SSH keys deployed in the device, can be both user or project SSH keys.
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// The state of the device.
	State   string `pulumi:"state"`
	Storage string `pulumi:"storage"`
	// Tags attached to the device.
	Tags []string `pulumi:"tags"`
}

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceResult, error) {
			args := v.(LookupDeviceArgs)
			r, err := LookupDevice(ctx, &args, opts...)
			var s LookupDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceResultOutput)
}

// A collection of arguments for invoking getDevice.
type LookupDeviceOutputArgs struct {
	// Device ID.
	//
	// > **NOTE:** You should pass either `deviceId`, or both `projectId` and `hostname`.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// The device name.
	Hostname pulumi.StringPtrInput `pulumi:"hostname"`
	// The id of the project in which the devices exists.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}

// A collection of values returned by getDevice.
type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

// The ipv4 private IP assigned to the device.
func (o LookupDeviceResultOutput) AccessPrivateIpv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.AccessPrivateIpv4 }).(pulumi.StringOutput)
}

// The ipv4 management IP assigned to the device.
func (o LookupDeviceResultOutput) AccessPublicIpv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.AccessPublicIpv4 }).(pulumi.StringOutput)
}

// The ipv6 management IP assigned to the device.
func (o LookupDeviceResultOutput) AccessPublicIpv6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.AccessPublicIpv6 }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) AlwaysPxe() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDeviceResult) bool { return v.AlwaysPxe }).(pulumi.BoolOutput)
}

// The billing cycle of the device (monthly or hourly).
func (o LookupDeviceResultOutput) BillingCycle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.BillingCycle }).(pulumi.StringOutput)
}

// Description string for the device.
func (o LookupDeviceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

// (**Deprecated**) The facility where the device is deployed. Use metro instead; read the facility to metro migration guide
//
// Deprecated: Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
func (o LookupDeviceResultOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Facility }).(pulumi.StringOutput)
}

// The id of hardware reservation which this device occupies.
func (o LookupDeviceResultOutput) HardwareReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.HardwareReservationId }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) IpxeScriptUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.IpxeScriptUrl }).(pulumi.StringOutput)
}

// The metro where the device is deployed
func (o LookupDeviceResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Metro }).(pulumi.StringOutput)
}

// L2 network type of the device, one of `layer3`, `layer2-bonded`,
// `layer2-individual`, `hybrid`.
func (o LookupDeviceResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

// The device's private and public IP (v4 and v6) network details. See
// Network Attribute below for more details.
func (o LookupDeviceResultOutput) Networks() GetDeviceNetworkArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []GetDeviceNetwork { return v.Networks }).(GetDeviceNetworkArrayOutput)
}

// The operating system running on the device.
func (o LookupDeviceResultOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.OperatingSystem }).(pulumi.StringOutput)
}

// The hardware config of the device.
func (o LookupDeviceResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Plan }).(pulumi.StringOutput)
}

// List of ports assigned to the device. See Ports Attribute below for
// more details.
func (o LookupDeviceResultOutput) Ports() GetDevicePortArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []GetDevicePort { return v.Ports }).(GetDevicePortArrayOutput)
}

func (o LookupDeviceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Root password to the server (if still available).
func (o LookupDeviceResultOutput) RootPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.RootPassword }).(pulumi.StringOutput)
}

// The hostname to use for [Serial over SSH](https://deploy.equinix.com/developers/docs/metal/resilience-recovery/serial-over-ssh/) access to the device
func (o LookupDeviceResultOutput) SosHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.SosHostname }).(pulumi.StringOutput)
}

// List of IDs of SSH keys deployed in the device, can be both user or project SSH keys.
func (o LookupDeviceResultOutput) SshKeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []string { return v.SshKeyIds }).(pulumi.StringArrayOutput)
}

// The state of the device.
func (o LookupDeviceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Storage }).(pulumi.StringOutput)
}

// Tags attached to the device.
func (o LookupDeviceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}
