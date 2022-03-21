// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		password, err := random.NewRandomPassword(ctx, "password", &random.RandomPasswordArgs{
// 			Length:          pulumi.Int(16),
// 			Special:         pulumi.Bool(true),
// 			OverrideSpecial: pulumi.String(fmt.Sprintf("%v%v%v%v", "!#", "$", "%", "&*()-_=+[]{}<>:?")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewInstance(ctx, "example", &rds.InstanceArgs{
// 			InstanceClass:    pulumi.String("db.t3.micro"),
// 			AllocatedStorage: pulumi.Int(64),
// 			Engine:           pulumi.String("mysql"),
// 			Username:         pulumi.String("someone"),
// 			Password:         password.Result,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # Random Password can be imported by specifying the value of the string
//
// ```sh
//  $ pulumi import random:index/randomPassword:RandomPassword password securepassword
// ```
type RandomPassword struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The length of the string desired.
	Length pulumi.IntOutput `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolPtrOutput `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntPtrOutput `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntPtrOutput `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntPtrOutput `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntPtrOutput `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`.
	Number pulumi.BoolPtrOutput `pulumi:"number"`
	// Supply your own list of special characters to use for string generation. This overrides the default character list in
	// the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
	// generation.
	OverrideSpecial pulumi.StringPtrOutput `pulumi:"overrideSpecial"`
	// The generated random string.
	Result pulumi.StringOutput `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrOutput `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrOutput `pulumi:"upper"`
}

// NewRandomPassword registers a new resource with the given unique name, arguments, and options.
func NewRandomPassword(ctx *pulumi.Context,
	name string, args *RandomPasswordArgs, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	var resource RandomPassword
	err := ctx.RegisterResource("random:index/randomPassword:RandomPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomPassword gets an existing RandomPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomPasswordState, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	var resource RandomPassword
	err := ctx.ReadResource("random:index/randomPassword:RandomPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomPassword resources.
type randomPasswordState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length of the string desired.
	Length *int `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`.
	Number *bool `pulumi:"number"`
	// Supply your own list of special characters to use for string generation. This overrides the default character list in
	// the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
	// generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// The generated random string.
	Result *string `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper *bool `pulumi:"upper"`
}

type RandomPasswordState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapInput
	// The length of the string desired.
	Length pulumi.IntPtrInput
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result. Default value is `true`.
	Number pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation. This overrides the default character list in
	// the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
	// generation.
	OverrideSpecial pulumi.StringPtrInput
	// The generated random string.
	Result pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrInput
}

func (RandomPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordState)(nil)).Elem()
}

type randomPasswordArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length of the string desired.
	Length int `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`.
	Number *bool `pulumi:"number"`
	// Supply your own list of special characters to use for string generation. This overrides the default character list in
	// the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
	// generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper *bool `pulumi:"upper"`
}

// The set of arguments for constructing a RandomPassword resource.
type RandomPasswordArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapInput
	// The length of the string desired.
	Length pulumi.IntInput
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result. Default value is `true`.
	Number pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation. This overrides the default character list in
	// the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
	// generation.
	OverrideSpecial pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrInput
}

func (RandomPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordArgs)(nil)).Elem()
}

type RandomPasswordInput interface {
	pulumi.Input

	ToRandomPasswordOutput() RandomPasswordOutput
	ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput
}

func (*RandomPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPassword)(nil)).Elem()
}

func (i *RandomPassword) ToRandomPasswordOutput() RandomPasswordOutput {
	return i.ToRandomPasswordOutputWithContext(context.Background())
}

func (i *RandomPassword) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordOutput)
}

// RandomPasswordArrayInput is an input type that accepts RandomPasswordArray and RandomPasswordArrayOutput values.
// You can construct a concrete instance of `RandomPasswordArrayInput` via:
//
//          RandomPasswordArray{ RandomPasswordArgs{...} }
type RandomPasswordArrayInput interface {
	pulumi.Input

	ToRandomPasswordArrayOutput() RandomPasswordArrayOutput
	ToRandomPasswordArrayOutputWithContext(context.Context) RandomPasswordArrayOutput
}

type RandomPasswordArray []RandomPasswordInput

func (RandomPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPassword)(nil)).Elem()
}

func (i RandomPasswordArray) ToRandomPasswordArrayOutput() RandomPasswordArrayOutput {
	return i.ToRandomPasswordArrayOutputWithContext(context.Background())
}

func (i RandomPasswordArray) ToRandomPasswordArrayOutputWithContext(ctx context.Context) RandomPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordArrayOutput)
}

// RandomPasswordMapInput is an input type that accepts RandomPasswordMap and RandomPasswordMapOutput values.
// You can construct a concrete instance of `RandomPasswordMapInput` via:
//
//          RandomPasswordMap{ "key": RandomPasswordArgs{...} }
type RandomPasswordMapInput interface {
	pulumi.Input

	ToRandomPasswordMapOutput() RandomPasswordMapOutput
	ToRandomPasswordMapOutputWithContext(context.Context) RandomPasswordMapOutput
}

type RandomPasswordMap map[string]RandomPasswordInput

func (RandomPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPassword)(nil)).Elem()
}

func (i RandomPasswordMap) ToRandomPasswordMapOutput() RandomPasswordMapOutput {
	return i.ToRandomPasswordMapOutputWithContext(context.Background())
}

func (i RandomPasswordMap) ToRandomPasswordMapOutputWithContext(ctx context.Context) RandomPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordMapOutput)
}

type RandomPasswordOutput struct{ *pulumi.OutputState }

func (RandomPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPassword)(nil)).Elem()
}

func (o RandomPasswordOutput) ToRandomPasswordOutput() RandomPasswordOutput {
	return o
}

func (o RandomPasswordOutput) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return o
}

type RandomPasswordArrayOutput struct{ *pulumi.OutputState }

func (RandomPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPassword)(nil)).Elem()
}

func (o RandomPasswordArrayOutput) ToRandomPasswordArrayOutput() RandomPasswordArrayOutput {
	return o
}

func (o RandomPasswordArrayOutput) ToRandomPasswordArrayOutputWithContext(ctx context.Context) RandomPasswordArrayOutput {
	return o
}

func (o RandomPasswordArrayOutput) Index(i pulumi.IntInput) RandomPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomPassword {
		return vs[0].([]*RandomPassword)[vs[1].(int)]
	}).(RandomPasswordOutput)
}

type RandomPasswordMapOutput struct{ *pulumi.OutputState }

func (RandomPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPassword)(nil)).Elem()
}

func (o RandomPasswordMapOutput) ToRandomPasswordMapOutput() RandomPasswordMapOutput {
	return o
}

func (o RandomPasswordMapOutput) ToRandomPasswordMapOutputWithContext(ctx context.Context) RandomPasswordMapOutput {
	return o
}

func (o RandomPasswordMapOutput) MapIndex(k pulumi.StringInput) RandomPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomPassword {
		return vs[0].(map[string]*RandomPassword)[vs[1].(string)]
	}).(RandomPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordInput)(nil)).Elem(), &RandomPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordArrayInput)(nil)).Elem(), RandomPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordMapInput)(nil)).Elem(), RandomPasswordMap{})
	pulumi.RegisterOutputType(RandomPasswordOutput{})
	pulumi.RegisterOutputType(RandomPasswordArrayOutput{})
	pulumi.RegisterOutputType(RandomPasswordMapOutput{})
}
