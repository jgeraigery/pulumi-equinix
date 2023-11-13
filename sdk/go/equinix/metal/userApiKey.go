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

// Use this resource to create Metal User API Key resources in Equinix Metal. Each API key contains a
// token which can be used for authentication in Equinix Metal HTTP API (in HTTP request header
// `X-Auth-Token`).
//
// Read-only keys only allow to list and view existing resources, read-write keys can also be used to
// create resources.
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
//			description := "An user level API Key"
//			if param := cfg.Get("description"); param != "" {
//				description = param
//			}
//			readOnly := false
//			if param := cfg.GetBool("readOnly"); param {
//				readOnly = param
//			}
//			apiKey, err := metal.NewUserApiKey(ctx, "apiKey", &metal.UserApiKeyArgs{
//				Description: pulumi.String(description),
//				ReadOnly:    pulumi.Bool(readOnly),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("apiKeyToken", apiKey.Token)
//			return nil
//		})
//	}
//
// ```
type UserApiKey struct {
	pulumi.CustomResourceState

	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringOutput `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients.
	Token pulumi.StringOutput `pulumi:"token"`
	// UUID of the owner of the API key.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserApiKey registers a new resource with the given unique name, arguments, and options.
func NewUserApiKey(ctx *pulumi.Context,
	name string, args *UserApiKeyArgs, opts ...pulumi.ResourceOption) (*UserApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ReadOnly == nil {
		return nil, errors.New("invalid value for required argument 'ReadOnly'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserApiKey
	err := ctx.RegisterResource("equinix:metal/userApiKey:UserApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserApiKey gets an existing UserApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserApiKeyState, opts ...pulumi.ResourceOption) (*UserApiKey, error) {
	var resource UserApiKey
	err := ctx.ReadResource("equinix:metal/userApiKey:UserApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserApiKey resources.
type userApiKeyState struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description *string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients.
	Token *string `pulumi:"token"`
	// UUID of the owner of the API key.
	UserId *string `pulumi:"userId"`
}

type UserApiKeyState struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringPtrInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolPtrInput
	// API token which can be used in Equinix Metal API clients.
	Token pulumi.StringPtrInput
	// UUID of the owner of the API key.
	UserId pulumi.StringPtrInput
}

func (UserApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userApiKeyState)(nil)).Elem()
}

type userApiKeyArgs struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly bool `pulumi:"readOnly"`
}

// The set of arguments for constructing a UserApiKey resource.
type UserApiKeyArgs struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolInput
}

func (UserApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userApiKeyArgs)(nil)).Elem()
}

type UserApiKeyInput interface {
	pulumi.Input

	ToUserApiKeyOutput() UserApiKeyOutput
	ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput
}

func (*UserApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UserApiKey)(nil)).Elem()
}

func (i *UserApiKey) ToUserApiKeyOutput() UserApiKeyOutput {
	return i.ToUserApiKeyOutputWithContext(context.Background())
}

func (i *UserApiKey) ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyOutput)
}

func (i *UserApiKey) ToOutput(ctx context.Context) pulumix.Output[*UserApiKey] {
	return pulumix.Output[*UserApiKey]{
		OutputState: i.ToUserApiKeyOutputWithContext(ctx).OutputState,
	}
}

// UserApiKeyArrayInput is an input type that accepts UserApiKeyArray and UserApiKeyArrayOutput values.
// You can construct a concrete instance of `UserApiKeyArrayInput` via:
//
//	UserApiKeyArray{ UserApiKeyArgs{...} }
type UserApiKeyArrayInput interface {
	pulumi.Input

	ToUserApiKeyArrayOutput() UserApiKeyArrayOutput
	ToUserApiKeyArrayOutputWithContext(context.Context) UserApiKeyArrayOutput
}

type UserApiKeyArray []UserApiKeyInput

func (UserApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserApiKey)(nil)).Elem()
}

func (i UserApiKeyArray) ToUserApiKeyArrayOutput() UserApiKeyArrayOutput {
	return i.ToUserApiKeyArrayOutputWithContext(context.Background())
}

func (i UserApiKeyArray) ToUserApiKeyArrayOutputWithContext(ctx context.Context) UserApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyArrayOutput)
}

func (i UserApiKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserApiKey] {
	return pulumix.Output[[]*UserApiKey]{
		OutputState: i.ToUserApiKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// UserApiKeyMapInput is an input type that accepts UserApiKeyMap and UserApiKeyMapOutput values.
// You can construct a concrete instance of `UserApiKeyMapInput` via:
//
//	UserApiKeyMap{ "key": UserApiKeyArgs{...} }
type UserApiKeyMapInput interface {
	pulumi.Input

	ToUserApiKeyMapOutput() UserApiKeyMapOutput
	ToUserApiKeyMapOutputWithContext(context.Context) UserApiKeyMapOutput
}

type UserApiKeyMap map[string]UserApiKeyInput

func (UserApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserApiKey)(nil)).Elem()
}

func (i UserApiKeyMap) ToUserApiKeyMapOutput() UserApiKeyMapOutput {
	return i.ToUserApiKeyMapOutputWithContext(context.Background())
}

func (i UserApiKeyMap) ToUserApiKeyMapOutputWithContext(ctx context.Context) UserApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyMapOutput)
}

func (i UserApiKeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserApiKey] {
	return pulumix.Output[map[string]*UserApiKey]{
		OutputState: i.ToUserApiKeyMapOutputWithContext(ctx).OutputState,
	}
}

type UserApiKeyOutput struct{ *pulumi.OutputState }

func (UserApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserApiKey)(nil)).Elem()
}

func (o UserApiKeyOutput) ToUserApiKeyOutput() UserApiKeyOutput {
	return o
}

func (o UserApiKeyOutput) ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput {
	return o
}

func (o UserApiKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*UserApiKey] {
	return pulumix.Output[*UserApiKey]{
		OutputState: o.OutputState,
	}
}

// Description string for the User API Key resource.
// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
func (o UserApiKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UserApiKey) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Flag indicating whether the API key shoud be read-only
func (o UserApiKeyOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserApiKey) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// API token which can be used in Equinix Metal API clients.
func (o UserApiKeyOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *UserApiKey) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// UUID of the owner of the API key.
func (o UserApiKeyOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserApiKey) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserApiKeyArrayOutput struct{ *pulumi.OutputState }

func (UserApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserApiKey)(nil)).Elem()
}

func (o UserApiKeyArrayOutput) ToUserApiKeyArrayOutput() UserApiKeyArrayOutput {
	return o
}

func (o UserApiKeyArrayOutput) ToUserApiKeyArrayOutputWithContext(ctx context.Context) UserApiKeyArrayOutput {
	return o
}

func (o UserApiKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserApiKey] {
	return pulumix.Output[[]*UserApiKey]{
		OutputState: o.OutputState,
	}
}

func (o UserApiKeyArrayOutput) Index(i pulumi.IntInput) UserApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserApiKey {
		return vs[0].([]*UserApiKey)[vs[1].(int)]
	}).(UserApiKeyOutput)
}

type UserApiKeyMapOutput struct{ *pulumi.OutputState }

func (UserApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserApiKey)(nil)).Elem()
}

func (o UserApiKeyMapOutput) ToUserApiKeyMapOutput() UserApiKeyMapOutput {
	return o
}

func (o UserApiKeyMapOutput) ToUserApiKeyMapOutputWithContext(ctx context.Context) UserApiKeyMapOutput {
	return o
}

func (o UserApiKeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserApiKey] {
	return pulumix.Output[map[string]*UserApiKey]{
		OutputState: o.OutputState,
	}
}

func (o UserApiKeyMapOutput) MapIndex(k pulumi.StringInput) UserApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserApiKey {
		return vs[0].(map[string]*UserApiKey)[vs[1].(string)]
	}).(UserApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserApiKeyInput)(nil)).Elem(), &UserApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserApiKeyArrayInput)(nil)).Elem(), UserApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserApiKeyMapInput)(nil)).Elem(), UserApiKeyMap{})
	pulumi.RegisterOutputType(UserApiKeyOutput{})
	pulumi.RegisterOutputType(UserApiKeyArrayOutput{})
	pulumi.RegisterOutputType(UserApiKeyMapOutput{})
}
