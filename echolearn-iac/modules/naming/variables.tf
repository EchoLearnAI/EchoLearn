variable "name" {
  description = "Name of the application or service the resource is deployed for."
  type        = string
}

variable "environment" {
  description = "The environment in which the resource is deployed, e.g. nonprd, prd."
  type        = string
}

variable "segment" {
  description = "(Optional) The segment to which this resource belongs to. Default: Petcare"
  type        = string
  default     = "Petcare"
}

variable "project" {
  description = "(Optional) The project to which this resource belongs to. Default: Growth Digital Platform"
  type        = string
  default     = "Growth Digital Platform"
}

variable "product" {
  description = "(Optional) The product to which this resource belongs to, e.g. CMS, CRM, SSO, CDP or DAM"
  type        = string
  default     = ""
}

variable "location" {
  description = "(Optional) The region in which the resource is deployed."
  type        = string
  default     = ""
}

variable "increment" {
  description = "(Optional) A 2-digit increment number"
  type        = number
  default     = null
  validation {
    condition     = var.increment != null ? (var.increment >= 0 && var.increment < 100) : true
    error_message = "Increment must be two digits or left empty."
  }
}
