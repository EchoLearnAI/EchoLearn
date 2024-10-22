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