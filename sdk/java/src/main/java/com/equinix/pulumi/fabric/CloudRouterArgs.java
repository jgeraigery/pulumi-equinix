// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric;

import com.equinix.pulumi.fabric.inputs.CloudRouterAccountArgs;
import com.equinix.pulumi.fabric.inputs.CloudRouterLocationArgs;
import com.equinix.pulumi.fabric.inputs.CloudRouterNotificationArgs;
import com.equinix.pulumi.fabric.inputs.CloudRouterOrderArgs;
import com.equinix.pulumi.fabric.inputs.CloudRouterPackageArgs;
import com.equinix.pulumi.fabric.inputs.CloudRouterProjectArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudRouterArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudRouterArgs Empty = new CloudRouterArgs();

    /**
     * Customer account information that is associated with this Fabric Cloud Router
     * 
     */
    @Import(name="account")
    private @Nullable Output<CloudRouterAccountArgs> account;

    /**
     * @return Customer account information that is associated with this Fabric Cloud Router
     * 
     */
    public Optional<Output<CloudRouterAccountArgs>> account() {
        return Optional.ofNullable(this.account);
    }

    /**
     * Customer-provided Fabric Cloud Router description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Customer-provided Fabric Cloud Router description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Fabric Cloud Router location
     * 
     */
    @Import(name="location", required=true)
    private Output<CloudRouterLocationArgs> location;

    /**
     * @return Fabric Cloud Router location
     * 
     */
    public Output<CloudRouterLocationArgs> location() {
        return this.location;
    }

    /**
     * Fabric Cloud Router name. An alpha-numeric 24 characters string which can include only hyphens and underscores
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Fabric Cloud Router name. An alpha-numeric 24 characters string which can include only hyphens and underscores
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Preferences for notifications on Fabric Cloud Router configuration or status changes
     * 
     */
    @Import(name="notifications", required=true)
    private Output<List<CloudRouterNotificationArgs>> notifications;

    /**
     * @return Preferences for notifications on Fabric Cloud Router configuration or status changes
     * 
     */
    public Output<List<CloudRouterNotificationArgs>> notifications() {
        return this.notifications;
    }

    /**
     * Order information related to this Fabric Cloud Router
     * 
     */
    @Import(name="order")
    private @Nullable Output<CloudRouterOrderArgs> order;

    /**
     * @return Order information related to this Fabric Cloud Router
     * 
     */
    public Optional<Output<CloudRouterOrderArgs>> order() {
        return Optional.ofNullable(this.order);
    }

    /**
     * Fabric Cloud Router package
     * 
     */
    @Import(name="package", required=true)
    private Output<CloudRouterPackageArgs> package_;

    /**
     * @return Fabric Cloud Router package
     * 
     */
    public Output<CloudRouterPackageArgs> package_() {
        return this.package_;
    }

    /**
     * Fabric Cloud Router project
     * 
     */
    @Import(name="project")
    private @Nullable Output<CloudRouterProjectArgs> project;

    /**
     * @return Fabric Cloud Router project
     * 
     */
    public Optional<Output<CloudRouterProjectArgs>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Notification Type - ALL,CONNECTION*APPROVAL,SALES*REP_NOTIFICATIONS, NOTIFICATIONS
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Notification Type - ALL,CONNECTION*APPROVAL,SALES*REP_NOTIFICATIONS, NOTIFICATIONS
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private CloudRouterArgs() {}

    private CloudRouterArgs(CloudRouterArgs $) {
        this.account = $.account;
        this.description = $.description;
        this.location = $.location;
        this.name = $.name;
        this.notifications = $.notifications;
        this.order = $.order;
        this.package_ = $.package_;
        this.project = $.project;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudRouterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudRouterArgs $;

        public Builder() {
            $ = new CloudRouterArgs();
        }

        public Builder(CloudRouterArgs defaults) {
            $ = new CloudRouterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param account Customer account information that is associated with this Fabric Cloud Router
         * 
         * @return builder
         * 
         */
        public Builder account(@Nullable Output<CloudRouterAccountArgs> account) {
            $.account = account;
            return this;
        }

        /**
         * @param account Customer account information that is associated with this Fabric Cloud Router
         * 
         * @return builder
         * 
         */
        public Builder account(CloudRouterAccountArgs account) {
            return account(Output.of(account));
        }

        /**
         * @param description Customer-provided Fabric Cloud Router description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Customer-provided Fabric Cloud Router description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param location Fabric Cloud Router location
         * 
         * @return builder
         * 
         */
        public Builder location(Output<CloudRouterLocationArgs> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Fabric Cloud Router location
         * 
         * @return builder
         * 
         */
        public Builder location(CloudRouterLocationArgs location) {
            return location(Output.of(location));
        }

        /**
         * @param name Fabric Cloud Router name. An alpha-numeric 24 characters string which can include only hyphens and underscores
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Fabric Cloud Router name. An alpha-numeric 24 characters string which can include only hyphens and underscores
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notifications Preferences for notifications on Fabric Cloud Router configuration or status changes
         * 
         * @return builder
         * 
         */
        public Builder notifications(Output<List<CloudRouterNotificationArgs>> notifications) {
            $.notifications = notifications;
            return this;
        }

        /**
         * @param notifications Preferences for notifications on Fabric Cloud Router configuration or status changes
         * 
         * @return builder
         * 
         */
        public Builder notifications(List<CloudRouterNotificationArgs> notifications) {
            return notifications(Output.of(notifications));
        }

        /**
         * @param notifications Preferences for notifications on Fabric Cloud Router configuration or status changes
         * 
         * @return builder
         * 
         */
        public Builder notifications(CloudRouterNotificationArgs... notifications) {
            return notifications(List.of(notifications));
        }

        /**
         * @param order Order information related to this Fabric Cloud Router
         * 
         * @return builder
         * 
         */
        public Builder order(@Nullable Output<CloudRouterOrderArgs> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order Order information related to this Fabric Cloud Router
         * 
         * @return builder
         * 
         */
        public Builder order(CloudRouterOrderArgs order) {
            return order(Output.of(order));
        }

        /**
         * @param package_ Fabric Cloud Router package
         * 
         * @return builder
         * 
         */
        public Builder package_(Output<CloudRouterPackageArgs> package_) {
            $.package_ = package_;
            return this;
        }

        /**
         * @param package_ Fabric Cloud Router package
         * 
         * @return builder
         * 
         */
        public Builder package_(CloudRouterPackageArgs package_) {
            return package_(Output.of(package_));
        }

        /**
         * @param project Fabric Cloud Router project
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<CloudRouterProjectArgs> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project Fabric Cloud Router project
         * 
         * @return builder
         * 
         */
        public Builder project(CloudRouterProjectArgs project) {
            return project(Output.of(project));
        }

        /**
         * @param type Notification Type - ALL,CONNECTION*APPROVAL,SALES*REP_NOTIFICATIONS, NOTIFICATIONS
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Notification Type - ALL,CONNECTION*APPROVAL,SALES*REP_NOTIFICATIONS, NOTIFICATIONS
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public CloudRouterArgs build() {
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            $.notifications = Objects.requireNonNull($.notifications, "expected parameter 'notifications' to be non-null");
            $.package_ = Objects.requireNonNull($.package_, "expected parameter 'package' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
