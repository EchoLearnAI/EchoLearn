module "naming_global" {
  source      = "../modules/naming"
  name        = var.name
  environment = var.environment
  location    = var.location
  increment   = var.increment
}

resource "azurerm_resource_group" "global" {
  name     = module.naming_global.resource_group.name
  location = var.location
  tags     = local.tags
}
