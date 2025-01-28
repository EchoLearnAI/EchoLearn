module "naming" {
  source      = "../modules/naming"
  name        = var.name
  environment = var.environment
  location    = var.location
  increment   = var.increment
}

resource "azurerm_resource_group" "tf" {
  name     = module.naming.resource_group.name
  location = var.location
  tags     = local.tags
}

resource "azurerm_storage_account" "tfstate" {
  name                             = module.naming.storage_account.name
  location                         = var.location
  resource_group_name              = azurerm_resource_group.tf.name
  account_tier                     = "Standard"
  account_replication_type         = "ZRS"
  account_kind                     = "StorageV2"
  allow_nested_items_to_be_public  = false
  cross_tenant_replication_enabled = false
  min_tls_version                  = "TLS1_2"

  blob_properties {
    versioning_enabled = true
  }

  identity {
    type = "SystemAssigned"
  }

  tags = local.tags

  lifecycle {
    ignore_changes = [
      # Managed by policy
      network_rules[0].private_link_access
    ]
  }
}

resource "azurerm_storage_container" "tfstate_init" {
  name                  = "tfstate-init"
  storage_account_id    = azurerm_storage_account.tfstate.id
  container_access_type = "private"
}

resource "azurerm_storage_container" "tfstate_us" {
  name                  = "tfstate-us"
  storage_account_id    = azurerm_storage_account.tfstate.id
  container_access_type = "private"
}

resource "azurerm_storage_container" "tfstate_eu" {
  name                  = "tfstate-eu"
  storage_account_id    = azurerm_storage_account.tfstate.id
  container_access_type = "private"
}

resource "azurerm_storage_container" "tfstate_global" {
  name                  = "tfstate-global"
  storage_account_id    = azurerm_storage_account.tfstate.id
  container_access_type = "private"
}