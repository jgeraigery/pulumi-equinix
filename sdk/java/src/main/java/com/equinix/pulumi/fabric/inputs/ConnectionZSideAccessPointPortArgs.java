// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.equinix.pulumi.fabric.inputs.ConnectionZSideAccessPointPortRedundancyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionZSideAccessPointPortArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionZSideAccessPointPortArgs Empty = new ConnectionZSideAccessPointPortArgs();

    /**
     * Unique Resource Identifier
     * 
     */
    @Import(name="href")
    private @Nullable Output<String> href;

    /**
     * @return Unique Resource Identifier
     * 
     */
    public Optional<Output<String>> href() {
        return Optional.ofNullable(this.href);
    }

    /**
     * Port name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Port name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Redundancy Information
     * 
     */
    @Import(name="redundancies")
    private @Nullable Output<List<ConnectionZSideAccessPointPortRedundancyArgs>> redundancies;

    /**
     * @return Redundancy Information
     * 
     */
    public Optional<Output<List<ConnectionZSideAccessPointPortRedundancyArgs>>> redundancies() {
        return Optional.ofNullable(this.redundancies);
    }

    /**
     * Equinix-assigned interface identifier
     * 
     */
    @Import(name="uuid")
    private @Nullable Output<String> uuid;

    /**
     * @return Equinix-assigned interface identifier
     * 
     */
    public Optional<Output<String>> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    private ConnectionZSideAccessPointPortArgs() {}

    private ConnectionZSideAccessPointPortArgs(ConnectionZSideAccessPointPortArgs $) {
        this.href = $.href;
        this.name = $.name;
        this.redundancies = $.redundancies;
        this.uuid = $.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionZSideAccessPointPortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionZSideAccessPointPortArgs $;

        public Builder() {
            $ = new ConnectionZSideAccessPointPortArgs();
        }

        public Builder(ConnectionZSideAccessPointPortArgs defaults) {
            $ = new ConnectionZSideAccessPointPortArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param href Unique Resource Identifier
         * 
         * @return builder
         * 
         */
        public Builder href(@Nullable Output<String> href) {
            $.href = href;
            return this;
        }

        /**
         * @param href Unique Resource Identifier
         * 
         * @return builder
         * 
         */
        public Builder href(String href) {
            return href(Output.of(href));
        }

        /**
         * @param name Port name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Port name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param redundancies Redundancy Information
         * 
         * @return builder
         * 
         */
        public Builder redundancies(@Nullable Output<List<ConnectionZSideAccessPointPortRedundancyArgs>> redundancies) {
            $.redundancies = redundancies;
            return this;
        }

        /**
         * @param redundancies Redundancy Information
         * 
         * @return builder
         * 
         */
        public Builder redundancies(List<ConnectionZSideAccessPointPortRedundancyArgs> redundancies) {
            return redundancies(Output.of(redundancies));
        }

        /**
         * @param redundancies Redundancy Information
         * 
         * @return builder
         * 
         */
        public Builder redundancies(ConnectionZSideAccessPointPortRedundancyArgs... redundancies) {
            return redundancies(List.of(redundancies));
        }

        /**
         * @param uuid Equinix-assigned interface identifier
         * 
         * @return builder
         * 
         */
        public Builder uuid(@Nullable Output<String> uuid) {
            $.uuid = uuid;
            return this;
        }

        /**
         * @param uuid Equinix-assigned interface identifier
         * 
         * @return builder
         * 
         */
        public Builder uuid(String uuid) {
            return uuid(Output.of(uuid));
        }

        public ConnectionZSideAccessPointPortArgs build() {
            return $;
        }
    }

}
