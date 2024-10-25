data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "main" {
  name                            = "el-kv-${random_string.main.result}"
  location                        = azurerm_resource_group.main.location
  resource_group_name             = azurerm_resource_group.main.name
  tenant_id                       = data.azurerm_subscription.main.tenant_id
  sku_name                        = "standard"
  soft_delete_retention_days      = 30
  purge_protection_enabled        = true
  enabled_for_disk_encryption     = true
  enabled_for_deployment          = true
  enabled_for_template_deployment = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    secret_permissions = ["Set", "Get", "Delete"]
  }
}
