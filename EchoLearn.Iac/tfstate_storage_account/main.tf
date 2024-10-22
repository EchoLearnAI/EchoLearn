provider "azurerm" {
  features {}
  # Set storage access to use Azure AD instead of storage key SAS
  storage_use_azuread = true
  subscription_id     = "50438a2a-9e6c-4222-8d8b-f2157a24ab96"
}

data "azurerm_subscription" "main" {}

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

  sas_policy { # Generating a SAS policy for the storage account, only 2 hours of access
    expiration_period = "00.02:00:00"
    expiration_action = "Log"
  }
}

resource "azurerm_storage_container" "main" {
  name                  = "tfstate"
  storage_account_name  = azurerm_storage_account.main.name
  container_access_type = "private"
}

resource "azurerm_role_definition" "main" {
  name        = "el-role-write-access"
  scope       = data.azurerm_subscription.main.id
  description = "Custom role definition allowing write access to storage account ${azurerm_storage_account.main.name}"

  permissions {
    actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/read",
      "Microsoft.Storage/storageAccounts/blobServices/generateUserDelegationKey/action"
    ]
    data_actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read",
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/write",
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/add/action"
    ]
  }

  assignable_scopes = [
    data.azurerm_subscription.main.id
  ]
}

# Create a service principal to assign the role to
data "azuread_client_config" "current" {}

resource "azuread_application" "main" {
  display_name = "el-app-tfstate"
  owners       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal" "main" {
  client_id                    = azuread_application.main.client_id
  app_role_assignment_required = false
  owners                       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal_password" "main" {
  service_principal_id = azuread_service_principal.main.id
}

resource "azurerm_role_assignment" "main" {
  scope                = azurerm_storage_account.main.id
  role_definition_name = azurerm_role_definition.main.name
  principal_id         = azuread_service_principal.main.object_id

  condition = templatefile("${path.module}/condition.tpl", {
    container_name = azurerm_storage_container.main.name
    state_path     = "echotf"
  })

  condition_version                = "2.0"
  skip_service_principal_aad_check = true

  depends_on = [azurerm_role_definition.main]
}

# Outputs
output "storage_account_name" {
  value = azurerm_storage_account.main.name
}

output "azurerm_storage_container" {
  value = azurerm_storage_container.main.name
}

output "role_name" {
  value = azurerm_role_definition.main.name
}

output "service_principal" {
  value = azuread_service_principal.main
}

output "subscription_id" {
  value = data.azurerm_subscription.main.subscription_id
}

output "service_principal_password" {
  value = nonsensitive(azuread_service_principal_password.main.value)
}