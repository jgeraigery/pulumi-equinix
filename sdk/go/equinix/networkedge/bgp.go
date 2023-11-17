// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkedge

import (
	"context"
	"reflect"

	"errors"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource `networkedge.Bgp` allows creation and management of Equinix Network
// Edge BGP peering configurations.
//
// ## Example Usage
// ```go
// package main
//
// import (
//
//	"github.com/equinix/pulumi-equinix/sdk/go/equinix/networkedge"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bgp, err := networkedge.NewBgp(ctx, "bgp", &networkedge.BgpArgs{
//				ConnectionId:      pulumi.String("54014acf-9730-4b55-a791-459283d05fb1"),
//				LocalIpAddress:    pulumi.String("10.1.1.1/30"),
//				LocalAsn:          pulumi.Int(12345),
//				RemoteIpAddress:   pulumi.String("10.1.1.2"),
//				RemoteAsn:         pulumi.Int(66123),
//				AuthenticationKey: pulumi.String("secret"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("state", bgp.State)
//			ctx.Export("provisioningStatus", bgp.ProvisioningStatus)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource can be imported using an existing ID: <break><break>```sh<break> $ pulumi import equinix:networkedge/bgp:Bgp example {existing_id} <break>```<break><break>
type Bgp struct {
	pulumi.CustomResourceState

	// shared key used for BGP peer authentication.
	AuthenticationKey pulumi.StringPtrOutput `pulumi:"authenticationKey"`
	// identifier of a connection established between.
	// network device and remote service provider that will be used for peering.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// unique identifier of a network device that is a local peer in a given BGP peering
	// configuration.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Local ASN number.
	LocalAsn pulumi.IntOutput `pulumi:"localAsn"`
	// IP address in CIDR format of a local device.
	LocalIpAddress pulumi.StringOutput `pulumi:"localIpAddress"`
	// BGP peering configuration provisioning status, one of `PROVISIONING`,
	// `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
	ProvisioningStatus pulumi.StringOutput `pulumi:"provisioningStatus"`
	// Remote ASN number.
	RemoteAsn pulumi.IntOutput `pulumi:"remoteAsn"`
	// IP address of remote peer.
	RemoteIpAddress pulumi.StringOutput `pulumi:"remoteIpAddress"`
	// BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
	// `Established`.
	State pulumi.StringOutput `pulumi:"state"`
	// BGP peering configuration unique identifier.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewBgp registers a new resource with the given unique name, arguments, and options.
func NewBgp(ctx *pulumi.Context,
	name string, args *BgpArgs, opts ...pulumi.ResourceOption) (*Bgp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.LocalAsn == nil {
		return nil, errors.New("invalid value for required argument 'LocalAsn'")
	}
	if args.LocalIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'LocalIpAddress'")
	}
	if args.RemoteAsn == nil {
		return nil, errors.New("invalid value for required argument 'RemoteAsn'")
	}
	if args.RemoteIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'RemoteIpAddress'")
	}
	if args.AuthenticationKey != nil {
		args.AuthenticationKey = pulumi.ToSecret(args.AuthenticationKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authenticationKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bgp
	err := ctx.RegisterResource("equinix:networkedge/bgp:Bgp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgp gets an existing Bgp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpState, opts ...pulumi.ResourceOption) (*Bgp, error) {
	var resource Bgp
	err := ctx.ReadResource("equinix:networkedge/bgp:Bgp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bgp resources.
type bgpState struct {
	// shared key used for BGP peer authentication.
	AuthenticationKey *string `pulumi:"authenticationKey"`
	// identifier of a connection established between.
	// network device and remote service provider that will be used for peering.
	ConnectionId *string `pulumi:"connectionId"`
	// unique identifier of a network device that is a local peer in a given BGP peering
	// configuration.
	DeviceId *string `pulumi:"deviceId"`
	// Local ASN number.
	LocalAsn *int `pulumi:"localAsn"`
	// IP address in CIDR format of a local device.
	LocalIpAddress *string `pulumi:"localIpAddress"`
	// BGP peering configuration provisioning status, one of `PROVISIONING`,
	// `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
	ProvisioningStatus *string `pulumi:"provisioningStatus"`
	// Remote ASN number.
	RemoteAsn *int `pulumi:"remoteAsn"`
	// IP address of remote peer.
	RemoteIpAddress *string `pulumi:"remoteIpAddress"`
	// BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
	// `Established`.
	State *string `pulumi:"state"`
	// BGP peering configuration unique identifier.
	Uuid *string `pulumi:"uuid"`
}

type BgpState struct {
	// shared key used for BGP peer authentication.
	AuthenticationKey pulumi.StringPtrInput
	// identifier of a connection established between.
	// network device and remote service provider that will be used for peering.
	ConnectionId pulumi.StringPtrInput
	// unique identifier of a network device that is a local peer in a given BGP peering
	// configuration.
	DeviceId pulumi.StringPtrInput
	// Local ASN number.
	LocalAsn pulumi.IntPtrInput
	// IP address in CIDR format of a local device.
	LocalIpAddress pulumi.StringPtrInput
	// BGP peering configuration provisioning status, one of `PROVISIONING`,
	// `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
	ProvisioningStatus pulumi.StringPtrInput
	// Remote ASN number.
	RemoteAsn pulumi.IntPtrInput
	// IP address of remote peer.
	RemoteIpAddress pulumi.StringPtrInput
	// BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
	// `Established`.
	State pulumi.StringPtrInput
	// BGP peering configuration unique identifier.
	Uuid pulumi.StringPtrInput
}

func (BgpState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpState)(nil)).Elem()
}

type bgpArgs struct {
	// shared key used for BGP peer authentication.
	AuthenticationKey *string `pulumi:"authenticationKey"`
	// identifier of a connection established between.
	// network device and remote service provider that will be used for peering.
	ConnectionId string `pulumi:"connectionId"`
	// Local ASN number.
	LocalAsn int `pulumi:"localAsn"`
	// IP address in CIDR format of a local device.
	LocalIpAddress string `pulumi:"localIpAddress"`
	// Remote ASN number.
	RemoteAsn int `pulumi:"remoteAsn"`
	// IP address of remote peer.
	RemoteIpAddress string `pulumi:"remoteIpAddress"`
}

// The set of arguments for constructing a Bgp resource.
type BgpArgs struct {
	// shared key used for BGP peer authentication.
	AuthenticationKey pulumi.StringPtrInput
	// identifier of a connection established between.
	// network device and remote service provider that will be used for peering.
	ConnectionId pulumi.StringInput
	// Local ASN number.
	LocalAsn pulumi.IntInput
	// IP address in CIDR format of a local device.
	LocalIpAddress pulumi.StringInput
	// Remote ASN number.
	RemoteAsn pulumi.IntInput
	// IP address of remote peer.
	RemoteIpAddress pulumi.StringInput
}

func (BgpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpArgs)(nil)).Elem()
}

type BgpInput interface {
	pulumi.Input

	ToBgpOutput() BgpOutput
	ToBgpOutputWithContext(ctx context.Context) BgpOutput
}

func (*Bgp) ElementType() reflect.Type {
	return reflect.TypeOf((**Bgp)(nil)).Elem()
}

func (i *Bgp) ToBgpOutput() BgpOutput {
	return i.ToBgpOutputWithContext(context.Background())
}

func (i *Bgp) ToBgpOutputWithContext(ctx context.Context) BgpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpOutput)
}

// BgpArrayInput is an input type that accepts BgpArray and BgpArrayOutput values.
// You can construct a concrete instance of `BgpArrayInput` via:
//
//	BgpArray{ BgpArgs{...} }
type BgpArrayInput interface {
	pulumi.Input

	ToBgpArrayOutput() BgpArrayOutput
	ToBgpArrayOutputWithContext(context.Context) BgpArrayOutput
}

type BgpArray []BgpInput

func (BgpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bgp)(nil)).Elem()
}

func (i BgpArray) ToBgpArrayOutput() BgpArrayOutput {
	return i.ToBgpArrayOutputWithContext(context.Background())
}

func (i BgpArray) ToBgpArrayOutputWithContext(ctx context.Context) BgpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpArrayOutput)
}

// BgpMapInput is an input type that accepts BgpMap and BgpMapOutput values.
// You can construct a concrete instance of `BgpMapInput` via:
//
//	BgpMap{ "key": BgpArgs{...} }
type BgpMapInput interface {
	pulumi.Input

	ToBgpMapOutput() BgpMapOutput
	ToBgpMapOutputWithContext(context.Context) BgpMapOutput
}

type BgpMap map[string]BgpInput

func (BgpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bgp)(nil)).Elem()
}

func (i BgpMap) ToBgpMapOutput() BgpMapOutput {
	return i.ToBgpMapOutputWithContext(context.Background())
}

func (i BgpMap) ToBgpMapOutputWithContext(ctx context.Context) BgpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpMapOutput)
}

type BgpOutput struct{ *pulumi.OutputState }

func (BgpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bgp)(nil)).Elem()
}

func (o BgpOutput) ToBgpOutput() BgpOutput {
	return o
}

func (o BgpOutput) ToBgpOutputWithContext(ctx context.Context) BgpOutput {
	return o
}

// shared key used for BGP peer authentication.
func (o BgpOutput) AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringPtrOutput { return v.AuthenticationKey }).(pulumi.StringPtrOutput)
}

// identifier of a connection established between.
// network device and remote service provider that will be used for peering.
func (o BgpOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// unique identifier of a network device that is a local peer in a given BGP peering
// configuration.
func (o BgpOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// Local ASN number.
func (o BgpOutput) LocalAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *Bgp) pulumi.IntOutput { return v.LocalAsn }).(pulumi.IntOutput)
}

// IP address in CIDR format of a local device.
func (o BgpOutput) LocalIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.LocalIpAddress }).(pulumi.StringOutput)
}

// BGP peering configuration provisioning status, one of `PROVISIONING`,
// `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
func (o BgpOutput) ProvisioningStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.ProvisioningStatus }).(pulumi.StringOutput)
}

// Remote ASN number.
func (o BgpOutput) RemoteAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *Bgp) pulumi.IntOutput { return v.RemoteAsn }).(pulumi.IntOutput)
}

// IP address of remote peer.
func (o BgpOutput) RemoteIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.RemoteIpAddress }).(pulumi.StringOutput)
}

// BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
// `Established`.
func (o BgpOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// BGP peering configuration unique identifier.
func (o BgpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Bgp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type BgpArrayOutput struct{ *pulumi.OutputState }

func (BgpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bgp)(nil)).Elem()
}

func (o BgpArrayOutput) ToBgpArrayOutput() BgpArrayOutput {
	return o
}

func (o BgpArrayOutput) ToBgpArrayOutputWithContext(ctx context.Context) BgpArrayOutput {
	return o
}

func (o BgpArrayOutput) Index(i pulumi.IntInput) BgpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bgp {
		return vs[0].([]*Bgp)[vs[1].(int)]
	}).(BgpOutput)
}

type BgpMapOutput struct{ *pulumi.OutputState }

func (BgpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bgp)(nil)).Elem()
}

func (o BgpMapOutput) ToBgpMapOutput() BgpMapOutput {
	return o
}

func (o BgpMapOutput) ToBgpMapOutputWithContext(ctx context.Context) BgpMapOutput {
	return o
}

func (o BgpMapOutput) MapIndex(k pulumi.StringInput) BgpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bgp {
		return vs[0].(map[string]*Bgp)[vs[1].(string)]
	}).(BgpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpInput)(nil)).Elem(), &Bgp{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpArrayInput)(nil)).Elem(), BgpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpMapInput)(nil)).Elem(), BgpMap{})
	pulumi.RegisterOutputType(BgpOutput{})
	pulumi.RegisterOutputType(BgpArrayOutput{})
	pulumi.RegisterOutputType(BgpMapOutput{})
}
