variable "name" {
  description = "The name of the Key Vault which is to be created."
  type        = string
}

variable "location" {
  description = "Location where the resource will be provisioned"
  type        = string
}

variable "resource_group_name" {
  description = "The resource group where Key Vault will be created."
  type        = string
}

variable "ipv4_networks_access" {
  description = "(Optional) IPV4 network access"
  type        = list(string)
  default     = []
}

variable "enabled_for_disk_encryption" {
  description = "(Optional) Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys."
  type        = bool
  default     = true
}

variable "purge_protection_enabled" {
  description = "(Optional) Enable purge protection on the Key Vault"
  type        = bool
  default     = true
}

variable "network_acls_default_action" {
  description = "(Optional) The Default Action to use when no rules match from ip_rules / virtual_network_subnet_ids. Possible values are Allow and Deny."
  type        = string
  default     = "Deny"
}

variable "bypass_network_rules" {
  description = "(Optional) Specifies which traffic can bypass the network rules. Possible values are AzureServices and None."
  type        = string
  default     = "AzureServices"
}

variable "sku" {
  description = "(Optional) Possible values are standard and premium."
  type        = string
  default     = "standard"
}

variable "access_policies" {
  description = "(Optional) Key Vault access policies, if not using Azure RBAC."
  type = list(object({
    object_id = string
    permissions = object({
      certificate_permissions = list(string)
      key_permissions         = list(string)
      secret_permissions      = list(string)
      storage_permissions     = list(string)
    })
  }))
  default = []
}

variable "enable_rbac_authorization" {
  description = "(Optional) Specify whether Key Vault uses Azure Role Based Access Control"
  type        = bool
  default     = true
}

variable "enabled_for_deployment" {
  description = "(Optional) Specifies whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the Key Vault."
  type        = bool
  default     = false
}

variable "soft_delete_retention_days" {
  description = "(Optional) The number of days that items should be retained for once soft-deleted. This value can be between 7 and 90 (the default) days."
  default     = 7
  type        = number
}

variable "enabled_for_template_deployment" {
  description = "(Optional) Specifies whether Azure Resource Manager is permitted to retrieve secrets from the Key Vault."
  type        = bool
  default     = false
}

variable "subnet_ids" {
  description = "(Optional) Network subnet ids for Key Vault acls"
  type        = list(string)
  default     = []
}

variable "tags" {
  description = "(Optional) List of Name-Value pairs tags for a resource"
  type        = map(string)
}
