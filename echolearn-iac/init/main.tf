resource "azurerm_resource_group" "tf" {
  name     = "rg-${var.location}-${var.project_name}-${var.environment}"
  location = var.location
  tags     = local.tags
}

resource "azurerm_storage_account" "tfstate" {
  name                             = "sa-tfstate-${var.location}-${var.project_name}-${var.environment}"
  location                         = azurerm_resource_group.tf.location
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

  lifecycle {
    ignore_changes = [network_rules[0].private_link_access]
  }

  tags = local.tags
}

resource "azurerm_storage_container" "tfstate_init" {
  name                  = "tfstate-init"
  storage_account_id    = azurerm_storage_account.tfstate.id
  container_access_type = "private"
}

# ------------------------------------------------------------------------------
# GITHUB TERRAFORM FEDERATED IDENTITY
# ------------------------------------------------------------------------------

resource "azurerm_user_assigned_identity" "github_terraform" {
  name = "uai-github-${var.location}-${var.project_name}-${var.environment}"
  location = azurerm_resource_group.tf.location
  resource_group_name = azurerm_resource_group.tf.name
  tags = local.tags
}

resource "azurerm_role_assignment" "github_terraform_sub_contributor" {
  scope = data.azurerm_subscription.current.id
  role_definition_name = "Contributor"
  principal_id = azurerm_user_assigned_identity.github_terraform.principal_id
}

resource "azurerm_role_assignment" "github_terraform_aks_admin" {
  scope = data.azurerm_subscription.current.id
  role_definition_name = "Azure Kubernetes Service RBAC Cluster Admin"
  principal_id = azurerm_user_assigned_identity.github_terraform.principal_id
}