// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.equinix.fabric.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPortPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPortPlainArgs Empty = new GetPortPlainArgs();

    /**
     * Equinix-assigned port identifier
     * 
     */
    @Import(name="uuid", required=true)
    private String uuid;

    /**
     * @return Equinix-assigned port identifier
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    private GetPortPlainArgs() {}

    private GetPortPlainArgs(GetPortPlainArgs $) {
        this.uuid = $.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPortPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPortPlainArgs $;

        public Builder() {
            $ = new GetPortPlainArgs();
        }

        public Builder(GetPortPlainArgs defaults) {
            $ = new GetPortPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param uuid Equinix-assigned port identifier
         * 
         * @return builder
         * 
         */
        public Builder uuid(String uuid) {
            $.uuid = uuid;
            return this;
        }

        public GetPortPlainArgs build() {
            $.uuid = Objects.requireNonNull($.uuid, "expected parameter 'uuid' to be non-null");
            return $;
        }
    }

}