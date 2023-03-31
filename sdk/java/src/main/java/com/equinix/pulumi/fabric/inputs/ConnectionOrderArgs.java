// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionOrderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionOrderArgs Empty = new ConnectionOrderArgs();

    /**
     * Billing tier for connection bandwidth
     * 
     */
    @Import(name="billingTier")
    private @Nullable Output<String> billingTier;

    /**
     * @return Billing tier for connection bandwidth
     * 
     */
    public Optional<Output<String>> billingTier() {
        return Optional.ofNullable(this.billingTier);
    }

    /**
     * Order Identification
     * 
     */
    @Import(name="orderId")
    private @Nullable Output<String> orderId;

    /**
     * @return Order Identification
     * 
     */
    public Optional<Output<String>> orderId() {
        return Optional.ofNullable(this.orderId);
    }

    /**
     * Order Reference Number
     * 
     */
    @Import(name="orderNumber")
    private @Nullable Output<String> orderNumber;

    /**
     * @return Order Reference Number
     * 
     */
    public Optional<Output<String>> orderNumber() {
        return Optional.ofNullable(this.orderNumber);
    }

    /**
     * Purchase order number
     * 
     */
    @Import(name="purchaseOrderNumber")
    private @Nullable Output<String> purchaseOrderNumber;

    /**
     * @return Purchase order number
     * 
     */
    public Optional<Output<String>> purchaseOrderNumber() {
        return Optional.ofNullable(this.purchaseOrderNumber);
    }

    private ConnectionOrderArgs() {}

    private ConnectionOrderArgs(ConnectionOrderArgs $) {
        this.billingTier = $.billingTier;
        this.orderId = $.orderId;
        this.orderNumber = $.orderNumber;
        this.purchaseOrderNumber = $.purchaseOrderNumber;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionOrderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionOrderArgs $;

        public Builder() {
            $ = new ConnectionOrderArgs();
        }

        public Builder(ConnectionOrderArgs defaults) {
            $ = new ConnectionOrderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param billingTier Billing tier for connection bandwidth
         * 
         * @return builder
         * 
         */
        public Builder billingTier(@Nullable Output<String> billingTier) {
            $.billingTier = billingTier;
            return this;
        }

        /**
         * @param billingTier Billing tier for connection bandwidth
         * 
         * @return builder
         * 
         */
        public Builder billingTier(String billingTier) {
            return billingTier(Output.of(billingTier));
        }

        /**
         * @param orderId Order Identification
         * 
         * @return builder
         * 
         */
        public Builder orderId(@Nullable Output<String> orderId) {
            $.orderId = orderId;
            return this;
        }

        /**
         * @param orderId Order Identification
         * 
         * @return builder
         * 
         */
        public Builder orderId(String orderId) {
            return orderId(Output.of(orderId));
        }

        /**
         * @param orderNumber Order Reference Number
         * 
         * @return builder
         * 
         */
        public Builder orderNumber(@Nullable Output<String> orderNumber) {
            $.orderNumber = orderNumber;
            return this;
        }

        /**
         * @param orderNumber Order Reference Number
         * 
         * @return builder
         * 
         */
        public Builder orderNumber(String orderNumber) {
            return orderNumber(Output.of(orderNumber));
        }

        /**
         * @param purchaseOrderNumber Purchase order number
         * 
         * @return builder
         * 
         */
        public Builder purchaseOrderNumber(@Nullable Output<String> purchaseOrderNumber) {
            $.purchaseOrderNumber = purchaseOrderNumber;
            return this;
        }

        /**
         * @param purchaseOrderNumber Purchase order number
         * 
         * @return builder
         * 
         */
        public Builder purchaseOrderNumber(String purchaseOrderNumber) {
            return purchaseOrderNumber(Output.of(purchaseOrderNumber));
        }

        public ConnectionOrderArgs build() {
            return $;
        }
    }

}