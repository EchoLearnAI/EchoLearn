resource "azurerm_cognitive_account" "el_ca_openai" {
  name                = "${var.org}-ca-openai-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  kind                = "OpenAI"
  sku_name            = "S0"

  identity {
    type = "SystemAssigned"
  }
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
