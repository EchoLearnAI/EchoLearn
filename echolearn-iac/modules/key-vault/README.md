# Key Vault

This module create an Azure Key Vault.

## Getting started

```hcl
module "kv" {
  source      = "../modules/key-vault"

  name                = "core"
  location            = "westeurope"
  resource_group_name = "rg"
}

output "key_vault_id" {
  value = module.kv.key_vault_id
}
```
<!-- BEGIN_TF_DOCS -->
## Inputs

| Name | Description | Type |
|------|-------------|------|
| <a name="input_location"></a> [location](#input\_location) | Location where the resource will be provisioned | `string` |
| <a name="input_name"></a> [name](#input\_name) | The name of the Key Vault which is to be created. | `string` |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | The resource group where Key Vault will be created. | `string` |
| <a name="input_tags"></a> [tags](#input\_tags) | (Optional) List of Name-Value pairs tags for a resource | `map(string)` |
| <a name="input_access_policies"></a> [access\_policies](#input\_access\_policies) | (Optional) Key Vault access policies, if not using Azure RBAC. | <pre>list(object({<br/>    object_id = string<br/>    permissions = object({<br/>      certificate_permissions = list(string)<br/>      key_permissions         = list(string)<br/>      secret_permissions      = list(string)<br/>      storage_permissions     = list(string)<br/>    })<br/>  }))</pre> |
| <a name="input_bypass_network_rules"></a> [bypass\_network\_rules](#input\_bypass\_network\_rules) | (Optional) Specifies which traffic can bypass the network rules. Possible values are AzureServices and None. | `string` |
| <a name="input_enable_rbac_authorization"></a> [enable\_rbac\_authorization](#input\_enable\_rbac\_authorization) | (Optional) Specify whether Key Vault uses Azure Role Based Access Control | `bool` |
| <a name="input_enabled_for_deployment"></a> [enabled\_for\_deployment](#input\_enabled\_for\_deployment) | (Optional) Specifies whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the Key Vault. | `bool` |
| <a name="input_enabled_for_disk_encryption"></a> [enabled\_for\_disk\_encryption](#input\_enabled\_for\_disk\_encryption) | (Optional) Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. | `bool` |
| <a name="input_enabled_for_template_deployment"></a> [enabled\_for\_template\_deployment](#input\_enabled\_for\_template\_deployment) | (Optional) Specifies whether Azure Resource Manager is permitted to retrieve secrets from the Key Vault. | `bool` |
| <a name="input_ipv4_networks_access"></a> [ipv4\_networks\_access](#input\_ipv4\_networks\_access) | (Optional) IPV4 network access | `list(string)` |
| <a name="input_network_acls_default_action"></a> [network\_acls\_default\_action](#input\_network\_acls\_default\_action) | (Optional) The Default Action to use when no rules match from ip\_rules / virtual\_network\_subnet\_ids. Possible values are Allow and Deny. | `string` |
| <a name="input_purge_protection_enabled"></a> [purge\_protection\_enabled](#input\_purge\_protection\_enabled) | (Optional) Enable purge protection on the Key Vault | `bool` |
| <a name="input_sku"></a> [sku](#input\_sku) | (Optional) Possible values are standard and premium. | `string` |
| <a name="input_soft_delete_retention_days"></a> [soft\_delete\_retention\_days](#input\_soft\_delete\_retention\_days) | (Optional) The number of days that items should be retained for once soft-deleted. This value can be between 7 and 90 (the default) days. | `number` |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | (Optional) Network subnet ids for Key Vault acls | `list(string)` |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_key_vault_id"></a> [key\_vault\_id](#output\_key\_vault\_id) | Id of the Key Vault |
| <a name="output_key_vault_name"></a> [key\_vault\_name](#output\_key\_vault\_name) | Name of the Key Vault |
| <a name="output_key_vault_uri"></a> [key\_vault\_uri](#output\_key\_vault\_uri) | URI of the Key Vault |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | Name of the Keyvault resource group |
<!-- END_TF_DOCS -->
