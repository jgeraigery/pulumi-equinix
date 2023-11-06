// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCloudRouterNotification {
    private List<String> emails;
    private String sendInterval;
    private String type;

    private GetCloudRouterNotification() {}
    public List<String> emails() {
        return this.emails;
    }
    public String sendInterval() {
        return this.sendInterval;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudRouterNotification defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> emails;
        private String sendInterval;
        private String type;
        public Builder() {}
        public Builder(GetCloudRouterNotification defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.emails = defaults.emails;
    	      this.sendInterval = defaults.sendInterval;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder emails(List<String> emails) {
            this.emails = Objects.requireNonNull(emails);
            return this;
        }
        public Builder emails(String... emails) {
            return emails(List.of(emails));
        }
        @CustomType.Setter
        public Builder sendInterval(String sendInterval) {
            this.sendInterval = Objects.requireNonNull(sendInterval);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetCloudRouterNotification build() {
            final var _resultValue = new GetCloudRouterNotification();
            _resultValue.emails = emails;
            _resultValue.sendInterval = sendInterval;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}