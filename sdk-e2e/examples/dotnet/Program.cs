using System;
using System.Collections.Generic;
using Pulumi;
using Pulumi.Xenorchestra;
using Pulumi.Xenorchestra.Inputs;

// Output keys are a contract shared by every language: see sdk-e2e/expectations.json.
static string Env(string k) => Environment.GetEnvironmentVariable(k) ?? "";

return await Deployment.RunAsync(() =>
{
    var pool = GetXoaPool.Invoke(new() { NameLabel = Env("XOA_POOL") });
    var poolId = pool.Apply(p => p.Id);
    var sr = GetXoaStorageRepository.Invoke(new() { NameLabel = Env("XOA_STORAGE"), PoolId = poolId });
    var net = GetXoaNetwork.Invoke(new() { NameLabel = Env("XOA_NETWORK"), PoolId = poolId });
    var tpl = GetXoaTemplate.Invoke(new() { NameLabel = Env("XOA_TEMPLATE"), PoolId = poolId });

    var vm = new Vm("vm", new VmArgs
    {
        NameLabel = Env("XOA_TEST_PREFIX") + "pulumi-e2e-dotnet",
        NameDescription = "pulumi sdk e2e test",
        Tags = { "pulumi-e2e" },
        Cpus = 1,
        MemoryMax = 1073741824,
        Template = tpl.Apply(t => t.Id),
        Disks = { new VmDiskArgs { SrId = sr.Apply(s => s.Id), NameLabel = "OS", Size = 8589934592 } },
        Networks = { new VmNetworkArgs { NetworkId = net.Apply(n => n.Id) } },
        PowerState = "Running",
    });

    return new Dictionary<string, object?>
    {
        ["poolId"] = poolId,
        ["srId"] = sr.Apply(s => s.Id),
        ["networkId"] = net.Apply(n => n.Id),
        ["templateId"] = tpl.Apply(t => t.Id),
        ["vmId"] = vm.Id,
        ["vmNameLabel"] = vm.NameLabel,
        ["vmPowerState"] = vm.PowerState,
        ["vmCpus"] = vm.Cpus,
        ["vmMemoryMax"] = vm.MemoryMax,
        ["vmTags"] = vm.Tags,
        ["vmIpv4Addresses"] = vm.Ipv4Addresses,
    };
});
