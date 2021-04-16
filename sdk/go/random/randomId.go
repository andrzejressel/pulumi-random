// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomId` generates random numbers that are intended to be
// used as unique identifiers for other resources.
//
// This resource *does* use a cryptographic random number generator in order
// to minimize the chance of collisions, making the results of this resource
// when a 16-byte identifier is requested of equivalent uniqueness to a
// type-4 UUID.
//
// This resource can be used in conjunction with resources that have
// the `createBeforeDestroy` lifecycle flag set to avoid conflicts with
// unique names during the brief period where both the old and new resources
// exist concurrently.
//
// ## Import
//
// Random Ids can be imported using the `b64_url` with an optional `prefix`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example with no prefix
//
// ```sh
//  $ pulumi import random:index/randomId:RandomId server p-9hUg
// ```
//
//  Example with prefix (prefix is separated by a `,`)
//
// ```sh
//  $ pulumi import random:index/randomId:RandomId server my-prefix-,p-9hUg
// ```
type RandomId struct {
	pulumi.CustomResourceState

	// The generated id presented in base64 without additional transformations.
	B64Std pulumi.StringOutput `pulumi:"b64Std"`
	// The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
	B64Url pulumi.StringOutput `pulumi:"b64Url"`
	// The number of random bytes to produce. The
	// minimum value is 1, which produces eight bits of randomness.
	ByteLength pulumi.IntOutput `pulumi:"byteLength"`
	// The generated id presented in non-padded decimal digits.
	Dec pulumi.StringOutput `pulumi:"dec"`
	// The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
	Hex pulumi.StringOutput `pulumi:"hex"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// Arbitrary string to prefix the output value with. This
	// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
	// base64 encoded.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
}

// NewRandomId registers a new resource with the given unique name, arguments, and options.
func NewRandomId(ctx *pulumi.Context,
	name string, args *RandomIdArgs, opts ...pulumi.ResourceOption) (*RandomId, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ByteLength == nil {
		return nil, errors.New("invalid value for required argument 'ByteLength'")
	}
	var resource RandomId
	err := ctx.RegisterResource("random:index/randomId:RandomId", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomId gets an existing RandomId resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomId(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomIdState, opts ...pulumi.ResourceOption) (*RandomId, error) {
	var resource RandomId
	err := ctx.ReadResource("random:index/randomId:RandomId", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomId resources.
type randomIdState struct {
	// The generated id presented in base64 without additional transformations.
	B64Std *string `pulumi:"b64Std"`
	// The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
	B64Url *string `pulumi:"b64Url"`
	// The number of random bytes to produce. The
	// minimum value is 1, which produces eight bits of randomness.
	ByteLength *int `pulumi:"byteLength"`
	// The generated id presented in non-padded decimal digits.
	Dec *string `pulumi:"dec"`
	// The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
	Hex *string `pulumi:"hex"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// Arbitrary string to prefix the output value with. This
	// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
	// base64 encoded.
	Prefix *string `pulumi:"prefix"`
}

type RandomIdState struct {
	// The generated id presented in base64 without additional transformations.
	B64Std pulumi.StringPtrInput
	// The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
	B64Url pulumi.StringPtrInput
	// The number of random bytes to produce. The
	// minimum value is 1, which produces eight bits of randomness.
	ByteLength pulumi.IntPtrInput
	// The generated id presented in non-padded decimal digits.
	Dec pulumi.StringPtrInput
	// The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
	Hex pulumi.StringPtrInput
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapInput
	// Arbitrary string to prefix the output value with. This
	// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
	// base64 encoded.
	Prefix pulumi.StringPtrInput
}

func (RandomIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIdState)(nil)).Elem()
}

type randomIdArgs struct {
	// The number of random bytes to produce. The
	// minimum value is 1, which produces eight bits of randomness.
	ByteLength int `pulumi:"byteLength"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// Arbitrary string to prefix the output value with. This
	// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
	// base64 encoded.
	Prefix *string `pulumi:"prefix"`
}

// The set of arguments for constructing a RandomId resource.
type RandomIdArgs struct {
	// The number of random bytes to produce. The
	// minimum value is 1, which produces eight bits of randomness.
	ByteLength pulumi.IntInput
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapInput
	// Arbitrary string to prefix the output value with. This
	// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
	// base64 encoded.
	Prefix pulumi.StringPtrInput
}

func (RandomIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIdArgs)(nil)).Elem()
}

type RandomIdInput interface {
	pulumi.Input

	ToRandomIdOutput() RandomIdOutput
	ToRandomIdOutputWithContext(ctx context.Context) RandomIdOutput
}

func (*RandomId) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomId)(nil))
}

func (i *RandomId) ToRandomIdOutput() RandomIdOutput {
	return i.ToRandomIdOutputWithContext(context.Background())
}

func (i *RandomId) ToRandomIdOutputWithContext(ctx context.Context) RandomIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIdOutput)
}

func (i *RandomId) ToRandomIdPtrOutput() RandomIdPtrOutput {
	return i.ToRandomIdPtrOutputWithContext(context.Background())
}

func (i *RandomId) ToRandomIdPtrOutputWithContext(ctx context.Context) RandomIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIdPtrOutput)
}

type RandomIdPtrInput interface {
	pulumi.Input

	ToRandomIdPtrOutput() RandomIdPtrOutput
	ToRandomIdPtrOutputWithContext(ctx context.Context) RandomIdPtrOutput
}

type randomIdPtrType RandomIdArgs

func (*randomIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomId)(nil))
}

func (i *randomIdPtrType) ToRandomIdPtrOutput() RandomIdPtrOutput {
	return i.ToRandomIdPtrOutputWithContext(context.Background())
}

func (i *randomIdPtrType) ToRandomIdPtrOutputWithContext(ctx context.Context) RandomIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIdPtrOutput)
}

// RandomIdArrayInput is an input type that accepts RandomIdArray and RandomIdArrayOutput values.
// You can construct a concrete instance of `RandomIdArrayInput` via:
//
//          RandomIdArray{ RandomIdArgs{...} }
type RandomIdArrayInput interface {
	pulumi.Input

	ToRandomIdArrayOutput() RandomIdArrayOutput
	ToRandomIdArrayOutputWithContext(context.Context) RandomIdArrayOutput
}

type RandomIdArray []RandomIdInput

func (RandomIdArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RandomId)(nil))
}

func (i RandomIdArray) ToRandomIdArrayOutput() RandomIdArrayOutput {
	return i.ToRandomIdArrayOutputWithContext(context.Background())
}

func (i RandomIdArray) ToRandomIdArrayOutputWithContext(ctx context.Context) RandomIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIdArrayOutput)
}

// RandomIdMapInput is an input type that accepts RandomIdMap and RandomIdMapOutput values.
// You can construct a concrete instance of `RandomIdMapInput` via:
//
//          RandomIdMap{ "key": RandomIdArgs{...} }
type RandomIdMapInput interface {
	pulumi.Input

	ToRandomIdMapOutput() RandomIdMapOutput
	ToRandomIdMapOutputWithContext(context.Context) RandomIdMapOutput
}

type RandomIdMap map[string]RandomIdInput

func (RandomIdMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RandomId)(nil))
}

func (i RandomIdMap) ToRandomIdMapOutput() RandomIdMapOutput {
	return i.ToRandomIdMapOutputWithContext(context.Background())
}

func (i RandomIdMap) ToRandomIdMapOutputWithContext(ctx context.Context) RandomIdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIdMapOutput)
}

type RandomIdOutput struct {
	*pulumi.OutputState
}

func (RandomIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomId)(nil))
}

func (o RandomIdOutput) ToRandomIdOutput() RandomIdOutput {
	return o
}

func (o RandomIdOutput) ToRandomIdOutputWithContext(ctx context.Context) RandomIdOutput {
	return o
}

func (o RandomIdOutput) ToRandomIdPtrOutput() RandomIdPtrOutput {
	return o.ToRandomIdPtrOutputWithContext(context.Background())
}

func (o RandomIdOutput) ToRandomIdPtrOutputWithContext(ctx context.Context) RandomIdPtrOutput {
	return o.ApplyT(func(v RandomId) *RandomId {
		return &v
	}).(RandomIdPtrOutput)
}

type RandomIdPtrOutput struct {
	*pulumi.OutputState
}

func (RandomIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomId)(nil))
}

func (o RandomIdPtrOutput) ToRandomIdPtrOutput() RandomIdPtrOutput {
	return o
}

func (o RandomIdPtrOutput) ToRandomIdPtrOutputWithContext(ctx context.Context) RandomIdPtrOutput {
	return o
}

type RandomIdArrayOutput struct{ *pulumi.OutputState }

func (RandomIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RandomId)(nil))
}

func (o RandomIdArrayOutput) ToRandomIdArrayOutput() RandomIdArrayOutput {
	return o
}

func (o RandomIdArrayOutput) ToRandomIdArrayOutputWithContext(ctx context.Context) RandomIdArrayOutput {
	return o
}

func (o RandomIdArrayOutput) Index(i pulumi.IntInput) RandomIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RandomId {
		return vs[0].([]RandomId)[vs[1].(int)]
	}).(RandomIdOutput)
}

type RandomIdMapOutput struct{ *pulumi.OutputState }

func (RandomIdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RandomId)(nil))
}

func (o RandomIdMapOutput) ToRandomIdMapOutput() RandomIdMapOutput {
	return o
}

func (o RandomIdMapOutput) ToRandomIdMapOutputWithContext(ctx context.Context) RandomIdMapOutput {
	return o
}

func (o RandomIdMapOutput) MapIndex(k pulumi.StringInput) RandomIdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RandomId {
		return vs[0].(map[string]RandomId)[vs[1].(string)]
	}).(RandomIdOutput)
}

func init() {
	pulumi.RegisterOutputType(RandomIdOutput{})
	pulumi.RegisterOutputType(RandomIdPtrOutput{})
	pulumi.RegisterOutputType(RandomIdArrayOutput{})
	pulumi.RegisterOutputType(RandomIdMapOutput{})
}
