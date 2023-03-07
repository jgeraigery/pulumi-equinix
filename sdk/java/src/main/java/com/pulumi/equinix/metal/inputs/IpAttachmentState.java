// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinix.metal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final IpAttachmentState Empty = new IpAttachmentState();

    @Import(name="address")
    private @Nullable Output<String> address;

    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Address family as integer. One of `4` or `6`.
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<Integer> addressFamily;

    /**
     * @return Address family as integer. One of `4` or `6`.
     * 
     */
    public Optional<Output<Integer>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * Length of CIDR prefix of the subnet as integer.
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<Integer> cidr;

    /**
     * @return Length of CIDR prefix of the subnet as integer.
     * 
     */
    public Optional<Output<Integer>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * CIDR notation of subnet from block reserved in the same project
     * and facility as the device.
     * 
     */
    @Import(name="cidrNotation")
    private @Nullable Output<String> cidrNotation;

    /**
     * @return CIDR notation of subnet from block reserved in the same project
     * and facility as the device.
     * 
     */
    public Optional<Output<String>> cidrNotation() {
        return Optional.ofNullable(this.cidrNotation);
    }

    /**
     * ID of device to which to assign the subnet.
     * 
     */
    @Import(name="deviceId")
    private @Nullable Output<String> deviceId;

    /**
     * @return ID of device to which to assign the subnet.
     * 
     */
    public Optional<Output<String>> deviceId() {
        return Optional.ofNullable(this.deviceId);
    }

    /**
     * IP address of gateway for the subnet.
     * 
     */
    @Import(name="gateway")
    private @Nullable Output<String> gateway;

    /**
     * @return IP address of gateway for the subnet.
     * 
     */
    public Optional<Output<String>> gateway() {
        return Optional.ofNullable(this.gateway);
    }

    /**
     * Flag indicating whether IP block is global, i.e. assignable in any location
     * 
     */
    @Import(name="global")
    private @Nullable Output<Boolean> global;

    /**
     * @return Flag indicating whether IP block is global, i.e. assignable in any location
     * 
     */
    public Optional<Output<Boolean>> global() {
        return Optional.ofNullable(this.global);
    }

    @Import(name="manageable")
    private @Nullable Output<Boolean> manageable;

    public Optional<Output<Boolean>> manageable() {
        return Optional.ofNullable(this.manageable);
    }

    @Import(name="management")
    private @Nullable Output<Boolean> management;

    public Optional<Output<Boolean>> management() {
        return Optional.ofNullable(this.management);
    }

    /**
     * Subnet mask in decimal notation, e.g., `255.255.255.0`.
     * 
     */
    @Import(name="netmask")
    private @Nullable Output<String> netmask;

    /**
     * @return Subnet mask in decimal notation, e.g., `255.255.255.0`.
     * 
     */
    public Optional<Output<String>> netmask() {
        return Optional.ofNullable(this.netmask);
    }

    /**
     * Subnet network address.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return Subnet network address.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * Boolean flag whether subnet is reachable from the Internet.
     * 
     */
    @Import(name="public")
    private @Nullable Output<Boolean> public_;

    /**
     * @return Boolean flag whether subnet is reachable from the Internet.
     * 
     */
    public Optional<Output<Boolean>> public_() {
        return Optional.ofNullable(this.public_);
    }

    @Import(name="vrfId")
    private @Nullable Output<String> vrfId;

    public Optional<Output<String>> vrfId() {
        return Optional.ofNullable(this.vrfId);
    }

    private IpAttachmentState() {}

    private IpAttachmentState(IpAttachmentState $) {
        this.address = $.address;
        this.addressFamily = $.addressFamily;
        this.cidr = $.cidr;
        this.cidrNotation = $.cidrNotation;
        this.deviceId = $.deviceId;
        this.gateway = $.gateway;
        this.global = $.global;
        this.manageable = $.manageable;
        this.management = $.management;
        this.netmask = $.netmask;
        this.network = $.network;
        this.public_ = $.public_;
        this.vrfId = $.vrfId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpAttachmentState $;

        public Builder() {
            $ = new IpAttachmentState();
        }

        public Builder(IpAttachmentState defaults) {
            $ = new IpAttachmentState(Objects.requireNonNull(defaults));
        }

        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param addressFamily Address family as integer. One of `4` or `6`.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<Integer> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address family as integer. One of `4` or `6`.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(Integer addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param cidr Length of CIDR prefix of the subnet as integer.
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<Integer> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr Length of CIDR prefix of the subnet as integer.
         * 
         * @return builder
         * 
         */
        public Builder cidr(Integer cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param cidrNotation CIDR notation of subnet from block reserved in the same project
         * and facility as the device.
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(@Nullable Output<String> cidrNotation) {
            $.cidrNotation = cidrNotation;
            return this;
        }

        /**
         * @param cidrNotation CIDR notation of subnet from block reserved in the same project
         * and facility as the device.
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(String cidrNotation) {
            return cidrNotation(Output.of(cidrNotation));
        }

        /**
         * @param deviceId ID of device to which to assign the subnet.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(@Nullable Output<String> deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param deviceId ID of device to which to assign the subnet.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(String deviceId) {
            return deviceId(Output.of(deviceId));
        }

        /**
         * @param gateway IP address of gateway for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder gateway(@Nullable Output<String> gateway) {
            $.gateway = gateway;
            return this;
        }

        /**
         * @param gateway IP address of gateway for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder gateway(String gateway) {
            return gateway(Output.of(gateway));
        }

        /**
         * @param global Flag indicating whether IP block is global, i.e. assignable in any location
         * 
         * @return builder
         * 
         */
        public Builder global(@Nullable Output<Boolean> global) {
            $.global = global;
            return this;
        }

        /**
         * @param global Flag indicating whether IP block is global, i.e. assignable in any location
         * 
         * @return builder
         * 
         */
        public Builder global(Boolean global) {
            return global(Output.of(global));
        }

        public Builder manageable(@Nullable Output<Boolean> manageable) {
            $.manageable = manageable;
            return this;
        }

        public Builder manageable(Boolean manageable) {
            return manageable(Output.of(manageable));
        }

        public Builder management(@Nullable Output<Boolean> management) {
            $.management = management;
            return this;
        }

        public Builder management(Boolean management) {
            return management(Output.of(management));
        }

        /**
         * @param netmask Subnet mask in decimal notation, e.g., `255.255.255.0`.
         * 
         * @return builder
         * 
         */
        public Builder netmask(@Nullable Output<String> netmask) {
            $.netmask = netmask;
            return this;
        }

        /**
         * @param netmask Subnet mask in decimal notation, e.g., `255.255.255.0`.
         * 
         * @return builder
         * 
         */
        public Builder netmask(String netmask) {
            return netmask(Output.of(netmask));
        }

        /**
         * @param network Subnet network address.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network Subnet network address.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param public_ Boolean flag whether subnet is reachable from the Internet.
         * 
         * @return builder
         * 
         */
        public Builder public_(@Nullable Output<Boolean> public_) {
            $.public_ = public_;
            return this;
        }

        /**
         * @param public_ Boolean flag whether subnet is reachable from the Internet.
         * 
         * @return builder
         * 
         */
        public Builder public_(Boolean public_) {
            return public_(Output.of(public_));
        }

        public Builder vrfId(@Nullable Output<String> vrfId) {
            $.vrfId = vrfId;
            return this;
        }

        public Builder vrfId(String vrfId) {
            return vrfId(Output.of(vrfId));
        }

        public IpAttachmentState build() {
            return $;
        }
    }

}