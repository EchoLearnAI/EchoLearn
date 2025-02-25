# ------------------------------------------------------------------------------
# ECHO AKS
# ------------------------------------------------------------------------------

module "kubernetes_echo" {
  source = "../services/kubernetes"

  name         = "echo"
  environment  = var.environment
  location     = var.location
  increment    = var.increment
  project = var.project
  landing_zone = "eastus"

  kubernetes_cluster_admin_group_ids = var.kubernetes_cluster_admin_group_ids
  kubernetes_version                 = var.kubernetes_version

  default_node_pool = {
    max_count                    = var.system_pool.max_count
    min_count                    = var.system_pool.min_count
    only_critical_addons_enabled = true
  }

  additional_node_pools = {
    apps = {
      min_count = var.app_pool.min_count
      max_count = var.app_pool.max_count
      node_labels = {
        workload = "apps-default"
      }
    }
  }

  network = {
    subnet_aks_id = azurerm_subnet.echo_aks.id
    subnet_ingress_id = azurerm_subnet.ingress.id
    subnet_private_link_id = azurerm_subnet.private_link.id
    vnet_resource_group_id = azurerm_resource_group.main.id
  }

  cluster_apps = {
    cert_manager = { enabled = true }
    external_dns = { enabled = true }
  }

  key_vault_private_dns_zone = azurerm_private_dns_zone.key_vault
  aks_lb_names = [ "kubernetes", "kubernetes-internal" ]
  monitor_action_group_id = var.monitor_action_group_id
}
