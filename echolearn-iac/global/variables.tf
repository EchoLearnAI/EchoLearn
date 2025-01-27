# Global naming
variable "name" { type = string }
variable "environment" { type = string }
variable "location" { type = string }
variable "increment" { type = number }

variable "subscription_id" { type = string }

variable "tags" {
  type    = map(string)
  default = {}
}
