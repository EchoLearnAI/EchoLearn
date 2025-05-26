resource "azurerm_key_vault" "main" {
  name                            = var.name
  location                        = var.location
  resource_group_name             = var.resource_group_name
  tenant_id                       = data.azurerm_client_config.current.tenant_id
  purge_protection_enabled        = var.purge_protection_enabled
  soft_delete_retention_days      = var.soft_delete_retention_days
  enabled_for_deployment          = var.enabled_for_deployment
  enabled_for_template_deployment = var.enabled_for_template_deployment
  enabled_for_disk_encryption     = var.enabled_for_disk_encryption
  enable_rbac_authorization       = var.enable_rbac_authorization
  sku_name                        = var.sku

  network_acls {
    # Add ignore comment due to special case of GitHub Actions pulling secrets
    #tfsec:ignore:azure-keyvault-specify-network-acl[default_action=Allow]
    default_action             = var.network_acls_default_action
    bypass                     = var.bypass_network_rules
    ip_rules                   = var.ipv4_networks_access
    virtual_network_subnet_ids = var.subnet_ids
  }

  dynamic "access_policy" {
    for_each = [for ap in var.access_policies : {
      object_id   = ap.object_id
      permissions = ap.permissions
    }]

    content {
      tenant_id               = data.azurerm_client_config.current.tenant_id
      object_id               = access_policy.value.object_id
      key_permissions         = access_policy.value.permissions.key_permissions
      secret_permissions      = access_policy.value.permissions.secret_permissions
      certificate_permissions = access_policy.value.permissions.certificate_permissions
      storage_permissions     = access_policy.value.permissions.storage_permissions
    }
  }

  tags = var.tags
}
