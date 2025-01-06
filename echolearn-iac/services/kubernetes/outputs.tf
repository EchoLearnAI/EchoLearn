output "identity" {
  value       = azurerm_user_assigned_identity.control_plane_identity
  description = <<-EOF
  The value of the entire cluster identity object (id, client_id and object_id)
  EOF
}

output "kubelet_identity" {
  value       = azurerm_kubernetes_cluster.main.kubelet_identity[0]
  description = <<-EOF
  The value of the entire kubelet identity object (client_id, object_id and
  user_assigned_identity_id)
  EOF
}

output "kube_config" {
  value       = azurerm_kubernetes_cluster.main.kube_config[0]
  description = <<-EOF
  The kube config file for the Azure Kubernetes Managed Cluster
  EOF
}

output "oidc_issuer_url" {
  value       = azurerm_kubernetes_cluster.main.oidc_issuer_url
  description = <<-EOF
  The OIDC issuer URL that is associated with the cluster.
  EOF
}

output "cluster_apps" {
  value = {
    ingress_nginx = {
      ip_addresses = module.ingress_nginx.ip_addresses
    }
  }
  description = "Cluster Apps related outputs"
}
