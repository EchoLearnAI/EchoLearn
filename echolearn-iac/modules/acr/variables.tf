variable "name" {
  description = "(Required) Specifies the name of the Container Registry. Changing this forces a new resource to be created. Alpha numeric characters only."
  type        = string
}

variable "resource_group_name" {
  description = "(Required) The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created."
  type        = string
}

variable "location" {
  description = "(Required) The location of the resource group in which the Container Registry should be created"
  type        = string
}

variable "sku" {
  description = "(Optional) The SKU name of the container registry. Possible values are Basic, Standard and Premium."
  type        = string
  default     = "Premium"
}

variable "identity_type" {
  description = "(Optional) Specifies the type of Managed Service Identity that should be configured on this Container Registry. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both)."
  type        = string
  default     = null
}

variable "identity_ids" {
  description = "(Optional) Specifies a list of User Assigned Managed Identity IDs to be assigned to this Container Registry."
  type        = list(string)
  default     = []
}

variable "zone_redundancy_enabled" {
  description = "(Optional) Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false. This is only supported on resources with the Premium SKU."
  type        = bool
  default     = false
}

variable "public_network_access_enabled" {
  description = "(Optional) Whether public network access is allowed for the container registry. Defaults to true."
  type        = bool
  default     = true
}

variable "quarantine_policy_enabled" {
  description = "(Optional) Boolean value that indicates whether quarantine policy is enabled. Defaults to false. This is only supported on resources with the Premium SKU."
  type        = bool
  default     = false
}

variable "export_policy_enabled" {
  description = "(Optional) Boolean value that indicates whether export policy is enabled. Defaults to true. In order to set it to false, make sure the public_network_access_enabled is also set to false. This is only supported on resources with the Premium SKU."
  type        = bool
  default     = true
}

variable "anonymous_pull_enabled" {
  description = "(Optional) Whether allows anonymous (unauthenticated) pull access to this Container Registry? Defaults to false. This is only supported on resources with the Standard or Premium SKU."
  type        = bool
  default     = false
}

variable "data_endpoint_enabled" {
  description = "(Optional) Whether to enable dedicated data endpoints for this Container Registry? Defaults to false. This is only supported on resources with the Premium SKU."
  type        = bool
  default     = false
}

variable "ipv4_networks_access" {
  description = "(Optional) List of CIDR block from which requests will be allowed. Allow all by default."
  type        = list(string)
  default     = []
}

variable "network_rule_bypass_option" {
  description = "(Optional) Whether to allow trusted Azure services to access a network restricted Container Registry? Possible values are None and AzureServices. Defaults to AzureServices."
  type        = string
  default     = "AzureServices"
}

variable "trust_policy_enabled" {
  description = "(Optional) Boolean value that indicates whether the policy is enabled. This is only supported on resources with the Premium SKU."
  type        = bool
  default     = false
}

variable "retention_policy_in_days" {
  description = "(Optional) The number of days to retain an untagged manifest after which it gets purged. Default is 7."
  type        = number
  default     = 7
}

variable "encryption_enabled" {
  description = "(Optional) Encrypt registry using a customer-managed key. Boolean value that indicates whether encryption is enabled."
  type        = bool
  default     = false
}

variable "encryption_key_vault_key_id" {
  description = "(Optional) The ID of the Key Vault Key. Required if encryption_enabled = true."
  type        = string
  default     = null
}

variable "encryption_identity_client_id" {
  description = "(Optional) The client ID of the managed identity associated with the encryption key. Required if encryption_enabled = true."
  type        = string
  default     = null
}

variable "admin_enabled" {
  description = "(Optional) Boolean value that indicates whether the admin user is enabled. Default false."
  type        = bool
  default     = false
}

variable "tags" {
  type        = map(string)
  description = "(Optional) A mapping of tags to assign to the resource."
  default     = {}
}
