# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RandomIntegerArgs', 'RandomInteger']

@pulumi.input_type
class RandomIntegerArgs:
    def __init__(__self__, *,
                 max: pulumi.Input[int],
                 min: pulumi.Input[int],
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 seed: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RandomInteger resource.
        :param pulumi.Input[int] max: The maximum inclusive value of the range.
        :param pulumi.Input[int] min: The minimum inclusive value of the range.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[str] seed: A custom seed to always produce the same value.
        """
        pulumi.set(__self__, "max", max)
        pulumi.set(__self__, "min", min)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if seed is not None:
            pulumi.set(__self__, "seed", seed)

    @property
    @pulumi.getter
    def max(self) -> pulumi.Input[int]:
        """
        The maximum inclusive value of the range.
        """
        return pulumi.get(self, "max")

    @max.setter
    def max(self, value: pulumi.Input[int]):
        pulumi.set(self, "max", value)

    @property
    @pulumi.getter
    def min(self) -> pulumi.Input[int]:
        """
        The minimum inclusive value of the range.
        """
        return pulumi.get(self, "min")

    @min.setter
    def min(self, value: pulumi.Input[int]):
        pulumi.set(self, "min", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will
        trigger a new id to be generated. See
        the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter
    def seed(self) -> Optional[pulumi.Input[str]]:
        """
        A custom seed to always produce the same value.
        """
        return pulumi.get(self, "seed")

    @seed.setter
    def seed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "seed", value)


@pulumi.input_type
class _RandomIntegerState:
    def __init__(__self__, *,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 max: Optional[pulumi.Input[int]] = None,
                 min: Optional[pulumi.Input[int]] = None,
                 result: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RandomInteger resources.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[int] max: The maximum inclusive value of the range.
        :param pulumi.Input[int] min: The minimum inclusive value of the range.
        :param pulumi.Input[int] result: (int) The random Integer result.
        :param pulumi.Input[str] seed: A custom seed to always produce the same value.
        """
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if max is not None:
            pulumi.set(__self__, "max", max)
        if min is not None:
            pulumi.set(__self__, "min", min)
        if result is not None:
            pulumi.set(__self__, "result", result)
        if seed is not None:
            pulumi.set(__self__, "seed", seed)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will
        trigger a new id to be generated. See
        the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter
    def max(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum inclusive value of the range.
        """
        return pulumi.get(self, "max")

    @max.setter
    def max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max", value)

    @property
    @pulumi.getter
    def min(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum inclusive value of the range.
        """
        return pulumi.get(self, "min")

    @min.setter
    def min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min", value)

    @property
    @pulumi.getter
    def result(self) -> Optional[pulumi.Input[int]]:
        """
        (int) The random Integer result.
        """
        return pulumi.get(self, "result")

    @result.setter
    def result(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "result", value)

    @property
    @pulumi.getter
    def seed(self) -> Optional[pulumi.Input[str]]:
        """
        A custom seed to always produce the same value.
        """
        return pulumi.get(self, "seed")

    @seed.setter
    def seed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "seed", value)


class RandomInteger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 max: Optional[pulumi.Input[int]] = None,
                 min: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The resource `RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.

        This resource can be used in conjunction with resources that have
        the `create_before_destroy` lifecycle flag set, to avoid conflicts with
        unique names during the brief period where both the old and new resources
        exist concurrently.

        ## Example Usage

        The following example shows how to generate a random priority between 1 and 50000 for
        a `aws_alb_listener_rule` resource:

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        priority = random.RandomInteger("priority",
            keepers={
                "listener_arn": var["listener_arn"],
            },
            max=50000,
            min=1)
        main = aws.alb.ListenerRule("main",
            actions=[aws.alb.ListenerRuleActionArgs(
                target_group_arn=var["target_group_arn"],
                type="forward",
            )],
            listener_arn=var["listener_arn"],
            priority=priority.result)
        ```

        The result of the above will set a random priority.

        ## Import

        Random integers can be imported using the `result`, `min`, and `max`, with an optional `seed`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a `,`)

        ```sh
         $ pulumi import random:index/randomInteger:RandomInteger priority 15390,1,50000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[int] max: The maximum inclusive value of the range.
        :param pulumi.Input[int] min: The minimum inclusive value of the range.
        :param pulumi.Input[str] seed: A custom seed to always produce the same value.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RandomIntegerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The resource `RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.

        This resource can be used in conjunction with resources that have
        the `create_before_destroy` lifecycle flag set, to avoid conflicts with
        unique names during the brief period where both the old and new resources
        exist concurrently.

        ## Example Usage

        The following example shows how to generate a random priority between 1 and 50000 for
        a `aws_alb_listener_rule` resource:

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        priority = random.RandomInteger("priority",
            keepers={
                "listener_arn": var["listener_arn"],
            },
            max=50000,
            min=1)
        main = aws.alb.ListenerRule("main",
            actions=[aws.alb.ListenerRuleActionArgs(
                target_group_arn=var["target_group_arn"],
                type="forward",
            )],
            listener_arn=var["listener_arn"],
            priority=priority.result)
        ```

        The result of the above will set a random priority.

        ## Import

        Random integers can be imported using the `result`, `min`, and `max`, with an optional `seed`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a `,`)

        ```sh
         $ pulumi import random:index/randomInteger:RandomInteger priority 15390,1,50000
        ```

        :param str resource_name: The name of the resource.
        :param RandomIntegerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RandomIntegerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 max: Optional[pulumi.Input[int]] = None,
                 min: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RandomIntegerArgs.__new__(RandomIntegerArgs)

            __props__.__dict__["keepers"] = keepers
            if max is None and not opts.urn:
                raise TypeError("Missing required property 'max'")
            __props__.__dict__["max"] = max
            if min is None and not opts.urn:
                raise TypeError("Missing required property 'min'")
            __props__.__dict__["min"] = min
            __props__.__dict__["seed"] = seed
            __props__.__dict__["result"] = None
        super(RandomInteger, __self__).__init__(
            'random:index/randomInteger:RandomInteger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            max: Optional[pulumi.Input[int]] = None,
            min: Optional[pulumi.Input[int]] = None,
            result: Optional[pulumi.Input[int]] = None,
            seed: Optional[pulumi.Input[str]] = None) -> 'RandomInteger':
        """
        Get an existing RandomInteger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[int] max: The maximum inclusive value of the range.
        :param pulumi.Input[int] min: The minimum inclusive value of the range.
        :param pulumi.Input[int] result: (int) The random Integer result.
        :param pulumi.Input[str] seed: A custom seed to always produce the same value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RandomIntegerState.__new__(_RandomIntegerState)

        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["max"] = max
        __props__.__dict__["min"] = min
        __props__.__dict__["result"] = result
        __props__.__dict__["seed"] = seed
        return RandomInteger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will
        trigger a new id to be generated. See
        the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter
    def max(self) -> pulumi.Output[int]:
        """
        The maximum inclusive value of the range.
        """
        return pulumi.get(self, "max")

    @property
    @pulumi.getter
    def min(self) -> pulumi.Output[int]:
        """
        The minimum inclusive value of the range.
        """
        return pulumi.get(self, "min")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[int]:
        """
        (int) The random Integer result.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def seed(self) -> pulumi.Output[Optional[str]]:
        """
        A custom seed to always produce the same value.
        """
        return pulumi.get(self, "seed")

