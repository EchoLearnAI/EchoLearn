variable "name" { type = string }
variable "increment" { type = string }

variable "location" {
  description = "The location of the resource"
  type        = string
}

variable "project" {
  description = "The project name or identifier"
  type        = string
}

variable "environment" {
  description = "The environment (e.g., stg, prd)"
  type        = string
}

variable "tags" {
  type    = map(string)
  default = {}
}

# ------------------------------------------------------------------------------
# KUBERNETES
# ------------------------------------------------------------------------------

variable "kubernetes_version" { type = string }
variable "kubernetes_cluster_admin_group_ids" { type = set(string) }

variable "private_cluster_enabled" {
  type        = bool
  description = "(Optional) If true cluster API server will be exposed only on internal IP address and available only in cluster VNET."
  default     = false
}

variable "private_dns_zone_id" {
  type        = string
  description = "(Optional) Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning. Changing this forces a new resource to be created."
  default     = null
}

variable "private_cluster_public_fqdn_enabled" {
  type        = bool
  description = "(Optional) Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`."
  default     = false
}

variable "default_node_pool" {
  description = "(Optional) Default node pool configuration"
  type = object({
    max_count                    = optional(number, 3)
    max_pods                     = optional(number, 110)
    min_count                    = optional(number, 1)
    os_disk_size_gb              = optional(number, 50)
    vm_size                      = optional(string, "Standard_D4ds_v5")
    only_critical_addons_enabled = optional(bool, false)
  })
  default = {}
}

variable "additional_node_pools" {
  description = "(Optional) Create additional node pools for the cluster"
  type = map(object({
    max_count       = optional(number, 3)
    max_pods        = optional(number, 110)
    min_count       = optional(number, 1)
    node_labels     = optional(map(string), {})
    node_taints     = optional(list(string), [])
    os_disk_size_gb = optional(number, 100)
    vm_size         = optional(string, "Standard_D8ds_v5")
  }))
  default = {}
}

# ------------------------------------------------------------------------------
# NETWORK
# ------------------------------------------------------------------------------

variable "network" {
  description = "Network config"
  type = object({
    subnet_aks_id          = string
    subnet_ingress_id      = string
    subnet_private_link_id = string
    vnet_resource_group_id = string
  })
}

variable "aks_lb_names" {
  description = "AKS Load Balancer names"
  type        = list(string)
}

# ------------------------------------------------------------------------------
# LOG ANALYTICS
# ------------------------------------------------------------------------------

variable "log_analytics_workspace_enable" {
  description = "(Optional) Enable the creation of log analytics"
  type        = bool
  default     = false
}

variable "monitor_action_group_id" {
  description = "The Action Group ID used for alerts"
  type        = string
}