# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RoutingProtocolArgs', 'RoutingProtocol']

@pulumi.input_type
class RoutingProtocolArgs:
    def __init__(__self__, *,
                 connection_uuid: pulumi.Input[str],
                 bfd: Optional[pulumi.Input['RoutingProtocolBfdArgs']] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_ipv4: Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']] = None,
                 bgp_ipv6: Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']] = None,
                 customer_asn: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direct_ipv4: Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']] = None,
                 direct_ipv6: Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RoutingProtocol resource.
        :param pulumi.Input[str] connection_uuid: Connection URI associated with Routing Protocol
        :param pulumi.Input['RoutingProtocolBfdArgs'] bfd: Bidirectional Forwarding Detection
        :param pulumi.Input[str] bgp_auth_key: BGP authorization key
        :param pulumi.Input['RoutingProtocolBgpIpv4Args'] bgp_ipv4: Routing Protocol BGP IPv4
        :param pulumi.Input['RoutingProtocolBgpIpv6Args'] bgp_ipv6: Routing Protocol BGP IPv6
        :param pulumi.Input[int] customer_asn: Customer-provided ASN
        :param pulumi.Input[str] description: Customer-provided Fabric Routing Protocol description
        :param pulumi.Input['RoutingProtocolDirectIpv4Args'] direct_ipv4: Routing Protocol Direct IPv4
        :param pulumi.Input['RoutingProtocolDirectIpv6Args'] direct_ipv6: Routing Protocol Direct IPv6
        :param pulumi.Input[str] name: Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        :param pulumi.Input[str] type: Defines the routing protocol type like BGP or DIRECT
        :param pulumi.Input[str] uuid: Equinix-assigned routing protocol identifier
        """
        pulumi.set(__self__, "connection_uuid", connection_uuid)
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if bgp_auth_key is not None:
            pulumi.set(__self__, "bgp_auth_key", bgp_auth_key)
        if bgp_ipv4 is not None:
            pulumi.set(__self__, "bgp_ipv4", bgp_ipv4)
        if bgp_ipv6 is not None:
            pulumi.set(__self__, "bgp_ipv6", bgp_ipv6)
        if customer_asn is not None:
            pulumi.set(__self__, "customer_asn", customer_asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if direct_ipv4 is not None:
            pulumi.set(__self__, "direct_ipv4", direct_ipv4)
        if direct_ipv6 is not None:
            pulumi.set(__self__, "direct_ipv6", direct_ipv6)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="connectionUuid")
    def connection_uuid(self) -> pulumi.Input[str]:
        """
        Connection URI associated with Routing Protocol
        """
        return pulumi.get(self, "connection_uuid")

    @connection_uuid.setter
    def connection_uuid(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_uuid", value)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input['RoutingProtocolBfdArgs']]:
        """
        Bidirectional Forwarding Detection
        """
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input['RoutingProtocolBfdArgs']]):
        pulumi.set(self, "bfd", value)

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        BGP authorization key
        """
        return pulumi.get(self, "bgp_auth_key")

    @bgp_auth_key.setter
    def bgp_auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_auth_key", value)

    @property
    @pulumi.getter(name="bgpIpv4")
    def bgp_ipv4(self) -> Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']]:
        """
        Routing Protocol BGP IPv4
        """
        return pulumi.get(self, "bgp_ipv4")

    @bgp_ipv4.setter
    def bgp_ipv4(self, value: Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']]):
        pulumi.set(self, "bgp_ipv4", value)

    @property
    @pulumi.getter(name="bgpIpv6")
    def bgp_ipv6(self) -> Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']]:
        """
        Routing Protocol BGP IPv6
        """
        return pulumi.get(self, "bgp_ipv6")

    @bgp_ipv6.setter
    def bgp_ipv6(self, value: Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']]):
        pulumi.set(self, "bgp_ipv6", value)

    @property
    @pulumi.getter(name="customerAsn")
    def customer_asn(self) -> Optional[pulumi.Input[int]]:
        """
        Customer-provided ASN
        """
        return pulumi.get(self, "customer_asn")

    @customer_asn.setter
    def customer_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "customer_asn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Customer-provided Fabric Routing Protocol description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="directIpv4")
    def direct_ipv4(self) -> Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']]:
        """
        Routing Protocol Direct IPv4
        """
        return pulumi.get(self, "direct_ipv4")

    @direct_ipv4.setter
    def direct_ipv4(self, value: Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']]):
        pulumi.set(self, "direct_ipv4", value)

    @property
    @pulumi.getter(name="directIpv6")
    def direct_ipv6(self) -> Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']]:
        """
        Routing Protocol Direct IPv6
        """
        return pulumi.get(self, "direct_ipv6")

    @direct_ipv6.setter
    def direct_ipv6(self, value: Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']]):
        pulumi.set(self, "direct_ipv6", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the routing protocol type like BGP or DIRECT
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Equinix-assigned routing protocol identifier
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


@pulumi.input_type
class _RoutingProtocolState:
    def __init__(__self__, *,
                 bfd: Optional[pulumi.Input['RoutingProtocolBfdArgs']] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_ipv4: Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']] = None,
                 bgp_ipv6: Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']] = None,
                 change_logs: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeLogArgs']]]] = None,
                 changes: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeArgs']]]] = None,
                 connection_uuid: Optional[pulumi.Input[str]] = None,
                 customer_asn: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direct_ipv4: Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']] = None,
                 direct_ipv6: Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']] = None,
                 equinix_asn: Optional[pulumi.Input[int]] = None,
                 href: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolOperationArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RoutingProtocol resources.
        :param pulumi.Input['RoutingProtocolBfdArgs'] bfd: Bidirectional Forwarding Detection
        :param pulumi.Input[str] bgp_auth_key: BGP authorization key
        :param pulumi.Input['RoutingProtocolBgpIpv4Args'] bgp_ipv4: Routing Protocol BGP IPv4
        :param pulumi.Input['RoutingProtocolBgpIpv6Args'] bgp_ipv6: Routing Protocol BGP IPv6
        :param pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeLogArgs']]] change_logs: Captures Routing Protocol lifecycle change information
        :param pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeArgs']]] changes: Routing Protocol configuration Changes
        :param pulumi.Input[str] connection_uuid: Connection URI associated with Routing Protocol
        :param pulumi.Input[int] customer_asn: Customer-provided ASN
        :param pulumi.Input[str] description: Customer-provided Fabric Routing Protocol description
        :param pulumi.Input['RoutingProtocolDirectIpv4Args'] direct_ipv4: Routing Protocol Direct IPv4
        :param pulumi.Input['RoutingProtocolDirectIpv6Args'] direct_ipv6: Routing Protocol Direct IPv6
        :param pulumi.Input[int] equinix_asn: Equinix ASN
        :param pulumi.Input[str] href: Routing Protocol URI information
        :param pulumi.Input[str] name: Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        :param pulumi.Input[Sequence[pulumi.Input['RoutingProtocolOperationArgs']]] operations: Routing Protocol type-specific operational data
        :param pulumi.Input[str] state: Routing Protocol overall state
        :param pulumi.Input[str] type: Defines the routing protocol type like BGP or DIRECT
        :param pulumi.Input[str] uuid: Equinix-assigned routing protocol identifier
        """
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if bgp_auth_key is not None:
            pulumi.set(__self__, "bgp_auth_key", bgp_auth_key)
        if bgp_ipv4 is not None:
            pulumi.set(__self__, "bgp_ipv4", bgp_ipv4)
        if bgp_ipv6 is not None:
            pulumi.set(__self__, "bgp_ipv6", bgp_ipv6)
        if change_logs is not None:
            pulumi.set(__self__, "change_logs", change_logs)
        if changes is not None:
            pulumi.set(__self__, "changes", changes)
        if connection_uuid is not None:
            pulumi.set(__self__, "connection_uuid", connection_uuid)
        if customer_asn is not None:
            pulumi.set(__self__, "customer_asn", customer_asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if direct_ipv4 is not None:
            pulumi.set(__self__, "direct_ipv4", direct_ipv4)
        if direct_ipv6 is not None:
            pulumi.set(__self__, "direct_ipv6", direct_ipv6)
        if equinix_asn is not None:
            pulumi.set(__self__, "equinix_asn", equinix_asn)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operations is not None:
            pulumi.set(__self__, "operations", operations)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input['RoutingProtocolBfdArgs']]:
        """
        Bidirectional Forwarding Detection
        """
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input['RoutingProtocolBfdArgs']]):
        pulumi.set(self, "bfd", value)

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        BGP authorization key
        """
        return pulumi.get(self, "bgp_auth_key")

    @bgp_auth_key.setter
    def bgp_auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_auth_key", value)

    @property
    @pulumi.getter(name="bgpIpv4")
    def bgp_ipv4(self) -> Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']]:
        """
        Routing Protocol BGP IPv4
        """
        return pulumi.get(self, "bgp_ipv4")

    @bgp_ipv4.setter
    def bgp_ipv4(self, value: Optional[pulumi.Input['RoutingProtocolBgpIpv4Args']]):
        pulumi.set(self, "bgp_ipv4", value)

    @property
    @pulumi.getter(name="bgpIpv6")
    def bgp_ipv6(self) -> Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']]:
        """
        Routing Protocol BGP IPv6
        """
        return pulumi.get(self, "bgp_ipv6")

    @bgp_ipv6.setter
    def bgp_ipv6(self, value: Optional[pulumi.Input['RoutingProtocolBgpIpv6Args']]):
        pulumi.set(self, "bgp_ipv6", value)

    @property
    @pulumi.getter(name="changeLogs")
    def change_logs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeLogArgs']]]]:
        """
        Captures Routing Protocol lifecycle change information
        """
        return pulumi.get(self, "change_logs")

    @change_logs.setter
    def change_logs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeLogArgs']]]]):
        pulumi.set(self, "change_logs", value)

    @property
    @pulumi.getter
    def changes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeArgs']]]]:
        """
        Routing Protocol configuration Changes
        """
        return pulumi.get(self, "changes")

    @changes.setter
    def changes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolChangeArgs']]]]):
        pulumi.set(self, "changes", value)

    @property
    @pulumi.getter(name="connectionUuid")
    def connection_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Connection URI associated with Routing Protocol
        """
        return pulumi.get(self, "connection_uuid")

    @connection_uuid.setter
    def connection_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_uuid", value)

    @property
    @pulumi.getter(name="customerAsn")
    def customer_asn(self) -> Optional[pulumi.Input[int]]:
        """
        Customer-provided ASN
        """
        return pulumi.get(self, "customer_asn")

    @customer_asn.setter
    def customer_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "customer_asn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Customer-provided Fabric Routing Protocol description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="directIpv4")
    def direct_ipv4(self) -> Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']]:
        """
        Routing Protocol Direct IPv4
        """
        return pulumi.get(self, "direct_ipv4")

    @direct_ipv4.setter
    def direct_ipv4(self, value: Optional[pulumi.Input['RoutingProtocolDirectIpv4Args']]):
        pulumi.set(self, "direct_ipv4", value)

    @property
    @pulumi.getter(name="directIpv6")
    def direct_ipv6(self) -> Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']]:
        """
        Routing Protocol Direct IPv6
        """
        return pulumi.get(self, "direct_ipv6")

    @direct_ipv6.setter
    def direct_ipv6(self, value: Optional[pulumi.Input['RoutingProtocolDirectIpv6Args']]):
        pulumi.set(self, "direct_ipv6", value)

    @property
    @pulumi.getter(name="equinixAsn")
    def equinix_asn(self) -> Optional[pulumi.Input[int]]:
        """
        Equinix ASN
        """
        return pulumi.get(self, "equinix_asn")

    @equinix_asn.setter
    def equinix_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "equinix_asn", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        """
        Routing Protocol URI information
        """
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolOperationArgs']]]]:
        """
        Routing Protocol type-specific operational data
        """
        return pulumi.get(self, "operations")

    @operations.setter
    def operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoutingProtocolOperationArgs']]]]):
        pulumi.set(self, "operations", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Routing Protocol overall state
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the routing protocol type like BGP or DIRECT
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Equinix-assigned routing protocol identifier
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class RoutingProtocol(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bfd: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBfdArgs']]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv4Args']]] = None,
                 bgp_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv6Args']]] = None,
                 connection_uuid: Optional[pulumi.Input[str]] = None,
                 customer_asn: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direct_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv4Args']]] = None,
                 direct_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv6Args']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ```python
        import pulumi
        import pulumi_equinix as equinix

        config = pulumi.Config()
        connection_id = config.require("connectionId")
        routing_protocol = equinix.fabric.RoutingProtocol("RoutingProtocol",
            connection_uuid=connection_id,
            name="My-Direct-route-1",
            type="DIRECT",
            direct_ipv4=equinix.fabric.RoutingProtocolDirectIpv4Args(
                equinix_iface_ip="192.168.100.1/30",
            ))
        pulumi.export("routingProtocolId", routing_protocol.id)
        pulumi.export("routingProtocolState", routing_protocol.state)
        pulumi.export("routingProtocolEquinixAsn", routing_protocol.equinix_asn)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBfdArgs']] bfd: Bidirectional Forwarding Detection
        :param pulumi.Input[str] bgp_auth_key: BGP authorization key
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv4Args']] bgp_ipv4: Routing Protocol BGP IPv4
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv6Args']] bgp_ipv6: Routing Protocol BGP IPv6
        :param pulumi.Input[str] connection_uuid: Connection URI associated with Routing Protocol
        :param pulumi.Input[int] customer_asn: Customer-provided ASN
        :param pulumi.Input[str] description: Customer-provided Fabric Routing Protocol description
        :param pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv4Args']] direct_ipv4: Routing Protocol Direct IPv4
        :param pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv6Args']] direct_ipv6: Routing Protocol Direct IPv6
        :param pulumi.Input[str] name: Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        :param pulumi.Input[str] type: Defines the routing protocol type like BGP or DIRECT
        :param pulumi.Input[str] uuid: Equinix-assigned routing protocol identifier
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoutingProtocolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ```python
        import pulumi
        import pulumi_equinix as equinix

        config = pulumi.Config()
        connection_id = config.require("connectionId")
        routing_protocol = equinix.fabric.RoutingProtocol("RoutingProtocol",
            connection_uuid=connection_id,
            name="My-Direct-route-1",
            type="DIRECT",
            direct_ipv4=equinix.fabric.RoutingProtocolDirectIpv4Args(
                equinix_iface_ip="192.168.100.1/30",
            ))
        pulumi.export("routingProtocolId", routing_protocol.id)
        pulumi.export("routingProtocolState", routing_protocol.state)
        pulumi.export("routingProtocolEquinixAsn", routing_protocol.equinix_asn)
        ```

        :param str resource_name: The name of the resource.
        :param RoutingProtocolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoutingProtocolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bfd: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBfdArgs']]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv4Args']]] = None,
                 bgp_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv6Args']]] = None,
                 connection_uuid: Optional[pulumi.Input[str]] = None,
                 customer_asn: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direct_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv4Args']]] = None,
                 direct_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv6Args']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoutingProtocolArgs.__new__(RoutingProtocolArgs)

            __props__.__dict__["bfd"] = bfd
            __props__.__dict__["bgp_auth_key"] = bgp_auth_key
            __props__.__dict__["bgp_ipv4"] = bgp_ipv4
            __props__.__dict__["bgp_ipv6"] = bgp_ipv6
            if connection_uuid is None and not opts.urn:
                raise TypeError("Missing required property 'connection_uuid'")
            __props__.__dict__["connection_uuid"] = connection_uuid
            __props__.__dict__["customer_asn"] = customer_asn
            __props__.__dict__["description"] = description
            __props__.__dict__["direct_ipv4"] = direct_ipv4
            __props__.__dict__["direct_ipv6"] = direct_ipv6
            __props__.__dict__["name"] = name
            __props__.__dict__["type"] = type
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["change_logs"] = None
            __props__.__dict__["changes"] = None
            __props__.__dict__["equinix_asn"] = None
            __props__.__dict__["href"] = None
            __props__.__dict__["operations"] = None
            __props__.__dict__["state"] = None
        super(RoutingProtocol, __self__).__init__(
            'equinix:fabric/routingProtocol:RoutingProtocol',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bfd: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBfdArgs']]] = None,
            bgp_auth_key: Optional[pulumi.Input[str]] = None,
            bgp_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv4Args']]] = None,
            bgp_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv6Args']]] = None,
            change_logs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolChangeLogArgs']]]]] = None,
            changes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolChangeArgs']]]]] = None,
            connection_uuid: Optional[pulumi.Input[str]] = None,
            customer_asn: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            direct_ipv4: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv4Args']]] = None,
            direct_ipv6: Optional[pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv6Args']]] = None,
            equinix_asn: Optional[pulumi.Input[int]] = None,
            href: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            operations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolOperationArgs']]]]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'RoutingProtocol':
        """
        Get an existing RoutingProtocol resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBfdArgs']] bfd: Bidirectional Forwarding Detection
        :param pulumi.Input[str] bgp_auth_key: BGP authorization key
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv4Args']] bgp_ipv4: Routing Protocol BGP IPv4
        :param pulumi.Input[pulumi.InputType['RoutingProtocolBgpIpv6Args']] bgp_ipv6: Routing Protocol BGP IPv6
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolChangeLogArgs']]]] change_logs: Captures Routing Protocol lifecycle change information
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolChangeArgs']]]] changes: Routing Protocol configuration Changes
        :param pulumi.Input[str] connection_uuid: Connection URI associated with Routing Protocol
        :param pulumi.Input[int] customer_asn: Customer-provided ASN
        :param pulumi.Input[str] description: Customer-provided Fabric Routing Protocol description
        :param pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv4Args']] direct_ipv4: Routing Protocol Direct IPv4
        :param pulumi.Input[pulumi.InputType['RoutingProtocolDirectIpv6Args']] direct_ipv6: Routing Protocol Direct IPv6
        :param pulumi.Input[int] equinix_asn: Equinix ASN
        :param pulumi.Input[str] href: Routing Protocol URI information
        :param pulumi.Input[str] name: Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutingProtocolOperationArgs']]]] operations: Routing Protocol type-specific operational data
        :param pulumi.Input[str] state: Routing Protocol overall state
        :param pulumi.Input[str] type: Defines the routing protocol type like BGP or DIRECT
        :param pulumi.Input[str] uuid: Equinix-assigned routing protocol identifier
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoutingProtocolState.__new__(_RoutingProtocolState)

        __props__.__dict__["bfd"] = bfd
        __props__.__dict__["bgp_auth_key"] = bgp_auth_key
        __props__.__dict__["bgp_ipv4"] = bgp_ipv4
        __props__.__dict__["bgp_ipv6"] = bgp_ipv6
        __props__.__dict__["change_logs"] = change_logs
        __props__.__dict__["changes"] = changes
        __props__.__dict__["connection_uuid"] = connection_uuid
        __props__.__dict__["customer_asn"] = customer_asn
        __props__.__dict__["description"] = description
        __props__.__dict__["direct_ipv4"] = direct_ipv4
        __props__.__dict__["direct_ipv6"] = direct_ipv6
        __props__.__dict__["equinix_asn"] = equinix_asn
        __props__.__dict__["href"] = href
        __props__.__dict__["name"] = name
        __props__.__dict__["operations"] = operations
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        __props__.__dict__["uuid"] = uuid
        return RoutingProtocol(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bfd(self) -> pulumi.Output[Optional['outputs.RoutingProtocolBfd']]:
        """
        Bidirectional Forwarding Detection
        """
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> pulumi.Output[Optional[str]]:
        """
        BGP authorization key
        """
        return pulumi.get(self, "bgp_auth_key")

    @property
    @pulumi.getter(name="bgpIpv4")
    def bgp_ipv4(self) -> pulumi.Output[Optional['outputs.RoutingProtocolBgpIpv4']]:
        """
        Routing Protocol BGP IPv4
        """
        return pulumi.get(self, "bgp_ipv4")

    @property
    @pulumi.getter(name="bgpIpv6")
    def bgp_ipv6(self) -> pulumi.Output[Optional['outputs.RoutingProtocolBgpIpv6']]:
        """
        Routing Protocol BGP IPv6
        """
        return pulumi.get(self, "bgp_ipv6")

    @property
    @pulumi.getter(name="changeLogs")
    def change_logs(self) -> pulumi.Output[Sequence['outputs.RoutingProtocolChangeLog']]:
        """
        Captures Routing Protocol lifecycle change information
        """
        return pulumi.get(self, "change_logs")

    @property
    @pulumi.getter
    def changes(self) -> pulumi.Output[Sequence['outputs.RoutingProtocolChange']]:
        """
        Routing Protocol configuration Changes
        """
        return pulumi.get(self, "changes")

    @property
    @pulumi.getter(name="connectionUuid")
    def connection_uuid(self) -> pulumi.Output[str]:
        """
        Connection URI associated with Routing Protocol
        """
        return pulumi.get(self, "connection_uuid")

    @property
    @pulumi.getter(name="customerAsn")
    def customer_asn(self) -> pulumi.Output[Optional[int]]:
        """
        Customer-provided ASN
        """
        return pulumi.get(self, "customer_asn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Customer-provided Fabric Routing Protocol description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="directIpv4")
    def direct_ipv4(self) -> pulumi.Output[Optional['outputs.RoutingProtocolDirectIpv4']]:
        """
        Routing Protocol Direct IPv4
        """
        return pulumi.get(self, "direct_ipv4")

    @property
    @pulumi.getter(name="directIpv6")
    def direct_ipv6(self) -> pulumi.Output[Optional['outputs.RoutingProtocolDirectIpv6']]:
        """
        Routing Protocol Direct IPv6
        """
        return pulumi.get(self, "direct_ipv6")

    @property
    @pulumi.getter(name="equinixAsn")
    def equinix_asn(self) -> pulumi.Output[int]:
        """
        Equinix ASN
        """
        return pulumi.get(self, "equinix_asn")

    @property
    @pulumi.getter
    def href(self) -> pulumi.Output[str]:
        """
        Routing Protocol URI information
        """
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operations(self) -> pulumi.Output[Sequence['outputs.RoutingProtocolOperation']]:
        """
        Routing Protocol type-specific operational data
        """
        return pulumi.get(self, "operations")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Routing Protocol overall state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Defines the routing protocol type like BGP or DIRECT
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Equinix-assigned routing protocol identifier
        """
        return pulumi.get(self, "uuid")
