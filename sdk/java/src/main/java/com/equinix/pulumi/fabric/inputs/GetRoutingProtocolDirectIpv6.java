// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRoutingProtocolDirectIpv6 extends com.pulumi.resources.InvokeArgs {

    public static final GetRoutingProtocolDirectIpv6 Empty = new GetRoutingProtocolDirectIpv6();

    /**
     * Equinix side Interface IP address
     * 
     */
    @Import(name="equinixIfaceIp")
    private @Nullable String equinixIfaceIp;

    /**
     * @return Equinix side Interface IP address
     * 
     */
    public Optional<String> equinixIfaceIp() {
        return Optional.ofNullable(this.equinixIfaceIp);
    }

    private GetRoutingProtocolDirectIpv6() {}

    private GetRoutingProtocolDirectIpv6(GetRoutingProtocolDirectIpv6 $) {
        this.equinixIfaceIp = $.equinixIfaceIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRoutingProtocolDirectIpv6 defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRoutingProtocolDirectIpv6 $;

        public Builder() {
            $ = new GetRoutingProtocolDirectIpv6();
        }

        public Builder(GetRoutingProtocolDirectIpv6 defaults) {
            $ = new GetRoutingProtocolDirectIpv6(Objects.requireNonNull(defaults));
        }

        /**
         * @param equinixIfaceIp Equinix side Interface IP address
         * 
         * @return builder
         * 
         */
        public Builder equinixIfaceIp(@Nullable String equinixIfaceIp) {
            $.equinixIfaceIp = equinixIfaceIp;
            return this;
        }

        public GetRoutingProtocolDirectIpv6 build() {
            return $;
        }
    }

}
