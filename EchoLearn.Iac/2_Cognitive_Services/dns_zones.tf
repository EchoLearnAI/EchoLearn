resource "azurerm_private_dns_zone" "app" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "privatelink.azurewebsites.net"
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
}

resource "azurerm_private_dns_zone" "key_vault" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "privatelink.vaultcore.azure.net"
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
}

resource "azurerm_private_dns_zone_virtual_network_link" "key_vault" {
  count = var.public_network_access_enabled ? 0 : 1

  name                  = "${azurerm_virtual_network.main[count.index].name}-link-to-${replace(azurerm_private_dns_zone.key_vault[count.index].name, ".", "-")}"
  resource_group_name   = azurerm_resource_group.el_rg_net[0].name
  private_dns_zone_name = azurerm_private_dns_zone.key_vault[count.index].name
  virtual_network_id    = azurerm_virtual_network.main[count.index].id
}

resource "azurerm_private_dns_zone_virtual_network_link" "openai" {
  count = var.public_network_access_enabled ? 0 : 1

  name                  = "${azurerm_virtual_network.main[count.index].name}-link-to-${replace(azurerm_private_dns_zone.openai[count.index].name, ".", "-")}"
  resource_group_name   = azurerm_resource_group.el_rg_net[0].name
  private_dns_zone_name = azurerm_private_dns_zone.openai[count.index].name
  virtual_network_id    = azurerm_virtual_network.main[count.index].id
}

resource "azurerm_private_dns_zone" "openai" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "privatelink.openai.azure.com"
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
}