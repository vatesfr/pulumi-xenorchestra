import * as pulumi from "@pulumi/pulumi";
import * as xenorchestra from "@vatesfr/xenorchestra";

const resource = new xenorchestra.Resource("Resource", { sampleAttribute: "attr" });

export const sampleAttribute = resource.sampleAttribute;
