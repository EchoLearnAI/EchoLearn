resource "azurerm_network_security_group" "main" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-nsg-${var.env}"
  location            = azurerm_resource_group.el_rg_net.location
  resource_group_name = azurerm_resource_group.el_rg_net.name
}

resource "azurerm_virtual_network" "main" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-vnet-${var.env}"
  location            = azurerm_resource_group.el_rg_net.location
  resource_group_name = azurerm_resource_group.el_rg_net.name
  address_space       = ["10.0.0.0/16"]
  dns_servers         = []
}

resource "azurerm_subnet" "private_endpoints" {
  count = var.public_network_access_enabled ? 0 : 1

  name                 = "${var.org}-subnet-pe-${var.env}"
  resource_group_name  = azurerm_resource_group.el_rg_net[0].name
  virtual_network_name = azurerm_virtual_network.main[0].name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_subnet" "admin" {
  count = var.deploy_virtual_machine && var.public_network_access_enabled ? 0 : 1

  name                 = "${var.org}-subnet-${var.env}"
  resource_group_name  = azurerm_resource_group.el_rg_net[0].name
  virtual_network_name = azurerm_virtual_network.main[0].name
  address_prefixes     = ["10.0.2.0/24"]
}
