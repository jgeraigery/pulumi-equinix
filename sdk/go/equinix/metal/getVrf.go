// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package metal

import (
	"context"
	"reflect"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve a VRF resource.
//
// > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
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
//			_, err := metal.LookupVrf(ctx, &metal.LookupVrfArgs{
//				VrfId: "48630899-9ff2-4ce6-a93f-50ff4ebcdf6e",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVrf(ctx *pulumi.Context, args *LookupVrfArgs, opts ...pulumi.InvokeOption) (*LookupVrfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVrfResult
	err := ctx.Invoke("equinix:metal/getVrf:getVrf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVrf.
type LookupVrfArgs struct {
	// ID of the VRF resource
	VrfId string `pulumi:"vrfId"`
}

// A collection of values returned by getVrf.
type LookupVrfResult struct {
	// Description of the VRF.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.
	IpRanges []string `pulumi:"ipRanges"`
	// The 4-byte ASN set on the VRF.
	LocalAsn int `pulumi:"localAsn"`
	// Metro ID or Code where the VRF will be deployed.
	Metro string `pulumi:"metro"`
	// User-supplied name of the VRF, unique to the project
	Name string `pulumi:"name"`
	// Project ID where the VRF will be deployed.
	ProjectId string `pulumi:"projectId"`
	VrfId     string `pulumi:"vrfId"`
}

func LookupVrfOutput(ctx *pulumi.Context, args LookupVrfOutputArgs, opts ...pulumi.InvokeOption) LookupVrfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVrfResult, error) {
			args := v.(LookupVrfArgs)
			r, err := LookupVrf(ctx, &args, opts...)
			var s LookupVrfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVrfResultOutput)
}

// A collection of arguments for invoking getVrf.
type LookupVrfOutputArgs struct {
	// ID of the VRF resource
	VrfId pulumi.StringInput `pulumi:"vrfId"`
}

func (LookupVrfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrfArgs)(nil)).Elem()
}

// A collection of values returned by getVrf.
type LookupVrfResultOutput struct{ *pulumi.OutputState }

func (LookupVrfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrfResult)(nil)).Elem()
}

func (o LookupVrfResultOutput) ToLookupVrfResultOutput() LookupVrfResultOutput {
	return o
}

func (o LookupVrfResultOutput) ToLookupVrfResultOutputWithContext(ctx context.Context) LookupVrfResultOutput {
	return o
}

func (o LookupVrfResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVrfResult] {
	return pulumix.Output[LookupVrfResult]{
		OutputState: o.OutputState,
	}
}

// Description of the VRF.
func (o LookupVrfResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVrfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.Id }).(pulumi.StringOutput)
}

// All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.
func (o LookupVrfResultOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVrfResult) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

// The 4-byte ASN set on the VRF.
func (o LookupVrfResultOutput) LocalAsn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVrfResult) int { return v.LocalAsn }).(pulumi.IntOutput)
}

// Metro ID or Code where the VRF will be deployed.
func (o LookupVrfResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.Metro }).(pulumi.StringOutput)
}

// User-supplied name of the VRF, unique to the project
func (o LookupVrfResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project ID where the VRF will be deployed.
func (o LookupVrfResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupVrfResultOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfResult) string { return v.VrfId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVrfResultOutput{})
}
