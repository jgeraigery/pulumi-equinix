// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.equinix.pulumi.fabric.outputs.GetConnectionZSideAccessPointPortRedundancy;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetConnectionZSideAccessPointPort {
    private String href;
    private String name;
    private List<GetConnectionZSideAccessPointPortRedundancy> redundancies;
    private String uuid;

    private GetConnectionZSideAccessPointPort() {}
    public String href() {
        return this.href;
    }
    public String name() {
        return this.name;
    }
    public List<GetConnectionZSideAccessPointPortRedundancy> redundancies() {
        return this.redundancies;
    }
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionZSideAccessPointPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String href;
        private String name;
        private List<GetConnectionZSideAccessPointPortRedundancy> redundancies;
        private String uuid;
        public Builder() {}
        public Builder(GetConnectionZSideAccessPointPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.href = defaults.href;
    	      this.name = defaults.name;
    	      this.redundancies = defaults.redundancies;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder href(String href) {
            this.href = Objects.requireNonNull(href);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder redundancies(List<GetConnectionZSideAccessPointPortRedundancy> redundancies) {
            this.redundancies = Objects.requireNonNull(redundancies);
            return this;
        }
        public Builder redundancies(GetConnectionZSideAccessPointPortRedundancy... redundancies) {
            return redundancies(List.of(redundancies));
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetConnectionZSideAccessPointPort build() {
            final var _resultValue = new GetConnectionZSideAccessPointPort();
            _resultValue.href = href;
            _resultValue.name = name;
            _resultValue.redundancies = redundancies;
            _resultValue.uuid = uuid;
            return _resultValue;
        }
    }
}
