// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random.RandomUuid` generates random uuid string that is intended to be
 * used as unique identifiers for other resources.
 *
 * This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
 * for use with services needed a unique string identifier.
 *
 * ## Example Usage
 *
 * The following example shows how to generate a unique name for an Azure Resource Group.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 *
 * const testRandomUuid = new random.RandomUuid("testRandomUuid", {});
 * const testResourceGroup = new azure.core.ResourceGroup("testResourceGroup", {location: "Central US"});
 * ```
 *
 * ## Import
 *
 * Random UUID's can be imported. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example
 *
 * ```sh
 *  $ pulumi import random:index/randomUuid:RandomUuid main aabbccdd-eeff-0011-2233-445566778899
 * ```
 */
export class RandomUuid extends pulumi.CustomResource {
    /**
     * Get an existing RandomUuid resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomUuidState, opts?: pulumi.CustomResourceOptions): RandomUuid {
        return new RandomUuid(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomUuid:RandomUuid';

    /**
     * Returns true if the given object is an instance of RandomUuid.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomUuid {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomUuid.__pulumiType;
    }

    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new uuid to be generated. See
     * the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The generated uuid presented in string format.
     */
    public /*out*/ readonly result!: pulumi.Output<string>;

    /**
     * Create a RandomUuid resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RandomUuidArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomUuidArgs | RandomUuidState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomUuidState | undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["result"] = state ? state.result : undefined;
        } else {
            const args = argsOrState as RandomUuidArgs | undefined;
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["result"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RandomUuid.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomUuid resources.
 */
export interface RandomUuidState {
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new uuid to be generated. See
     * the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The generated uuid presented in string format.
     */
    result?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomUuid resource.
 */
export interface RandomUuidArgs {
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new uuid to be generated. See
     * the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
}
