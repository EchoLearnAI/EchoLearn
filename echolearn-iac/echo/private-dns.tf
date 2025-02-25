# ------------------------------------------------------------------------------
# KEYVAULT PRIVATE DNS
# ------------------------------------------------------------------------------

resource "azurerm_private_dns_zone" "key_vault" {
  name                = "privatelink.valtcore.core.windows.net"
  resource_group_name = azurerm_resource_group.main.name

  tags = local.tags
}

resource "azurerm_private_dns_zone_virtual_network_link" "keyvault_main" {
  name                  = "keyvault-core-link"
  resource_group_name   = azurerm_resource_group.main.name
  private_dns_zone_name = azurerm_private_dns_zone.key_vault.name
  virtual_network_id    = azurerm_virtual_network.main.id
}

# ------------------------------------------------------------------------------
# STORAGE ACCOUNT PRIVATE DNS
# ------------------------------------------------------------------------------

resource "azurerm_private_dns_zone" "blob" {
  name                = "privatelink.blob.core.windows.net"
  resource_group_name = azurerm_resource_group.main.name

  tags = local.tags
}

resource "azurerm_private_dns_zone_virtual_network_link" "blob_main" {
  name                  = "blob-core-link"
  resource_group_name   = azurerm_resource_group.main.name
  private_dns_zone_name = azurerm_private_dns_zone.blob.name
  virtual_network_id    = azurerm_virtual_network.main.id
}

resource "azurerm_private_dns_zone" "dfs" {
  name                = "privatelink.dfs.core.windows.net"
  resource_group_name = azurerm_resource_group.main.name

  tags = local.tags
}

resource "azurerm_private_dns_zone_virtual_network_link" "dfs_main" {
  name                  = "dfs-core-link"
  resource_group_name   = azurerm_resource_group.main.name
  private_dns_zone_name = azurerm_private_dns_zone.dfs.name
  virtual_network_id    = azurerm_virtual_network.main.id
}

# ------------------------------------------------------------------------------
# AKS MANAGEMENT
# ------------------------------------------------------------------------------

resource "azurerm_private_dns_zone" "aks" {
  name                = "privatelink.${azurerm_resource_group.main.location}.azmk8s.io"
  resource_group_name = azurerm_resource_group.main.name

  tags = local.tags
}

resource "azurerm_private_dns_zone_virtual_network_link" "aks_main" {
  name                  = "als-${azurerm_resource_group.main.location}-link"
  resource_group_name   = azurerm_resource_group.main.name
  private_dns_zone_name = azurerm_private_dns_zone.aks.name
  virtual_network_id    = azurerm_virtual_network.main.id
}