variable "location" {
  type        = string
  description = "Azure Region where all these resources will be provisioned"
  default     = "eastus2"
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

variable "client_id" {
  type        = string
  description = "Azure AD Client ID"
}

variable "client_secret" {
  type        = string
  description = "Azure AD Client Secret"
}

variable "tenant_id" {
  type        = string
  description = "Azure AD Tenant ID"
}

variable "subscription_id" {
  type        = string
  description = "Azure Subscription"
}