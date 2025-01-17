// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRoutingProtocolBgpIpv4 extends com.pulumi.resources.InvokeArgs {

    public static final GetRoutingProtocolBgpIpv4 Empty = new GetRoutingProtocolBgpIpv4();

    /**
     * Customer side peering ip
     * 
     */
    @Import(name="customerPeerIp", required=true)
    private String customerPeerIp;

    /**
     * @return Customer side peering ip
     * 
     */
    public String customerPeerIp() {
        return this.customerPeerIp;
    }

    /**
     * Admin status for the BGP session
     * 
     */
    @Import(name="enabled")
    private @Nullable Boolean enabled;

    /**
     * @return Admin status for the BGP session
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Equinix side peering ip
     * 
     */
    @Import(name="equinixPeerIp", required=true)
    private String equinixPeerIp;

    /**
     * @return Equinix side peering ip
     * 
     */
    public String equinixPeerIp() {
        return this.equinixPeerIp;
    }

    private GetRoutingProtocolBgpIpv4() {}

    private GetRoutingProtocolBgpIpv4(GetRoutingProtocolBgpIpv4 $) {
        this.customerPeerIp = $.customerPeerIp;
        this.enabled = $.enabled;
        this.equinixPeerIp = $.equinixPeerIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRoutingProtocolBgpIpv4 defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRoutingProtocolBgpIpv4 $;

        public Builder() {
            $ = new GetRoutingProtocolBgpIpv4();
        }

        public Builder(GetRoutingProtocolBgpIpv4 defaults) {
            $ = new GetRoutingProtocolBgpIpv4(Objects.requireNonNull(defaults));
        }

        /**
         * @param customerPeerIp Customer side peering ip
         * 
         * @return builder
         * 
         */
        public Builder customerPeerIp(String customerPeerIp) {
            $.customerPeerIp = customerPeerIp;
            return this;
        }

        /**
         * @param enabled Admin status for the BGP session
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Boolean enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param equinixPeerIp Equinix side peering ip
         * 
         * @return builder
         * 
         */
        public Builder equinixPeerIp(String equinixPeerIp) {
            $.equinixPeerIp = equinixPeerIp;
            return this;
        }

        public GetRoutingProtocolBgpIpv4 build() {
            $.customerPeerIp = Objects.requireNonNull($.customerPeerIp, "expected parameter 'customerPeerIp' to be non-null");
            $.equinixPeerIp = Objects.requireNonNull($.equinixPeerIp, "expected parameter 'equinixPeerIp' to be non-null");
            return $;
        }
    }

}
