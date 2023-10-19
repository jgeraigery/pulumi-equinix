# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BgpSessionArgs', 'BgpSession']

@pulumi.input_type
class BgpSessionArgs:
    def __init__(__self__, *,
                 address_family: pulumi.Input[str],
                 device_id: pulumi.Input[str],
                 default_route: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BgpSession resource.
        :param pulumi.Input[str] address_family: `ipv4` or `ipv6`.
        :param pulumi.Input[str] device_id: ID of device.
        :param pulumi.Input[bool] default_route: Boolean flag to set the default route policy. False by default.
        """
        BgpSessionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address_family=address_family,
            device_id=device_id,
            default_route=default_route,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address_family: pulumi.Input[str],
             device_id: pulumi.Input[str],
             default_route: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'addressFamily' in kwargs:
            address_family = kwargs['addressFamily']
        if 'deviceId' in kwargs:
            device_id = kwargs['deviceId']
        if 'defaultRoute' in kwargs:
            default_route = kwargs['defaultRoute']

        _setter("address_family", address_family)
        _setter("device_id", device_id)
        if default_route is not None:
            _setter("default_route", default_route)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Input[str]:
        """
        `ipv4` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        ID of device.
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="defaultRoute")
    def default_route(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag to set the default route policy. False by default.
        """
        return pulumi.get(self, "default_route")

    @default_route.setter
    def default_route(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_route", value)


@pulumi.input_type
class _BgpSessionState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 default_route: Optional[pulumi.Input[bool]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpSession resources.
        :param pulumi.Input[str] address_family: `ipv4` or `ipv6`.
        :param pulumi.Input[bool] default_route: Boolean flag to set the default route policy. False by default.
        :param pulumi.Input[str] device_id: ID of device.
        :param pulumi.Input[str] status: Status of the session - `up` or `down`
        """
        _BgpSessionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address_family=address_family,
            default_route=default_route,
            device_id=device_id,
            status=status,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address_family: Optional[pulumi.Input[str]] = None,
             default_route: Optional[pulumi.Input[bool]] = None,
             device_id: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'addressFamily' in kwargs:
            address_family = kwargs['addressFamily']
        if 'defaultRoute' in kwargs:
            default_route = kwargs['defaultRoute']
        if 'deviceId' in kwargs:
            device_id = kwargs['deviceId']

        if address_family is not None:
            _setter("address_family", address_family)
        if default_route is not None:
            _setter("default_route", default_route)
        if device_id is not None:
            _setter("device_id", device_id)
        if status is not None:
            _setter("status", status)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        `ipv4` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter(name="defaultRoute")
    def default_route(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag to set the default route policy. False by default.
        """
        return pulumi.get(self, "default_route")

    @default_route.setter
    def default_route(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_route", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of device.
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the session - `up` or `down`
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BgpSession(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 default_route: Optional[pulumi.Input[bool]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage BGP sessions in Equinix Metal Host. Refer to [Equinix Metal BGP documentation](https://metal.equinix.com/developers/docs/networking/local-global-bgp/) for more details.

        You need to have BGP config enabled in your project.

        BGP session must be linked to a device running [BIRD](https://bird.network.cz) or other BGP routing daemon which will control route advertisements via the session to Equinix Metal's upstream routers.

        ## Example Usage
        ```python
        import pulumi
        import pulumi_equinix as equinix

        config = pulumi.Config()
        device_id = config.require("deviceId")
        bgp = equinix.metal.BgpSession("bgp",
            device_id=device_id,
            address_family="ipv4")
        pulumi.export("bgpSessionStatus", bgp.status)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: `ipv4` or `ipv6`.
        :param pulumi.Input[bool] default_route: Boolean flag to set the default route policy. False by default.
        :param pulumi.Input[str] device_id: ID of device.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpSessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage BGP sessions in Equinix Metal Host. Refer to [Equinix Metal BGP documentation](https://metal.equinix.com/developers/docs/networking/local-global-bgp/) for more details.

        You need to have BGP config enabled in your project.

        BGP session must be linked to a device running [BIRD](https://bird.network.cz) or other BGP routing daemon which will control route advertisements via the session to Equinix Metal's upstream routers.

        ## Example Usage
        ```python
        import pulumi
        import pulumi_equinix as equinix

        config = pulumi.Config()
        device_id = config.require("deviceId")
        bgp = equinix.metal.BgpSession("bgp",
            device_id=device_id,
            address_family="ipv4")
        pulumi.export("bgpSessionStatus", bgp.status)
        ```

        :param str resource_name: The name of the resource.
        :param BgpSessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpSessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            BgpSessionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 default_route: Optional[pulumi.Input[bool]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpSessionArgs.__new__(BgpSessionArgs)

            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            __props__.__dict__["default_route"] = default_route
            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["status"] = None
        super(BgpSession, __self__).__init__(
            'equinix:metal/bgpSession:BgpSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            default_route: Optional[pulumi.Input[bool]] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BgpSession':
        """
        Get an existing BgpSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: `ipv4` or `ipv6`.
        :param pulumi.Input[bool] default_route: Boolean flag to set the default route policy. False by default.
        :param pulumi.Input[str] device_id: ID of device.
        :param pulumi.Input[str] status: Status of the session - `up` or `down`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpSessionState.__new__(_BgpSessionState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["default_route"] = default_route
        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["status"] = status
        return BgpSession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        `ipv4` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter(name="defaultRoute")
    def default_route(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag to set the default route policy. False by default.
        """
        return pulumi.get(self, "default_route")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        ID of device.
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the session - `up` or `down`
        """
        return pulumi.get(self, "status")

