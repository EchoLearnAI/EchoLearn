output "acr_id" {
  description = "Azure Container Registry identifier"
  value       = azurerm_container_registry.main.id
}

output "acr_name" {
  description = "Name of the Key Vault"
  value       = azurerm_container_registry.main.name
}

output "acr_identity" {
  description = "Container registry identity"
  value       = azurerm_container_registry.main.identity
}

output "resource_group_name" {
  description = "Resource group name"
  value       = var.resource_group_name
}
