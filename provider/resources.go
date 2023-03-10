// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package equinix

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/equinix/pulumi-equinix/provider/pkg/version"
	"github.com/equinix/terraform-provider-equinix/equinix"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	pulumiSchema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	equinixPkg = "equinix"
	// modules:
	equinixMod 		= "index" 		// the equinix root module
	fabricMod  		= "Fabric"  	// Equinix Fabric
	metalMod		= "Metal"  		// Equinix Metal
	networkEdgeMod 	= "NetworkEdge"	// Equinix Network Edge

)

var namespaceMap = map[string]string{
	"equinix": "Equinix",
}

// makeEquinixResource
func makeEquinixResource(moduleTitle, mem string) tokens.Type {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	return tfbridge.MakeResource(equinixPkg, moduleName, mem)
}

// makeEquinixDataSource 
func makeEquinixDataSource(moduleTitle, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	return tfbridge.MakeDataSource(equinixPkg, moduleName, "get"+mem)
}

// makeEquinixType
func makeEquinixType(moduleTitle, res string) tokens.Type {
	return tokens.Type(tokens.ModuleMember(makeEquinixToken(moduleTitle, res)))
}

// makeEquinixToken
func makeEquinixToken(moduleTitle, res string) string {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return strings.Join([]string{equinixPkg, strings.ToLower(moduleTitle) + "/" + fn, res}, ":")
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(equinix.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "equinix",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Equinix",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Equinix",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/equinix/pulumi-equinix",
		Description:       "A Pulumi package for creating and managing equinix cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "equinix", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://deploy.equinix.com/",
		Repository: "https://github.com/equinix/pulumi-equinix",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "equinix",
		Config:    map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		// IgnoreMappings is a list of TF resources and data sources to ignore in mappings errors
		IgnoreMappings: []string{
			"equinix_ecx_l2_connection", // to be deprecated in terraform in favor of Fabric v4 equinix_fabric_connection
			"equinix_ecx_l2_connection_accepter", // deprecated in terraform
			"equinix_ecx_l2_serviceprofile", // to be deprecated in terraform in favor of Fabric v4 equinix_fabric_service_profile
			"equinix_ecx_l2_sellerprofile", // to be deprecated in terraform in favor of Fabric v4 equinix_fabric_service_profile datasource
			"equinix_ecx_l2_sellerprofiles", // to be deprecated in terraform in favor of Fabric v4 equinix_fabric_service_profiles datasource
			"equinix_ecx_port", // to be deprecated in terraform in favor of Fabric v4 equinix_fabric_port datasource
		},
		Resources:            map[string]*tfbridge.ResourceInfo{
			// Equinix Fabric v4
			"equinix_fabric_connection":      {
				Tok: makeEquinixResource(fabricMod, "Connection"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ConnectionType")},
					},
					"a_side": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"access_point": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"peering_type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointPeeringType")},
											},
											"type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointType")},
											},
											"location": {
												Elem: &tfbridge.SchemaInfo{
													Fields: map[string]*tfbridge.SchemaInfo{
														"metro_code": {
															Type:     "string",
															AltTypes: []tokens.Type{makeEquinixType(equinixMod, "Metro")},
														},
													},
												},
											},
											"link_protocol": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointLinkProtocolType")},
											},
											"profile": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileType")},
											},
										},
									},
								},
								"service_token": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ServiceTokenType")},
											},
										},
									},
								},
							},
						},
					},
					"z_side": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"access_point": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"peering_type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointPeeringType")},
											},
											"type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointType")},
											},
											"location": {
												Elem: &tfbridge.SchemaInfo{
													Fields: map[string]*tfbridge.SchemaInfo{
														"metro_code": {
															Type:     "string",
															AltTypes: []tokens.Type{makeEquinixType(equinixMod, "Metro")},
														},
													},
												},
											},
											"link_protocol": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "AccessPointLinkProtocolType")},
											},
											"profile": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileType")},
											},
										},
									},
								},
								"service_token": {
									Elem: &tfbridge.SchemaInfo{
										Fields: map[string]*tfbridge.SchemaInfo{
											"type": {
												Type:     "string",
												AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ServiceTokenType")},
											},
										},
									},
								},
							},
						},
					},
					"notifications": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"type": {
									Type:     "string",
									AltTypes: []tokens.Type{makeEquinixType(fabricMod, "NotificationsType")},
								},
							},
						},
					},
				},
			},
			"equinix_fabric_service_profile": {
				Tok: makeEquinixResource(fabricMod, "ServiceProfile"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileType")},
					},
					"state": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileState")},
					},
					"visibility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileVisibility")},
					},
					"access_point_type_configs": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"type": {
									Type:     "string",
									AltTypes: []tokens.Type{makeEquinixType(fabricMod, "ProfileAccessPointType")},
								},
							},
						},
					},
					"notifications": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"type": {
									Type:     "string",
									AltTypes: []tokens.Type{makeEquinixType(fabricMod, "NotificationsType")},
								},
							},
						},
					},
				},
			},
			// Equinix Metal v1
			"equinix_metal_bgp_session":          {Tok: makeEquinixResource(metalMod, "BgpSession")},
			"equinix_metal_connection":           {Tok: makeEquinixResource(metalMod, "Connection")},
			"equinix_metal_device":      		  {
				Tok: makeEquinixResource(metalMod, "Device"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"billing_cycle": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "BillingCycle")},
					},
					"facilities": {
						Elem: &tfbridge.SchemaInfo{
							Type:     "string",
							AltTypes: []tokens.Type{makeEquinixType(metalMod, "Facility")},
						},
					},
					"network_type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "NetworkType")},
					},
					"operating_system": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "OperatingSystem")},
					},
					"plan": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "Plan")},
					},
				},
			},
			"equinix_metal_device_network_type":  {Tok: makeEquinixResource(metalMod, "DeviceNetworkType")},
			"equinix_metal_gateway":      		  {Tok: makeEquinixResource(metalMod, "Gateway")},
			"equinix_metal_ip_attachment":        {Tok: makeEquinixResource(metalMod, "IpAttachment")},
			"equinix_metal_organization":      	  {Tok: makeEquinixResource(metalMod, "Organization")},
			"equinix_metal_organization_member":  {Tok: makeEquinixResource(metalMod, "OrganizationMember")},
			"equinix_metal_port":      			  {Tok: makeEquinixResource(metalMod, "Port")},
			"equinix_metal_port_vlan_attachment": {Tok: makeEquinixResource(metalMod, "PortVlanAttachment")},
			"equinix_metal_project":      		  {Tok: makeEquinixResource(metalMod, "Project")},
			"equinix_metal_project_api_key":      {Tok: makeEquinixResource(metalMod, "ProjectApiKey")},
			"equinix_metal_project_ssh_key":      {Tok: makeEquinixResource(metalMod, "ProjectSshKey")},
			"equinix_metal_reserved_ip_block": {
				Tok: makeEquinixResource(metalMod, "ReservedIpBlock"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"facility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "Facility")},
					},
					"type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "IpBlockType")},
					},
				},
			},
			"equinix_metal_spot_market_request":  {Tok: makeEquinixResource(metalMod, "SpotMarketRequest")},
			"equinix_metal_ssh_key":      		  {Tok: makeEquinixResource(metalMod, "SshKey")},
			"equinix_metal_user_api_key":      	  {Tok: makeEquinixResource(metalMod, "UserApiKey")},
			"equinix_metal_virtual_circuit":      {Tok: makeEquinixResource(metalMod, "VirtualCircuit")},
			"equinix_metal_vlan":      			  {
				Tok: makeEquinixResource(metalMod, "Vlan"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"facility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(metalMod, "Facility")},
					},
				},
			},
			"equinix_metal_vrf":      			  {Tok: makeEquinixResource(metalMod, "Vrf")},
			// Network Edge v1
			"equinix_network_acl_template": {
				Tok: makeEquinixResource(networkEdgeMod, "AclTemplate"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"inbound_rule": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"protocol": {
									Type:     "string",
									AltTypes: []tokens.Type{makeEquinixType(networkEdgeMod, "AclProtocolType")},
								},
							},
						},
					},
				},
			},
			"equinix_network_bgp":     	 	{Tok: makeEquinixResource(networkEdgeMod, "Bgp")},
			"equinix_network_device":       {
				Tok: makeEquinixResource(networkEdgeMod, "Device"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"throughput_unit": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(networkEdgeMod, "ThroughputUnit")},
					},
				},
			},
			"equinix_network_device_link":  {Tok: makeEquinixResource(networkEdgeMod, "DeviceLink")},
			"equinix_network_ssh_key":      {Tok: makeEquinixResource(networkEdgeMod, "SshKey")},
			"equinix_network_ssh_user":     {Tok: makeEquinixResource(networkEdgeMod, "SshUser")},
			"equinix_network_file":         {
				Tok: makeEquinixResource(networkEdgeMod, "NetworkFile"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"metro_code": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(equinixMod, "Metro")},
					},
					"process_type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeEquinixType(networkEdgeMod, "FileType")},
					},
				},
			},
		},
		ExtraTypes: map[string]pulumiSchema.ComplexTypeSpec{
			makeEquinixToken(equinixMod, "Metro"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Amsterdam", Value: "AM"},
					{Name: "Ashburn", Value: "DC"},
					{Name: "Atlanta", Value: "AT"},
					{Name: "Barcelona", Value: "BA"},
					{Name: "Bogota", Value: "BG"},
					{Name: "Bordeaux", Value: "BX"},
					{Name: "Boston", Value: "BO"},
					{Name: "Brussels", Value: "BL"},
					{Name: "Calgary", Value: "CL"},
					{Name: "Canberra", Value: "CA"},
					{Name: "Chicago", Value: "CH"},
					{Name: "Dallas", Value: "DA"},
					{Name: "Denver", Value: "DE"},
					{Name: "Dubai", Value: "DX"},
					{Name: "Dublin", Value: "DB"},
					{Name: "Frankfurt", Value: "FR"},
					{Name: "Geneva", Value: "GV"},
					{Name: "Hamburg", Value: "HH"},
					{Name: "Helsinki", Value: "HE"},
					{Name: "HongKong", Value: "HK"},
					{Name: "Istanbul", Value: "IL"},
					{Name: "Kamloops", Value: "KA"},
					{Name: "Lisbon", Value: "LS"},
					{Name: "London", Value: "LD"},
					{Name: "LosAngeles", Value: "LA"},
					{Name: "Madrid", Value: "MD"},
					{Name: "Manchester", Value: "MA"},
					{Name: "Melbourne", Value: "ME"},
					{Name: "MexicoCity", Value: "MX"},
					{Name: "Miami", Value: "MI"},
					{Name: "Milan", Value: "ML"},
					{Name: "Montreal", Value: "MT"},
					{Name: "Mumbai", Value: "MB"},
					{Name: "Munich", Value: "MU"},
					{Name: "NewYork", Value: "NY"},
					{Name: "Osaka", Value: "OS"},
					{Name: "Paris", Value: "PA"},
					{Name: "Perth", Value: "PE"},
					{Name: "Philadelphia", Value: "PH"},
					{Name: "RioDeJaneiro", Value: "RJ"},
					{Name: "SaoPaulo", Value: "SP"},
					{Name: "Seattle", Value: "SE"},
					{Name: "Seoul", Value: "SL"},
					{Name: "SiliconValley", Value: "SV"},
					{Name: "Singapore", Value: "SG"},
					{Name: "Sofia", Value: "SO"},
					{Name: "Stockholm", Value: "SK"},
					{Name: "Sydney", Value: "SY"},
					{Name: "Tokyo", Value: "TY"},
					{Name: "Toronto", Value: "TR"},
					{Name: "Vancouver", Value: "VA"},
					{Name: "Warsaw", Value: "WA"},
					{Name: "Winnipeg", Value: "WI"},
					{Name: "Zurich", Value: "ZH"},
				},
			},
			makeEquinixToken(fabricMod, "ServiceTokenType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "VCToken", Value: "VC_TOKEN"},
				},
			},
			makeEquinixToken(fabricMod, "AccessPointLinkProtocolType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Untagged", Value: "UNTAGGED"},
					{Name: "Dot1q", Value: "DOT1Q"},
					{Name: "QinQ", Value: "QINQ"},
					{Name: "EVPN_VXLAN", Value: "EVPN_VXLAN"},
				},
			},
			makeEquinixToken(fabricMod, "AccessPointType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Colo", Value: "COLO", Description: "Colocation"},
					{Name: "VD", Value: "VD", Description: "Virtual Device"},
					{Name: "SP", Value: "SP", Description: "Service Profile"},
					{Name: "IGW", Value: "IGW", Description: "Internet Gateway"},
					{Name: "Subnet", Value: "SUBNET", Description: "Subnet"},
					{Name: "GW", Value: "GW", Description: "Gateway"},
					{Name: "Network", Value: "NETWORK", Description: "Network"},
				},
			},
			makeEquinixToken(fabricMod, "AccessPointPeeringType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Private", Value: "PRIVATE"},
					{Name: "Microsoft", Value: "MICROSOFT"},
					{Name: "Public", Value: "PUBLIC"},
				},
			},
			makeEquinixToken(fabricMod, "ConnectionType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "VG", Value: "VG_VC", Description: "Virtual Gateway"},
					{Name: "EVPL", Value: "EVPL_VC", Description: "Ethernet Virtual Private Line"},
					{Name: "EPL", Value: "EPL_VC", Description: "Ethernet Private Line"},
					{Name: "GW", Value: "GW_VC", Description: "Fabric Gateway virtual connection"},
					{Name: "AccessEPL", Value: "ACCESS_EPL_VC", Description: "E-access, layer 2 connection between a QINQ port and an EPL port."},
				},
			},
			makeEquinixToken(fabricMod, "NotificationsType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "All", Value: "ALL"},
					{Name: "ConnectionApproval", Value: "CONNECTION_APPROVAL"},
					{Name: "SalesNotifications", Value: "SALES_REP_NOTIFICATIONS"},
					{Name: "Notifications", Value: "NOTIFICATIONS"},
				},
			},
			makeEquinixToken(fabricMod, "ProfileType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "L2Profile", Value: "L2_PROFILE"},
					{Name: "L3Profile", Value: "L3_PROFILE"},
				},
			},
			makeEquinixToken(fabricMod, "ProfileState"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Active", Value: "ACTIVE"},
					{Name: "PendingApproval", Value: "PENDING_APPROVAL"},
					{Name: "Deleted", Value: "DELETED"},
					{Name: "Rejected", Value: "REJECTED"},
				},
			},
			makeEquinixToken(fabricMod, "ProfileVisibility"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Public", Value: "PUBLIC"},
					{Name: "Private", Value: "PRIVATE"},
				},
			},
			makeEquinixToken(fabricMod, "ProfileAccessPointType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Colo", Value: "COLO", Description: "Colocation"},
					{Name: "VD", Value: "VD", Description: "Virtual Device"},
				},
			},
			makeEquinixToken(networkEdgeMod, "AclProtocolType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "IP", Value: "IP"},
					{Name: "TCP", Value: "TCP"},
					{Name: "UDP", Value: "UDP"},
				},
			},
			makeEquinixToken(networkEdgeMod, "ThroughputUnit"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Mbps", Value: "Mbps"},
					{Name: "Gbps", Value: "Gbps"},
				},
			},
			makeEquinixToken(networkEdgeMod, "FileType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "License", Value: "LICENSE"},
					{Name: "CloudInit", Value: "CLOUD_INIT"},
				},
			},
			makeEquinixToken(metalMod, "BillingCycle"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Hourly", Value: "hourly"},
					{Name: "Monthly", Value: "monthly"},
				},
			},
			makeEquinixToken(metalMod, "IpBlockType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "GlobalIPv4", Value: "global_ipv4"},
					{Name: "PublicIPv4", Value: "public_ipv4"},
				},
			},
			makeEquinixToken(metalMod, "NetworkType"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Layer3", Value: "layer3"},
					{Name: "Layer2Individual", Value: "layer2-individual"},
					{Name: "Layer2Bonded", Value: "layer2-bonded"},
					{Name: "Hybrid", Value: "hybrid"},
				},
			},
			makeEquinixToken(metalMod, "Plan"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
					Description: "See https://deploy.equinix.com/developers/api/metal/#tag/Plans/operation/findPlans",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "A3LargeX86", Value: "a3.large.x86"},
					{Name: "C2LargeARM", Value: "c2.large.arm"},
					{Name: "C2MediumX86", Value: "c2.medium.x86"},
					{Name: "C3LargeARM", Value: "c3.large.arm64"},
					{Name: "C3MediumX86", Value: "c3.medium.x86"},
					{Name: "C3SmallX86", Value: "c3.small.x86"},
					{Name: "F3LargeX86", Value: "f3.large.x86"},
					{Name: "F3MediumX86", Value: "f3.medium.x86"},
					{Name: "G2LargeX86", Value: "g2.large.x86"},
					{Name: "M2XLargeX86", Value: "m2.xlarge.x86"},
					{Name: "M3LargeX86", Value: "m3.large.x86"},
					{Name: "M3SmallX86", Value: "m3.small.x86"},
					{Name: "N2XLargeX86", Value: "n2.xlarge.x86"},
					{Name: "N3XLargeX86", Value: "n3.xlarge.x86"},
					{Name: "S3XLargeX86", Value: "s3.xlarge.x86"},
					{Name: "T3SmallX86", Value: "t3.small.x86"},
					{Name: "X2XLargeX86", Value: "x2.xlarge.x86"},
					{Name: "X3XLargeX86", Value: "x3.xlarge.x86"},
				},
			},
			makeEquinixToken(metalMod, "OperatingSystem"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
					Description: "See https://deploy.equinix.com/developers/api/metal/#tag/OperatingSystems/operation/findOperatingSystems",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "Alma8", Value: "alma_8"},
					{Name: "Alma9", Value: "alma_9"},
					{Name: "Alpine3", Value: "alpine_3"},
					{Name: "CentOS6", Value: "centos_6"},
					{Name: "CentOS7", Value: "centos_7"},
					{Name: "CentOS8", Value: "centos_8"},
					{Name: "CoreOSAlpha", Value: "coreos_alpha"},
					{Name: "CoreOSBeta", Value: "coreos_beta"},
					{Name: "CoreOSStable", Value: "coreos_stable"},
					{Name: "CustomIPXE", Value: "custom_ipxe"},
					{Name: "Debian10", Value: "debian_10"},
					{Name: "Debian11", Value: "debian_11"},
					{Name: "Debian7", Value: "debian_7"},
					{Name: "Debian8", Value: "debian_8"},
					{Name: "Debian9", Value: "debian_9"},
					{Name: "Fedora31", Value: "fedora_31"},
					{Name: "Fedora34", Value: "fedora_34"},
					{Name: "FlatcarAlpha", Value: "flatcar_alpha"},
					{Name: "FlatcarBeta", Value: "flatcar_beta"},
					{Name: "FlatcarEdge", Value: "flatcar_edge"},
					{Name: "FlatcarLTS", Value: "flatcar_lts"},
					{Name: "FlatcarStable", Value: "flatcar_stable"},
					{Name: "FreeBSD10_3", Value: "freebsd_10_3" },
					{Name: "FreeBSD10_4", Value: "freebsd_10_4"},
					{Name: "FreeBSD11_0", Value: "freebsd_11_0"},
					{Name: "FreeBSD11_1", Value: "freebsd_11_1"},
					{Name: "FreeBSD11_2", Value: "freebsd_11_2"},
					{Name: "FreeBSD11_3", Value: "freebsd_11_3"},
					{Name: "FreeBSD11_4", Value: "freebsd_11_4"},
					{Name: "FreeBSD12Testing", Value: "freebsd_12_testing"},
					{Name: "FreeBSD12_1", Value: "freebsd_12_1"},
					{Name: "FreeBSD13_0", Value: "freebsd_13_0"},
					{Name: "FreeBSD13_1", Value: "freebsd_13_1"},
					{Name: "FreeBSD13_testing", Value: "freebsd_13_testing"},
					{Name: "HookEsxi7", Value: "hook_esxi7"},
					{Name: "NixOS17_03", Value: "nixos_17_03"},
					{Name: "NixOS18_03", Value: "nixos_18_03"},
					{Name: "NixOS19_03", Value: "nixos_19_03"},
					{Name: "NixOS20_09", Value: "nixos_20_09"},
					{Name: "NixOS21_11", Value: "nixos_21_11"},
					{Name: "NixOS22_05", Value: "nixos_22_05"},
					{Name: "NixOS22_11", Value: "nixos_22_11"},
					{Name: "Nixos21_05", Value: "nixos_21_05"},
					{Name: "NutanixCentos", Value: "nutanix_centos"},
					{Name: "NutanixLTS5_19_1", Value: "nutanix_lts_5_19_1"},
					{Name: "NutanixLTS5_20", Value: "nutanix_lts_5_20"},
					{Name: "OpenSUSE42_3", Value: "opensuse_42_3"},
					{Name: "RHEL7", Value: "rhel_7"},
					{Name: "RHEL8", Value: "rhel_8"},
					{Name: "RHEL9", Value: "rhel_9"},
					{Name: "RancherOS", Value: "rancher"},
					{Name: "Rocky8", Value: "rocky_8"},
					{Name: "Rocky9", Value: "rocky_9"},
					{Name: "SLES12SP3", Value: "suse_sles12_sp3"},
					{Name: "ScientificLinux6", Value: "scientific_6"},
					{Name: "TalosV1", Value: "talos_v1"},
					{Name: "Ubuntu1710", Value: "ubuntu_17_10"},
					{Name: "Ubuntu18_04", Value: "ubuntu_18_04"},
					{Name: "Ubuntu1904", Value: "ubuntu_19_04"},
					{Name: "Ubuntu19_10", Value: "ubuntu_19_10"},
					{Name: "Ubuntu20_04", Value: "ubuntu_20_04"},
					{Name: "Ubuntu20_04_legacy", Value: "ubuntu_20_04_legacy"},
					{Name: "Ubuntu20_10", Value: "ubuntu_20_10"},
					{Name: "Ubuntu2204", Value: "ubuntu_22_04"},
					{Name: "VMWareAlloyVcf", Value: "vmware_alloy_vcf"},
					{Name: "VMWareEsxi5_5", Value: "vmware_esxi_5_5"},
					{Name: "VMWareEsxi6_0", Value: "vmware_esxi_6_0"},
					{Name: "VMWareEsxi6_5", Value: "vmware_esxi_6_5"},
					{Name: "VMWareEsxi6_7", Value: "vmware_esxi_6_7"},
					{Name: "VMWareEsxi6_7_vcf", Value: "vmware_esxi_6_7_vcf"},
					{Name: "VMWareEsxi7_0", Value: "vmware_esxi_7_0"},
					{Name: "VMWareEsxi7_0U2a", Value: "vmware_esxi_7_0U2a"},
					{Name: "VMWareEsxi7_0_vcf", Value: "vmware_esxi_7_0_vcf"},
					{Name: "VMWareNsx2_5_0", Value: "vmware_nsx_2_5_0"},
					{Name: "VMWareNsx3_0_0", Value: "vmware_nsx_3_0_0"},
					{Name: "Virtuozzo7", Value: "virtuozzo_7"},
					{Name: "Windows2012R2", Value: "windows_2012_r2"},
					{Name: "Windows2016", Value: "windows_2016"},
					{Name: "Windows2019", Value: "windows_2019"},
					{Name: "Windows2022", Value: "windows_2022"},
				},
			},
			makeEquinixToken(metalMod, "Facility"): {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
					Description: "See https://deploy.equinix.com/developers/api/metal/#tag/Facilities/operation/findFacilities",
				},
				Enum: []pulumiSchema.EnumValueSpec{
					{Name: "AM2", Value: "am2", Description: "Amsterdam 2"},
					{Name: "AM6", Value: "am6", Description: "Amsterdam 6"},
					{Name: "MA5", Value: "ma5", Description: "Manchester 5"},
					{Name: "NRT1", Value: "nrt1", Description: "Tokio 1"},
					{Name: "PA4", Value: "pa4", Description: "Paris 4"},
					{Name: "SK2", Value: "sk2", Description: "Stockholm 2"},
					{Name: "ME2", Value: "me2", Description: "Melbourne 2"},
					{Name: "HK2", Value: "hk2", Description: "Hong Kong 2"},
					{Name: "TY11", Value: "ty11", Description: "Tokyo 11"},
					{Name: "LA4", Value: "la4", Description: "Los Angeles 4"},
					{Name: "DA6", Value: "da6", Description: "Dallas 6"},
					{Name: "DA11", Value: "da11", Description: "Dallas 11"},
					{Name: "DA3", Value: "da3", Description: "Dallas 3"},
					{Name: "SP4", Value: "sp4", Description: "Sao Paulo 4"},
					{Name: "MT1", Value: "mt1", Description: "Montreal 1"},
					{Name: "SV16", Value: "sv16", Description: "Silicon Valley 16"},
					{Name: "SJC1", Value: "sjc1", Description: "Sunnyvale, CA 1"},
					{Name: "FRA2", Value: "fra2", Description: "Frankfurt 2"},
					{Name: "FRA8", Value: "fr8", Description: "Frankfurt 8"},
					{Name: "NY5", Value: "ny5", Description: "New York 5"},
					{Name: "NY6", Value: "ny6", Description: "New York 6"},
					{Name: "NY7", Value: "ny7", Description: "New York 7"},
					{Name: "CH3", Value: "ch3", Description: "Chicago 3"},
					{Name: "SL1", Value: "sl1", Description: "Seoul 1"},
					{Name: "SY5", Value: "sy5", Description: "Sydney 5"},
					{Name: "OS3", Value: "os3", Description: "Osaka 3"},
					{Name: "LD7", Value: "ld7", Description: "London 7"},
					{Name: "DC10", Value: "dc10", Description: "Washington DC 10"},
					{Name: "AMS1", Value: "ams1", Description: "Amsterdam 1"},
					{Name: "SG4", Value: "sg4", Description: "Singapore 4"},
					{Name: "SE4", Value: "se4", Description: "Seattle 4"},
					{Name: "SY4", Value: "sy4", Description: "Sydney 4"},
					{Name: "AT4", Value: "at4", Description: "Atlanta 4"},
					{Name: "DFW2", Value: "dfw2", Description: "Dallas 2"},
					{Name: "TR2", Value: "tr2", Description: "Toronto"},
					{Name: "DC13", Value: "dc13", Description: "Washington DC"},
					{Name: "HE7", Value: "he7", Description: "Helsinki"},
					{Name: "EWR1", Value: "ewr1", Description: "Parsippany, NJ 1"},
					{Name: "SG5", Value: "sg5", Description: "Singapore 5"},
					{Name: "SG1", Value: "sg1", Description: "Singapore 1"},
					{Name: "MD2", Value: "md2", Description: "Madrid 2"},
					{Name: "SV15", Value: "sv15", Description: "Silicon Valley 15"},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Equinix Fabric v4
			"equinix_fabric_connection":       {Tok: makeEquinixDataSource(fabricMod, "Connection")},
			"equinix_fabric_port":             {Tok: makeEquinixDataSource(fabricMod, "Port")},
			"equinix_fabric_ports":            {Tok: makeEquinixDataSource(fabricMod, "Ports")},
			"equinix_fabric_service_profile":  {Tok: makeEquinixDataSource(fabricMod, "ServiceProfile")},
			"equinix_fabric_service_profiles": {Tok: makeEquinixDataSource(fabricMod, "ServiceProfiles")},
			// Equinix Metal v1
			"equinix_metal_connection":           {Tok: makeEquinixDataSource(metalMod, "Connection")},
			"equinix_metal_device":               {Tok: makeEquinixDataSource(metalMod, "Device")},
			"equinix_metal_device_bgp_neighbors": {Tok: makeEquinixDataSource(metalMod, "DeviceBgpNeighbors")},
			"equinix_metal_facility":             {Tok: makeEquinixDataSource(metalMod, "Facility")},
			"equinix_metal_gateway":              {Tok: makeEquinixDataSource(metalMod, "Gateway")},
			"equinix_metal_hardware_reservation": {Tok: makeEquinixDataSource(metalMod, "HardwareReservation")},
			"equinix_metal_ip_block_ranges":      {Tok: makeEquinixDataSource(metalMod, "IpBlockRanges")},
			"equinix_metal_metro":                {Tok: makeEquinixDataSource(metalMod, "Metro")},
			"equinix_metal_operating_system":     {Tok: makeEquinixDataSource(metalMod, "OperatingSystem")},
			"equinix_metal_organization":         {Tok: makeEquinixDataSource(metalMod, "Organization")},
			"equinix_metal_plans":                {Tok: makeEquinixDataSource(metalMod, "Plans")},
			"equinix_metal_port":                 {Tok: makeEquinixDataSource(metalMod, "Port")},
			"equinix_metal_precreated_ip_block":  {Tok: makeEquinixDataSource(metalMod, "PrecreatedIpBlock")},
			"equinix_metal_project":              {Tok: makeEquinixDataSource(metalMod, "Project")},
			"equinix_metal_project_ssh_key":      {Tok: makeEquinixDataSource(metalMod, "ProjectSshKey")},
			"equinix_metal_reserved_ip_block":    {Tok: makeEquinixDataSource(metalMod, "ReservedIpBlock")},
			"equinix_metal_spot_market_price":    {Tok: makeEquinixDataSource(metalMod, "SpotMarketPrice")},
			"equinix_metal_spot_market_request":  {Tok: makeEquinixDataSource(metalMod, "SpotMarketRequest")},
			"equinix_metal_virtual_circuit":      {Tok: makeEquinixDataSource(metalMod, "VirtualCircuit")},
			"equinix_metal_vlan":                 {Tok: makeEquinixDataSource(metalMod, "Vlan")},
			"equinix_metal_vrf":                  {Tok: makeEquinixDataSource(metalMod, "Vrf")},
			// Network Edge v1
			"equinix_network_account":         {Tok: makeEquinixDataSource(networkEdgeMod, "Account")},
			"equinix_network_device":          {Tok: makeEquinixDataSource(networkEdgeMod, "Device")},
			"equinix_network_device_platform": {Tok: makeEquinixDataSource(networkEdgeMod, "DevicePlatform")},
			"equinix_network_device_software": {Tok: makeEquinixDataSource(networkEdgeMod, "DeviceSoftware")},
			"equinix_network_device_type":     {Tok: makeEquinixDataSource(networkEdgeMod, "DeviceType")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@equinix/pulumi-equinix",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_equinix",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/equinix/pulumi-%[1]s/sdk/", equinixPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				equinixPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumi",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.equinix.pulumi",
			BuildFiles: "gradle",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
