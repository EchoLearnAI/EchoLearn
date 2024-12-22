data "azurerm_subscription" "current" {}

locals {
  tags = merge(var.tags, data.azurerm_subscription.current.tags)
}