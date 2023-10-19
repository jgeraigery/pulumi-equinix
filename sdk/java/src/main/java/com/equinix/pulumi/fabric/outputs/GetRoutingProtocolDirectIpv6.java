// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRoutingProtocolDirectIpv6 {
    /**
     * @return Equinix side Interface IP address
     * 
     */
    private @Nullable String equinixIfaceIp;

    private GetRoutingProtocolDirectIpv6() {}
    /**
     * @return Equinix side Interface IP address
     * 
     */
    public Optional<String> equinixIfaceIp() {
        return Optional.ofNullable(this.equinixIfaceIp);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRoutingProtocolDirectIpv6 defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String equinixIfaceIp;
        public Builder() {}
        public Builder(GetRoutingProtocolDirectIpv6 defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.equinixIfaceIp = defaults.equinixIfaceIp;
        }

        @CustomType.Setter
        public Builder equinixIfaceIp(@Nullable String equinixIfaceIp) {
            this.equinixIfaceIp = equinixIfaceIp;
            return this;
        }
        public GetRoutingProtocolDirectIpv6 build() {
            final var o = new GetRoutingProtocolDirectIpv6();
            o.equinixIfaceIp = equinixIfaceIp;
            return o;
        }
    }
}
