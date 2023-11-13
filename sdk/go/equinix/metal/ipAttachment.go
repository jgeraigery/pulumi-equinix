// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package metal

import (
	"context"
	"reflect"

	"errors"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to attach elastic IP subnets to devices.
//
// To attach an IP subnet from a reserved block to a provisioned device, you must derive a subnet CIDR
// belonging to one of your reserved blocks in the same project and metro as the target device.
//
// For example, you have reserved IPv4 address block `147.229.10.152/30`, you can choose to assign
// either the whole block as one subnet to a device; or 2 subnets with CIDRs `147.229.10.152/31` and
// `147.229.10.154/31`; or 4 subnets with mask prefix length `32`. More about the elastic IP subnets
// is [here](https://metal.equinix.com/developers/docs/networking/elastic-ips/).
//
// Device and reserved block must be in the same metro.
//
// ## Example Usage
// ```go
// package main
//
// import (
//
//	"github.com/equinix/pulumi-equinix/sdk/go/equinix/metal"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			deviceId := cfg.Require("deviceId")
//			subnetCidr := "147.229.10.152/31"
//			if param := cfg.Get("subnetCidr"); param != "" {
//				subnetCidr = param
//			}
//			ipAttachResource, err := metal.NewIpAttachment(ctx, "ipAttach", &metal.IpAttachmentArgs{
//				DeviceId:     pulumi.String(deviceId),
//				CidrNotation: pulumi.String(subnetCidr),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("ipAttach", ipAttachResource.ID())
//			ctx.Export("ipNetmask", ipAttachResource.Netmask)
//			return nil
//		})
//	}
//
// ```
type IpAttachment struct {
	pulumi.CustomResourceState

	Address pulumi.StringOutput `pulumi:"address"`
	// Address family as integer. One of `4` or `6`.
	AddressFamily pulumi.IntOutput `pulumi:"addressFamily"`
	// Length of CIDR prefix of the subnet as integer.
	Cidr pulumi.IntOutput `pulumi:"cidr"`
	// CIDR notation of subnet from block reserved in the same project
	// and metro as the device.
	CidrNotation pulumi.StringOutput `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// IP address of gateway for the subnet.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Flag indicating whether IP block is global, i.e. assignable in any location
	Global     pulumi.BoolOutput `pulumi:"global"`
	Manageable pulumi.BoolOutput `pulumi:"manageable"`
	Management pulumi.BoolOutput `pulumi:"management"`
	// Subnet mask in decimal notation, e.g., `255.255.255.0`.
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// Subnet network address.
	Network pulumi.StringOutput `pulumi:"network"`
	// Boolean flag whether subnet is reachable from the Internet.
	Public pulumi.BoolOutput   `pulumi:"public"`
	VrfId  pulumi.StringOutput `pulumi:"vrfId"`
}

// NewIpAttachment registers a new resource with the given unique name, arguments, and options.
func NewIpAttachment(ctx *pulumi.Context,
	name string, args *IpAttachmentArgs, opts ...pulumi.ResourceOption) (*IpAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrNotation == nil {
		return nil, errors.New("invalid value for required argument 'CidrNotation'")
	}
	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpAttachment
	err := ctx.RegisterResource("equinix:metal/ipAttachment:IpAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAttachment gets an existing IpAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAttachmentState, opts ...pulumi.ResourceOption) (*IpAttachment, error) {
	var resource IpAttachment
	err := ctx.ReadResource("equinix:metal/ipAttachment:IpAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAttachment resources.
type ipAttachmentState struct {
	Address *string `pulumi:"address"`
	// Address family as integer. One of `4` or `6`.
	AddressFamily *int `pulumi:"addressFamily"`
	// Length of CIDR prefix of the subnet as integer.
	Cidr *int `pulumi:"cidr"`
	// CIDR notation of subnet from block reserved in the same project
	// and metro as the device.
	CidrNotation *string `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet.
	DeviceId *string `pulumi:"deviceId"`
	// IP address of gateway for the subnet.
	Gateway *string `pulumi:"gateway"`
	// Flag indicating whether IP block is global, i.e. assignable in any location
	Global     *bool `pulumi:"global"`
	Manageable *bool `pulumi:"manageable"`
	Management *bool `pulumi:"management"`
	// Subnet mask in decimal notation, e.g., `255.255.255.0`.
	Netmask *string `pulumi:"netmask"`
	// Subnet network address.
	Network *string `pulumi:"network"`
	// Boolean flag whether subnet is reachable from the Internet.
	Public *bool   `pulumi:"public"`
	VrfId  *string `pulumi:"vrfId"`
}

type IpAttachmentState struct {
	Address pulumi.StringPtrInput
	// Address family as integer. One of `4` or `6`.
	AddressFamily pulumi.IntPtrInput
	// Length of CIDR prefix of the subnet as integer.
	Cidr pulumi.IntPtrInput
	// CIDR notation of subnet from block reserved in the same project
	// and metro as the device.
	CidrNotation pulumi.StringPtrInput
	// ID of device to which to assign the subnet.
	DeviceId pulumi.StringPtrInput
	// IP address of gateway for the subnet.
	Gateway pulumi.StringPtrInput
	// Flag indicating whether IP block is global, i.e. assignable in any location
	Global     pulumi.BoolPtrInput
	Manageable pulumi.BoolPtrInput
	Management pulumi.BoolPtrInput
	// Subnet mask in decimal notation, e.g., `255.255.255.0`.
	Netmask pulumi.StringPtrInput
	// Subnet network address.
	Network pulumi.StringPtrInput
	// Boolean flag whether subnet is reachable from the Internet.
	Public pulumi.BoolPtrInput
	VrfId  pulumi.StringPtrInput
}

func (IpAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAttachmentState)(nil)).Elem()
}

type ipAttachmentArgs struct {
	// CIDR notation of subnet from block reserved in the same project
	// and metro as the device.
	CidrNotation string `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet.
	DeviceId string `pulumi:"deviceId"`
}

// The set of arguments for constructing a IpAttachment resource.
type IpAttachmentArgs struct {
	// CIDR notation of subnet from block reserved in the same project
	// and metro as the device.
	CidrNotation pulumi.StringInput
	// ID of device to which to assign the subnet.
	DeviceId pulumi.StringInput
}

func (IpAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAttachmentArgs)(nil)).Elem()
}

type IpAttachmentInput interface {
	pulumi.Input

	ToIpAttachmentOutput() IpAttachmentOutput
	ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput
}

func (*IpAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAttachment)(nil)).Elem()
}

func (i *IpAttachment) ToIpAttachmentOutput() IpAttachmentOutput {
	return i.ToIpAttachmentOutputWithContext(context.Background())
}

func (i *IpAttachment) ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentOutput)
}

func (i *IpAttachment) ToOutput(ctx context.Context) pulumix.Output[*IpAttachment] {
	return pulumix.Output[*IpAttachment]{
		OutputState: i.ToIpAttachmentOutputWithContext(ctx).OutputState,
	}
}

// IpAttachmentArrayInput is an input type that accepts IpAttachmentArray and IpAttachmentArrayOutput values.
// You can construct a concrete instance of `IpAttachmentArrayInput` via:
//
//	IpAttachmentArray{ IpAttachmentArgs{...} }
type IpAttachmentArrayInput interface {
	pulumi.Input

	ToIpAttachmentArrayOutput() IpAttachmentArrayOutput
	ToIpAttachmentArrayOutputWithContext(context.Context) IpAttachmentArrayOutput
}

type IpAttachmentArray []IpAttachmentInput

func (IpAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAttachment)(nil)).Elem()
}

func (i IpAttachmentArray) ToIpAttachmentArrayOutput() IpAttachmentArrayOutput {
	return i.ToIpAttachmentArrayOutputWithContext(context.Background())
}

func (i IpAttachmentArray) ToIpAttachmentArrayOutputWithContext(ctx context.Context) IpAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentArrayOutput)
}

func (i IpAttachmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*IpAttachment] {
	return pulumix.Output[[]*IpAttachment]{
		OutputState: i.ToIpAttachmentArrayOutputWithContext(ctx).OutputState,
	}
}

// IpAttachmentMapInput is an input type that accepts IpAttachmentMap and IpAttachmentMapOutput values.
// You can construct a concrete instance of `IpAttachmentMapInput` via:
//
//	IpAttachmentMap{ "key": IpAttachmentArgs{...} }
type IpAttachmentMapInput interface {
	pulumi.Input

	ToIpAttachmentMapOutput() IpAttachmentMapOutput
	ToIpAttachmentMapOutputWithContext(context.Context) IpAttachmentMapOutput
}

type IpAttachmentMap map[string]IpAttachmentInput

func (IpAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAttachment)(nil)).Elem()
}

func (i IpAttachmentMap) ToIpAttachmentMapOutput() IpAttachmentMapOutput {
	return i.ToIpAttachmentMapOutputWithContext(context.Background())
}

func (i IpAttachmentMap) ToIpAttachmentMapOutputWithContext(ctx context.Context) IpAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentMapOutput)
}

func (i IpAttachmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IpAttachment] {
	return pulumix.Output[map[string]*IpAttachment]{
		OutputState: i.ToIpAttachmentMapOutputWithContext(ctx).OutputState,
	}
}

type IpAttachmentOutput struct{ *pulumi.OutputState }

func (IpAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAttachment)(nil)).Elem()
}

func (o IpAttachmentOutput) ToIpAttachmentOutput() IpAttachmentOutput {
	return o
}

func (o IpAttachmentOutput) ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput {
	return o
}

func (o IpAttachmentOutput) ToOutput(ctx context.Context) pulumix.Output[*IpAttachment] {
	return pulumix.Output[*IpAttachment]{
		OutputState: o.OutputState,
	}
}

func (o IpAttachmentOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Address family as integer. One of `4` or `6`.
func (o IpAttachmentOutput) AddressFamily() pulumi.IntOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.IntOutput { return v.AddressFamily }).(pulumi.IntOutput)
}

// Length of CIDR prefix of the subnet as integer.
func (o IpAttachmentOutput) Cidr() pulumi.IntOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.IntOutput { return v.Cidr }).(pulumi.IntOutput)
}

// CIDR notation of subnet from block reserved in the same project
// and metro as the device.
func (o IpAttachmentOutput) CidrNotation() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.CidrNotation }).(pulumi.StringOutput)
}

// ID of device to which to assign the subnet.
func (o IpAttachmentOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// IP address of gateway for the subnet.
func (o IpAttachmentOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Flag indicating whether IP block is global, i.e. assignable in any location
func (o IpAttachmentOutput) Global() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.BoolOutput { return v.Global }).(pulumi.BoolOutput)
}

func (o IpAttachmentOutput) Manageable() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.BoolOutput { return v.Manageable }).(pulumi.BoolOutput)
}

func (o IpAttachmentOutput) Management() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.BoolOutput { return v.Management }).(pulumi.BoolOutput)
}

// Subnet mask in decimal notation, e.g., `255.255.255.0`.
func (o IpAttachmentOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// Subnet network address.
func (o IpAttachmentOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// Boolean flag whether subnet is reachable from the Internet.
func (o IpAttachmentOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.BoolOutput { return v.Public }).(pulumi.BoolOutput)
}

func (o IpAttachmentOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAttachment) pulumi.StringOutput { return v.VrfId }).(pulumi.StringOutput)
}

type IpAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IpAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAttachment)(nil)).Elem()
}

func (o IpAttachmentArrayOutput) ToIpAttachmentArrayOutput() IpAttachmentArrayOutput {
	return o
}

func (o IpAttachmentArrayOutput) ToIpAttachmentArrayOutputWithContext(ctx context.Context) IpAttachmentArrayOutput {
	return o
}

func (o IpAttachmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IpAttachment] {
	return pulumix.Output[[]*IpAttachment]{
		OutputState: o.OutputState,
	}
}

func (o IpAttachmentArrayOutput) Index(i pulumi.IntInput) IpAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpAttachment {
		return vs[0].([]*IpAttachment)[vs[1].(int)]
	}).(IpAttachmentOutput)
}

type IpAttachmentMapOutput struct{ *pulumi.OutputState }

func (IpAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAttachment)(nil)).Elem()
}

func (o IpAttachmentMapOutput) ToIpAttachmentMapOutput() IpAttachmentMapOutput {
	return o
}

func (o IpAttachmentMapOutput) ToIpAttachmentMapOutputWithContext(ctx context.Context) IpAttachmentMapOutput {
	return o
}

func (o IpAttachmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IpAttachment] {
	return pulumix.Output[map[string]*IpAttachment]{
		OutputState: o.OutputState,
	}
}

func (o IpAttachmentMapOutput) MapIndex(k pulumi.StringInput) IpAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpAttachment {
		return vs[0].(map[string]*IpAttachment)[vs[1].(string)]
	}).(IpAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpAttachmentInput)(nil)).Elem(), &IpAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAttachmentArrayInput)(nil)).Elem(), IpAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAttachmentMapInput)(nil)).Elem(), IpAttachmentMap{})
	pulumi.RegisterOutputType(IpAttachmentOutput{})
	pulumi.RegisterOutputType(IpAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IpAttachmentMapOutput{})
}
