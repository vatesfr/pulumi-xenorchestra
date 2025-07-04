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
 * const eth0 = xenorchestra.getXoaPif({
 *     device: "eth0",
 *     vlan: -1,
 * });
 * const demo_vm = new xenorchestra.Vm("demo-vm", {networks: [{
 *     networkId: eth0.then(eth0 => eth0.network),
 * }]});
 * ```
 */
export function getXoaPif(args: GetXoaPifArgs, opts?: pulumi.InvokeOptions): Promise<GetXoaPifResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("xenorchestra:index/getXoaPif:getXoaPif", {
        "device": args.device,
        "hostId": args.hostId,
        "vlan": args.vlan,
    }, opts);
}

/**
 * A collection of arguments for invoking getXoaPif.
 */
export interface GetXoaPifArgs {
    /**
     * The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
     */
    device: string;
    /**
     * The ID of the host that the PIF belongs to.
     */
    hostId?: string;
    /**
     * The VLAN the PIF belongs to.
     */
    vlan: number;
}

/**
 * A collection of values returned by getXoaPif.
 */
export interface GetXoaPifResult {
    /**
     * If the PIF is attached to the network.
     */
    readonly attached: boolean;
    /**
     * In case of a bond slave, the uuid of the bond master.
     */
    readonly bondMaster: string;
    /**
     * In case of a bond master, the PIFs (uuid) that are used for this bond.
     */
    readonly bondSlaves: string[];
    /**
     * The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
     */
    readonly device: string;
    /**
     * The host the PIF is associated with.
     */
    readonly host: string;
    /**
     * The ID of the host that the PIF belongs to.
     */
    readonly hostId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * True if this PIF is a bond master.
     */
    readonly isBondMaster: boolean;
    /**
     * True if this PIF is a bond slave.
     */
    readonly isBondSlave: boolean;
    /**
     * The network the PIF is associated with.
     */
    readonly network: string;
    /**
     * The pool the PIF is associated with.
     */
    readonly poolId: string;
    /**
     * The uuid of the PIF.
     */
    readonly uuid: string;
    /**
     * The VLAN the PIF belongs to.
     */
    readonly vlan: number;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as xenorchestra from "@pulumi/xenorchestra";
 * import * as xenorchestra from "@vates/pulumi-xenorchestra";
 *
 * const eth0 = xenorchestra.getXoaPif({
 *     device: "eth0",
 *     vlan: -1,
 * });
 * const demo_vm = new xenorchestra.Vm("demo-vm", {networks: [{
 *     networkId: eth0.then(eth0 => eth0.network),
 * }]});
 * ```
 */
export function getXoaPifOutput(args: GetXoaPifOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetXoaPifResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("xenorchestra:index/getXoaPif:getXoaPif", {
        "device": args.device,
        "hostId": args.hostId,
        "vlan": args.vlan,
    }, opts);
}

/**
 * A collection of arguments for invoking getXoaPif.
 */
export interface GetXoaPifOutputArgs {
    /**
     * The name of the network device. Examples include eth0, eth1, etc. See `ifconfig` for possible devices.
     */
    device: pulumi.Input<string>;
    /**
     * The ID of the host that the PIF belongs to.
     */
    hostId?: pulumi.Input<string>;
    /**
     * The VLAN the PIF belongs to.
     */
    vlan: pulumi.Input<number>;
}
