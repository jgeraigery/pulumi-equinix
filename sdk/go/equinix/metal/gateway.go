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

// Use this resource to create Metal Gateway resources in Equinix Metal.
//
// > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
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
//			projectId := cfg.Require("projectId")
//			vlanId := cfg.Require("vlanId")
//			gateway, err := metal.NewGateway(ctx, "gateway", &metal.GatewayArgs{
//				ProjectId:             pulumi.String(projectId),
//				VlanId:                pulumi.String(vlanId),
//				PrivateIpv4SubnetSize: pulumi.Int(8),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("gatewayState", gateway.State)
//			return nil
//		})
//	}
//
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
	IpReservationId pulumi.StringPtrOutput `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
	PrivateIpv4SubnetSize pulumi.IntOutput `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Status of the gateway resource.
	State pulumi.StringOutput `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to.
	VlanId pulumi.StringOutput `pulumi:"vlanId"`
	// UUID of the VRF associated with the IP Reservation
	VrfId pulumi.StringOutput `pulumi:"vrfId"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("equinix:metal/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("equinix:metal/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
	IpReservationId *string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
	PrivateIpv4SubnetSize *int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to.
	ProjectId *string `pulumi:"projectId"`
	// Status of the gateway resource.
	State *string `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to.
	VlanId *string `pulumi:"vlanId"`
	// UUID of the VRF associated with the IP Reservation
	VrfId *string `pulumi:"vrfId"`
}

type GatewayState struct {
	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
	IpReservationId pulumi.StringPtrInput
	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
	PrivateIpv4SubnetSize pulumi.IntPtrInput
	// UUID of the project where the gateway is scoped to.
	ProjectId pulumi.StringPtrInput
	// Status of the gateway resource.
	State pulumi.StringPtrInput
	// UUID of the VLAN where the gateway is scoped to.
	VlanId pulumi.StringPtrInput
	// UUID of the VRF associated with the IP Reservation
	VrfId pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
	IpReservationId *string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
	PrivateIpv4SubnetSize *int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to.
	ProjectId string `pulumi:"projectId"`
	// UUID of the VLAN where the gateway is scoped to.
	VlanId string `pulumi:"vlanId"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
	IpReservationId pulumi.StringPtrInput
	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
	PrivateIpv4SubnetSize pulumi.IntPtrInput
	// UUID of the project where the gateway is scoped to.
	ProjectId pulumi.StringInput
	// UUID of the VLAN where the gateway is scoped to.
	VlanId pulumi.StringInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

func (i *Gateway) ToOutput(ctx context.Context) pulumix.Output[*Gateway] {
	return pulumix.Output[*Gateway]{
		OutputState: i.ToGatewayOutputWithContext(ctx).OutputState,
	}
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//	GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

func (i GatewayArray) ToOutput(ctx context.Context) pulumix.Output[[]*Gateway] {
	return pulumix.Output[[]*Gateway]{
		OutputState: i.ToGatewayArrayOutputWithContext(ctx).OutputState,
	}
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//	GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

func (i GatewayMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Gateway] {
	return pulumix.Output[map[string]*Gateway]{
		OutputState: i.ToGatewayMapOutputWithContext(ctx).OutputState,
	}
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*Gateway] {
	return pulumix.Output[*Gateway]{
		OutputState: o.OutputState,
	}
}

// UUID of Public or VRF IP Reservation to associate with the gateway, the
// reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`.
func (o GatewayOutput) IpReservationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.IpReservationId }).(pulumi.StringPtrOutput)
}

// Size of the private IPv4 subnet to create for this metal
// gateway, must be one of `8`, `16`, `32`, `64`, `128`. Conflicts with `ipReservationId`.
func (o GatewayOutput) PrivateIpv4SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.PrivateIpv4SubnetSize }).(pulumi.IntOutput)
}

// UUID of the project where the gateway is scoped to.
func (o GatewayOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Status of the gateway resource.
func (o GatewayOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// UUID of the VLAN where the gateway is scoped to.
func (o GatewayOutput) VlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VlanId }).(pulumi.StringOutput)
}

// UUID of the VRF associated with the IP Reservation
func (o GatewayOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VrfId }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Gateway] {
	return pulumix.Output[[]*Gateway]{
		OutputState: o.OutputState,
	}
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Gateway] {
	return pulumix.Output[map[string]*Gateway]{
		OutputState: o.OutputState,
	}
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
