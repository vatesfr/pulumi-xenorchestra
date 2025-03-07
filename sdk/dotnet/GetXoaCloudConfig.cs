// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra
{
    public static class GetXoaCloudConfig
    {
        /// <summary>
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
        ///     var cloudConfig = Xenorchestra.GetXoaCloudConfig.Invoke(new()
        ///     {
        ///         Name = "Name of cloud config",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetXoaCloudConfigResult> InvokeAsync(GetXoaCloudConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetXoaCloudConfigResult>("xenorchestra:index/getXoaCloudConfig:getXoaCloudConfig", args ?? new GetXoaCloudConfigArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var cloudConfig = Xenorchestra.GetXoaCloudConfig.Invoke(new()
        ///     {
        ///         Name = "Name of cloud config",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetXoaCloudConfigResult> Invoke(GetXoaCloudConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetXoaCloudConfigResult>("xenorchestra:index/getXoaCloudConfig:getXoaCloudConfig", args ?? new GetXoaCloudConfigInvokeArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var cloudConfig = Xenorchestra.GetXoaCloudConfig.Invoke(new()
        ///     {
        ///         Name = "Name of cloud config",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetXoaCloudConfigResult> Invoke(GetXoaCloudConfigInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetXoaCloudConfigResult>("xenorchestra:index/getXoaCloudConfig:getXoaCloudConfig", args ?? new GetXoaCloudConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetXoaCloudConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cloud config you want to look up.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetXoaCloudConfigArgs()
        {
        }
        public static new GetXoaCloudConfigArgs Empty => new GetXoaCloudConfigArgs();
    }

    public sealed class GetXoaCloudConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cloud config you want to look up.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetXoaCloudConfigInvokeArgs()
        {
        }
        public static new GetXoaCloudConfigInvokeArgs Empty => new GetXoaCloudConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetXoaCloudConfigResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the cloud config you want to look up.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The contents of the cloud-config.
        /// </summary>
        public readonly string Template;

        [OutputConstructor]
        private GetXoaCloudConfigResult(
            string id,

            string name,

            string template)
        {
            Id = id;
            Name = name;
            Template = template;
        }
    }
}
