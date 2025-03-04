// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xenorchestra.Inputs
{

    public sealed class VmNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the device should be attached to the VM.
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("expectedIpCidr")]
        public Input<string>? ExpectedIpCidr { get; set; }

        [Input("ipv4Addresses")]
        private InputList<string>? _ipv4Addresses;
        public InputList<string> Ipv4Addresses
        {
            get => _ipv4Addresses ?? (_ipv4Addresses = new InputList<string>());
            set => _ipv4Addresses = value;
        }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the network the VM will be on.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        public VmNetworkGetArgs()
        {
        }
        public static new VmNetworkGetArgs Empty => new VmNetworkGetArgs();
    }
}
