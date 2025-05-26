# Global naming
variable "name" { type = string }
variable "project" { type = string }
variable "environment" { type = string }
variable "location" { type = string }
variable "increment" { type = number }

variable "subscription_id" { type = string }

variable "tags" {
  type    = map(string)
  default = {}
}


# ------------------------------------------------------------------------------
# AZURE KUBERNETES
# ------------------------------------------------------------------------------

variable "kubernetes_version" { type = string }
variable "kubernetes_cluster_admin_group_ids" { type = set(string) }
variable "system_pool" {
  type = object({
    min_count = number
    max_count = number
  })
}
variable "app_pool" {
  type = object({
    min_count = number
    max_count = number
  })
}

# ------------------------------------------------------------------------------
# MONITOR ALERTS
# ------------------------------------------------------------------------------
variable "monitor_action_group_id" {
  description = "The Action Group ID used for alerts"
  type        = string
}