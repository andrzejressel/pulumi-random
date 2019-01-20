# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from . import utilities, tables

class RandomShuffle(pulumi.CustomResource):
    inputs: pulumi.Output[list]
    """
    The list of strings to shuffle.
    """
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new id to be generated. See
    the main provider documentation for more information.
    """
    results: pulumi.Output[list]
    """
    Random permutation of the list of strings given in `input`.
    """
    result_count: pulumi.Output[int]
    """
    The number of results to return. Defaults to
    the number of items in the `input` list. If fewer items are requested,
    some elements will be excluded from the result. If more items are requested,
    items will be repeated in the result but not more frequently than the number
    of items in the input list.
    """
    seed: pulumi.Output[str]
    """
    Arbitrary string with which to seed the random number
    generator, in order to produce less-volatile permutations of the list.
    **Important:** Even with an identical seed, it is not guaranteed that the
    same permutation will be produced across different versions of Terraform.
    This argument causes the result to be *less volatile*, but not fixed for
    all time.
    """
    def __init__(__self__, __name__, __opts__=None, inputs=None, keepers=None, result_count=None, seed=None):
        """
        The resource `random_shuffle` generates a random permutation of a list
        of strings given as an argument.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] inputs: The list of strings to shuffle.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[int] result_count: The number of results to return. Defaults to
               the number of items in the `input` list. If fewer items are requested,
               some elements will be excluded from the result. If more items are requested,
               items will be repeated in the result but not more frequently than the number
               of items in the input list.
        :param pulumi.Input[str] seed: Arbitrary string with which to seed the random number
               generator, in order to produce less-volatile permutations of the list.
               **Important:** Even with an identical seed, it is not guaranteed that the
               same permutation will be produced across different versions of Terraform.
               This argument causes the result to be *less volatile*, but not fixed for
               all time.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not inputs:
            raise TypeError('Missing required property inputs')
        __props__['inputs'] = inputs

        __props__['keepers'] = keepers

        __props__['result_count'] = result_count

        __props__['seed'] = seed

        __props__['results'] = None

        super(RandomShuffle, __self__).__init__(
            'random:index/randomShuffle:RandomShuffle',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

