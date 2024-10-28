resource "azurerm_cognitive_account" "el_ca_openai" {
  name                = "${var.org}-ca-openai-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  kind                = "OpenAI"
  custom_subdomain_name = "${var.org}-pv-openai"
  sku_name            = "S0"

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_private_endpoint" "pe_openai" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-pe-openai-${var.env}"
  location            = azurerm_resource_group.el_rg_net[0].location
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
  subnet_id           = azurerm_subnet.private_endpoints[0].id

  private_service_connection {
    name                           = "${var.org}-pe-psc-openai-${var.env}"
    private_connection_resource_id = azurerm_cognitive_account.el_ca_openai.id
    subresource_names              = ["account"]
    is_manual_connection           = false
  }
}

resource "azurerm_private_dns_a_record" "pv_dns_openai" {
  count = var.public_network_access_enabled ? 0 : 1

  name                = azurerm_cognitive_account.el_ca_openai.name
  zone_name = azurerm_private_dns_zone.openai[count.index].name
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
  ttl = 300
  records = [ azurerm_private_endpoint.pe_openai[count.index].private_service_connection[0].private_ip_address ]
}

resource "azurerm_key_vault_secret" "kv_s_openai" {
  name         = "openai-api-key"
  value = azurerm_cognitive_account.el_ca_openai.primary_access_key
  key_vault_id = azurerm_key_vault.main.id

  # Ensure the role assignment is created before the secret
  depends_on = [ azurerm_role_assignment.key_vault_secrets_officer ]
}

resource "azurerm_cognitive_deployment" "el_ca_openai_deployment" {
  name                 = "${var.org}-cd-openai-${var.env}"
  cognitive_account_id = azurerm_cognitive_account.el_ca_openai.id
  model {
    format  = "OpenAI"
    name    = "gpt-4o-realtime-preview"
    version = "2024-10-01"
  }

  sku {
    name = "GlobalStandard"
  }
}

resource "azurerm_log_analytics_workspace" "el_la_openai" {
  name                = "${var.org}-la-openai-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_monitor_diagnostic_setting" "example" {
  name                       = "${var.org}-montior-cd-${var.env}"
  target_resource_id         = azurerm_cognitive_account.el_ca_openai.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.el_la_openai.id

  enabled_log {
    category = "AuditLogs"
  }

  metric {
    category = "AllMetrics"
  }
}
