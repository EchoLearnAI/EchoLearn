resource "azurerm_resource_group" "main" {
  name     = "rg-${module.naming.resource_name}"
  location = var.location
  tags     = var.tags
}