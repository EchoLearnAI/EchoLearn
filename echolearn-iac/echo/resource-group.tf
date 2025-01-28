module "naming_main" {
  source      = "../modules/naming"
  name        = "main"
  environment = var.environment
  location    = var.location
  increment   = var.increment
}

resource "azurerm_resource_group" "main" {
  name     = module.naming_main.resource_group.name
  location = var.location
  tags     = local.tags
}
