// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides information about a Xen Orchestra user. If the Xen Orchestra user account you are using is not an admin, see the `searchInSession` parameter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as xenorchestra from "@pulumi/xenorchestra";
 *
 * const user = xenorchestra.getXoaUser({
 *     username: "my-username",
 * });
 * ```
 */
export function getXoaUser(args: GetXoaUserArgs, opts?: pulumi.InvokeOptions): Promise<GetXoaUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("xenorchestra:index/getXoaUser:getXoaUser", {
        "searchInSession": args.searchInSession,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getXoaUser.
 */
export interface GetXoaUserArgs {
    /**
     * A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
     */
    searchInSession?: boolean;
    /**
     * The username of the XO user.
     */
    username: string;
}

/**
 * A collection of values returned by getXoaUser.
 */
export interface GetXoaUserResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
     */
    readonly searchInSession?: boolean;
    /**
     * The username of the XO user.
     */
    readonly username: string;
}
/**
 * Provides information about a Xen Orchestra user. If the Xen Orchestra user account you are using is not an admin, see the `searchInSession` parameter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as xenorchestra from "@pulumi/xenorchestra";
 *
 * const user = xenorchestra.getXoaUser({
 *     username: "my-username",
 * });
 * ```
 */
export function getXoaUserOutput(args: GetXoaUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetXoaUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("xenorchestra:index/getXoaUser:getXoaUser", {
        "searchInSession": args.searchInSession,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getXoaUser.
 */
export interface GetXoaUserOutputArgs {
    /**
     * A boolean which will search for the user in the current session (`session.getUser` Xen Orchestra RPC call). This allows a non admin user to look up their own user account.
     */
    searchInSession?: pulumi.Input<boolean>;
    /**
     * The username of the XO user.
     */
    username: pulumi.Input<string>;
}
