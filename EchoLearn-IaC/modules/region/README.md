# Region module

This module intend to provide a consistent naming convention of Azure regions.

## Getting started

```hcl
module "region" {
  source = "./region"
  name = "westeurope"
}

output "region_short" {
  value = module.region.region_short // weeu
}

output "region" {
  value = module.region.region // West Europe
}
```

<!-- BEGIN_TF_DOCS -->
## Inputs

| Name | Description | Type |
|------|-------------|------|
| <a name="input_name"></a> [name](#input\_name) | Azure Region standard name, CLI name or slug format | `string` |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_region"></a> [region](#output\_region) | Azure region in standard format |
| <a name="output_region_short"></a> [region\_short](#output\_region\_short) | Azure region in short format for resource naming purpose |
<!-- END_TF_DOCS -->