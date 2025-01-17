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


public final class GetRoutingProtocolBfdArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetRoutingProtocolBfdArgs Empty = new GetRoutingProtocolBfdArgs();

    /**
     * Bidirectional Forwarding Detection enablement
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Bidirectional Forwarding Detection enablement
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Interval range between the received BFD control packets
     * 
     */
    @Import(name="interval")
    private @Nullable Output<String> interval;

    /**
     * @return Interval range between the received BFD control packets
     * 
     */
    public Optional<Output<String>> interval() {
        return Optional.ofNullable(this.interval);
    }

    private GetRoutingProtocolBfdArgs() {}

    private GetRoutingProtocolBfdArgs(GetRoutingProtocolBfdArgs $) {
        this.enabled = $.enabled;
        this.interval = $.interval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRoutingProtocolBfdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRoutingProtocolBfdArgs $;

        public Builder() {
            $ = new GetRoutingProtocolBfdArgs();
        }

        public Builder(GetRoutingProtocolBfdArgs defaults) {
            $ = new GetRoutingProtocolBfdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Bidirectional Forwarding Detection enablement
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Bidirectional Forwarding Detection enablement
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param interval Interval range between the received BFD control packets
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<String> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval Interval range between the received BFD control packets
         * 
         * @return builder
         * 
         */
        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        public GetRoutingProtocolBfdArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            return $;
        }
    }

}
