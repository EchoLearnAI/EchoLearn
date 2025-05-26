module "acr" {
  source              = "../modules/acr"
  name                = module.naming_global.container_registry.name
  resource_group_name = azurerm_resource_group.global.name
  location            = azurerm_resource_group.global.location
}

# ------------------------------------------------------------------------------
# ACR ROLES
# ------------------------------------------------------------------------------

resource "azurerm_role_assignment" "acr_pull" {
  for_each = var.acr_pull

  scope              = module.acr.acr_id
  role_definition_id = "AcrPull"
  principal_id       = each.value
}

# ------------------------------------------------------------------------------
# ACR CACHE RULES
# ------------------------------------------------------------------------------

resource "azurerm_container_registry_cache_rule" "registry_k8s_io" {
  name                  = "registry-k8s-io"
  container_registry_id = module.acr.acr_id
  target_repo           = "registry-k8s-io/*"
  source_repo           = "registry.k8s.io/*"
}

resource "azurerm_container_registry_cache_rule" "aci_external_secrets_io" {
  name                  = "oci-external-secrets-io"
  container_registry_id = module.acr.acr_id
  target_repo           = "aci-external-secrets-io/external-secrets/*"
  source_repo           = "ghcr.io/external-secrets/*"
}
