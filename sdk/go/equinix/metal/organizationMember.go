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

// Manage the membership of existing and new invitees within an Equinix Metal organization and its projects.
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
//			organizationId := cfg.Require("organizationId")
//			projectId := cfg.Require("projectId")
//			userEmailAddress := cfg.Require("userEmailAddress")
//			member, err := metal.NewOrganizationMember(ctx, "member", &metal.OrganizationMemberArgs{
//				Invitee: pulumi.String(userEmailAddress),
//				Roles: pulumi.StringArray{
//					pulumi.String("limited_collaborator"),
//				},
//				ProjectsIds: pulumi.StringArray{
//					pulumi.String(projectId),
//				},
//				OrganizationId: pulumi.String(organizationId),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("memberId", member.ID())
//			ctx.Export("memberState", member.State)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource can be imported using the `invitee` and `organization_id` as colon separated arguments: <break><break>```sh<break> $ pulumi import equinix:metal/organizationMember:OrganizationMember resource_name {invitee}:{organization_id} <break>```<break><break>
type OrganizationMember struct {
	pulumi.CustomResourceState

	// When the invitation was created (only known in the invitation stage)
	Created pulumi.StringOutput `pulumi:"created"`
	// The userId of the user that sent the invitation (only known in the invitation stage)
	InvitedBy pulumi.StringOutput `pulumi:"invitedBy"`
	// The email address of the user to invite
	Invitee pulumi.StringOutput `pulumi:"invitee"`
	// A message to include in the emailed invitation.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// The nonce for the invitation (only known in the invitation stage)
	Nonce pulumi.StringOutput `pulumi:"nonce"`
	// The organization to invite the user to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	ProjectsIds pulumi.StringArrayOutput `pulumi:"projectsIds"`
	// Organization roles (admin, collaborator, limited_collaborator, billing)
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
	State pulumi.StringOutput `pulumi:"state"`
	// When the invitation was updated (only known in the invitation stage)
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewOrganizationMember registers a new resource with the given unique name, arguments, and options.
func NewOrganizationMember(ctx *pulumi.Context,
	name string, args *OrganizationMemberArgs, opts ...pulumi.ResourceOption) (*OrganizationMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Invitee == nil {
		return nil, errors.New("invalid value for required argument 'Invitee'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.ProjectsIds == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsIds'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationMember
	err := ctx.RegisterResource("equinix:metal/organizationMember:OrganizationMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationMember gets an existing OrganizationMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationMemberState, opts ...pulumi.ResourceOption) (*OrganizationMember, error) {
	var resource OrganizationMember
	err := ctx.ReadResource("equinix:metal/organizationMember:OrganizationMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationMember resources.
type organizationMemberState struct {
	// When the invitation was created (only known in the invitation stage)
	Created *string `pulumi:"created"`
	// The userId of the user that sent the invitation (only known in the invitation stage)
	InvitedBy *string `pulumi:"invitedBy"`
	// The email address of the user to invite
	Invitee *string `pulumi:"invitee"`
	// A message to include in the emailed invitation.
	Message *string `pulumi:"message"`
	// The nonce for the invitation (only known in the invitation stage)
	Nonce *string `pulumi:"nonce"`
	// The organization to invite the user to
	OrganizationId *string `pulumi:"organizationId"`
	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	ProjectsIds []string `pulumi:"projectsIds"`
	// Organization roles (admin, collaborator, limited_collaborator, billing)
	Roles []string `pulumi:"roles"`
	// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
	State *string `pulumi:"state"`
	// When the invitation was updated (only known in the invitation stage)
	Updated *string `pulumi:"updated"`
}

type OrganizationMemberState struct {
	// When the invitation was created (only known in the invitation stage)
	Created pulumi.StringPtrInput
	// The userId of the user that sent the invitation (only known in the invitation stage)
	InvitedBy pulumi.StringPtrInput
	// The email address of the user to invite
	Invitee pulumi.StringPtrInput
	// A message to include in the emailed invitation.
	Message pulumi.StringPtrInput
	// The nonce for the invitation (only known in the invitation stage)
	Nonce pulumi.StringPtrInput
	// The organization to invite the user to
	OrganizationId pulumi.StringPtrInput
	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	ProjectsIds pulumi.StringArrayInput
	// Organization roles (admin, collaborator, limited_collaborator, billing)
	Roles pulumi.StringArrayInput
	// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
	State pulumi.StringPtrInput
	// When the invitation was updated (only known in the invitation stage)
	Updated pulumi.StringPtrInput
}

func (OrganizationMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationMemberState)(nil)).Elem()
}

type organizationMemberArgs struct {
	// The email address of the user to invite
	Invitee string `pulumi:"invitee"`
	// A message to include in the emailed invitation.
	Message *string `pulumi:"message"`
	// The organization to invite the user to
	OrganizationId string `pulumi:"organizationId"`
	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	ProjectsIds []string `pulumi:"projectsIds"`
	// Organization roles (admin, collaborator, limited_collaborator, billing)
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a OrganizationMember resource.
type OrganizationMemberArgs struct {
	// The email address of the user to invite
	Invitee pulumi.StringInput
	// A message to include in the emailed invitation.
	Message pulumi.StringPtrInput
	// The organization to invite the user to
	OrganizationId pulumi.StringInput
	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	ProjectsIds pulumi.StringArrayInput
	// Organization roles (admin, collaborator, limited_collaborator, billing)
	Roles pulumi.StringArrayInput
}

func (OrganizationMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationMemberArgs)(nil)).Elem()
}

type OrganizationMemberInput interface {
	pulumi.Input

	ToOrganizationMemberOutput() OrganizationMemberOutput
	ToOrganizationMemberOutputWithContext(ctx context.Context) OrganizationMemberOutput
}

func (*OrganizationMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationMember)(nil)).Elem()
}

func (i *OrganizationMember) ToOrganizationMemberOutput() OrganizationMemberOutput {
	return i.ToOrganizationMemberOutputWithContext(context.Background())
}

func (i *OrganizationMember) ToOrganizationMemberOutputWithContext(ctx context.Context) OrganizationMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMemberOutput)
}

func (i *OrganizationMember) ToOutput(ctx context.Context) pulumix.Output[*OrganizationMember] {
	return pulumix.Output[*OrganizationMember]{
		OutputState: i.ToOrganizationMemberOutputWithContext(ctx).OutputState,
	}
}

// OrganizationMemberArrayInput is an input type that accepts OrganizationMemberArray and OrganizationMemberArrayOutput values.
// You can construct a concrete instance of `OrganizationMemberArrayInput` via:
//
//	OrganizationMemberArray{ OrganizationMemberArgs{...} }
type OrganizationMemberArrayInput interface {
	pulumi.Input

	ToOrganizationMemberArrayOutput() OrganizationMemberArrayOutput
	ToOrganizationMemberArrayOutputWithContext(context.Context) OrganizationMemberArrayOutput
}

type OrganizationMemberArray []OrganizationMemberInput

func (OrganizationMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationMember)(nil)).Elem()
}

func (i OrganizationMemberArray) ToOrganizationMemberArrayOutput() OrganizationMemberArrayOutput {
	return i.ToOrganizationMemberArrayOutputWithContext(context.Background())
}

func (i OrganizationMemberArray) ToOrganizationMemberArrayOutputWithContext(ctx context.Context) OrganizationMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMemberArrayOutput)
}

func (i OrganizationMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrganizationMember] {
	return pulumix.Output[[]*OrganizationMember]{
		OutputState: i.ToOrganizationMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// OrganizationMemberMapInput is an input type that accepts OrganizationMemberMap and OrganizationMemberMapOutput values.
// You can construct a concrete instance of `OrganizationMemberMapInput` via:
//
//	OrganizationMemberMap{ "key": OrganizationMemberArgs{...} }
type OrganizationMemberMapInput interface {
	pulumi.Input

	ToOrganizationMemberMapOutput() OrganizationMemberMapOutput
	ToOrganizationMemberMapOutputWithContext(context.Context) OrganizationMemberMapOutput
}

type OrganizationMemberMap map[string]OrganizationMemberInput

func (OrganizationMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationMember)(nil)).Elem()
}

func (i OrganizationMemberMap) ToOrganizationMemberMapOutput() OrganizationMemberMapOutput {
	return i.ToOrganizationMemberMapOutputWithContext(context.Background())
}

func (i OrganizationMemberMap) ToOrganizationMemberMapOutputWithContext(ctx context.Context) OrganizationMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMemberMapOutput)
}

func (i OrganizationMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrganizationMember] {
	return pulumix.Output[map[string]*OrganizationMember]{
		OutputState: i.ToOrganizationMemberMapOutputWithContext(ctx).OutputState,
	}
}

type OrganizationMemberOutput struct{ *pulumi.OutputState }

func (OrganizationMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationMember)(nil)).Elem()
}

func (o OrganizationMemberOutput) ToOrganizationMemberOutput() OrganizationMemberOutput {
	return o
}

func (o OrganizationMemberOutput) ToOrganizationMemberOutputWithContext(ctx context.Context) OrganizationMemberOutput {
	return o
}

func (o OrganizationMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*OrganizationMember] {
	return pulumix.Output[*OrganizationMember]{
		OutputState: o.OutputState,
	}
}

// When the invitation was created (only known in the invitation stage)
func (o OrganizationMemberOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The userId of the user that sent the invitation (only known in the invitation stage)
func (o OrganizationMemberOutput) InvitedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.InvitedBy }).(pulumi.StringOutput)
}

// The email address of the user to invite
func (o OrganizationMemberOutput) Invitee() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.Invitee }).(pulumi.StringOutput)
}

// A message to include in the emailed invitation.
func (o OrganizationMemberOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// The nonce for the invitation (only known in the invitation stage)
func (o OrganizationMemberOutput) Nonce() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.Nonce }).(pulumi.StringOutput)
}

// The organization to invite the user to
func (o OrganizationMemberOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
func (o OrganizationMemberOutput) ProjectsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringArrayOutput { return v.ProjectsIds }).(pulumi.StringArrayOutput)
}

// Organization roles (admin, collaborator, limited_collaborator, billing)
func (o OrganizationMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
func (o OrganizationMemberOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// When the invitation was updated (only known in the invitation stage)
func (o OrganizationMemberOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationMember) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

type OrganizationMemberArrayOutput struct{ *pulumi.OutputState }

func (OrganizationMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationMember)(nil)).Elem()
}

func (o OrganizationMemberArrayOutput) ToOrganizationMemberArrayOutput() OrganizationMemberArrayOutput {
	return o
}

func (o OrganizationMemberArrayOutput) ToOrganizationMemberArrayOutputWithContext(ctx context.Context) OrganizationMemberArrayOutput {
	return o
}

func (o OrganizationMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrganizationMember] {
	return pulumix.Output[[]*OrganizationMember]{
		OutputState: o.OutputState,
	}
}

func (o OrganizationMemberArrayOutput) Index(i pulumi.IntInput) OrganizationMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationMember {
		return vs[0].([]*OrganizationMember)[vs[1].(int)]
	}).(OrganizationMemberOutput)
}

type OrganizationMemberMapOutput struct{ *pulumi.OutputState }

func (OrganizationMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationMember)(nil)).Elem()
}

func (o OrganizationMemberMapOutput) ToOrganizationMemberMapOutput() OrganizationMemberMapOutput {
	return o
}

func (o OrganizationMemberMapOutput) ToOrganizationMemberMapOutputWithContext(ctx context.Context) OrganizationMemberMapOutput {
	return o
}

func (o OrganizationMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrganizationMember] {
	return pulumix.Output[map[string]*OrganizationMember]{
		OutputState: o.OutputState,
	}
}

func (o OrganizationMemberMapOutput) MapIndex(k pulumi.StringInput) OrganizationMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationMember {
		return vs[0].(map[string]*OrganizationMember)[vs[1].(string)]
	}).(OrganizationMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMemberInput)(nil)).Elem(), &OrganizationMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMemberArrayInput)(nil)).Elem(), OrganizationMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMemberMapInput)(nil)).Elem(), OrganizationMemberMap{})
	pulumi.RegisterOutputType(OrganizationMemberOutput{})
	pulumi.RegisterOutputType(OrganizationMemberArrayOutput{})
	pulumi.RegisterOutputType(OrganizationMemberMapOutput{})
}
