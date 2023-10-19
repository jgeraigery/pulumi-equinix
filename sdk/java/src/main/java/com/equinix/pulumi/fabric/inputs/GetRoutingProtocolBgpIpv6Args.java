// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRoutingProtocolBgpIpv6Args extends com.pulumi.resources.ResourceArgs {

    public static final GetRoutingProtocolBgpIpv6Args Empty = new GetRoutingProtocolBgpIpv6Args();

    /**
     * Customer side peering ip
     * 
     */
    @Import(name="customerPeerIp", required=true)
    private Output<String> customerPeerIp;

    /**
     * @return Customer side peering ip
     * 
     */
    public Output<String> customerPeerIp() {
        return this.customerPeerIp;
    }

    /**
     * Admin status for the BGP session
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Admin status for the BGP session
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Equinix side peering ip
     * 
     */
    @Import(name="equinixPeerIp", required=true)
    private Output<String> equinixPeerIp;

    /**
     * @return Equinix side peering ip
     * 
     */
    public Output<String> equinixPeerIp() {
        return this.equinixPeerIp;
    }

    private GetRoutingProtocolBgpIpv6Args() {}

    private GetRoutingProtocolBgpIpv6Args(GetRoutingProtocolBgpIpv6Args $) {
        this.customerPeerIp = $.customerPeerIp;
        this.enabled = $.enabled;
        this.equinixPeerIp = $.equinixPeerIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRoutingProtocolBgpIpv6Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRoutingProtocolBgpIpv6Args $;

        public Builder() {
            $ = new GetRoutingProtocolBgpIpv6Args();
        }

        public Builder(GetRoutingProtocolBgpIpv6Args defaults) {
            $ = new GetRoutingProtocolBgpIpv6Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param customerPeerIp Customer side peering ip
         * 
         * @return builder
         * 
         */
        public Builder customerPeerIp(Output<String> customerPeerIp) {
            $.customerPeerIp = customerPeerIp;
            return this;
        }

        /**
         * @param customerPeerIp Customer side peering ip
         * 
         * @return builder
         * 
         */
        public Builder customerPeerIp(String customerPeerIp) {
            return customerPeerIp(Output.of(customerPeerIp));
        }

        /**
         * @param enabled Admin status for the BGP session
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Admin status for the BGP session
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param equinixPeerIp Equinix side peering ip
         * 
         * @return builder
         * 
         */
        public Builder equinixPeerIp(Output<String> equinixPeerIp) {
            $.equinixPeerIp = equinixPeerIp;
            return this;
        }

        /**
         * @param equinixPeerIp Equinix side peering ip
         * 
         * @return builder
         * 
         */
        public Builder equinixPeerIp(String equinixPeerIp) {
            return equinixPeerIp(Output.of(equinixPeerIp));
        }

        public GetRoutingProtocolBgpIpv6Args build() {
            $.customerPeerIp = Objects.requireNonNull($.customerPeerIp, "expected parameter 'customerPeerIp' to be non-null");
            $.equinixPeerIp = Objects.requireNonNull($.equinixPeerIp, "expected parameter 'equinixPeerIp' to be non-null");
            return $;
        }
    }

}
