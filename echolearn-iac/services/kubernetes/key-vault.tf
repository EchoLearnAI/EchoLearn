module "key_vault" {
  source              = "../../modules/key-vault"
  name                = module.naming.key_vault.name
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  subnet_ids = [var.network.subnet_aks_id]

  tags = var.tags
}

resource "azurerm_role_assignment" "admins" {
  for_each = var.kubernetes_cluster_admin_group_ids

  principal_id         = each.value
  role_definition_name = "Key Vault Administrator"
  scope                = module.key_vault.key_vault_id
}

resource "azurerm_role_assignment" "key_vault_secrets_user" {
  principal_id         = azurerm_user_assigned_identity.kubelet_identity.principal_id
  role_definition_name = "Key Vault Secrets User"
  scope                = module.key_vault.key_vault_id
}

resource "azurerm_private_endpoint" "key_vault" {
  count = var.key_vault_private_dns_zone == null ? 0 : 1

  name                = module.naming.private_endpoint.name
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  subnet_id           = var.network.subnet_private_link_id

  private_service_connection {
    name                           = module.naming.private_service_connection.name
    private_connection_resource_id = module.key_vault.key_vault_id
    is_manual_connection           = false
    subresource_names              = ["Vault"]
  }

  private_dns_zone_group {
    name                 = module.naming.private_dns_zone_group.name
    private_dns_zone_ids = [var.key_vault_private_dns_zone.id]
  }

  tags = var.tags
}