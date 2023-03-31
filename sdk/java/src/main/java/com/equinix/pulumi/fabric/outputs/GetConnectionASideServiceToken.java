// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetConnectionASideServiceToken {
    /**
     * @return Customer-provided connection description
     * 
     */
    private String description;
    /**
     * @return Connection URI information
     * 
     */
    private String href;
    /**
     * @return Defines the connection type like VG*VC, EVPL*VC, EPL*VC, EC*VC, GW*VC, ACCESS*EPL_VC
     * 
     */
    private String type;
    /**
     * @return Equinix-assigned connection identifier
     * 
     */
    private String uuid;

    private GetConnectionASideServiceToken() {}
    /**
     * @return Customer-provided connection description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Connection URI information
     * 
     */
    public String href() {
        return this.href;
    }
    /**
     * @return Defines the connection type like VG*VC, EVPL*VC, EPL*VC, EC*VC, GW*VC, ACCESS*EPL_VC
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return Equinix-assigned connection identifier
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionASideServiceToken defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String href;
        private String type;
        private String uuid;
        public Builder() {}
        public Builder(GetConnectionASideServiceToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.href = defaults.href;
    	      this.type = defaults.type;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder href(String href) {
            this.href = Objects.requireNonNull(href);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetConnectionASideServiceToken build() {
            final var o = new GetConnectionASideServiceToken();
            o.description = description;
            o.href = href;
            o.type = type;
            o.uuid = uuid;
            return o;
        }
    }
}