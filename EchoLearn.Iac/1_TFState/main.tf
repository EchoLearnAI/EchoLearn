provider "azurerm" {
  features {}
  storage_use_azuread = true
  subscription_id     = "50438a2a-9e6c-4222-8d8b-f2157a24ab96"
}

data "azurerm_subscription" "main" {}
data "azurerm_client_config" "main" {}

resource "random_string" "main" {
  length  = 6
  special = false
  upper   = false
}

resource "azurerm_resource_group" "main" {
  name     = "el-rg-tfstate"
  location = "westeurope"
}

resource "azurerm_storage_account" "main" {
  name                = "elsa${random_string.main.result}"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  account_tier                      = "Standard"
  account_kind                      = "StorageV2"
  account_replication_type          = "GRS"
  https_traffic_only_enabled        = true
  min_tls_version                   = "TLS1_2"
  shared_access_key_enabled         = false # Desabling Storage account keys, using Azure EntraID for access
  default_to_oauth_authentication   = true
  infrastructure_encryption_enabled = false

  blob_properties {
    versioning_enabled            = true
    change_feed_enabled           = true
    change_feed_retention_in_days = 90
    last_access_time_enabled      = true

    delete_retention_policy {
      days = 30
    }

    container_delete_retention_policy {
      days = 30
    }
  }
}

resource "azurerm_storage_container" "main" {
  name                  = "tfstate"
  storage_account_name  = azurerm_storage_account.main.name
  container_access_type = "private"
}

resource "azurerm_role_assignment" "user_blob_contributor" {
  scope                = azurerm_storage_account.main.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azurerm_client_config.main.object_id
}