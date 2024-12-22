variable "location" { type = string }
variable "project_name" { type = string }
variable "environment" { type = string }
variable "subscription_id" { type = string }

variable "tags" {
  type    = map(string)
  default = {}
}
