// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.metal.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpBlockRangesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpBlockRangesPlainArgs Empty = new GetIpBlockRangesPlainArgs();

    /**
     * Facility code filtering the IP blocks. Global IPv4 blocks will be listed
     * anyway. If you omit this and metro, all the block from the project will be listed.   Use metro instead; read the facility to metro migration guide
     * 
     * @deprecated
     * Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
     * 
     */
    @Deprecated /* Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices */
    @Import(name="facility")
    private @Nullable String facility;

    /**
     * @return Facility code filtering the IP blocks. Global IPv4 blocks will be listed
     * anyway. If you omit this and metro, all the block from the project will be listed.   Use metro instead; read the facility to metro migration guide
     * 
     * @deprecated
     * Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
     * 
     */
    @Deprecated /* Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices */
    public Optional<String> facility() {
        return Optional.ofNullable(this.facility);
    }

    /**
     * Metro code filtering the IP blocks. Global IPv4 blocks will be listed
     * anyway. If you omit this and facility, all the block from the project will be listed.
     * 
     */
    @Import(name="metro")
    private @Nullable String metro;

    /**
     * @return Metro code filtering the IP blocks. Global IPv4 blocks will be listed
     * anyway. If you omit this and facility, all the block from the project will be listed.
     * 
     */
    public Optional<String> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * ID of the project from which to list the blocks.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return ID of the project from which to list the blocks.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetIpBlockRangesPlainArgs() {}

    private GetIpBlockRangesPlainArgs(GetIpBlockRangesPlainArgs $) {
        this.facility = $.facility;
        this.metro = $.metro;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpBlockRangesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpBlockRangesPlainArgs $;

        public Builder() {
            $ = new GetIpBlockRangesPlainArgs();
        }

        public Builder(GetIpBlockRangesPlainArgs defaults) {
            $ = new GetIpBlockRangesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param facility Facility code filtering the IP blocks. Global IPv4 blocks will be listed
         * anyway. If you omit this and metro, all the block from the project will be listed.   Use metro instead; read the facility to metro migration guide
         * 
         * @return builder
         * 
         * @deprecated
         * Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices
         * 
         */
        @Deprecated /* Use metro instead of facility.  For more information, read the migration guide: https://registry.terraform.io/providers/equinix/equinix/latest/docs/guides/migration_guide_facilities_to_metros_devices */
        public Builder facility(@Nullable String facility) {
            $.facility = facility;
            return this;
        }

        /**
         * @param metro Metro code filtering the IP blocks. Global IPv4 blocks will be listed
         * anyway. If you omit this and facility, all the block from the project will be listed.
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable String metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param projectId ID of the project from which to list the blocks.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetIpBlockRangesPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
