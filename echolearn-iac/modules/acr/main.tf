resource "azurerm_container_registry" "main" {
  name = var.name

  location            = var.location
  resource_group_name = var.resource_group_name

  sku           = var.sku
  admin_enabled = var.admin_enabled

  dynamic "network_rule_set" {
    for_each = var.sku == "Premium" && length(var.ipv4_networks_access) > 0 ? ["network_rule_set"] : []
    content {
      default_action = "Deny"
      dynamic "ip_rule" {
        for_each = var.ipv4_networks_access
        content {
          action   = "Allow"
          ip_range = ip_rule.value
        }
      }
    }
  }

  network_rule_bypass_option    = var.network_rule_bypass_option
  public_network_access_enabled = var.public_network_access_enabled
  quarantine_policy_enabled     = var.quarantine_policy_enabled
  retention_policy_in_days      = var.retention_policy_in_days
  trust_policy_enabled          = var.trust_policy_enabled
  zone_redundancy_enabled       = var.zone_redundancy_enabled
  export_policy_enabled         = var.export_policy_enabled

  dynamic "identity" {
    for_each = (
      var.identity_type == null
      ? []
      : ["identity"]
    )

    content {
      type = var.identity_type
      identity_ids = (
        var.identity_type == "UserAssigned"
        ? var.identity_ids
        : null
      )
    }
  }

  dynamic "encryption" {
    for_each = var.encryption_key_vault_key_id != null && var.encryption_identity_client_id != null ? ["encryption"] : []
    content {
      key_vault_key_id   = var.encryption_key_vault_key_id
      identity_client_id = var.encryption_identity_client_id
    }
  }

  anonymous_pull_enabled = var.anonymous_pull_enabled
  data_endpoint_enabled  = var.data_endpoint_enabled

  tags = var.tags
}
