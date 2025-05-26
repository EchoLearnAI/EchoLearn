# ACR (Azure Container Registry)

Terraform module to setup Azure container registry. It is based on the official
[azurerm_container_registry][azurerm_container_registry] block.

[azurerm_container_registry]: https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry

## Usage

### Minimal configuration

```hcl
module "acr" {
  source = "path/to/acr/module"

  name                = "my-registry"
  resource_group_name = "resource-group"
  location            = "location"
}
```

### Login to ACR

```hcl
az acr login -n <registry name>
```

<!-- BEGIN_TF_DOCS -->
## Inputs

| Name | Description | Type |
|------|-------------|------|
| <a name="input_location"></a> [location](#input\_location) | (Required) The location of the resource group in which the Container Registry should be created | `string` |
| <a name="input_name"></a> [name](#input\_name) | (Required) Specifies the name of the Container Registry. Changing this forces a new resource to be created. Alpha numeric characters only. | `string` |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | (Required) The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created. | `string` |
| <a name="input_admin_enabled"></a> [admin\_enabled](#input\_admin\_enabled) | (Optional) Boolean value that indicates whether the admin user is enabled. Default false. | `bool` |
| <a name="input_anonymous_pull_enabled"></a> [anonymous\_pull\_enabled](#input\_anonymous\_pull\_enabled) | (Optional) Whether allows anonymous (unauthenticated) pull access to this Container Registry? Defaults to false. This is only supported on resources with the Standard or Premium SKU. | `bool` |
| <a name="input_data_endpoint_enabled"></a> [data\_endpoint\_enabled](#input\_data\_endpoint\_enabled) | (Optional) Whether to enable dedicated data endpoints for this Container Registry? Defaults to false. This is only supported on resources with the Premium SKU. | `bool` |
| <a name="input_encryption_enabled"></a> [encryption\_enabled](#input\_encryption\_enabled) | (Optional) Encrypt registry using a customer-managed key. Boolean value that indicates whether encryption is enabled. | `bool` |
| <a name="input_encryption_identity_client_id"></a> [encryption\_identity\_client\_id](#input\_encryption\_identity\_client\_id) | (Optional) The client ID of the managed identity associated with the encryption key. Required if encryption\_enabled = true. | `string` |
| <a name="input_encryption_key_vault_key_id"></a> [encryption\_key\_vault\_key\_id](#input\_encryption\_key\_vault\_key\_id) | (Optional) The ID of the Key Vault Key. Required if encryption\_enabled = true. | `string` |
| <a name="input_export_policy_enabled"></a> [export\_policy\_enabled](#input\_export\_policy\_enabled) | (Optional) Boolean value that indicates whether export policy is enabled. Defaults to true. In order to set it to false, make sure the public\_network\_access\_enabled is also set to false. This is only supported on resources with the Premium SKU. | `bool` |
| <a name="input_identity_ids"></a> [identity\_ids](#input\_identity\_ids) | (Optional) Specifies a list of User Assigned Managed Identity IDs to be assigned to this Container Registry. | `list(string)` |
| <a name="input_identity_type"></a> [identity\_type](#input\_identity\_type) | (Optional) Specifies the type of Managed Service Identity that should be configured on this Container Registry. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). | `string` |
| <a name="input_ipv4_networks_access"></a> [ipv4\_networks\_access](#input\_ipv4\_networks\_access) | (Optional) List of CIDR block from which requests will be allowed. Allow all by default. | `list(string)` |
| <a name="input_network_rule_bypass_option"></a> [network\_rule\_bypass\_option](#input\_network\_rule\_bypass\_option) | (Optional) Whether to allow trusted Azure services to access a network restricted Container Registry? Possible values are None and AzureServices. Defaults to AzureServices. | `string` |
| <a name="input_public_network_access_enabled"></a> [public\_network\_access\_enabled](#input\_public\_network\_access\_enabled) | (Optional) Whether public network access is allowed for the container registry. Defaults to true. | `bool` |
| <a name="input_quarantine_policy_enabled"></a> [quarantine\_policy\_enabled](#input\_quarantine\_policy\_enabled) | (Optional) Boolean value that indicates whether quarantine policy is enabled. Defaults to false. This is only supported on resources with the Premium SKU. | `bool` |
| <a name="input_retention_policy_in_days"></a> [retention\_policy\_in\_days](#input\_retention\_policy\_in\_days) | (Optional) The number of days to retain an untagged manifest after which it gets purged. Default is 7. | `number` |
| <a name="input_sku"></a> [sku](#input\_sku) | (Optional) The SKU name of the container registry. Possible values are Basic, Standard and Premium. | `string` |
| <a name="input_tags"></a> [tags](#input\_tags) | (Optional) A mapping of tags to assign to the resource. | `map(string)` |
| <a name="input_trust_policy_enabled"></a> [trust\_policy\_enabled](#input\_trust\_policy\_enabled) | (Optional) Boolean value that indicates whether the policy is enabled. This is only supported on resources with the Premium SKU. | `bool` |
| <a name="input_zone_redundancy_enabled"></a> [zone\_redundancy\_enabled](#input\_zone\_redundancy\_enabled) | (Optional) Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false. This is only supported on resources with the Premium SKU. | `bool` |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_acr_id"></a> [acr\_id](#output\_acr\_id) | Azure Container Registry identifier |
| <a name="output_acr_identity"></a> [acr\_identity](#output\_acr\_identity) | Container registry identity |
| <a name="output_acr_name"></a> [acr\_name](#output\_acr\_name) | Name of the Key Vault |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | Resource group name |
<!-- END_TF_DOCS -->
