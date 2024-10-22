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