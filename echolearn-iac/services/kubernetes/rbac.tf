# ------------------------------------------------------------------------------
# RBAC
# ------------------------------------------------------------------------------

# Control Plane managed identity
resource "azurerm_user_assigned_identity" "control_plane_identity" {
  name = "${module.naming.user_assigned_identity.name}-control-plane"
  resource_group_name = azurerm_resource_group.main.name
  location = var.location
  tags = var.tags
}