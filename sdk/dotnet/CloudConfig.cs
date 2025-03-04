// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra
{
    /// <summary>
    /// Creates a Xen Orchestra cloud config resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Xenorchestra = Pulumi.Xenorchestra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var demo = new Xenorchestra.CloudConfig("demo", new()
    ///     {
    ///         Name = "cloud config name",
    ///         Template = @"#cloud-config
    /// 
    /// runcmd:
    ///  - [ ls, -l, / ]
    ///  - [ sh, -xc, ""echo $(date) ': hello world!'"" ]
    ///  - [ sh, -c, echo ""=========hello world'========="" ]
    ///  - ls -l /root
    /// ",
    ///     });
    /// 
    ///     var bar = new Xenorchestra.Vm("bar", new()
    ///     {
    ///         CloudConfig = demo.Template,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [XenorchestraResourceType("xenorchestra:index/cloudConfig:CloudConfig")]
    public partial class CloudConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the cloud config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;


        /// <summary>
        /// Create a CloudConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudConfig(string name, CloudConfigArgs args, CustomResourceOptions? options = null)
            : base("xenorchestra:index/cloudConfig:CloudConfig", name, args ?? new CloudConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudConfig(string name, Input<string> id, CloudConfigState? state = null, CustomResourceOptions? options = null)
            : base("xenorchestra:index/cloudConfig:CloudConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/vatesfr/pulumi-xenorchestra",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudConfig Get(string name, Input<string> id, CloudConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudConfig(name, id, state, options);
        }
    }

    public sealed class CloudConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cloud config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public CloudConfigArgs()
        {
        }
        public static new CloudConfigArgs Empty => new CloudConfigArgs();
    }

    public sealed class CloudConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cloud config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The cloud init config. See the cloud init docs for more [information](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public CloudConfigState()
        {
        }
        public static new CloudConfigState Empty => new CloudConfigState();
    }
}
