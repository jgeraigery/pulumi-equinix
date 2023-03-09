// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.equinix.metal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetConnectionServiceToken {
    /**
     * @return Expiration date of the service token.
     * 
     */
    private String expiresAt;
    /**
     * @return Port UUID.
     * 
     */
    private String id;
    /**
     * @return Maximum allowed speed for the service token, string like in the `speed` attribute.
     * 
     */
    private String maxAllowedSpeed;
    /**
     * @return Port role - primary or secondary.
     * 
     */
    private String role;
    private String state;
    /**
     * @return Token type, `a_side` or `z_side`.
     * 
     */
    private String type;

    private GetConnectionServiceToken() {}
    /**
     * @return Expiration date of the service token.
     * 
     */
    public String expiresAt() {
        return this.expiresAt;
    }
    /**
     * @return Port UUID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Maximum allowed speed for the service token, string like in the `speed` attribute.
     * 
     */
    public String maxAllowedSpeed() {
        return this.maxAllowedSpeed;
    }
    /**
     * @return Port role - primary or secondary.
     * 
     */
    public String role() {
        return this.role;
    }
    public String state() {
        return this.state;
    }
    /**
     * @return Token type, `a_side` or `z_side`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionServiceToken defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String expiresAt;
        private String id;
        private String maxAllowedSpeed;
        private String role;
        private String state;
        private String type;
        public Builder() {}
        public Builder(GetConnectionServiceToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expiresAt = defaults.expiresAt;
    	      this.id = defaults.id;
    	      this.maxAllowedSpeed = defaults.maxAllowedSpeed;
    	      this.role = defaults.role;
    	      this.state = defaults.state;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder expiresAt(String expiresAt) {
            this.expiresAt = Objects.requireNonNull(expiresAt);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder maxAllowedSpeed(String maxAllowedSpeed) {
            this.maxAllowedSpeed = Objects.requireNonNull(maxAllowedSpeed);
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetConnectionServiceToken build() {
            final var o = new GetConnectionServiceToken();
            o.expiresAt = expiresAt;
            o.id = id;
            o.maxAllowedSpeed = maxAllowedSpeed;
            o.role = role;
            o.state = state;
            o.type = type;
            return o;
        }
    }
}