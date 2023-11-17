// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package metal

import (
	"context"
	"reflect"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this datasource to retrieve attributes of a Project SSH Key API resource.
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
//			_, err := metal.LookupProjectSshKey(ctx, &metal.LookupProjectSshKeyArgs{
//				Search:    pulumi.StringRef("username@hostname"),
//				ProjectId: local.Project_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProjectSshKey(ctx *pulumi.Context, args *LookupProjectSshKeyArgs, opts ...pulumi.InvokeOption) (*LookupProjectSshKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectSshKeyResult
	err := ctx.Invoke("equinix:metal/getProjectSshKey:getProjectSshKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectSshKey.
type LookupProjectSshKeyArgs struct {
	// The id of the SSH Key to search for in the Equinix Metal project.
	Id *string `pulumi:"id"`
	// The Equinix Metal project id of the Equinix Metal SSH Key.
	//
	// > **NOTE:** One of either `search` or `id` must be provided along with `projectId`.
	ProjectId string `pulumi:"projectId"`
	// The name, fingerprint, or publicKey of the SSH Key to search for
	// in the Equinix Metal project.
	Search *string `pulumi:"search"`
}

// A collection of values returned by getProjectSshKey.
type LookupProjectSshKeyResult struct {
	// The timestamp for when the SSH key was created.
	Created string `pulumi:"created"`
	// The fingerprint of the SSH key.
	Fingerprint string `pulumi:"fingerprint"`
	// The unique ID of the key.
	Id string `pulumi:"id"`
	// The name of the SSH key.
	Name string `pulumi:"name"`
	// The ID of parent project (same as project_id).
	OwnerId string `pulumi:"ownerId"`
	// The ID of parent project.
	ProjectId string `pulumi:"projectId"`
	// The text of the public key.
	PublicKey string  `pulumi:"publicKey"`
	Search    *string `pulumi:"search"`
	// The timestamp for the last time the SSH key was updated.
	Updated string `pulumi:"updated"`
}

func LookupProjectSshKeyOutput(ctx *pulumi.Context, args LookupProjectSshKeyOutputArgs, opts ...pulumi.InvokeOption) LookupProjectSshKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectSshKeyResult, error) {
			args := v.(LookupProjectSshKeyArgs)
			r, err := LookupProjectSshKey(ctx, &args, opts...)
			var s LookupProjectSshKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectSshKeyResultOutput)
}

// A collection of arguments for invoking getProjectSshKey.
type LookupProjectSshKeyOutputArgs struct {
	// The id of the SSH Key to search for in the Equinix Metal project.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The Equinix Metal project id of the Equinix Metal SSH Key.
	//
	// > **NOTE:** One of either `search` or `id` must be provided along with `projectId`.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The name, fingerprint, or publicKey of the SSH Key to search for
	// in the Equinix Metal project.
	Search pulumi.StringPtrInput `pulumi:"search"`
}

func (LookupProjectSshKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectSshKeyArgs)(nil)).Elem()
}

// A collection of values returned by getProjectSshKey.
type LookupProjectSshKeyResultOutput struct{ *pulumi.OutputState }

func (LookupProjectSshKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectSshKeyResult)(nil)).Elem()
}

func (o LookupProjectSshKeyResultOutput) ToLookupProjectSshKeyResultOutput() LookupProjectSshKeyResultOutput {
	return o
}

func (o LookupProjectSshKeyResultOutput) ToLookupProjectSshKeyResultOutputWithContext(ctx context.Context) LookupProjectSshKeyResultOutput {
	return o
}

// The timestamp for when the SSH key was created.
func (o LookupProjectSshKeyResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.Created }).(pulumi.StringOutput)
}

// The fingerprint of the SSH key.
func (o LookupProjectSshKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The unique ID of the key.
func (o LookupProjectSshKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the SSH key.
func (o LookupProjectSshKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of parent project (same as project_id).
func (o LookupProjectSshKeyResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// The ID of parent project.
func (o LookupProjectSshKeyResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The text of the public key.
func (o LookupProjectSshKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupProjectSshKeyResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// The timestamp for the last time the SSH key was updated.
func (o LookupProjectSshKeyResultOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectSshKeyResult) string { return v.Updated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectSshKeyResultOutput{})
}
