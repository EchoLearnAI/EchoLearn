# ------------------------------------------------------------------------------
# STORAGE ACCOUNT
# ------------------------------------------------------------------------------

module "naming_observability" {
  source      = "../../modules/naming"
  name        = "observability"
  environment = var.environment
  location    = var.location
  increment   = var.increment
}

resource "azurerm_resource_group" "observability" {
  name     = module.naming_observability.resource_group.name
  location = var.location
  tags     = var.tags
}

resource "azurerm_storage_account" "observability" {
  name                = module.naming_observability.storage_account.name
  location            = azurerm_resource_group.observability.location
  resource_group_name = azurerm_resource_group.observability.name

  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"
  min_tls_version          = "TLS1_2"

  allow_nested_items_to_be_public  = false
  cross_tenant_replication_enabled = false
  shared_access_key_enabled        = false
  is_hns_enabled                   = false

  identity {
    type = "SystemAssgined"
  }

  tags = var.tags

  lifecycle {
    ignore_changes = [network_rules[0].private_link_access]
  }
}

resource "azurerm_storage_account_network_rules" "observability" {
  storage_account_id = azurerm_storage_account.observability.id

  default_action             = "Deny"
  virtual_network_subnet_ids = [azurerm_subnet.echo_aks.id]
  bypass                     = ["AzureServices"]
}

resource "azurerm_private_endpoint" "observability" {
  name                = module.naming_observability.private_endpoint.name
  location            = azurerm_resource_group.observability.location
  resource_group_name = azurerm_resource_group.observability.name
  subnet_id           = azurerm_subnet.private_link.id

  private_service_connection {
    name                 = module.naming_observability.private_service_connection.name
    is_manual_connection = false
    subresource_names    = ["Blob"]
  }

  private_dns_zone_group {
    name                 = module.naming_observability.private_dns_zone_group.name
    private_dns_zone_ids = [azurerm_private_dns_zone.blob.id]
  }

  tags = var.tags
}