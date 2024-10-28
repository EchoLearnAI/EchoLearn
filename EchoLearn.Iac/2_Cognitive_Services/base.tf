resource "random_string" "main" {
  length  = 6
  special = false
  upper   = false
}

data "azurerm_client_config" "current" {}

# App components resource group
resource "azurerm_resource_group" "el_rg" {
  name     = "${var.org}-rg-${var.env}"
  location = var.location
}

# Network resource group
resource "azurerm_resource_group" "el_rg_net" {
  count    = var.public_network_access_enabled ? 0 : 1 # Only create if public network access is disabled
  name     = "${var.org}-rg-net-${var.env}"
  location = var.location
}