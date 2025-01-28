# ------------------------------------------------------------------------------
# MAIN
# ------------------------------------------------------------------------------

resource "azurerm_virtual_network" "main" {
  name                = module.naming_main.virtual_network.name
  address_space       = ["10.240.0.0/16"]
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
}