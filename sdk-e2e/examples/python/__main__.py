"""SDK e2e test (python). Output keys are a contract shared by every language: see sdk-e2e/expectations.json."""

import os

import pulumi
import pulumi_xenorchestra as xo

env = os.environ.get

pool = xo.get_xoa_pool(name_label=env("XOA_POOL"))
sr = xo.get_xoa_storage_repository(name_label=env("XOA_STORAGE"), pool_id=pool.id)
net = xo.get_xoa_network(name_label=env("XOA_NETWORK"), pool_id=pool.id)
tpl = xo.get_xoa_template(name_label=env("XOA_TEMPLATE"), pool_id=pool.id)

vm = xo.Vm(
    "vm",
    name_label=env("XOA_TEST_PREFIX", "") + "pulumi-e2e-python",
    name_description="pulumi sdk e2e test",
    tags=["pulumi-e2e"],
    cpus=1,
    memory_max=1073741824,
    template=tpl.id,
    disks=[xo.VmDiskArgs(sr_id=sr.id, name_label="OS", size=8589934592)],
    networks=[xo.VmNetworkArgs(network_id=net.id)],
    power_state="Running",
)

pulumi.export("poolId", pool.id)
pulumi.export("srId", sr.id)
pulumi.export("networkId", net.id)
pulumi.export("templateId", tpl.id)
pulumi.export("vmId", vm.id)
pulumi.export("vmNameLabel", vm.name_label)
pulumi.export("vmPowerState", vm.power_state)
pulumi.export("vmCpus", vm.cpus)
pulumi.export("vmMemoryMax", vm.memory_max)
pulumi.export("vmTags", vm.tags)
pulumi.export("vmIpv4Addresses", vm.ipv4_addresses)
