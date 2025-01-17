// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.metal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDevicesSort {
    /**
     * @return The attribute used to filter. Filter attributes are case-sensitive
     * 
     */
    private String attribute;
    private @Nullable String direction;

    private GetDevicesSort() {}
    /**
     * @return The attribute used to filter. Filter attributes are case-sensitive
     * 
     */
    public String attribute() {
        return this.attribute;
    }
    public Optional<String> direction() {
        return Optional.ofNullable(this.direction);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDevicesSort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String attribute;
        private @Nullable String direction;
        public Builder() {}
        public Builder(GetDevicesSort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attribute = defaults.attribute;
    	      this.direction = defaults.direction;
        }

        @CustomType.Setter
        public Builder attribute(String attribute) {
            this.attribute = Objects.requireNonNull(attribute);
            return this;
        }
        @CustomType.Setter
        public Builder direction(@Nullable String direction) {
            this.direction = direction;
            return this;
        }
        public GetDevicesSort build() {
            final var _resultValue = new GetDevicesSort();
            _resultValue.attribute = attribute;
            _resultValue.direction = direction;
            return _resultValue;
        }
    }
}
