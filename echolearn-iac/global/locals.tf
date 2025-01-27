locals {
  tags = merge(var.tags, data.azurerm_subscription.current.tags)
}
