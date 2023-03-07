// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinix.fabric.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinix.fabric.inputs.ServiceProfileMarketingInfoProcessStepArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceProfileMarketingInfoArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceProfileMarketingInfoArgs Empty = new ServiceProfileMarketingInfoArgs();

    /**
     * Logo
     * 
     */
    @Import(name="logo")
    private @Nullable Output<String> logo;

    /**
     * @return Logo
     * 
     */
    public Optional<Output<String>> logo() {
        return Optional.ofNullable(this.logo);
    }

    /**
     * Process Step
     * 
     */
    @Import(name="processSteps")
    private @Nullable Output<List<ServiceProfileMarketingInfoProcessStepArgs>> processSteps;

    /**
     * @return Process Step
     * 
     */
    public Optional<Output<List<ServiceProfileMarketingInfoProcessStepArgs>>> processSteps() {
        return Optional.ofNullable(this.processSteps);
    }

    /**
     * Promotion
     * 
     */
    @Import(name="promotion")
    private @Nullable Output<Boolean> promotion;

    /**
     * @return Promotion
     * 
     */
    public Optional<Output<Boolean>> promotion() {
        return Optional.ofNullable(this.promotion);
    }

    private ServiceProfileMarketingInfoArgs() {}

    private ServiceProfileMarketingInfoArgs(ServiceProfileMarketingInfoArgs $) {
        this.logo = $.logo;
        this.processSteps = $.processSteps;
        this.promotion = $.promotion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceProfileMarketingInfoArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceProfileMarketingInfoArgs $;

        public Builder() {
            $ = new ServiceProfileMarketingInfoArgs();
        }

        public Builder(ServiceProfileMarketingInfoArgs defaults) {
            $ = new ServiceProfileMarketingInfoArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param logo Logo
         * 
         * @return builder
         * 
         */
        public Builder logo(@Nullable Output<String> logo) {
            $.logo = logo;
            return this;
        }

        /**
         * @param logo Logo
         * 
         * @return builder
         * 
         */
        public Builder logo(String logo) {
            return logo(Output.of(logo));
        }

        /**
         * @param processSteps Process Step
         * 
         * @return builder
         * 
         */
        public Builder processSteps(@Nullable Output<List<ServiceProfileMarketingInfoProcessStepArgs>> processSteps) {
            $.processSteps = processSteps;
            return this;
        }

        /**
         * @param processSteps Process Step
         * 
         * @return builder
         * 
         */
        public Builder processSteps(List<ServiceProfileMarketingInfoProcessStepArgs> processSteps) {
            return processSteps(Output.of(processSteps));
        }

        /**
         * @param processSteps Process Step
         * 
         * @return builder
         * 
         */
        public Builder processSteps(ServiceProfileMarketingInfoProcessStepArgs... processSteps) {
            return processSteps(List.of(processSteps));
        }

        /**
         * @param promotion Promotion
         * 
         * @return builder
         * 
         */
        public Builder promotion(@Nullable Output<Boolean> promotion) {
            $.promotion = promotion;
            return this;
        }

        /**
         * @param promotion Promotion
         * 
         * @return builder
         * 
         */
        public Builder promotion(Boolean promotion) {
            return promotion(Output.of(promotion));
        }

        public ServiceProfileMarketingInfoArgs build() {
            return $;
        }
    }

}