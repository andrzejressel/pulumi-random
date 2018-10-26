# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class RandomUuid(pulumi.CustomResource):
    """
    The resource `random_uuid` generates random uuid string that is intended to be
    used as unique identifiers for other resources.
    
    This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
    for use with services needed a unique string identifier.
    
    """
    def __init__(__self__, __name__, __opts__=None, keepers=None):
        """Create a RandomUuid resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if keepers and not isinstance(keepers, dict):
            raise TypeError('Expected property keepers to be a dict')
        __self__.keepers = keepers
        """
        Arbitrary map of values that, when changed, will
        trigger a new uuid to be generated. See
        the main provider documentation for more information.
        """
        __props__['keepers'] = keepers

        __self__.result = pulumi.runtime.UNKNOWN
        """
        The generated uuid presented in string format.
        """

        super(RandomUuid, __self__).__init__(
            'random:index/randomUuid:RandomUuid',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'keepers' in outs:
            self.keepers = outs['keepers']
        if 'result' in outs:
            self.result = outs['result']
