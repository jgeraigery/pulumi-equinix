# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_equinix.config as __config
    config = __config
    import pulumi_equinix.fabric as __fabric
    fabric = __fabric
    import pulumi_equinix.metal as __metal
    metal = __metal
    import pulumi_equinix.networkedge as __networkedge
    networkedge = __networkedge
else:
    config = _utilities.lazy_import('pulumi_equinix.config')
    fabric = _utilities.lazy_import('pulumi_equinix.fabric')
    metal = _utilities.lazy_import('pulumi_equinix.metal')
    networkedge = _utilities.lazy_import('pulumi_equinix.networkedge')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "equinix",
  "mod": "fabric/connection",
  "fqn": "pulumi_equinix.fabric",
  "classes": {
   "equinix:fabric/connection:Connection": "Connection"
  }
 },
 {
  "pkg": "equinix",
  "mod": "fabric/serviceProfile",
  "fqn": "pulumi_equinix.fabric",
  "classes": {
   "equinix:fabric/serviceProfile:ServiceProfile": "ServiceProfile"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/bgpSession",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/bgpSession:BgpSession": "BgpSession"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/connection",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/connection:Connection": "Connection"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/device",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/device:Device": "Device"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/deviceNetworkType",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/deviceNetworkType:DeviceNetworkType": "DeviceNetworkType"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/gateway",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/gateway:Gateway": "Gateway"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/ipAttachment",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/ipAttachment:IpAttachment": "IpAttachment"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/organization",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/organization:Organization": "Organization"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/organizationMember",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/organizationMember:OrganizationMember": "OrganizationMember"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/port",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/port:Port": "Port"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/portVlanAttachment",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/portVlanAttachment:PortVlanAttachment": "PortVlanAttachment"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/project",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/project:Project": "Project"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/projectApiKey",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/projectApiKey:ProjectApiKey": "ProjectApiKey"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/projectSshKey",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/projectSshKey:ProjectSshKey": "ProjectSshKey"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/reservedIpBlock",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/reservedIpBlock:ReservedIpBlock": "ReservedIpBlock"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/spotMarketRequest",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/spotMarketRequest:SpotMarketRequest": "SpotMarketRequest"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/sshKey",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/sshKey:SshKey": "SshKey"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/userApiKey",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/userApiKey:UserApiKey": "UserApiKey"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/virtualCircuit",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/virtualCircuit:VirtualCircuit": "VirtualCircuit"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/vlan",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/vlan:Vlan": "Vlan"
  }
 },
 {
  "pkg": "equinix",
  "mod": "metal/vrf",
  "fqn": "pulumi_equinix.metal",
  "classes": {
   "equinix:metal/vrf:Vrf": "Vrf"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/aclTemplate",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/aclTemplate:AclTemplate": "AclTemplate"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/bgp",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/bgp:Bgp": "Bgp"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/device",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/device:Device": "Device"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/deviceLink",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/deviceLink:DeviceLink": "DeviceLink"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/sshKey",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/sshKey:SshKey": "SshKey"
  }
 },
 {
  "pkg": "equinix",
  "mod": "networkedge/sshUser",
  "fqn": "pulumi_equinix.networkedge",
  "classes": {
   "equinix:networkedge/sshUser:SshUser": "SshUser"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "equinix",
  "token": "pulumi:providers:equinix",
  "fqn": "pulumi_equinix",
  "class": "Provider"
 }
]
"""
)
