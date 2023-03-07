// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinix.networkedge;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.equinix.Utilities;
import com.pulumi.equinix.networkedge.BgpArgs;
import com.pulumi.equinix.networkedge.inputs.BgpState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource `equinix.networkedge.Bgp` allows creation and management of Equinix Network
 * Edge BGP peering configurations.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.equinix.networkedge.Bgp;
 * import com.pulumi.equinix.networkedge.BgpArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test = new Bgp(&#34;test&#34;, BgpArgs.builder()        
 *             .authenticationKey(&#34;secret&#34;)
 *             .connectionId(&#34;54014acf-9730-4b55-a791-459283d05fb1&#34;)
 *             .localAsn(12345)
 *             .localIpAddress(&#34;10.1.1.1/30&#34;)
 *             .remoteAsn(66123)
 *             .remoteIpAddress(&#34;10.1.1.2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using an existing ID
 * 
 * ```sh
 *  $ pulumi import equinix:networkedge/bgp:Bgp example {existing_id}
 * ```
 * 
 */
@ResourceType(type="equinix:networkedge/bgp:Bgp")
public class Bgp extends com.pulumi.resources.CustomResource {
    /**
     * shared key used for BGP peer authentication.
     * 
     */
    @Export(name="authenticationKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationKey;

    /**
     * @return shared key used for BGP peer authentication.
     * 
     */
    public Output<Optional<String>> authenticationKey() {
        return Codegen.optional(this.authenticationKey);
    }
    /**
     * identifier of a connection established between.
     * network device and remote service provider that will be used for peering.
     * 
     */
    @Export(name="connectionId", refs={String.class}, tree="[0]")
    private Output<String> connectionId;

    /**
     * @return identifier of a connection established between.
     * network device and remote service provider that will be used for peering.
     * 
     */
    public Output<String> connectionId() {
        return this.connectionId;
    }
    /**
     * unique identifier of a network device that is a local peer in a given BGP peering
     * configuration.
     * 
     */
    @Export(name="deviceId", refs={String.class}, tree="[0]")
    private Output<String> deviceId;

    /**
     * @return unique identifier of a network device that is a local peer in a given BGP peering
     * configuration.
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }
    /**
     * Local ASN number.
     * 
     */
    @Export(name="localAsn", refs={Integer.class}, tree="[0]")
    private Output<Integer> localAsn;

    /**
     * @return Local ASN number.
     * 
     */
    public Output<Integer> localAsn() {
        return this.localAsn;
    }
    /**
     * IP address in CIDR format of a local device.
     * 
     */
    @Export(name="localIpAddress", refs={String.class}, tree="[0]")
    private Output<String> localIpAddress;

    /**
     * @return IP address in CIDR format of a local device.
     * 
     */
    public Output<String> localIpAddress() {
        return this.localIpAddress;
    }
    /**
     * BGP peering configuration provisioning status, one of `PROVISIONING`,
     * `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
     * 
     */
    @Export(name="provisioningStatus", refs={String.class}, tree="[0]")
    private Output<String> provisioningStatus;

    /**
     * @return BGP peering configuration provisioning status, one of `PROVISIONING`,
     * `PENDING_UPDATE`, `PROVISIONED`, `FAILED`.
     * 
     */
    public Output<String> provisioningStatus() {
        return this.provisioningStatus;
    }
    /**
     * Remote ASN number.
     * 
     */
    @Export(name="remoteAsn", refs={Integer.class}, tree="[0]")
    private Output<Integer> remoteAsn;

    /**
     * @return Remote ASN number.
     * 
     */
    public Output<Integer> remoteAsn() {
        return this.remoteAsn;
    }
    /**
     * IP address of remote peer.
     * 
     */
    @Export(name="remoteIpAddress", refs={String.class}, tree="[0]")
    private Output<String> remoteIpAddress;

    /**
     * @return IP address of remote peer.
     * 
     */
    public Output<String> remoteIpAddress() {
        return this.remoteIpAddress;
    }
    /**
     * BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
     * `Established`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return BGP peer state, one of `Idle`, `Connect`, `Active`, `OpenSent`, `OpenConfirm`,
     * `Established`.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * BGP peering configuration unique identifier.
     * 
     */
    @Export(name="uuid", refs={String.class}, tree="[0]")
    private Output<String> uuid;

    /**
     * @return BGP peering configuration unique identifier.
     * 
     */
    public Output<String> uuid() {
        return this.uuid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Bgp(String name) {
        this(name, BgpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Bgp(String name, BgpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Bgp(String name, BgpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix:networkedge/bgp:Bgp", name, args == null ? BgpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Bgp(String name, Output<String> id, @Nullable BgpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix:networkedge/bgp:Bgp", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Bgp get(String name, Output<String> id, @Nullable BgpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Bgp(name, id, state, options);
    }
}