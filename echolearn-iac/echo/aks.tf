# ------------------------------------------------------------------------------
# ECHO AKS
# ------------------------------------------------------------------------------

module "kubernetes_echo" {
  source = "../services/kubernetes"

  name         = "echo"
  environment  = var.environment
  location     = var.location
  increment    = var.increment
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
}