// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra
{
    [Obsolete(@"xenorchestra.index/gettemplate.getTemplate has been deprecated in favor of xenorchestra.index/getxoatemplate.getXoaTemplate")]
    public static class GetTemplate
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
        ///     var template = Xenorchestra.GetXoaTemplate.Invoke(new()
        ///     {
        ///         NameLabel = "Ubuntu Bionic Beaver 18.04",
        ///     });
        /// 
        ///     var demo_vm = new Xenorchestra.Vm("demo-vm", new()
        ///     {
        ///         Template = template.Apply(getXoaTemplateResult =&gt; getXoaTemplateResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTemplateResult> InvokeAsync(GetTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateResult>("xenorchestra:index/getTemplate:getTemplate", args ?? new GetTemplateArgs(), options.WithDefaults());

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
        ///     var template = Xenorchestra.GetXoaTemplate.Invoke(new()
        ///     {
        ///         NameLabel = "Ubuntu Bionic Beaver 18.04",
        ///     });
        /// 
        ///     var demo_vm = new Xenorchestra.Vm("demo-vm", new()
        ///     {
        ///         Template = template.Apply(getXoaTemplateResult =&gt; getXoaTemplateResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("xenorchestra:index/getTemplate:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());

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
        ///     var template = Xenorchestra.GetXoaTemplate.Invoke(new()
        ///     {
        ///         NameLabel = "Ubuntu Bionic Beaver 18.04",
        ///     });
        /// 
        ///     var demo_vm = new Xenorchestra.Vm("demo-vm", new()
        ///     {
        ///         Template = template.Apply(getXoaTemplateResult =&gt; getXoaTemplateResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("xenorchestra:index/getTemplate:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the template to look up.
        /// </summary>
        [Input("nameLabel", required: true)]
        public string NameLabel { get; set; } = null!;

        /// <summary>
        /// The id of the pool that the template belongs to.
        /// </summary>
        [Input("poolId")]
        public string? PoolId { get; set; }

        public GetTemplateArgs()
        {
        }
        public static new GetTemplateArgs Empty => new GetTemplateArgs();
    }

    public sealed class GetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the template to look up.
        /// </summary>
        [Input("nameLabel", required: true)]
        public Input<string> NameLabel { get; set; } = null!;

        /// <summary>
        /// The id of the pool that the template belongs to.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        public GetTemplateInvokeArgs()
        {
        }
        public static new GetTemplateInvokeArgs Empty => new GetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the template to look up.
        /// </summary>
        public readonly string NameLabel;
        /// <summary>
        /// The id of the pool that the template belongs to.
        /// </summary>
        public readonly string? PoolId;
        /// <summary>
        /// The uuid of the template.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetTemplateResult(
            string id,

            string nameLabel,

            string? poolId,

            string uuid)
        {
            Id = id;
            NameLabel = nameLabel;
            PoolId = poolId;
            Uuid = uuid;
        }
    }
}
