// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra
{
    public static class GetXoaHost
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
        ///     var host1 = Xenorchestra.GetXoaHost.Invoke(new()
        ///     {
        ///         NameLabel = "Your host",
        ///     });
        /// 
        ///     var node = new Xenorchestra.Vm("node", new()
        ///     {
        ///         AffinityHost = host1.Apply(getXoaHostResult =&gt; getXoaHostResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetXoaHostResult> InvokeAsync(GetXoaHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetXoaHostResult>("xenorchestra:index/getXoaHost:getXoaHost", args ?? new GetXoaHostArgs(), options.WithDefaults());

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
        ///     var host1 = Xenorchestra.GetXoaHost.Invoke(new()
        ///     {
        ///         NameLabel = "Your host",
        ///     });
        /// 
        ///     var node = new Xenorchestra.Vm("node", new()
        ///     {
        ///         AffinityHost = host1.Apply(getXoaHostResult =&gt; getXoaHostResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetXoaHostResult> Invoke(GetXoaHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetXoaHostResult>("xenorchestra:index/getXoaHost:getXoaHost", args ?? new GetXoaHostInvokeArgs(), options.WithDefaults());

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
        ///     var host1 = Xenorchestra.GetXoaHost.Invoke(new()
        ///     {
        ///         NameLabel = "Your host",
        ///     });
        /// 
        ///     var node = new Xenorchestra.Vm("node", new()
        ///     {
        ///         AffinityHost = host1.Apply(getXoaHostResult =&gt; getXoaHostResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetXoaHostResult> Invoke(GetXoaHostInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetXoaHostResult>("xenorchestra:index/getXoaHost:getXoaHost", args ?? new GetXoaHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetXoaHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name label of the host.
        /// </summary>
        [Input("nameLabel", required: true)]
        public string NameLabel { get; set; } = null!;

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags (labels) applied to the given entity.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetXoaHostArgs()
        {
        }
        public static new GetXoaHostArgs Empty => new GetXoaHostArgs();
    }

    public sealed class GetXoaHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name label of the host.
        /// </summary>
        [Input("nameLabel", required: true)]
        public Input<string> NameLabel { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags (labels) applied to the given entity.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GetXoaHostInvokeArgs()
        {
        }
        public static new GetXoaHostInvokeArgs Empty => new GetXoaHostInvokeArgs();
    }


    [OutputType]
    public sealed class GetXoaHostResult
    {
        /// <summary>
        /// CPU information about the host. The 'cores' key will contain the number of cpu cores and the 'sockets' key will contain the number of sockets.
        /// </summary>
        public readonly ImmutableDictionary<string, int> Cpus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The memory size of the host.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The memory usage of the host.
        /// </summary>
        public readonly int MemoryUsage;
        /// <summary>
        /// The name label of the host.
        /// </summary>
        public readonly string NameLabel;
        /// <summary>
        /// Id of the pool that the host belongs to.
        /// </summary>
        public readonly string PoolId;
        /// <summary>
        /// The tags (labels) applied to the given entity.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetXoaHostResult(
            ImmutableDictionary<string, int> cpus,

            string id,

            int memory,

            int memoryUsage,

            string nameLabel,

            string poolId,

            ImmutableArray<string> tags)
        {
            Cpus = cpus;
            Id = id;
            Memory = memory;
            MemoryUsage = memoryUsage;
            NameLabel = nameLabel;
            PoolId = poolId;
            Tags = tags;
        }
    }
}
