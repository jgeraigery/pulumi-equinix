// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinix.metal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPlansPlan {
    private List<String> availableInMetros;
    private List<String> availableIns;
    private String class_;
    private List<String> deploymentTypes;
    private String description;
    /**
     * @return id of the plan
     * 
     */
    private String id;
    private Boolean legacy;
    private String line;
    /**
     * @return name of the plan
     * - `slug`- plan slug
     * - `description`- description of the plan
     * - `line`- plan line, e.g. baremetal
     * - `legacy`- flag showing if it&#39;s a legacy plan
     * - `class`- plan class
     * - `pricing_hour`- plan hourly price
     * - `pricing_month`- plan monthly price
     * - `deployment_types`- list of deployment types, e.g. on_demand, spot_market
     * - `available_in`- list of facilities where the plan is available
     * - `available_in_metros`- list of facilities where the plan is available
     * 
     */
    private String name;
    private Double pricingHour;
    private Double pricingMonth;
    private String slug;

    private GetPlansPlan() {}
    public List<String> availableInMetros() {
        return this.availableInMetros;
    }
    public List<String> availableIns() {
        return this.availableIns;
    }
    public String class_() {
        return this.class_;
    }
    public List<String> deploymentTypes() {
        return this.deploymentTypes;
    }
    public String description() {
        return this.description;
    }
    /**
     * @return id of the plan
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean legacy() {
        return this.legacy;
    }
    public String line() {
        return this.line;
    }
    /**
     * @return name of the plan
     * - `slug`- plan slug
     * - `description`- description of the plan
     * - `line`- plan line, e.g. baremetal
     * - `legacy`- flag showing if it&#39;s a legacy plan
     * - `class`- plan class
     * - `pricing_hour`- plan hourly price
     * - `pricing_month`- plan monthly price
     * - `deployment_types`- list of deployment types, e.g. on_demand, spot_market
     * - `available_in`- list of facilities where the plan is available
     * - `available_in_metros`- list of facilities where the plan is available
     * 
     */
    public String name() {
        return this.name;
    }
    public Double pricingHour() {
        return this.pricingHour;
    }
    public Double pricingMonth() {
        return this.pricingMonth;
    }
    public String slug() {
        return this.slug;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPlansPlan defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> availableInMetros;
        private List<String> availableIns;
        private String class_;
        private List<String> deploymentTypes;
        private String description;
        private String id;
        private Boolean legacy;
        private String line;
        private String name;
        private Double pricingHour;
        private Double pricingMonth;
        private String slug;
        public Builder() {}
        public Builder(GetPlansPlan defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availableInMetros = defaults.availableInMetros;
    	      this.availableIns = defaults.availableIns;
    	      this.class_ = defaults.class_;
    	      this.deploymentTypes = defaults.deploymentTypes;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.legacy = defaults.legacy;
    	      this.line = defaults.line;
    	      this.name = defaults.name;
    	      this.pricingHour = defaults.pricingHour;
    	      this.pricingMonth = defaults.pricingMonth;
    	      this.slug = defaults.slug;
        }

        @CustomType.Setter
        public Builder availableInMetros(List<String> availableInMetros) {
            this.availableInMetros = Objects.requireNonNull(availableInMetros);
            return this;
        }
        public Builder availableInMetros(String... availableInMetros) {
            return availableInMetros(List.of(availableInMetros));
        }
        @CustomType.Setter
        public Builder availableIns(List<String> availableIns) {
            this.availableIns = Objects.requireNonNull(availableIns);
            return this;
        }
        public Builder availableIns(String... availableIns) {
            return availableIns(List.of(availableIns));
        }
        @CustomType.Setter("class")
        public Builder class_(String class_) {
            this.class_ = Objects.requireNonNull(class_);
            return this;
        }
        @CustomType.Setter
        public Builder deploymentTypes(List<String> deploymentTypes) {
            this.deploymentTypes = Objects.requireNonNull(deploymentTypes);
            return this;
        }
        public Builder deploymentTypes(String... deploymentTypes) {
            return deploymentTypes(List.of(deploymentTypes));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder legacy(Boolean legacy) {
            this.legacy = Objects.requireNonNull(legacy);
            return this;
        }
        @CustomType.Setter
        public Builder line(String line) {
            this.line = Objects.requireNonNull(line);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder pricingHour(Double pricingHour) {
            this.pricingHour = Objects.requireNonNull(pricingHour);
            return this;
        }
        @CustomType.Setter
        public Builder pricingMonth(Double pricingMonth) {
            this.pricingMonth = Objects.requireNonNull(pricingMonth);
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }
        public GetPlansPlan build() {
            final var o = new GetPlansPlan();
            o.availableInMetros = availableInMetros;
            o.availableIns = availableIns;
            o.class_ = class_;
            o.deploymentTypes = deploymentTypes;
            o.description = description;
            o.id = id;
            o.legacy = legacy;
            o.line = line;
            o.name = name;
            o.pricingHour = pricingHour;
            o.pricingMonth = pricingMonth;
            o.slug = slug;
            return o;
        }
    }
}