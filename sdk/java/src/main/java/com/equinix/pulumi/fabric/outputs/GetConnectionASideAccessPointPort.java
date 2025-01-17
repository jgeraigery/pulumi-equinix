// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.equinix.pulumi.fabric.outputs;

import com.equinix.pulumi.fabric.outputs.GetConnectionASideAccessPointPortRedundancy;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetConnectionASideAccessPointPort {
    private String href;
    private String name;
    private List<GetConnectionASideAccessPointPortRedundancy> redundancies;
    private String uuid;

    private GetConnectionASideAccessPointPort() {}
    public String href() {
        return this.href;
    }
    public String name() {
        return this.name;
    }
    public List<GetConnectionASideAccessPointPortRedundancy> redundancies() {
        return this.redundancies;
    }
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionASideAccessPointPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String href;
        private String name;
        private List<GetConnectionASideAccessPointPortRedundancy> redundancies;
        private String uuid;
        public Builder() {}
        public Builder(GetConnectionASideAccessPointPort defaults) {
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
        public Builder redundancies(List<GetConnectionASideAccessPointPortRedundancy> redundancies) {
            this.redundancies = Objects.requireNonNull(redundancies);
            return this;
        }
        public Builder redundancies(GetConnectionASideAccessPointPortRedundancy... redundancies) {
            return redundancies(List.of(redundancies));
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetConnectionASideAccessPointPort build() {
            final var _resultValue = new GetConnectionASideAccessPointPort();
            _resultValue.href = href;
            _resultValue.name = name;
            _resultValue.redundancies = redundancies;
            _resultValue.uuid = uuid;
            return _resultValue;
        }
    }
}
