// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.metal.inputs;

import com.equinix.pulumi.metal.enums.Facility;
import com.equinix.pulumi.metal.enums.IpBlockType;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReservedIpBlockState extends com.pulumi.resources.ResourceArgs {

    public static final ReservedIpBlockState Empty = new ReservedIpBlockState();

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
     * Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<Integer> cidr;

    /**
     * @return Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
     * 
     */
    public Optional<Output<Integer>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
     * 
     */
    @Import(name="cidrNotation")
    private @Nullable Output<String> cidrNotation;

    /**
     * @return Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
     * 
     */
    public Optional<Output<String>> cidrNotation() {
        return Optional.ofNullable(this.cidrNotation);
    }

    /**
     * Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
     * be helpful for self-managed IPAM. The object must be valid JSON.
     * 
     */
    @Import(name="customData")
    private @Nullable Output<String> customData;

    /**
     * @return Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
     * be helpful for self-managed IPAM. The object must be valid JSON.
     * 
     */
    public Optional<Output<String>> customData() {
        return Optional.ofNullable(this.customData);
    }

    /**
     * Arbitrary description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Arbitrary description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Facility where to allocate the public IP address block, makes sense only
     * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
     * 
     */
    @Import(name="facility")
    private @Nullable Output<Either<String,Facility>> facility;

    /**
     * @return Facility where to allocate the public IP address block, makes sense only
     * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
     * 
     */
    public Optional<Output<Either<String,Facility>>> facility() {
        return Optional.ofNullable(this.facility);
    }

    @Import(name="gateway")
    private @Nullable Output<String> gateway;

    public Optional<Output<String>> gateway() {
        return Optional.ofNullable(this.gateway);
    }

    /**
     * Boolean flag whether addresses from a block are global (i.e. can be assigned in any
     * facility).
     * 
     */
    @Import(name="global")
    private @Nullable Output<Boolean> global;

    /**
     * @return Boolean flag whether addresses from a block are global (i.e. can be assigned in any
     * facility).
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
     * Metro where to allocate the public IP address block, makes sense only
     * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
     * 
     */
    @Import(name="metro")
    private @Nullable Output<String> metro;

    /**
     * @return Metro where to allocate the public IP address block, makes sense only
     * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
     * 
     */
    public Optional<Output<String>> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * Mask in decimal notation, e.g. `255.255.255.0`.
     * 
     */
    @Import(name="netmask")
    private @Nullable Output<String> netmask;

    /**
     * @return Mask in decimal notation, e.g. `255.255.255.0`.
     * 
     */
    public Optional<Output<String>> netmask() {
        return Optional.ofNullable(this.netmask);
    }

    /**
     * Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The metal project ID where to allocate the address block.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The metal project ID where to allocate the address block.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Boolean flag whether addresses from a block are public.
     * 
     */
    @Import(name="public")
    private @Nullable Output<Boolean> public_;

    /**
     * @return Boolean flag whether addresses from a block are public.
     * 
     */
    public Optional<Output<Boolean>> public_() {
        return Optional.ofNullable(this.public_);
    }

    /**
     * The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<Integer> quantity;

    /**
     * @return The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
     * 
     */
    public Optional<Output<Integer>> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    /**
     * String list of tags.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return String list of tags.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
     * compatibility.
     * 
     */
    @Import(name="type")
    private @Nullable Output<Either<String,IpBlockType>> type;

    /**
     * @return One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
     * compatibility.
     * 
     */
    public Optional<Output<Either<String,IpBlockType>>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
     * 
     */
    @Import(name="vrfId")
    private @Nullable Output<String> vrfId;

    /**
     * @return Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
     * 
     */
    public Optional<Output<String>> vrfId() {
        return Optional.ofNullable(this.vrfId);
    }

    /**
     * Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
     * 
     */
    @Import(name="waitForState")
    private @Nullable Output<String> waitForState;

    /**
     * @return Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
     * 
     */
    public Optional<Output<String>> waitForState() {
        return Optional.ofNullable(this.waitForState);
    }

    private ReservedIpBlockState() {}

    private ReservedIpBlockState(ReservedIpBlockState $) {
        this.address = $.address;
        this.addressFamily = $.addressFamily;
        this.cidr = $.cidr;
        this.cidrNotation = $.cidrNotation;
        this.customData = $.customData;
        this.description = $.description;
        this.facility = $.facility;
        this.gateway = $.gateway;
        this.global = $.global;
        this.manageable = $.manageable;
        this.management = $.management;
        this.metro = $.metro;
        this.netmask = $.netmask;
        this.network = $.network;
        this.projectId = $.projectId;
        this.public_ = $.public_;
        this.quantity = $.quantity;
        this.tags = $.tags;
        this.type = $.type;
        this.vrfId = $.vrfId;
        this.waitForState = $.waitForState;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReservedIpBlockState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReservedIpBlockState $;

        public Builder() {
            $ = new ReservedIpBlockState();
        }

        public Builder(ReservedIpBlockState defaults) {
            $ = new ReservedIpBlockState(Objects.requireNonNull(defaults));
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
         * @param cidr Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<Integer> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr Only valid as an argument and required when `type` is `vrf`. The size of the network to reserve from an existing VRF ip_range. `cidr` can only be specified with `vrf_id`. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range.
         * 
         * @return builder
         * 
         */
        public Builder cidr(Integer cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param cidrNotation Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(@Nullable Output<String> cidrNotation) {
            $.cidrNotation = cidrNotation;
            return this;
        }

        /**
         * @param cidrNotation Address and mask in CIDR notation, e.g. `147.229.15.30/31`.
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(String cidrNotation) {
            return cidrNotation(Output.of(cidrNotation));
        }

        /**
         * @param customData Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
         * be helpful for self-managed IPAM. The object must be valid JSON.
         * 
         * @return builder
         * 
         */
        public Builder customData(@Nullable Output<String> customData) {
            $.customData = customData;
            return this;
        }

        /**
         * @param customData Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may
         * be helpful for self-managed IPAM. The object must be valid JSON.
         * 
         * @return builder
         * 
         */
        public Builder customData(String customData) {
            return customData(Output.of(customData));
        }

        /**
         * @param description Arbitrary description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Arbitrary description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
         * 
         * @return builder
         * 
         */
        public Builder facility(@Nullable Output<Either<String,Facility>> facility) {
            $.facility = facility;
            return this;
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
         * 
         * @return builder
         * 
         */
        public Builder facility(Either<String,Facility> facility) {
            return facility(Output.of(facility));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
         * 
         * @return builder
         * 
         */
        public Builder facility(String facility) {
            return facility(Either.ofLeft(facility));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `metro`.
         * 
         * @return builder
         * 
         */
        public Builder facility(Facility facility) {
            return facility(Either.ofRight(facility));
        }

        public Builder gateway(@Nullable Output<String> gateway) {
            $.gateway = gateway;
            return this;
        }

        public Builder gateway(String gateway) {
            return gateway(Output.of(gateway));
        }

        /**
         * @param global Boolean flag whether addresses from a block are global (i.e. can be assigned in any
         * facility).
         * 
         * @return builder
         * 
         */
        public Builder global(@Nullable Output<Boolean> global) {
            $.global = global;
            return this;
        }

        /**
         * @param global Boolean flag whether addresses from a block are global (i.e. can be assigned in any
         * facility).
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
         * @param metro Metro where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable Output<String> metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param metro Metro where to allocate the public IP address block, makes sense only
         * if type is `public_ipv4` and must be empty if type is `global_ipv4`. Conflicts with `facility`.
         * 
         * @return builder
         * 
         */
        public Builder metro(String metro) {
            return metro(Output.of(metro));
        }

        /**
         * @param netmask Mask in decimal notation, e.g. `255.255.255.0`.
         * 
         * @return builder
         * 
         */
        public Builder netmask(@Nullable Output<String> netmask) {
            $.netmask = netmask;
            return this;
        }

        /**
         * @param netmask Mask in decimal notation, e.g. `255.255.255.0`.
         * 
         * @return builder
         * 
         */
        public Builder netmask(String netmask) {
            return netmask(Output.of(netmask));
        }

        /**
         * @param network Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network Only valid as an argument and required when `type` is `vrf`. An unreserved network address from an existing `ip_range` in the specified VRF.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param projectId The metal project ID where to allocate the address block.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The metal project ID where to allocate the address block.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param public_ Boolean flag whether addresses from a block are public.
         * 
         * @return builder
         * 
         */
        public Builder public_(@Nullable Output<Boolean> public_) {
            $.public_ = public_;
            return this;
        }

        /**
         * @param public_ Boolean flag whether addresses from a block are public.
         * 
         * @return builder
         * 
         */
        public Builder public_(Boolean public_) {
            return public_(Output.of(public_));
        }

        /**
         * @param quantity The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<Integer> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity The number of allocated `/32` addresses, a power of 2. Required when `type` is not `vrf`.
         * 
         * @return builder
         * 
         */
        public Builder quantity(Integer quantity) {
            return quantity(Output.of(quantity));
        }

        /**
         * @param tags String list of tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags String list of tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags String list of tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
         * compatibility.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<Either<String,IpBlockType>> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
         * compatibility.
         * 
         * @return builder
         * 
         */
        public Builder type(Either<String,IpBlockType> type) {
            return type(Output.of(type));
        }

        /**
         * @param type One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
         * compatibility.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Either.ofLeft(type));
        }

        /**
         * @param type One of `global_ipv4`, `public_ipv4`, or `vrf`. Defaults to `public_ipv4` for backward
         * compatibility.
         * 
         * @return builder
         * 
         */
        public Builder type(IpBlockType type) {
            return type(Either.ofRight(type));
        }

        /**
         * @param vrfId Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
         * 
         * @return builder
         * 
         */
        public Builder vrfId(@Nullable Output<String> vrfId) {
            $.vrfId = vrfId;
            return this;
        }

        /**
         * @param vrfId Only valid and required when `type` is `vrf`. VRF ID for type=vrf reservations.
         * 
         * @return builder
         * 
         */
        public Builder vrfId(String vrfId) {
            return vrfId(Output.of(vrfId));
        }

        /**
         * @param waitForState Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
         * 
         * @return builder
         * 
         */
        public Builder waitForState(@Nullable Output<String> waitForState) {
            $.waitForState = waitForState;
            return this;
        }

        /**
         * @param waitForState Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
         * 
         * @return builder
         * 
         */
        public Builder waitForState(String waitForState) {
            return waitForState(Output.of(waitForState));
        }

        public ReservedIpBlockState build() {
            return $;
        }
    }

}