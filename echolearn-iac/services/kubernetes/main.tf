resource "azurerm_kubernetes_cluster" "main" {
  name                = "aks-${module.naming.resource_name}"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  dns_prefix                          = "aks-${module.naming.resource_name}"
  azure_policy_enabled                = true
  kubernetes_version                  = var.kubernetes_version
  local_account_disabled              = true
  workload_identity_enabled           = true
  oidc_issuer_enabled                 = true
  node_resource_group                 = "norg-${module.naming.resource_name}"
  private_cluster_enabled             = true
  private_dns_zone_id                 = var.private_dns_zone_id
  private_cluster_public_fqdn_enabled = true

  sku_tier = "Standard"

  default_node_pool {
    name                         = "systempool"
    vm_size                      = var.default_node_pool.vm_size
    zones                        = ["1", "2", "3"]
    auto_scaling_enabled         = true
    max_pods                     = var.default_node_pool.max_pods
    os_disk_size_gb              = 50
    os_disk_type                 = "Ephemeral"
    vnet_subnet_id               = var.network.subnet_aks_id
    max_count                    = var.default_node_pool.max_count
    min_count                    = var.default_node_pool.min_count
    node_count                   = var.default_node_pool.min_count
    temporary_name_for_rotation  = "sprotation"
    only_critical_addons_enabled = var.default_node_pool.only_critical_addons_enabled
    tags                         = var.tags

    upgrade_settings {
      drain_timeout_in_minutes = 30
      max_surge                = 1
    }
  }

  azure_active_directory_role_based_access_control {
    admin_group_object_ids = var.kubernetes_cluster_admin_group_ids
    azure_rbac_enabled     = true
  }

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.control_plane_identity.id]
  }

  kubelet_identity {
    user_assigned_identity_id = azurerm_user_assigned_identity.kubelet_identity.id
    client_id                 = azurerm_user_assigned_identity.kubelet_identity.client_id
    object_id                 = azurerm_user_assigned_identity.kubelet_identity.principal_id
  }

  network_profile {
    network_plugin = "none"
  }

  storage_profile {
    blob_driver_enabled = true
  }

  dynamic "oms_agent" {
    for_each = var.log_analytics_workspace_enable ? ["oms_agent"] : []
    content {
      log_analytics_workspace_id = var.log_analytics_workspace_enable ? azurerm_log_analytics_workspace.main.0.id : null
    }
  }

  lifecycle {
    ignore_changes = [
      # ignore node_count in case we are using AutoScaling
      default_node_pool[0].node_count,
      # MSG auto attach an analytics workspace upon creation
      microsoft_defender,
      oms_agent
    ]
  }

  depends_on = [
    azurerm_role_assignment.control_plane_identity_to_kubelet_identity,
    azurerm_role_assignment.control_plane_identity_to_vnet_resource_group
  ]

  tags = var.tags
}

resource "azurerm_kubernetes_cluster_node_pool" "main" {
  for_each = var.additional_node_pools

  kubernetes_cluster_id = azurerm_kubernetes_cluster.main.id
  vnet_subnet_id        = var.network.subnet_aks_id

  name     = each.key
  priority = "Regular"

  vm_size              = each.value.vm_size
  zones                = ["1", "2", "3"]
  auto_scaling_enabled = true
  max_pods             = each.value.max_pods
  os_disk_size_gb      = each.value.os_disk_size_gb
  os_disk_type         = "Ephemeral"
  max_count            = each.value.max_count
  min_count            = each.value.min_count
  node_count           = each.value.min_count
  node_labels          = each.value.node_labels
  node_taints          = each.value.node_taints

  upgrade_settings {
    drain_timeout_in_minutes = 30
    max_surge                = 1
  }

  lifecycle {
    ignore_changes = [
      node_count,
      tags
    ]
  }

  tags = var.tags
}
