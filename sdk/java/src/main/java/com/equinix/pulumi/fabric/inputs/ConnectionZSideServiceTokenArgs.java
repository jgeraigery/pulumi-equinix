// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.equinix.pulumi.fabric.enums.ServiceTokenType;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionZSideServiceTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionZSideServiceTokenArgs Empty = new ConnectionZSideServiceTokenArgs();

    /**
     * User-provided service description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User-provided service description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

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
     * Interface type
     * 
     */
    @Import(name="type")
    private @Nullable Output<Either<String,ServiceTokenType>> type;

    /**
     * @return Interface type
     * 
     */
    public Optional<Output<Either<String,ServiceTokenType>>> type() {
        return Optional.ofNullable(this.type);
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

    private ConnectionZSideServiceTokenArgs() {}

    private ConnectionZSideServiceTokenArgs(ConnectionZSideServiceTokenArgs $) {
        this.description = $.description;
        this.href = $.href;
        this.type = $.type;
        this.uuid = $.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionZSideServiceTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionZSideServiceTokenArgs $;

        public Builder() {
            $ = new ConnectionZSideServiceTokenArgs();
        }

        public Builder(ConnectionZSideServiceTokenArgs defaults) {
            $ = new ConnectionZSideServiceTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description User-provided service description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User-provided service description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param type Interface type
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<Either<String,ServiceTokenType>> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Interface type
         * 
         * @return builder
         * 
         */
        public Builder type(Either<String,ServiceTokenType> type) {
            return type(Output.of(type));
        }

        /**
         * @param type Interface type
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Either.ofLeft(type));
        }

        /**
         * @param type Interface type
         * 
         * @return builder
         * 
         */
        public Builder type(ServiceTokenType type) {
            return type(Either.ofRight(type));
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

        public ConnectionZSideServiceTokenArgs build() {
            return $;
        }
    }

}
