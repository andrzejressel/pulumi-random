// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomShuffle` generates a random permutation of a list of strings given as an argument.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elb"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			az, err := random.NewRandomShuffle(ctx, "az", &random.RandomShuffleArgs{
//				Inputs: pulumi.StringArray{
//					pulumi.String("us-west-1a"),
//					pulumi.String("us-west-1c"),
//					pulumi.String("us-west-1d"),
//					pulumi.String("us-west-1e"),
//				},
//				ResultCount: pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elb.NewLoadBalancer(ctx, "example", &elb.LoadBalancerArgs{
//				AvailabilityZones: pulumi.StringArray{
//					az.Results,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RandomShuffle struct {
	pulumi.CustomResourceState

	// The list of strings to shuffle.
	Inputs pulumi.StringArrayOutput `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapOutput `pulumi:"keepers"`
	// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
	ResultCount pulumi.IntPtrOutput `pulumi:"resultCount"`
	// Random permutation of the list of strings given in `input`.
	Results pulumi.StringArrayOutput `pulumi:"results"`
	// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
	Seed pulumi.StringPtrOutput `pulumi:"seed"`
}

// NewRandomShuffle registers a new resource with the given unique name, arguments, and options.
func NewRandomShuffle(ctx *pulumi.Context,
	name string, args *RandomShuffleArgs, opts ...pulumi.ResourceOption) (*RandomShuffle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Inputs == nil {
		return nil, errors.New("invalid value for required argument 'Inputs'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RandomShuffle
	err := ctx.RegisterResource("random:index/randomShuffle:RandomShuffle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomShuffle gets an existing RandomShuffle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomShuffle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomShuffleState, opts ...pulumi.ResourceOption) (*RandomShuffle, error) {
	var resource RandomShuffle
	err := ctx.ReadResource("random:index/randomShuffle:RandomShuffle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomShuffle resources.
type randomShuffleState struct {
	// The list of strings to shuffle.
	Inputs []string `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
	ResultCount *int `pulumi:"resultCount"`
	// Random permutation of the list of strings given in `input`.
	Results []string `pulumi:"results"`
	// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
	Seed *string `pulumi:"seed"`
}

type RandomShuffleState struct {
	// The list of strings to shuffle.
	Inputs pulumi.StringArrayInput
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
	ResultCount pulumi.IntPtrInput
	// Random permutation of the list of strings given in `input`.
	Results pulumi.StringArrayInput
	// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
	Seed pulumi.StringPtrInput
}

func (RandomShuffleState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomShuffleState)(nil)).Elem()
}

type randomShuffleArgs struct {
	// The list of strings to shuffle.
	Inputs []string `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
	ResultCount *int `pulumi:"resultCount"`
	// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
	Seed *string `pulumi:"seed"`
}

// The set of arguments for constructing a RandomShuffle resource.
type RandomShuffleArgs struct {
	// The list of strings to shuffle.
	Inputs pulumi.StringArrayInput
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
	ResultCount pulumi.IntPtrInput
	// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
	Seed pulumi.StringPtrInput
}

func (RandomShuffleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomShuffleArgs)(nil)).Elem()
}

type RandomShuffleInput interface {
	pulumi.Input

	ToRandomShuffleOutput() RandomShuffleOutput
	ToRandomShuffleOutputWithContext(ctx context.Context) RandomShuffleOutput
}

func (*RandomShuffle) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomShuffle)(nil)).Elem()
}

func (i *RandomShuffle) ToRandomShuffleOutput() RandomShuffleOutput {
	return i.ToRandomShuffleOutputWithContext(context.Background())
}

func (i *RandomShuffle) ToRandomShuffleOutputWithContext(ctx context.Context) RandomShuffleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomShuffleOutput)
}

// RandomShuffleArrayInput is an input type that accepts RandomShuffleArray and RandomShuffleArrayOutput values.
// You can construct a concrete instance of `RandomShuffleArrayInput` via:
//
//	RandomShuffleArray{ RandomShuffleArgs{...} }
type RandomShuffleArrayInput interface {
	pulumi.Input

	ToRandomShuffleArrayOutput() RandomShuffleArrayOutput
	ToRandomShuffleArrayOutputWithContext(context.Context) RandomShuffleArrayOutput
}

type RandomShuffleArray []RandomShuffleInput

func (RandomShuffleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomShuffle)(nil)).Elem()
}

func (i RandomShuffleArray) ToRandomShuffleArrayOutput() RandomShuffleArrayOutput {
	return i.ToRandomShuffleArrayOutputWithContext(context.Background())
}

func (i RandomShuffleArray) ToRandomShuffleArrayOutputWithContext(ctx context.Context) RandomShuffleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomShuffleArrayOutput)
}

// RandomShuffleMapInput is an input type that accepts RandomShuffleMap and RandomShuffleMapOutput values.
// You can construct a concrete instance of `RandomShuffleMapInput` via:
//
//	RandomShuffleMap{ "key": RandomShuffleArgs{...} }
type RandomShuffleMapInput interface {
	pulumi.Input

	ToRandomShuffleMapOutput() RandomShuffleMapOutput
	ToRandomShuffleMapOutputWithContext(context.Context) RandomShuffleMapOutput
}

type RandomShuffleMap map[string]RandomShuffleInput

func (RandomShuffleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomShuffle)(nil)).Elem()
}

func (i RandomShuffleMap) ToRandomShuffleMapOutput() RandomShuffleMapOutput {
	return i.ToRandomShuffleMapOutputWithContext(context.Background())
}

func (i RandomShuffleMap) ToRandomShuffleMapOutputWithContext(ctx context.Context) RandomShuffleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomShuffleMapOutput)
}

type RandomShuffleOutput struct{ *pulumi.OutputState }

func (RandomShuffleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomShuffle)(nil)).Elem()
}

func (o RandomShuffleOutput) ToRandomShuffleOutput() RandomShuffleOutput {
	return o
}

func (o RandomShuffleOutput) ToRandomShuffleOutputWithContext(ctx context.Context) RandomShuffleOutput {
	return o
}

// The list of strings to shuffle.
func (o RandomShuffleOutput) Inputs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RandomShuffle) pulumi.StringArrayOutput { return v.Inputs }).(pulumi.StringArrayOutput)
}

// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
func (o RandomShuffleOutput) Keepers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RandomShuffle) pulumi.StringMapOutput { return v.Keepers }).(pulumi.StringMapOutput)
}

// The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
func (o RandomShuffleOutput) ResultCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RandomShuffle) pulumi.IntPtrOutput { return v.ResultCount }).(pulumi.IntPtrOutput)
}

// Random permutation of the list of strings given in `input`.
func (o RandomShuffleOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RandomShuffle) pulumi.StringArrayOutput { return v.Results }).(pulumi.StringArrayOutput)
}

// Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
func (o RandomShuffleOutput) Seed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RandomShuffle) pulumi.StringPtrOutput { return v.Seed }).(pulumi.StringPtrOutput)
}

type RandomShuffleArrayOutput struct{ *pulumi.OutputState }

func (RandomShuffleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomShuffle)(nil)).Elem()
}

func (o RandomShuffleArrayOutput) ToRandomShuffleArrayOutput() RandomShuffleArrayOutput {
	return o
}

func (o RandomShuffleArrayOutput) ToRandomShuffleArrayOutputWithContext(ctx context.Context) RandomShuffleArrayOutput {
	return o
}

func (o RandomShuffleArrayOutput) Index(i pulumi.IntInput) RandomShuffleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomShuffle {
		return vs[0].([]*RandomShuffle)[vs[1].(int)]
	}).(RandomShuffleOutput)
}

type RandomShuffleMapOutput struct{ *pulumi.OutputState }

func (RandomShuffleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomShuffle)(nil)).Elem()
}

func (o RandomShuffleMapOutput) ToRandomShuffleMapOutput() RandomShuffleMapOutput {
	return o
}

func (o RandomShuffleMapOutput) ToRandomShuffleMapOutputWithContext(ctx context.Context) RandomShuffleMapOutput {
	return o
}

func (o RandomShuffleMapOutput) MapIndex(k pulumi.StringInput) RandomShuffleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomShuffle {
		return vs[0].(map[string]*RandomShuffle)[vs[1].(string)]
	}).(RandomShuffleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomShuffleInput)(nil)).Elem(), &RandomShuffle{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomShuffleArrayInput)(nil)).Elem(), RandomShuffleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomShuffleMapInput)(nil)).Elem(), RandomShuffleMap{})
	pulumi.RegisterOutputType(RandomShuffleOutput{})
	pulumi.RegisterOutputType(RandomShuffleArrayOutput{})
	pulumi.RegisterOutputType(RandomShuffleMapOutput{})
}
