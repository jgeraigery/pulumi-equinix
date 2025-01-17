// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointAccount;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointGateway;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointInterface;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointLinkProtocol;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointLocation;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointPort;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointProfile;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointRouter;
import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointVirtualDevice;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetConnectionZSideAccessPoint {
    private List<GetConnectionZSideAccessPointAccount> accounts;
    private String authenticationKey;
    /**
     * @deprecated
     * router attribute will be returned instead
     * 
     */
    @Deprecated /* router attribute will be returned instead */
    private List<GetConnectionZSideAccessPointGateway> gateways;
    private List<GetConnectionZSideAccessPointInterface> interfaces;
    private List<GetConnectionZSideAccessPointLinkProtocol> linkProtocols;
    private List<GetConnectionZSideAccessPointLocation> locations;
    private String peeringType;
    private List<GetConnectionZSideAccessPointPort> ports;
    private List<GetConnectionZSideAccessPointProfile> profiles;
    private String providerConnectionId;
    /**
     * @return CloudRouter; Replaces `gateway` attribute (Set of Object)
     * 
     */
    private List<GetConnectionZSideAccessPointRouter> routers;
    private String sellerRegion;
    private String type;
    private List<GetConnectionZSideAccessPointVirtualDevice> virtualDevices;

    private GetConnectionZSideAccessPoint() {}
    public List<GetConnectionZSideAccessPointAccount> accounts() {
        return this.accounts;
    }
    public String authenticationKey() {
        return this.authenticationKey;
    }
    /**
     * @deprecated
     * router attribute will be returned instead
     * 
     */
    @Deprecated /* router attribute will be returned instead */
    public List<GetConnectionZSideAccessPointGateway> gateways() {
        return this.gateways;
    }
    public List<GetConnectionZSideAccessPointInterface> interfaces() {
        return this.interfaces;
    }
    public List<GetConnectionZSideAccessPointLinkProtocol> linkProtocols() {
        return this.linkProtocols;
    }
    public List<GetConnectionZSideAccessPointLocation> locations() {
        return this.locations;
    }
    public String peeringType() {
        return this.peeringType;
    }
    public List<GetConnectionZSideAccessPointPort> ports() {
        return this.ports;
    }
    public List<GetConnectionZSideAccessPointProfile> profiles() {
        return this.profiles;
    }
    public String providerConnectionId() {
        return this.providerConnectionId;
    }
    /**
     * @return CloudRouter; Replaces `gateway` attribute (Set of Object)
     * 
     */
    public List<GetConnectionZSideAccessPointRouter> routers() {
        return this.routers;
    }
    public String sellerRegion() {
        return this.sellerRegion;
    }
    public String type() {
        return this.type;
    }
    public List<GetConnectionZSideAccessPointVirtualDevice> virtualDevices() {
        return this.virtualDevices;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionZSideAccessPoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetConnectionZSideAccessPointAccount> accounts;
        private String authenticationKey;
        private List<GetConnectionZSideAccessPointGateway> gateways;
        private List<GetConnectionZSideAccessPointInterface> interfaces;
        private List<GetConnectionZSideAccessPointLinkProtocol> linkProtocols;
        private List<GetConnectionZSideAccessPointLocation> locations;
        private String peeringType;
        private List<GetConnectionZSideAccessPointPort> ports;
        private List<GetConnectionZSideAccessPointProfile> profiles;
        private String providerConnectionId;
        private List<GetConnectionZSideAccessPointRouter> routers;
        private String sellerRegion;
        private String type;
        private List<GetConnectionZSideAccessPointVirtualDevice> virtualDevices;
        public Builder() {}
        public Builder(GetConnectionZSideAccessPoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accounts = defaults.accounts;
    	      this.authenticationKey = defaults.authenticationKey;
    	      this.gateways = defaults.gateways;
    	      this.interfaces = defaults.interfaces;
    	      this.linkProtocols = defaults.linkProtocols;
    	      this.locations = defaults.locations;
    	      this.peeringType = defaults.peeringType;
    	      this.ports = defaults.ports;
    	      this.profiles = defaults.profiles;
    	      this.providerConnectionId = defaults.providerConnectionId;
    	      this.routers = defaults.routers;
    	      this.sellerRegion = defaults.sellerRegion;
    	      this.type = defaults.type;
    	      this.virtualDevices = defaults.virtualDevices;
        }

        @CustomType.Setter
        public Builder accounts(List<GetConnectionZSideAccessPointAccount> accounts) {
            this.accounts = Objects.requireNonNull(accounts);
            return this;
        }
        public Builder accounts(GetConnectionZSideAccessPointAccount... accounts) {
            return accounts(List.of(accounts));
        }
        @CustomType.Setter
        public Builder authenticationKey(String authenticationKey) {
            this.authenticationKey = Objects.requireNonNull(authenticationKey);
            return this;
        }
        @CustomType.Setter
        public Builder gateways(List<GetConnectionZSideAccessPointGateway> gateways) {
            this.gateways = Objects.requireNonNull(gateways);
            return this;
        }
        public Builder gateways(GetConnectionZSideAccessPointGateway... gateways) {
            return gateways(List.of(gateways));
        }
        @CustomType.Setter
        public Builder interfaces(List<GetConnectionZSideAccessPointInterface> interfaces) {
            this.interfaces = Objects.requireNonNull(interfaces);
            return this;
        }
        public Builder interfaces(GetConnectionZSideAccessPointInterface... interfaces) {
            return interfaces(List.of(interfaces));
        }
        @CustomType.Setter
        public Builder linkProtocols(List<GetConnectionZSideAccessPointLinkProtocol> linkProtocols) {
            this.linkProtocols = Objects.requireNonNull(linkProtocols);
            return this;
        }
        public Builder linkProtocols(GetConnectionZSideAccessPointLinkProtocol... linkProtocols) {
            return linkProtocols(List.of(linkProtocols));
        }
        @CustomType.Setter
        public Builder locations(List<GetConnectionZSideAccessPointLocation> locations) {
            this.locations = Objects.requireNonNull(locations);
            return this;
        }
        public Builder locations(GetConnectionZSideAccessPointLocation... locations) {
            return locations(List.of(locations));
        }
        @CustomType.Setter
        public Builder peeringType(String peeringType) {
            this.peeringType = Objects.requireNonNull(peeringType);
            return this;
        }
        @CustomType.Setter
        public Builder ports(List<GetConnectionZSideAccessPointPort> ports) {
            this.ports = Objects.requireNonNull(ports);
            return this;
        }
        public Builder ports(GetConnectionZSideAccessPointPort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder profiles(List<GetConnectionZSideAccessPointProfile> profiles) {
            this.profiles = Objects.requireNonNull(profiles);
            return this;
        }
        public Builder profiles(GetConnectionZSideAccessPointProfile... profiles) {
            return profiles(List.of(profiles));
        }
        @CustomType.Setter
        public Builder providerConnectionId(String providerConnectionId) {
            this.providerConnectionId = Objects.requireNonNull(providerConnectionId);
            return this;
        }
        @CustomType.Setter
        public Builder routers(List<GetConnectionZSideAccessPointRouter> routers) {
            this.routers = Objects.requireNonNull(routers);
            return this;
        }
        public Builder routers(GetConnectionZSideAccessPointRouter... routers) {
            return routers(List.of(routers));
        }
        @CustomType.Setter
        public Builder sellerRegion(String sellerRegion) {
            this.sellerRegion = Objects.requireNonNull(sellerRegion);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder virtualDevices(List<GetConnectionZSideAccessPointVirtualDevice> virtualDevices) {
            this.virtualDevices = Objects.requireNonNull(virtualDevices);
            return this;
        }
        public Builder virtualDevices(GetConnectionZSideAccessPointVirtualDevice... virtualDevices) {
            return virtualDevices(List.of(virtualDevices));
        }
        public GetConnectionZSideAccessPoint build() {
            final var _resultValue = new GetConnectionZSideAccessPoint();
            _resultValue.accounts = accounts;
            _resultValue.authenticationKey = authenticationKey;
            _resultValue.gateways = gateways;
            _resultValue.interfaces = interfaces;
            _resultValue.linkProtocols = linkProtocols;
            _resultValue.locations = locations;
            _resultValue.peeringType = peeringType;
            _resultValue.ports = ports;
            _resultValue.profiles = profiles;
            _resultValue.providerConnectionId = providerConnectionId;
            _resultValue.routers = routers;
            _resultValue.sellerRegion = sellerRegion;
            _resultValue.type = type;
            _resultValue.virtualDevices = virtualDevices;
            return _resultValue;
        }
    }
}
