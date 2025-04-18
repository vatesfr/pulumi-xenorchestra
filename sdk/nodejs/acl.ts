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
 * const pool = xenorchestra.getXoaPool({
 *     nameLabel: "Your pool",
 * });
 * const user = xenorchestra.getXoaUser({
 *     username: "my-username",
 * });
 * const acl = new xenorchestra.Acl("acl", {
 *     subject: user.then(user => user.id),
 *     object: pool.then(pool => pool.id),
 *     action: "operator",
 * });
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'xenorchestra:index/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * Must be one of admin, operator, viewer. See the [Xen orchestra docs](https://xen-orchestra.com/docs/acls.html) on ACLs for more details.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The id of the object that will be able to be used by the subject.
     */
    public readonly object!: pulumi.Output<string>;
    /**
     * The uuid of the user account that the acl will apply to.
     */
    public readonly subject!: pulumi.Output<string>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.object === undefined) && !opts.urn) {
                throw new Error("Missing required property 'object'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["object"] = args ? args.object : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Acl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * Must be one of admin, operator, viewer. See the [Xen orchestra docs](https://xen-orchestra.com/docs/acls.html) on ACLs for more details.
     */
    action?: pulumi.Input<string>;
    /**
     * The id of the object that will be able to be used by the subject.
     */
    object?: pulumi.Input<string>;
    /**
     * The uuid of the user account that the acl will apply to.
     */
    subject?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * Must be one of admin, operator, viewer. See the [Xen orchestra docs](https://xen-orchestra.com/docs/acls.html) on ACLs for more details.
     */
    action: pulumi.Input<string>;
    /**
     * The id of the object that will be able to be used by the subject.
     */
    object: pulumi.Input<string>;
    /**
     * The uuid of the user account that the acl will apply to.
     */
    subject: pulumi.Input<string>;
}
