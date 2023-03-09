// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.equinix.fabric.outputs;

import com.equinix.pulumi.equinix.fabric.outputs.GetServiceProfilesDataMarketingInfoProcessStep;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetServiceProfilesDataMarketingInfo {
    private String logo;
    private List<GetServiceProfilesDataMarketingInfoProcessStep> processSteps;
    private Boolean promotion;

    private GetServiceProfilesDataMarketingInfo() {}
    public String logo() {
        return this.logo;
    }
    public List<GetServiceProfilesDataMarketingInfoProcessStep> processSteps() {
        return this.processSteps;
    }
    public Boolean promotion() {
        return this.promotion;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceProfilesDataMarketingInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String logo;
        private List<GetServiceProfilesDataMarketingInfoProcessStep> processSteps;
        private Boolean promotion;
        public Builder() {}
        public Builder(GetServiceProfilesDataMarketingInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.logo = defaults.logo;
    	      this.processSteps = defaults.processSteps;
    	      this.promotion = defaults.promotion;
        }

        @CustomType.Setter
        public Builder logo(String logo) {
            this.logo = Objects.requireNonNull(logo);
            return this;
        }
        @CustomType.Setter
        public Builder processSteps(List<GetServiceProfilesDataMarketingInfoProcessStep> processSteps) {
            this.processSteps = Objects.requireNonNull(processSteps);
            return this;
        }
        public Builder processSteps(GetServiceProfilesDataMarketingInfoProcessStep... processSteps) {
            return processSteps(List.of(processSteps));
        }
        @CustomType.Setter
        public Builder promotion(Boolean promotion) {
            this.promotion = Objects.requireNonNull(promotion);
            return this;
        }
        public GetServiceProfilesDataMarketingInfo build() {
            final var o = new GetServiceProfilesDataMarketingInfo();
            o.logo = logo;
            o.processSteps = processSteps;
            o.promotion = promotion;
            return o;
        }
    }
}