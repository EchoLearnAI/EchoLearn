# ------------------------------------------------------------------------------
# RBAC
# ------------------------------------------------------------------------------

# Control Plane managed identity
resource "azurerm_user_assigned_identity" "control_plane_identity" {
  name                = "${module.naming.user_assigned_identity.name}-control-plane"
  resource_group_name = azurerm_resource_group.main.name
  location            = var.location
  tags                = var.tags
}

# Kubernetes Managed Identity
resource "azurerm_user_assigned_identity" "kubelet_identity" {
  name                = "${module.naming.user_assigned_identity.name}-kubelet"
  resource_group_name = azurerm_resource_group.main.name
  location            = var.location
  tags                = var.tags
}

# Assign Managed Identity Operator permissions for the control plane on kubelet
resource "azurerm_role_assignment" "control_plane_identity_to_kubelet_identity" {
  scope                = azurerm_user_assigned_identity.kubelet_identity.id
  role_definition_name = "Managed Identity Operator"
  principal_id         = azurerm_user_assigned_identity.control_plane_identity.principal_id
}

# Assign cluster admin permissions to provided clusters administrators
resource "azurerm_role_assignment" "aks_cluster_admin" {
  for_each = var.kubernetes_cluster_admin_group_ids

  principal_id         = each.value
  role_definition_name = "Azure Kubernetes Service RBAC Cluster Admin"
  scope                = azurerm_kubernetes_cluster.main.id
}

# Assign Network Contributor permissions for the control plane on VNET
# Assigning the role at the RG level to allow the Azure Storage NFS agents to interact
# with the security groups assigned to the VNET
resource "azurerm_role_assignment" "control_plane_identity_to_vnet_rg" {
  scope                = var.network.vnet_rg_id
  role_definition_name = "Network Contributor"
  principal_id         = azurerm_user_assigned_identity.control_plane_identity.principal_id
}