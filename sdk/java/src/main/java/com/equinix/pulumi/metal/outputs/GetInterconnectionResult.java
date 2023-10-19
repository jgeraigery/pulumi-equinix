// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.metal.outputs;

import com.equinix.pulumi.metal.outputs.GetInterconnectionPort;
import com.equinix.pulumi.metal.outputs.GetInterconnectionServiceToken;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInterconnectionResult {
    private String connectionId;
    /**
     * @return The preferred email used for communication and notifications about the Equinix Fabric interconnection.
     * 
     */
    private String contactEmail;
    /**
     * @return Description of the connection resource.
     * 
     */
    private String description;
    /**
     * @return (**Deprecated**) Slug of a facility to which the connection belongs. Use metro instead; read the facility to metro migration guide
     * 
     * @deprecated
     * Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
     * 
     */
    @Deprecated /* Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices */
    private String facility;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Slug of a metro to which the connection belongs.
     * 
     */
    private String metro;
    /**
     * @return Mode for connections in IBX facilities with the dedicated type - standard or tunnel.
     * 
     */
    private String mode;
    /**
     * @return Port name.
     * 
     */
    private String name;
    /**
     * @return ID of the organization where the connection is scoped to.
     * 
     */
    private String organizationId;
    /**
     * @return List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
     * 
     */
    private List<GetInterconnectionPort> ports;
    /**
     * @return ID of project to which the connection belongs.
     * 
     */
    private String projectId;
    /**
     * @return Connection redundancy, reduntant or primary.
     * 
     */
    private String redundancy;
    /**
     * @return Type of service token, a_side or z_side. One available in shared connection.
     * 
     */
    private String serviceTokenType;
    /**
     * @return List of connection service tokens with attributes
     * 
     */
    private List<GetInterconnectionServiceToken> serviceTokens;
    /**
     * @return Port speed in bits per second.
     * 
     */
    private String speed;
    /**
     * @return Port status.
     * 
     */
    private String status;
    /**
     * @return String list of tags.
     * 
     */
    private List<String> tags;
    /**
     * @return (Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the equinix_ecx_l2_connection resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). If your organization already has connection service tokens enabled, use `service_tokens` instead.
     * 
     * @deprecated
     * If your organization already has connection service tokens enabled, use `service_tokens` instead
     * 
     */
    @Deprecated /* If your organization already has connection service tokens enabled, use `service_tokens` instead */
    private String token;
    /**
     * @return Token type, `a_side` or `z_side`.
     * 
     */
    private String type;
    /**
     * @return Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection.
     * 
     */
    private List<Integer> vlans;

    private GetInterconnectionResult() {}
    public String connectionId() {
        return this.connectionId;
    }
    /**
     * @return The preferred email used for communication and notifications about the Equinix Fabric interconnection.
     * 
     */
    public String contactEmail() {
        return this.contactEmail;
    }
    /**
     * @return Description of the connection resource.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return (**Deprecated**) Slug of a facility to which the connection belongs. Use metro instead; read the facility to metro migration guide
     * 
     * @deprecated
     * Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
     * 
     */
    @Deprecated /* Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices */
    public String facility() {
        return this.facility;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Slug of a metro to which the connection belongs.
     * 
     */
    public String metro() {
        return this.metro;
    }
    /**
     * @return Mode for connections in IBX facilities with the dedicated type - standard or tunnel.
     * 
     */
    public String mode() {
        return this.mode;
    }
    /**
     * @return Port name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return ID of the organization where the connection is scoped to.
     * 
     */
    public String organizationId() {
        return this.organizationId;
    }
    /**
     * @return List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
     * 
     */
    public List<GetInterconnectionPort> ports() {
        return this.ports;
    }
    /**
     * @return ID of project to which the connection belongs.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Connection redundancy, reduntant or primary.
     * 
     */
    public String redundancy() {
        return this.redundancy;
    }
    /**
     * @return Type of service token, a_side or z_side. One available in shared connection.
     * 
     */
    public String serviceTokenType() {
        return this.serviceTokenType;
    }
    /**
     * @return List of connection service tokens with attributes
     * 
     */
    public List<GetInterconnectionServiceToken> serviceTokens() {
        return this.serviceTokens;
    }
    /**
     * @return Port speed in bits per second.
     * 
     */
    public String speed() {
        return this.speed;
    }
    /**
     * @return Port status.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return String list of tags.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return (Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the equinix_ecx_l2_connection resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). If your organization already has connection service tokens enabled, use `service_tokens` instead.
     * 
     * @deprecated
     * If your organization already has connection service tokens enabled, use `service_tokens` instead
     * 
     */
    @Deprecated /* If your organization already has connection service tokens enabled, use `service_tokens` instead */
    public String token() {
        return this.token;
    }
    /**
     * @return Token type, `a_side` or `z_side`.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection.
     * 
     */
    public List<Integer> vlans() {
        return this.vlans;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInterconnectionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String connectionId;
        private String contactEmail;
        private String description;
        private String facility;
        private String id;
        private String metro;
        private String mode;
        private String name;
        private String organizationId;
        private List<GetInterconnectionPort> ports;
        private String projectId;
        private String redundancy;
        private String serviceTokenType;
        private List<GetInterconnectionServiceToken> serviceTokens;
        private String speed;
        private String status;
        private List<String> tags;
        private String token;
        private String type;
        private List<Integer> vlans;
        public Builder() {}
        public Builder(GetInterconnectionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionId = defaults.connectionId;
    	      this.contactEmail = defaults.contactEmail;
    	      this.description = defaults.description;
    	      this.facility = defaults.facility;
    	      this.id = defaults.id;
    	      this.metro = defaults.metro;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.organizationId = defaults.organizationId;
    	      this.ports = defaults.ports;
    	      this.projectId = defaults.projectId;
    	      this.redundancy = defaults.redundancy;
    	      this.serviceTokenType = defaults.serviceTokenType;
    	      this.serviceTokens = defaults.serviceTokens;
    	      this.speed = defaults.speed;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.token = defaults.token;
    	      this.type = defaults.type;
    	      this.vlans = defaults.vlans;
        }

        @CustomType.Setter
        public Builder connectionId(String connectionId) {
            this.connectionId = Objects.requireNonNull(connectionId);
            return this;
        }
        @CustomType.Setter
        public Builder contactEmail(String contactEmail) {
            this.contactEmail = Objects.requireNonNull(contactEmail);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder facility(String facility) {
            this.facility = Objects.requireNonNull(facility);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder metro(String metro) {
            this.metro = Objects.requireNonNull(metro);
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder ports(List<GetInterconnectionPort> ports) {
            this.ports = Objects.requireNonNull(ports);
            return this;
        }
        public Builder ports(GetInterconnectionPort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder redundancy(String redundancy) {
            this.redundancy = Objects.requireNonNull(redundancy);
            return this;
        }
        @CustomType.Setter
        public Builder serviceTokenType(String serviceTokenType) {
            this.serviceTokenType = Objects.requireNonNull(serviceTokenType);
            return this;
        }
        @CustomType.Setter
        public Builder serviceTokens(List<GetInterconnectionServiceToken> serviceTokens) {
            this.serviceTokens = Objects.requireNonNull(serviceTokens);
            return this;
        }
        public Builder serviceTokens(GetInterconnectionServiceToken... serviceTokens) {
            return serviceTokens(List.of(serviceTokens));
        }
        @CustomType.Setter
        public Builder speed(String speed) {
            this.speed = Objects.requireNonNull(speed);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder vlans(List<Integer> vlans) {
            this.vlans = Objects.requireNonNull(vlans);
            return this;
        }
        public Builder vlans(Integer... vlans) {
            return vlans(List.of(vlans));
        }
        public GetInterconnectionResult build() {
            final var o = new GetInterconnectionResult();
            o.connectionId = connectionId;
            o.contactEmail = contactEmail;
            o.description = description;
            o.facility = facility;
            o.id = id;
            o.metro = metro;
            o.mode = mode;
            o.name = name;
            o.organizationId = organizationId;
            o.ports = ports;
            o.projectId = projectId;
            o.redundancy = redundancy;
            o.serviceTokenType = serviceTokenType;
            o.serviceTokens = serviceTokens;
            o.speed = speed;
            o.status = status;
            o.tags = tags;
            o.token = token;
            o.type = type;
            o.vlans = vlans;
            return o;
        }
    }
}
