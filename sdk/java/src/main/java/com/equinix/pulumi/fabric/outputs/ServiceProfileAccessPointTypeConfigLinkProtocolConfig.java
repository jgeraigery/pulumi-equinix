// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceProfileAccessPointTypeConfigLinkProtocolConfig {
    /**
     * @return Port Encapsulation
     * 
     */
    private @Nullable String encapsulation;
    /**
     * @return Encapsulation strategy
     * 
     */
    private @Nullable String encapsulationStrategy;
    /**
     * @return Reuse vlan sTag
     * 
     */
    private @Nullable Boolean reuseVlanSTag;

    private ServiceProfileAccessPointTypeConfigLinkProtocolConfig() {}
    /**
     * @return Port Encapsulation
     * 
     */
    public Optional<String> encapsulation() {
        return Optional.ofNullable(this.encapsulation);
    }
    /**
     * @return Encapsulation strategy
     * 
     */
    public Optional<String> encapsulationStrategy() {
        return Optional.ofNullable(this.encapsulationStrategy);
    }
    /**
     * @return Reuse vlan sTag
     * 
     */
    public Optional<Boolean> reuseVlanSTag() {
        return Optional.ofNullable(this.reuseVlanSTag);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceProfileAccessPointTypeConfigLinkProtocolConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String encapsulation;
        private @Nullable String encapsulationStrategy;
        private @Nullable Boolean reuseVlanSTag;
        public Builder() {}
        public Builder(ServiceProfileAccessPointTypeConfigLinkProtocolConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.encapsulation = defaults.encapsulation;
    	      this.encapsulationStrategy = defaults.encapsulationStrategy;
    	      this.reuseVlanSTag = defaults.reuseVlanSTag;
        }

        @CustomType.Setter
        public Builder encapsulation(@Nullable String encapsulation) {
            this.encapsulation = encapsulation;
            return this;
        }
        @CustomType.Setter
        public Builder encapsulationStrategy(@Nullable String encapsulationStrategy) {
            this.encapsulationStrategy = encapsulationStrategy;
            return this;
        }
        @CustomType.Setter
        public Builder reuseVlanSTag(@Nullable Boolean reuseVlanSTag) {
            this.reuseVlanSTag = reuseVlanSTag;
            return this;
        }
        public ServiceProfileAccessPointTypeConfigLinkProtocolConfig build() {
            final var _resultValue = new ServiceProfileAccessPointTypeConfigLinkProtocolConfig();
            _resultValue.encapsulation = encapsulation;
            _resultValue.encapsulationStrategy = encapsulationStrategy;
            _resultValue.reuseVlanSTag = reuseVlanSTag;
            return _resultValue;
        }
    }
}
