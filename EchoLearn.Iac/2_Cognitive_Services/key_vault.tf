data "azurerm_role_definition" "key_vault_secrets_officer" {
  name = "Key Vault Secrets Officer"
}

data "azurerm_role_definition" "key_vault_secrets_user" {
  name = "Key Vault Secrets User"
}

resource "azurerm_key_vault" "main" {
  name                            = "${var.org}-kv-${random_string.main.result}"
  location                        = azurerm_resource_group.el_rg.location
  resource_group_name             = azurerm_resource_group.el_rg.name
  tenant_id                       = data.azurerm_subscription.main.tenant_id
  sku_name                        = "standard"
  soft_delete_retention_days      = 30
  purge_protection_enabled        = false
  enabled_for_disk_encryption     = true
  enabled_for_deployment          = true
  enabled_for_template_deployment = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    secret_permissions = ["Set", "Get", "Delete"]
  }
}

# Ensure the principal running Terraform access to the Key Vault
resource "azurerm_role_assignment" "key_vault_secrets_officer" {
  scope                = azurerm_key_vault.main.id
  role_definition_name = data.azurerm_role_definition.key_vault_secrets_officer.name
  principal_id         = data.azurerm_client_config.current.object_id
}