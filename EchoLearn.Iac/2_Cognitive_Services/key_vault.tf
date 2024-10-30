data "azurerm_role_definition" "key_vault_secrets_officer" {
  name = "Key Vault Secrets Officer"
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
  enable_rbac_authorization       = true
}

# Ensure the principal running Terraform access to the Key Vault
resource "azurerm_role_assignment" "key_vault_secrets_officer" {
  scope                = azurerm_key_vault.main.id
  role_definition_name = data.azurerm_role_definition.key_vault_secrets_officer.name
  principal_id         = data.azurerm_client_config.current.object_id
}

# Connectivity to the VM to the Key Vault
resource "azurerm_private_endpoint" "pe_key_vault" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-pe-kv-${var.env}"
  location            = azurerm_resource_group.el_rg_net[0].location
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
  subnet_id           = azurerm_subnet.admin[0].id

  private_service_connection {
    name                           = "${var.org}-pe-psc-kv-${var.env}"
    private_connection_resource_id = azurerm_key_vault.main.id
    subresource_names              = ["vault"]
    is_manual_connection           = false
  }
}

resource "azurerm_private_dns_a_record" "dns_key_vault" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = azurerm_key_vault.main.name
  zone_name           = azurerm_private_endpoint.pe_key_vault[count.index].name
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
  ttl                 = 300
  records             = [azurerm_private_endpoint.pe_key_vault[count.index].private_service_connection[0].private_ip_address]
}