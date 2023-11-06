// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fabric

import (
	"context"
	"reflect"

	"errors"
	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ```go
// package main
//
// import (
//
//	"github.com/equinix/pulumi-equinix/sdk/go/equinix/fabric"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			connectionId := cfg.Require("connectionId")
//			routingProtocol, err := fabric.NewRoutingProtocol(ctx, "RoutingProtocol", &fabric.RoutingProtocolArgs{
//				ConnectionUuid: pulumi.String(connectionId),
//				Name:           pulumi.String("My-Direct-route-1"),
//				Type:           pulumi.String("DIRECT"),
//				DirectIpv4: &fabric.RoutingProtocolDirectIpv4Args{
//					EquinixIfaceIp: pulumi.String("192.168.100.1/30"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("routingProtocolId", routingProtocol.ID())
//			ctx.Export("routingProtocolState", routingProtocol.State)
//			ctx.Export("routingProtocolEquinixAsn", routingProtocol.EquinixAsn)
//			return nil
//		})
//	}
//
// ```
type RoutingProtocol struct {
	pulumi.CustomResourceState

	// Bidirectional Forwarding Detection
	Bfd RoutingProtocolBfdPtrOutput `pulumi:"bfd"`
	// BGP authorization key
	BgpAuthKey pulumi.StringPtrOutput `pulumi:"bgpAuthKey"`
	// Routing Protocol BGP IPv4
	BgpIpv4 RoutingProtocolBgpIpv4PtrOutput `pulumi:"bgpIpv4"`
	// Routing Protocol BGP IPv6
	BgpIpv6 RoutingProtocolBgpIpv6PtrOutput `pulumi:"bgpIpv6"`
	// Captures Routing Protocol lifecycle change information
	ChangeLogs RoutingProtocolChangeLogArrayOutput `pulumi:"changeLogs"`
	// Routing Protocol configuration Changes
	Changes RoutingProtocolChangeArrayOutput `pulumi:"changes"`
	// Connection URI associated with Routing Protocol
	ConnectionUuid pulumi.StringOutput `pulumi:"connectionUuid"`
	// Customer-provided ASN
	CustomerAsn pulumi.IntPtrOutput `pulumi:"customerAsn"`
	// Customer-provided Fabric Routing Protocol description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Routing Protocol Direct IPv4
	DirectIpv4 RoutingProtocolDirectIpv4PtrOutput `pulumi:"directIpv4"`
	// Routing Protocol Direct IPv6
	DirectIpv6 RoutingProtocolDirectIpv6PtrOutput `pulumi:"directIpv6"`
	// Equinix ASN
	EquinixAsn pulumi.IntOutput `pulumi:"equinixAsn"`
	// Routing Protocol URI information
	Href pulumi.StringOutput `pulumi:"href"`
	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name pulumi.StringOutput `pulumi:"name"`
	// Routing Protocol type-specific operational data
	Operations RoutingProtocolOperationArrayOutput `pulumi:"operations"`
	// Routing Protocol overall state
	State pulumi.StringOutput `pulumi:"state"`
	// Defines the routing protocol type like BGP or DIRECT
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Equinix-assigned routing protocol identifier
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewRoutingProtocol registers a new resource with the given unique name, arguments, and options.
func NewRoutingProtocol(ctx *pulumi.Context,
	name string, args *RoutingProtocolArgs, opts ...pulumi.ResourceOption) (*RoutingProtocol, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionUuid == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionUuid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoutingProtocol
	err := ctx.RegisterResource("equinix:fabric/routingProtocol:RoutingProtocol", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutingProtocol gets an existing RoutingProtocol resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutingProtocol(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutingProtocolState, opts ...pulumi.ResourceOption) (*RoutingProtocol, error) {
	var resource RoutingProtocol
	err := ctx.ReadResource("equinix:fabric/routingProtocol:RoutingProtocol", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoutingProtocol resources.
type routingProtocolState struct {
	// Bidirectional Forwarding Detection
	Bfd *RoutingProtocolBfd `pulumi:"bfd"`
	// BGP authorization key
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// Routing Protocol BGP IPv4
	BgpIpv4 *RoutingProtocolBgpIpv4 `pulumi:"bgpIpv4"`
	// Routing Protocol BGP IPv6
	BgpIpv6 *RoutingProtocolBgpIpv6 `pulumi:"bgpIpv6"`
	// Captures Routing Protocol lifecycle change information
	ChangeLogs []RoutingProtocolChangeLog `pulumi:"changeLogs"`
	// Routing Protocol configuration Changes
	Changes []RoutingProtocolChange `pulumi:"changes"`
	// Connection URI associated with Routing Protocol
	ConnectionUuid *string `pulumi:"connectionUuid"`
	// Customer-provided ASN
	CustomerAsn *int `pulumi:"customerAsn"`
	// Customer-provided Fabric Routing Protocol description
	Description *string `pulumi:"description"`
	// Routing Protocol Direct IPv4
	DirectIpv4 *RoutingProtocolDirectIpv4 `pulumi:"directIpv4"`
	// Routing Protocol Direct IPv6
	DirectIpv6 *RoutingProtocolDirectIpv6 `pulumi:"directIpv6"`
	// Equinix ASN
	EquinixAsn *int `pulumi:"equinixAsn"`
	// Routing Protocol URI information
	Href *string `pulumi:"href"`
	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name *string `pulumi:"name"`
	// Routing Protocol type-specific operational data
	Operations []RoutingProtocolOperation `pulumi:"operations"`
	// Routing Protocol overall state
	State *string `pulumi:"state"`
	// Defines the routing protocol type like BGP or DIRECT
	Type *string `pulumi:"type"`
	// Equinix-assigned routing protocol identifier
	Uuid *string `pulumi:"uuid"`
}

type RoutingProtocolState struct {
	// Bidirectional Forwarding Detection
	Bfd RoutingProtocolBfdPtrInput
	// BGP authorization key
	BgpAuthKey pulumi.StringPtrInput
	// Routing Protocol BGP IPv4
	BgpIpv4 RoutingProtocolBgpIpv4PtrInput
	// Routing Protocol BGP IPv6
	BgpIpv6 RoutingProtocolBgpIpv6PtrInput
	// Captures Routing Protocol lifecycle change information
	ChangeLogs RoutingProtocolChangeLogArrayInput
	// Routing Protocol configuration Changes
	Changes RoutingProtocolChangeArrayInput
	// Connection URI associated with Routing Protocol
	ConnectionUuid pulumi.StringPtrInput
	// Customer-provided ASN
	CustomerAsn pulumi.IntPtrInput
	// Customer-provided Fabric Routing Protocol description
	Description pulumi.StringPtrInput
	// Routing Protocol Direct IPv4
	DirectIpv4 RoutingProtocolDirectIpv4PtrInput
	// Routing Protocol Direct IPv6
	DirectIpv6 RoutingProtocolDirectIpv6PtrInput
	// Equinix ASN
	EquinixAsn pulumi.IntPtrInput
	// Routing Protocol URI information
	Href pulumi.StringPtrInput
	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name pulumi.StringPtrInput
	// Routing Protocol type-specific operational data
	Operations RoutingProtocolOperationArrayInput
	// Routing Protocol overall state
	State pulumi.StringPtrInput
	// Defines the routing protocol type like BGP or DIRECT
	Type pulumi.StringPtrInput
	// Equinix-assigned routing protocol identifier
	Uuid pulumi.StringPtrInput
}

func (RoutingProtocolState) ElementType() reflect.Type {
	return reflect.TypeOf((*routingProtocolState)(nil)).Elem()
}

type routingProtocolArgs struct {
	// Bidirectional Forwarding Detection
	Bfd *RoutingProtocolBfd `pulumi:"bfd"`
	// BGP authorization key
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// Routing Protocol BGP IPv4
	BgpIpv4 *RoutingProtocolBgpIpv4 `pulumi:"bgpIpv4"`
	// Routing Protocol BGP IPv6
	BgpIpv6 *RoutingProtocolBgpIpv6 `pulumi:"bgpIpv6"`
	// Connection URI associated with Routing Protocol
	ConnectionUuid string `pulumi:"connectionUuid"`
	// Customer-provided ASN
	CustomerAsn *int `pulumi:"customerAsn"`
	// Customer-provided Fabric Routing Protocol description
	Description *string `pulumi:"description"`
	// Routing Protocol Direct IPv4
	DirectIpv4 *RoutingProtocolDirectIpv4 `pulumi:"directIpv4"`
	// Routing Protocol Direct IPv6
	DirectIpv6 *RoutingProtocolDirectIpv6 `pulumi:"directIpv6"`
	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name *string `pulumi:"name"`
	// Defines the routing protocol type like BGP or DIRECT
	Type *string `pulumi:"type"`
	// Equinix-assigned routing protocol identifier
	Uuid *string `pulumi:"uuid"`
}

// The set of arguments for constructing a RoutingProtocol resource.
type RoutingProtocolArgs struct {
	// Bidirectional Forwarding Detection
	Bfd RoutingProtocolBfdPtrInput
	// BGP authorization key
	BgpAuthKey pulumi.StringPtrInput
	// Routing Protocol BGP IPv4
	BgpIpv4 RoutingProtocolBgpIpv4PtrInput
	// Routing Protocol BGP IPv6
	BgpIpv6 RoutingProtocolBgpIpv6PtrInput
	// Connection URI associated with Routing Protocol
	ConnectionUuid pulumi.StringInput
	// Customer-provided ASN
	CustomerAsn pulumi.IntPtrInput
	// Customer-provided Fabric Routing Protocol description
	Description pulumi.StringPtrInput
	// Routing Protocol Direct IPv4
	DirectIpv4 RoutingProtocolDirectIpv4PtrInput
	// Routing Protocol Direct IPv6
	DirectIpv6 RoutingProtocolDirectIpv6PtrInput
	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name pulumi.StringPtrInput
	// Defines the routing protocol type like BGP or DIRECT
	Type pulumi.StringPtrInput
	// Equinix-assigned routing protocol identifier
	Uuid pulumi.StringPtrInput
}

func (RoutingProtocolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routingProtocolArgs)(nil)).Elem()
}

type RoutingProtocolInput interface {
	pulumi.Input

	ToRoutingProtocolOutput() RoutingProtocolOutput
	ToRoutingProtocolOutputWithContext(ctx context.Context) RoutingProtocolOutput
}

func (*RoutingProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProtocol)(nil)).Elem()
}

func (i *RoutingProtocol) ToRoutingProtocolOutput() RoutingProtocolOutput {
	return i.ToRoutingProtocolOutputWithContext(context.Background())
}

func (i *RoutingProtocol) ToRoutingProtocolOutputWithContext(ctx context.Context) RoutingProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingProtocolOutput)
}

func (i *RoutingProtocol) ToOutput(ctx context.Context) pulumix.Output[*RoutingProtocol] {
	return pulumix.Output[*RoutingProtocol]{
		OutputState: i.ToRoutingProtocolOutputWithContext(ctx).OutputState,
	}
}

// RoutingProtocolArrayInput is an input type that accepts RoutingProtocolArray and RoutingProtocolArrayOutput values.
// You can construct a concrete instance of `RoutingProtocolArrayInput` via:
//
//	RoutingProtocolArray{ RoutingProtocolArgs{...} }
type RoutingProtocolArrayInput interface {
	pulumi.Input

	ToRoutingProtocolArrayOutput() RoutingProtocolArrayOutput
	ToRoutingProtocolArrayOutputWithContext(context.Context) RoutingProtocolArrayOutput
}

type RoutingProtocolArray []RoutingProtocolInput

func (RoutingProtocolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoutingProtocol)(nil)).Elem()
}

func (i RoutingProtocolArray) ToRoutingProtocolArrayOutput() RoutingProtocolArrayOutput {
	return i.ToRoutingProtocolArrayOutputWithContext(context.Background())
}

func (i RoutingProtocolArray) ToRoutingProtocolArrayOutputWithContext(ctx context.Context) RoutingProtocolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingProtocolArrayOutput)
}

func (i RoutingProtocolArray) ToOutput(ctx context.Context) pulumix.Output[[]*RoutingProtocol] {
	return pulumix.Output[[]*RoutingProtocol]{
		OutputState: i.ToRoutingProtocolArrayOutputWithContext(ctx).OutputState,
	}
}

// RoutingProtocolMapInput is an input type that accepts RoutingProtocolMap and RoutingProtocolMapOutput values.
// You can construct a concrete instance of `RoutingProtocolMapInput` via:
//
//	RoutingProtocolMap{ "key": RoutingProtocolArgs{...} }
type RoutingProtocolMapInput interface {
	pulumi.Input

	ToRoutingProtocolMapOutput() RoutingProtocolMapOutput
	ToRoutingProtocolMapOutputWithContext(context.Context) RoutingProtocolMapOutput
}

type RoutingProtocolMap map[string]RoutingProtocolInput

func (RoutingProtocolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoutingProtocol)(nil)).Elem()
}

func (i RoutingProtocolMap) ToRoutingProtocolMapOutput() RoutingProtocolMapOutput {
	return i.ToRoutingProtocolMapOutputWithContext(context.Background())
}

func (i RoutingProtocolMap) ToRoutingProtocolMapOutputWithContext(ctx context.Context) RoutingProtocolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingProtocolMapOutput)
}

func (i RoutingProtocolMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RoutingProtocol] {
	return pulumix.Output[map[string]*RoutingProtocol]{
		OutputState: i.ToRoutingProtocolMapOutputWithContext(ctx).OutputState,
	}
}

type RoutingProtocolOutput struct{ *pulumi.OutputState }

func (RoutingProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProtocol)(nil)).Elem()
}

func (o RoutingProtocolOutput) ToRoutingProtocolOutput() RoutingProtocolOutput {
	return o
}

func (o RoutingProtocolOutput) ToRoutingProtocolOutputWithContext(ctx context.Context) RoutingProtocolOutput {
	return o
}

func (o RoutingProtocolOutput) ToOutput(ctx context.Context) pulumix.Output[*RoutingProtocol] {
	return pulumix.Output[*RoutingProtocol]{
		OutputState: o.OutputState,
	}
}

// Bidirectional Forwarding Detection
func (o RoutingProtocolOutput) Bfd() RoutingProtocolBfdPtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolBfdPtrOutput { return v.Bfd }).(RoutingProtocolBfdPtrOutput)
}

// BGP authorization key
func (o RoutingProtocolOutput) BgpAuthKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringPtrOutput { return v.BgpAuthKey }).(pulumi.StringPtrOutput)
}

// Routing Protocol BGP IPv4
func (o RoutingProtocolOutput) BgpIpv4() RoutingProtocolBgpIpv4PtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolBgpIpv4PtrOutput { return v.BgpIpv4 }).(RoutingProtocolBgpIpv4PtrOutput)
}

// Routing Protocol BGP IPv6
func (o RoutingProtocolOutput) BgpIpv6() RoutingProtocolBgpIpv6PtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolBgpIpv6PtrOutput { return v.BgpIpv6 }).(RoutingProtocolBgpIpv6PtrOutput)
}

// Captures Routing Protocol lifecycle change information
func (o RoutingProtocolOutput) ChangeLogs() RoutingProtocolChangeLogArrayOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolChangeLogArrayOutput { return v.ChangeLogs }).(RoutingProtocolChangeLogArrayOutput)
}

// Routing Protocol configuration Changes
func (o RoutingProtocolOutput) Changes() RoutingProtocolChangeArrayOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolChangeArrayOutput { return v.Changes }).(RoutingProtocolChangeArrayOutput)
}

// Connection URI associated with Routing Protocol
func (o RoutingProtocolOutput) ConnectionUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringOutput { return v.ConnectionUuid }).(pulumi.StringOutput)
}

// Customer-provided ASN
func (o RoutingProtocolOutput) CustomerAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.IntPtrOutput { return v.CustomerAsn }).(pulumi.IntPtrOutput)
}

// Customer-provided Fabric Routing Protocol description
func (o RoutingProtocolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Routing Protocol Direct IPv4
func (o RoutingProtocolOutput) DirectIpv4() RoutingProtocolDirectIpv4PtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolDirectIpv4PtrOutput { return v.DirectIpv4 }).(RoutingProtocolDirectIpv4PtrOutput)
}

// Routing Protocol Direct IPv6
func (o RoutingProtocolOutput) DirectIpv6() RoutingProtocolDirectIpv6PtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolDirectIpv6PtrOutput { return v.DirectIpv6 }).(RoutingProtocolDirectIpv6PtrOutput)
}

// Equinix ASN
func (o RoutingProtocolOutput) EquinixAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.IntOutput { return v.EquinixAsn }).(pulumi.IntOutput)
}

// Routing Protocol URI information
func (o RoutingProtocolOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringOutput { return v.Href }).(pulumi.StringOutput)
}

// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
func (o RoutingProtocolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Routing Protocol type-specific operational data
func (o RoutingProtocolOutput) Operations() RoutingProtocolOperationArrayOutput {
	return o.ApplyT(func(v *RoutingProtocol) RoutingProtocolOperationArrayOutput { return v.Operations }).(RoutingProtocolOperationArrayOutput)
}

// Routing Protocol overall state
func (o RoutingProtocolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Defines the routing protocol type like BGP or DIRECT
func (o RoutingProtocolOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Equinix-assigned routing protocol identifier
func (o RoutingProtocolOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProtocol) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type RoutingProtocolArrayOutput struct{ *pulumi.OutputState }

func (RoutingProtocolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoutingProtocol)(nil)).Elem()
}

func (o RoutingProtocolArrayOutput) ToRoutingProtocolArrayOutput() RoutingProtocolArrayOutput {
	return o
}

func (o RoutingProtocolArrayOutput) ToRoutingProtocolArrayOutputWithContext(ctx context.Context) RoutingProtocolArrayOutput {
	return o
}

func (o RoutingProtocolArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RoutingProtocol] {
	return pulumix.Output[[]*RoutingProtocol]{
		OutputState: o.OutputState,
	}
}

func (o RoutingProtocolArrayOutput) Index(i pulumi.IntInput) RoutingProtocolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoutingProtocol {
		return vs[0].([]*RoutingProtocol)[vs[1].(int)]
	}).(RoutingProtocolOutput)
}

type RoutingProtocolMapOutput struct{ *pulumi.OutputState }

func (RoutingProtocolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoutingProtocol)(nil)).Elem()
}

func (o RoutingProtocolMapOutput) ToRoutingProtocolMapOutput() RoutingProtocolMapOutput {
	return o
}

func (o RoutingProtocolMapOutput) ToRoutingProtocolMapOutputWithContext(ctx context.Context) RoutingProtocolMapOutput {
	return o
}

func (o RoutingProtocolMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RoutingProtocol] {
	return pulumix.Output[map[string]*RoutingProtocol]{
		OutputState: o.OutputState,
	}
}

func (o RoutingProtocolMapOutput) MapIndex(k pulumi.StringInput) RoutingProtocolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoutingProtocol {
		return vs[0].(map[string]*RoutingProtocol)[vs[1].(string)]
	}).(RoutingProtocolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingProtocolInput)(nil)).Elem(), &RoutingProtocol{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingProtocolArrayInput)(nil)).Elem(), RoutingProtocolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingProtocolMapInput)(nil)).Elem(), RoutingProtocolMap{})
	pulumi.RegisterOutputType(RoutingProtocolOutput{})
	pulumi.RegisterOutputType(RoutingProtocolArrayOutput{})
	pulumi.RegisterOutputType(RoutingProtocolMapOutput{})
}