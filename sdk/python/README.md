# Xen Orchestra Provider

The provider on top of Xen Orchestra to manage XCP-ng resources.
The releases of this package are hosted on the DESY Gitlab [here](https://gitlab.desy.de/cloud-public/pulumi-xenorchestra). 
Pull Requests from dependabot are managed here and then synced back to our Gitlab.
Only Python packages are released for the provider since we cannot test it for other languages.

## Installing 

If you want to use this provider in your python Code you need to install the python package
from the Gitlab pypi registry:
```
pip install pulumi-xenorchestra --index-url https://gitlab.desy.de/api/v4/projects/14393/packages/pypi/simple
```
Then you can run the following command to install the plugin (the golang binary) from the gitlab releaes page:
```
pulumi plugin install resource xenorchestra 1.2.0 --server https://gitlab.desy.de/cloud-public/pulumi-xenorchestra/-/releases/v1.2.0/downloads/
````

## Usage

in your `__main__.py` set up the Xen-Orchestra provider for pulumi in the following way:
```
import pulumi_xenorchestra as xoa

xoa_provider = xoa.Provider(
    "xenorchestra",
    xoa.ProviderArgs(
        url=Output.secret(secret_data["xoa_url"]),
        token=Output.secret(secret_data["xoa_token"]),
        insecure=False,
)
```

You can then une the provider to create a Virtual Machine:

```
xoa.Vm(
    resource_name=machine.name,
    name_label=machine.name,
    name_description="XOA VM",
    cpus=4,
    memory_max=8590000000,
    template=template.id,
    tags=["pulumi"],
    cloud_config=pathlib.Path("./config/cloudinit-xen-static.yaml").read_text()
    ),
    disks=[
        xoa.VmDiskArgs(
            sr_id=machine.storage_repository_id,
            name_label="OS",
            size=34361835520,
        ),
    ],
    networks=[
        xoa.VmNetworkArgs(
            network_id=network.id,
            mac_address=map_ip_to_mac(
                machine.ipv4_address,
                separator=":",
            ),
        ),
    ],
    opts=pulumi.ResourceOptions(
        provider=xoa_provider,
        # Do not recreate node with new credentials
        ignore_changes=["cloud_config"],
        parent=self,
    ),
)
```