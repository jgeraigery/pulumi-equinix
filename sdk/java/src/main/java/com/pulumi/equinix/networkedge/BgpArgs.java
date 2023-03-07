// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinix.networkedge;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpArgs extends com.pulumi.resources.ResourceArgs {

    public static final BgpArgs Empty = new BgpArgs();

    /**
     * shared key used for BGP peer authentication.
     * 
     */
    @Import(name="authenticationKey")
    private @Nullable Output<String> authenticationKey;

    /**
     * @return shared key used for BGP peer authentication.
     * 
     */
    public Optional<Output<String>> authenticationKey() {
        return Optional.ofNullable(this.authenticationKey);
    }

    /**
     * identifier of a connection established between.
     * network device and remote service provider that will be used for peering.
     * 
     */
    @Import(name="connectionId", required=true)
    private Output<String> connectionId;

    /**
     * @return identifier of a connection established between.
     * network device and remote service provider that will be used for peering.
     * 
     */
    public Output<String> connectionId() {
        return this.connectionId;
    }

    /**
     * Local ASN number.
     * 
     */
    @Import(name="localAsn", required=true)
    private Output<Integer> localAsn;

    /**
     * @return Local ASN number.
     * 
     */
    public Output<Integer> localAsn() {
        return this.localAsn;
    }

    /**
     * IP address in CIDR format of a local device.
     * 
     */
    @Import(name="localIpAddress", required=true)
    private Output<String> localIpAddress;

    /**
     * @return IP address in CIDR format of a local device.
     * 
     */
    public Output<String> localIpAddress() {
        return this.localIpAddress;
    }

    /**
     * Remote ASN number.
     * 
     */
    @Import(name="remoteAsn", required=true)
    private Output<Integer> remoteAsn;

    /**
     * @return Remote ASN number.
     * 
     */
    public Output<Integer> remoteAsn() {
        return this.remoteAsn;
    }

    /**
     * IP address of remote peer.
     * 
     */
    @Import(name="remoteIpAddress", required=true)
    private Output<String> remoteIpAddress;

    /**
     * @return IP address of remote peer.
     * 
     */
    public Output<String> remoteIpAddress() {
        return this.remoteIpAddress;
    }

    private BgpArgs() {}

    private BgpArgs(BgpArgs $) {
        this.authenticationKey = $.authenticationKey;
        this.connectionId = $.connectionId;
        this.localAsn = $.localAsn;
        this.localIpAddress = $.localIpAddress;
        this.remoteAsn = $.remoteAsn;
        this.remoteIpAddress = $.remoteIpAddress;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpArgs $;

        public Builder() {
            $ = new BgpArgs();
        }

        public Builder(BgpArgs defaults) {
            $ = new BgpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authenticationKey shared key used for BGP peer authentication.
         * 
         * @return builder
         * 
         */
        public Builder authenticationKey(@Nullable Output<String> authenticationKey) {
            $.authenticationKey = authenticationKey;
            return this;
        }

        /**
         * @param authenticationKey shared key used for BGP peer authentication.
         * 
         * @return builder
         * 
         */
        public Builder authenticationKey(String authenticationKey) {
            return authenticationKey(Output.of(authenticationKey));
        }

        /**
         * @param connectionId identifier of a connection established between.
         * network device and remote service provider that will be used for peering.
         * 
         * @return builder
         * 
         */
        public Builder connectionId(Output<String> connectionId) {
            $.connectionId = connectionId;
            return this;
        }

        /**
         * @param connectionId identifier of a connection established between.
         * network device and remote service provider that will be used for peering.
         * 
         * @return builder
         * 
         */
        public Builder connectionId(String connectionId) {
            return connectionId(Output.of(connectionId));
        }

        /**
         * @param localAsn Local ASN number.
         * 
         * @return builder
         * 
         */
        public Builder localAsn(Output<Integer> localAsn) {
            $.localAsn = localAsn;
            return this;
        }

        /**
         * @param localAsn Local ASN number.
         * 
         * @return builder
         * 
         */
        public Builder localAsn(Integer localAsn) {
            return localAsn(Output.of(localAsn));
        }

        /**
         * @param localIpAddress IP address in CIDR format of a local device.
         * 
         * @return builder
         * 
         */
        public Builder localIpAddress(Output<String> localIpAddress) {
            $.localIpAddress = localIpAddress;
            return this;
        }

        /**
         * @param localIpAddress IP address in CIDR format of a local device.
         * 
         * @return builder
         * 
         */
        public Builder localIpAddress(String localIpAddress) {
            return localIpAddress(Output.of(localIpAddress));
        }

        /**
         * @param remoteAsn Remote ASN number.
         * 
         * @return builder
         * 
         */
        public Builder remoteAsn(Output<Integer> remoteAsn) {
            $.remoteAsn = remoteAsn;
            return this;
        }

        /**
         * @param remoteAsn Remote ASN number.
         * 
         * @return builder
         * 
         */
        public Builder remoteAsn(Integer remoteAsn) {
            return remoteAsn(Output.of(remoteAsn));
        }

        /**
         * @param remoteIpAddress IP address of remote peer.
         * 
         * @return builder
         * 
         */
        public Builder remoteIpAddress(Output<String> remoteIpAddress) {
            $.remoteIpAddress = remoteIpAddress;
            return this;
        }

        /**
         * @param remoteIpAddress IP address of remote peer.
         * 
         * @return builder
         * 
         */
        public Builder remoteIpAddress(String remoteIpAddress) {
            return remoteIpAddress(Output.of(remoteIpAddress));
        }

        public BgpArgs build() {
            $.connectionId = Objects.requireNonNull($.connectionId, "expected parameter 'connectionId' to be non-null");
            $.localAsn = Objects.requireNonNull($.localAsn, "expected parameter 'localAsn' to be non-null");
            $.localIpAddress = Objects.requireNonNull($.localIpAddress, "expected parameter 'localIpAddress' to be non-null");
            $.remoteAsn = Objects.requireNonNull($.remoteAsn, "expected parameter 'remoteAsn' to be non-null");
            $.remoteIpAddress = Objects.requireNonNull($.remoteIpAddress, "expected parameter 'remoteIpAddress' to be non-null");
            return $;
        }
    }

}