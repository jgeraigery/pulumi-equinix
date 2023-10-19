# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPortResult',
    'AwaitableGetPortResult',
    'get_port',
    'get_port_output',
]

@pulumi.output_type
class GetPortResult:
    """
    A collection of values returned by getPort.
    """
    def __init__(__self__, bond_id=None, bond_name=None, bonded=None, device_id=None, disbond_supported=None, id=None, layer2=None, mac=None, name=None, native_vlan_id=None, network_type=None, port_id=None, type=None, vlan_ids=None, vxlan_ids=None):
        if bond_id and not isinstance(bond_id, str):
            raise TypeError("Expected argument 'bond_id' to be a str")
        pulumi.set(__self__, "bond_id", bond_id)
        if bond_name and not isinstance(bond_name, str):
            raise TypeError("Expected argument 'bond_name' to be a str")
        pulumi.set(__self__, "bond_name", bond_name)
        if bonded and not isinstance(bonded, bool):
            raise TypeError("Expected argument 'bonded' to be a bool")
        pulumi.set(__self__, "bonded", bonded)
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if disbond_supported and not isinstance(disbond_supported, bool):
            raise TypeError("Expected argument 'disbond_supported' to be a bool")
        pulumi.set(__self__, "disbond_supported", disbond_supported)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if layer2 and not isinstance(layer2, bool):
            raise TypeError("Expected argument 'layer2' to be a bool")
        pulumi.set(__self__, "layer2", layer2)
        if mac and not isinstance(mac, str):
            raise TypeError("Expected argument 'mac' to be a str")
        pulumi.set(__self__, "mac", mac)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if native_vlan_id and not isinstance(native_vlan_id, str):
            raise TypeError("Expected argument 'native_vlan_id' to be a str")
        pulumi.set(__self__, "native_vlan_id", native_vlan_id)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vlan_ids and not isinstance(vlan_ids, list):
            raise TypeError("Expected argument 'vlan_ids' to be a list")
        pulumi.set(__self__, "vlan_ids", vlan_ids)
        if vxlan_ids and not isinstance(vxlan_ids, list):
            raise TypeError("Expected argument 'vxlan_ids' to be a list")
        pulumi.set(__self__, "vxlan_ids", vxlan_ids)

    @property
    @pulumi.getter(name="bondId")
    def bond_id(self) -> str:
        """
        UUID of the bond port.
        """
        return pulumi.get(self, "bond_id")

    @property
    @pulumi.getter(name="bondName")
    def bond_name(self) -> str:
        """
        Name of the bond port.
        """
        return pulumi.get(self, "bond_name")

    @property
    @pulumi.getter
    def bonded(self) -> bool:
        """
        Flag indicating whether the port is bonded.
        """
        return pulumi.get(self, "bonded")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[str]:
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="disbondSupported")
    def disbond_supported(self) -> bool:
        """
        Flag indicating whether the port can be removed from a bond.
        """
        return pulumi.get(self, "disbond_supported")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def layer2(self) -> bool:
        return pulumi.get(self, "layer2")

    @property
    @pulumi.getter
    def mac(self) -> str:
        """
        MAC address of the port.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nativeVlanId")
    def native_vlan_id(self) -> str:
        """
        UUID of native VLAN of the port.
        """
        return pulumi.get(self, "native_vlan_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        One of `layer2-bonded`, `layer2-individual`, `layer3`, `hybrid`, `hybrid-bonded`.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[str]:
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type is either `NetworkBondPort` for bond ports or `NetworkPort` for bondable ethernet ports.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vlanIds")
    def vlan_ids(self) -> Sequence[str]:
        """
        UUIDs of attached VLANs.
        """
        return pulumi.get(self, "vlan_ids")

    @property
    @pulumi.getter(name="vxlanIds")
    def vxlan_ids(self) -> Sequence[int]:
        """
        VXLAN ids of attached VLANs.
        """
        return pulumi.get(self, "vxlan_ids")


class AwaitableGetPortResult(GetPortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortResult(
            bond_id=self.bond_id,
            bond_name=self.bond_name,
            bonded=self.bonded,
            device_id=self.device_id,
            disbond_supported=self.disbond_supported,
            id=self.id,
            layer2=self.layer2,
            mac=self.mac,
            name=self.name,
            native_vlan_id=self.native_vlan_id,
            network_type=self.network_type,
            port_id=self.port_id,
            type=self.type,
            vlan_ids=self.vlan_ids,
            vxlan_ids=self.vxlan_ids)


def get_port(device_id: Optional[str] = None,
             name: Optional[str] = None,
             port_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortResult:
    """
    Use this data source to read ports of existing devices. You can read port by either its UUID,
    or by a device UUID and port name.

    ## Example Usage

    Create a device and read it's eth0 port to the datasource.

    ```python
    import pulumi
    import pulumi_equinix as equinix

    project_id = "<UUID_of_your_project>"
    test_device = equinix.metal.Device("testDevice",
        hostname="tfacc-test-device-port",
        plan="c3.medium.x86",
        metro="sv",
        operating_system="ubuntu_20_04",
        billing_cycle="hourly",
        project_id=project_id)
    test_port = equinix.metal.get_port_output(device_id=test_device.id,
        name="eth0")
    ```


    :param str device_id: Device UUID where to lookup the port.
    :param str name: Name of the port to look up, i.e. `bond0`, `eth1`.
    :param str port_id: ID of the port to read, conflicts with `device_id`.
    """
    __args__ = dict()
    __args__['deviceId'] = device_id
    __args__['name'] = name
    __args__['portId'] = port_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('equinix:metal/getPort:getPort', __args__, opts=opts, typ=GetPortResult).value

    return AwaitableGetPortResult(
        bond_id=pulumi.get(__ret__, 'bond_id'),
        bond_name=pulumi.get(__ret__, 'bond_name'),
        bonded=pulumi.get(__ret__, 'bonded'),
        device_id=pulumi.get(__ret__, 'device_id'),
        disbond_supported=pulumi.get(__ret__, 'disbond_supported'),
        id=pulumi.get(__ret__, 'id'),
        layer2=pulumi.get(__ret__, 'layer2'),
        mac=pulumi.get(__ret__, 'mac'),
        name=pulumi.get(__ret__, 'name'),
        native_vlan_id=pulumi.get(__ret__, 'native_vlan_id'),
        network_type=pulumi.get(__ret__, 'network_type'),
        port_id=pulumi.get(__ret__, 'port_id'),
        type=pulumi.get(__ret__, 'type'),
        vlan_ids=pulumi.get(__ret__, 'vlan_ids'),
        vxlan_ids=pulumi.get(__ret__, 'vxlan_ids'))


@_utilities.lift_output_func(get_port)
def get_port_output(device_id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    port_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortResult]:
    """
    Use this data source to read ports of existing devices. You can read port by either its UUID,
    or by a device UUID and port name.

    ## Example Usage

    Create a device and read it's eth0 port to the datasource.

    ```python
    import pulumi
    import pulumi_equinix as equinix

    project_id = "<UUID_of_your_project>"
    test_device = equinix.metal.Device("testDevice",
        hostname="tfacc-test-device-port",
        plan="c3.medium.x86",
        metro="sv",
        operating_system="ubuntu_20_04",
        billing_cycle="hourly",
        project_id=project_id)
    test_port = equinix.metal.get_port_output(device_id=test_device.id,
        name="eth0")
    ```


    :param str device_id: Device UUID where to lookup the port.
    :param str name: Name of the port to look up, i.e. `bond0`, `eth1`.
    :param str port_id: ID of the port to read, conflicts with `device_id`.
    """
    ...
