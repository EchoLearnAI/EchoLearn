variable "location" {
  type        = string
  description = "Azure Region where all these resources will be provisioned"
  default     = "westeurope"
}

variable "org" {
  type        = string
  description = "Organization name"
  default     = "el"
}

variable "env" {
  type        = string
  description = "Environment name"
  default     = "dev"
}

variable "public_network_access_enabled" {
  type        = bool
  description = "Should public network access be enabled for the application. Defaults to true."
  default     = true
}

# Virtual Machine
variable "deploy_virtual_machine" {
  type        = bool
  description = "Should a virtual machine be deployed to the resource group."
  default     = false
}

variable "enable_automatic_vm_shutdown" {
  type        = bool
  description = "Should the virtual machine be automatically shutdown at a specific time."
  default     = true
}