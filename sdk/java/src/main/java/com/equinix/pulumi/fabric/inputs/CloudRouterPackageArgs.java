// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class CloudRouterPackageArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudRouterPackageArgs Empty = new CloudRouterPackageArgs();

    /**
     * Fabric Cloud Router package code
     * 
     */
    @Import(name="code", required=true)
    private Output<String> code;

    /**
     * @return Fabric Cloud Router package code
     * 
     */
    public Output<String> code() {
        return this.code;
    }

    private CloudRouterPackageArgs() {}

    private CloudRouterPackageArgs(CloudRouterPackageArgs $) {
        this.code = $.code;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudRouterPackageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudRouterPackageArgs $;

        public Builder() {
            $ = new CloudRouterPackageArgs();
        }

        public Builder(CloudRouterPackageArgs defaults) {
            $ = new CloudRouterPackageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param code Fabric Cloud Router package code
         * 
         * @return builder
         * 
         */
        public Builder code(Output<String> code) {
            $.code = code;
            return this;
        }

        /**
         * @param code Fabric Cloud Router package code
         * 
         * @return builder
         * 
         */
        public Builder code(String code) {
            return code(Output.of(code));
        }

        public CloudRouterPackageArgs build() {
            $.code = Objects.requireNonNull($.code, "expected parameter 'code' to be non-null");
            return $;
        }
    }

}
