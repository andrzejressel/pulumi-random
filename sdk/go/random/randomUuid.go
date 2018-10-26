// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The resource `random_uuid` generates random uuid string that is intended to be
// used as unique identifiers for other resources.
// 
// This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
// for use with services needed a unique string identifier.
// 
type RandomUuid struct {
	s *pulumi.ResourceState
}

// NewRandomUuid registers a new resource with the given unique name, arguments, and options.
func NewRandomUuid(ctx *pulumi.Context,
	name string, args *RandomUuidArgs, opts ...pulumi.ResourceOpt) (*RandomUuid, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keepers"] = nil
	} else {
		inputs["keepers"] = args.Keepers
	}
	inputs["result"] = nil
	s, err := ctx.RegisterResource("random:index/randomUuid:RandomUuid", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomUuid{s: s}, nil
}

// GetRandomUuid gets an existing RandomUuid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomUuid(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RandomUuidState, opts ...pulumi.ResourceOpt) (*RandomUuid, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["keepers"] = state.Keepers
		inputs["result"] = state.Result
	}
	s, err := ctx.ReadResource("random:index/randomUuid:RandomUuid", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomUuid{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RandomUuid) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RandomUuid) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Arbitrary map of values that, when changed, will
// trigger a new uuid to be generated. See
// the main provider documentation for more information.
func (r *RandomUuid) Keepers() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["keepers"])
}

// The generated uuid presented in string format.
func (r *RandomUuid) Result() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["result"])
}

// Input properties used for looking up and filtering RandomUuid resources.
type RandomUuidState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated. See
	// the main provider documentation for more information.
	Keepers interface{}
	// The generated uuid presented in string format.
	Result interface{}
}

// The set of arguments for constructing a RandomUuid resource.
type RandomUuidArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated. See
	// the main provider documentation for more information.
	Keepers interface{}
}
