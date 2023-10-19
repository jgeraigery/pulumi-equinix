// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.metal.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDevicePlainArgs Empty = new GetDevicePlainArgs();

    /**
     * Device ID.
     * 
     * &gt; **NOTE:** You should pass either `device_id`, or both `project_id` and `hostname`.
     * 
     */
    @Import(name="deviceId")
    private @Nullable String deviceId;

    /**
     * @return Device ID.
     * 
     * &gt; **NOTE:** You should pass either `device_id`, or both `project_id` and `hostname`.
     * 
     */
    public Optional<String> deviceId() {
        return Optional.ofNullable(this.deviceId);
    }

    /**
     * The device name.
     * 
     */
    @Import(name="hostname")
    private @Nullable String hostname;

    /**
     * @return The device name.
     * 
     */
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * The id of the project in which the devices exists.
     * 
     */
    @Import(name="projectId")
    private @Nullable String projectId;

    /**
     * @return The id of the project in which the devices exists.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private GetDevicePlainArgs() {}

    private GetDevicePlainArgs(GetDevicePlainArgs $) {
        this.deviceId = $.deviceId;
        this.hostname = $.hostname;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDevicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDevicePlainArgs $;

        public Builder() {
            $ = new GetDevicePlainArgs();
        }

        public Builder(GetDevicePlainArgs defaults) {
            $ = new GetDevicePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deviceId Device ID.
         * 
         * &gt; **NOTE:** You should pass either `device_id`, or both `project_id` and `hostname`.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(@Nullable String deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param hostname The device name.
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable String hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param projectId The id of the project in which the devices exists.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetDevicePlainArgs build() {
            return $;
        }
    }

}
