import * as xoa from "@vates/pulumi-xenorchestra";

// Output keys are a contract shared by every language: see sdk-e2e/expectations.json.
const env = (k: string) => process.env[k] ?? "";

const pool = xoa.getXoaPoolOutput({ nameLabel: env("XOA_POOL") });
const sr = xoa.getXoaStorageRepositoryOutput({ nameLabel: env("XOA_STORAGE"), poolId: pool.id });
const net = xoa.getXoaNetworkOutput({ nameLabel: env("XOA_NETWORK"), poolId: pool.id });
const tpl = xoa.getXoaTemplateOutput({ nameLabel: env("XOA_TEMPLATE"), poolId: pool.id });

const vm = new xoa.Vm("vm", {
    nameLabel: env("XOA_TEST_PREFIX") + "pulumi-e2e-nodejs",
    nameDescription: "pulumi sdk e2e test",
    tags: ["pulumi-e2e"],
    cpus: 1,
    memoryMax: 1073741824,
    template: tpl.id,
    disks: [{ srId: sr.id, nameLabel: "OS", size: 8589934592 }],
    networks: [{ networkId: net.id }],
    powerState: "Running",
});

export const poolId = pool.id;
export const srId = sr.id;
export const networkId = net.id;
export const templateId = tpl.id;
export const vmId = vm.id;
export const vmNameLabel = vm.nameLabel;
export const vmPowerState = vm.powerState;
export const vmCpus = vm.cpus;
export const vmMemoryMax = vm.memoryMax;
export const vmTags = vm.tags;
export const vmIpv4Addresses = vm.ipv4Addresses;
