// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra.Inputs
{

    public sealed class ResourceSetLimitGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The numerical limit for the given type.
        /// </summary>
        [Input("quantity", required: true)]
        public Input<int> Quantity { get; set; } = null!;

        /// <summary>
        /// The type of resource set limit. Must be cpus, memory or disk.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourceSetLimitGetArgs()
        {
        }
        public static new ResourceSetLimitGetArgs Empty => new ResourceSetLimitGetArgs();
    }
}
