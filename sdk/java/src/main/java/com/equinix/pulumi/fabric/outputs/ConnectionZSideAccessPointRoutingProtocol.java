// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConnectionZSideAccessPointRoutingProtocol {
    /**
     * @return Routing protocol instance state
     * 
     */
    private @Nullable String state;
    /**
     * @return Interface type
     * 
     */
    private @Nullable String type;
    /**
     * @return Equinix-assigned interface identifier
     * 
     */
    private @Nullable String uuid;

    private ConnectionZSideAccessPointRoutingProtocol() {}
    /**
     * @return Routing protocol instance state
     * 
     */
    public Optional<String> state() {
        return Optional.ofNullable(this.state);
    }
    /**
     * @return Interface type
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return Equinix-assigned interface identifier
     * 
     */
    public Optional<String> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectionZSideAccessPointRoutingProtocol defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String state;
        private @Nullable String type;
        private @Nullable String uuid;
        public Builder() {}
        public Builder(ConnectionZSideAccessPointRoutingProtocol defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.state = defaults.state;
    	      this.type = defaults.type;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder state(@Nullable String state) {
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder uuid(@Nullable String uuid) {
            this.uuid = uuid;
            return this;
        }
        public ConnectionZSideAccessPointRoutingProtocol build() {
            final var o = new ConnectionZSideAccessPointRoutingProtocol();
            o.state = state;
            o.type = type;
            o.uuid = uuid;
            return o;
        }
    }
}
