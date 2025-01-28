# ------------------------------------------------------------------------------
# MAIN
# ------------------------------------------------------------------------------

resource "azurerm_virtual_network" "main" {
  name                = module.naming_main.virtual_network.name
  address_space       = ["10.240.0.0/16"]
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
}

# ------------------------------------------------------------------------------
# MAIN VPN
# ------------------------------------------------------------------------------

module "naming_main_vpn" {
  source      = "../modules/naming"
  name        = "main-vpn"
  environment = var.environment
  increment   = var.increment
  location    = var.location
}

resource "azurerm_subnet" "main_vpn" {
  name                 = "GatewaySubnet"
  resource_group_name  = azurerm_virtual_network.main.resource_group_name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.240.5.0/24"]

  private_endpoint_network_policies = "Enabled"
}

# ------------------------------------------------------------------------------
# MAIN AKS
# ------------------------------------------------------------------------------

module "naming_main_aks" {
  source      = "../modules/naming"
  name        = "echo-aks"
  environment = var.environment
  increment   = var.increment
  location    = var.location
}

resource "azurerm_subnet" "echo_aks" {
  name                 = module.naming_main_aks.subnet.name
  resource_group_name  = azurerm_virtual_network.main.resource_group_name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.240.0.0/22"]

  private_endpoint_network_policies = "Enabled"

  service_endpoints = [
    "Microsoft.KeyVault",
    "Microsoft.Sql",
    "Microsoft.Storage",
    "Microsoft.EventHub"
  ]
}

resource "azurerm_network_security_group" "echo_aks" {
  name                = module.naming_main_aks.network_security_group.name
  resource_group_name = azurerm_virtual_network.main.resource_group_name
  location            = azurerm_virtual_network.main.location

  tags = local.tags
}

resource "azurerm_subnet_network_security_group_association" "echo_aks" {
  subnet_id                 = azurerm_subnet.echo_aks.id
  network_security_group_id = azurerm_network_security_group.echo_aks.id
}

# ------------------------------------------------------------------------------
# MAIN INGRESS
# ------------------------------------------------------------------------------

module "naming_ingress" {
  source      = "../../modules/naming"
  name        = "echo-ingress"
  environment = var.environment
  increment   = var.increment
  location    = var.location
}

resource "azurerm_subnet" "ingress" {
  name                 = module.naming_ingress.subnet.name
  resource_group_name  = azurerm_virtual_network.main.resource_group_name
  virtual_network_name = azurerm_virtual_network.main.location
  address_prefixes     = ["10.240.4.0/28"]

  private_endpoint_network_policies = "Enabled"

  service_endpoints = ["Microsoft.Storage"]
}

resource "azurerm_network_security_group" "ingress" {
  name                = module.naming_ingress.network_security_group.name
  resource_group_name = azurerm_virtual_network.main.resource_group_name
  location            = azurerm_virtual_network.main.location

  tags = local.tags
}

resource "azurerm_subnet_network_security_group_association" "ingress" {
  subnet_id                 = azurerm_subnet.ingress.id
  network_security_group_id = azurerm_network_security_group.ingress.id
}

# ------------------------------------------------------------------------------
# MAIN PRIVATE LINK
# ------------------------------------------------------------------------------

module "naming_private_link" {
  source      = "../../modules/naming"
  name        = "echo-private-link"
  environment = var.environment
  increment   = var.increment
  location    = var.location
}

resource "azurerm_subnet" "private_link" {
  name                 = module.naming_private_link.subnet.name
  resource_group_name  = azurerm_virtual_network.main.resource_group_name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.240.4.16/28"]

  private_endpoint_network_policies             = "Enabled"
  private_link_service_network_policies_enabled = false

  service_endpoints = [
    "Microsoft.KeyVault",
    "Microsoft.Sql",
    "Microsoft.Storage",
    "Microsoft.EventHub"
  ]
}

resource "azurerm_network_security_group" "private_link" {
  name                = module.naming_private_link.network_security_group.name
  resource_group_name = azurerm_virtual_network.main.resource_group_name
  location            = azurerm_virtual_network.main.location

  tags = local.tags
}

resource "azurerm_subnet_network_security_group_association" "private_link" {
  subnet_id                 = azurerm_subnet.private_link.id
  network_security_group_id = azurerm_network_security_group.private_link.id
}
