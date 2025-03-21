// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as xenorchestra from "@pulumi/xenorchestra";
 * import * as xenorchestra from "@vates/pulumi-xenorchestra";
 *
 * const net = xenorchestra.getXoaNetwork({
 *     nameLabel: "Pool-wide network associated with eth0",
 * });
 * const demo_vm = new xenorchestra.Vm("demo-vm", {networks: [{
 *     networkId: net.then(net => net.id),
 * }]});
 * ```
 */
/** @deprecated xenorchestra.index/getnetwork.getNetwork has been deprecated in favor of xenorchestra.index/getxoanetwork.getXoaNetwork */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    pulumi.log.warn("getNetwork is deprecated: xenorchestra.index/getnetwork.getNetwork has been deprecated in favor of xenorchestra.index/getxoanetwork.getXoaNetwork")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("xenorchestra:index/getNetwork:getNetwork", {
        "bridge": args.bridge,
        "nameLabel": args.nameLabel,
        "poolId": args.poolId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The name of the bridge network interface.
     */
    bridge?: string;
    /**
     * The name of the network.
     */
    nameLabel: string;
    /**
     * The pool the network is associated with.
     */
    poolId?: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * The name of the bridge network interface.
     */
    readonly bridge: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the network.
     */
    readonly nameLabel: string;
    /**
     * The pool the network is associated with.
     */
    readonly poolId?: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as xenorchestra from "@pulumi/xenorchestra";
 * import * as xenorchestra from "@vates/pulumi-xenorchestra";
 *
 * const net = xenorchestra.getXoaNetwork({
 *     nameLabel: "Pool-wide network associated with eth0",
 * });
 * const demo_vm = new xenorchestra.Vm("demo-vm", {networks: [{
 *     networkId: net.then(net => net.id),
 * }]});
 * ```
 */
/** @deprecated xenorchestra.index/getnetwork.getNetwork has been deprecated in favor of xenorchestra.index/getxoanetwork.getXoaNetwork */
export function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNetworkResult> {
    pulumi.log.warn("getNetwork is deprecated: xenorchestra.index/getnetwork.getNetwork has been deprecated in favor of xenorchestra.index/getxoanetwork.getXoaNetwork")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("xenorchestra:index/getNetwork:getNetwork", {
        "bridge": args.bridge,
        "nameLabel": args.nameLabel,
        "poolId": args.poolId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkOutputArgs {
    /**
     * The name of the bridge network interface.
     */
    bridge?: pulumi.Input<string>;
    /**
     * The name of the network.
     */
    nameLabel: pulumi.Input<string>;
    /**
     * The pool the network is associated with.
     */
    poolId?: pulumi.Input<string>;
}
