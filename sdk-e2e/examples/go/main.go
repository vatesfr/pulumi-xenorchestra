package main

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	xoa "github.com/vatesfr/pulumi-xenorchestra/sdk/v2/go/xenorchestra"
)

// Output keys are a contract shared by every language: see sdk-e2e/expectations.json.
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := xoa.GetXoaPool(ctx, &xoa.GetXoaPoolArgs{NameLabel: os.Getenv("XOA_POOL")})
		if err != nil {
			return err
		}
		sr, err := xoa.GetXoaStorageRepository(ctx, &xoa.GetXoaStorageRepositoryArgs{
			NameLabel: os.Getenv("XOA_STORAGE"), PoolId: &pool.Id})
		if err != nil {
			return err
		}
		net, err := xoa.LookupXoaNetwork(ctx, &xoa.LookupXoaNetworkArgs{
			NameLabel: os.Getenv("XOA_NETWORK"), PoolId: &pool.Id})
		if err != nil {
			return err
		}
		tpl, err := xoa.GetXoaTemplate(ctx, &xoa.GetXoaTemplateArgs{
			NameLabel: os.Getenv("XOA_TEMPLATE"), PoolId: &pool.Id})
		if err != nil {
			return err
		}

		vm, err := xoa.NewVm(ctx, "vm", &xoa.VmArgs{
			NameLabel:       pulumi.String(os.Getenv("XOA_TEST_PREFIX") + "pulumi-e2e-go"),
			NameDescription: pulumi.String("pulumi sdk e2e test"),
			Tags:            pulumi.StringArray{pulumi.String("pulumi-e2e")},
			Cpus:            pulumi.Int(1),
			MemoryMax:       pulumi.Float64(1073741824),
			Template:        pulumi.String(tpl.Id),
			Disks: xoa.VmDiskArray{xoa.VmDiskArgs{
				SrId: pulumi.String(sr.Id), NameLabel: pulumi.String("OS"), Size: pulumi.Float64(8589934592),
			}},
			Networks:   xoa.VmNetworkArray{xoa.VmNetworkArgs{NetworkId: pulumi.String(net.Id)}},
			PowerState: pulumi.String("Running"),
		})
		if err != nil {
			return err
		}

		ctx.Export("poolId", pulumi.String(pool.Id))
		ctx.Export("srId", pulumi.String(sr.Id))
		ctx.Export("networkId", pulumi.String(net.Id))
		ctx.Export("templateId", pulumi.String(tpl.Id))
		ctx.Export("vmId", vm.ID())
		ctx.Export("vmNameLabel", vm.NameLabel)
		ctx.Export("vmPowerState", vm.PowerState)
		ctx.Export("vmCpus", vm.Cpus)
		ctx.Export("vmMemoryMax", vm.MemoryMax)
		ctx.Export("vmTags", vm.Tags)
		ctx.Export("vmIpv4Addresses", vm.Ipv4Addresses)
		return nil
	})
}
