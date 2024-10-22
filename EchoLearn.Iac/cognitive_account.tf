terraform {
  backend "azurerm" {
    storage_account_name = "elsadjz1o7"
    container_name       = "tfstate"
    key                  = "echotf/terraform.tfstate"
    use_azuread_auth     = true
    client_id            = ""
    client_secret        = ""
    tenant_id            = ""
    subscription_id      = ""
  }
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>4.6.0"
    }
  }
}

provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = false
    }
  }
  storage_use_azuread = true
  subscription_id     = ""
}

resource "azurerm_resource_group" "el_rg" {
  name     = "${var.org}-rg-${var.env}"
  location = var.location
}

resource "azurerm_cognitive_account" "ca_speech" {
  name                = "${var.org}-ca-speech-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  kind                = "SpeechServices"

  sku_name = "S0"
}

resource "azurerm_cognitive_account" "ca_openai" {
  name                = "${var.org}-ca-openai-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  kind                = "OpenAI"
  sku_name            = "S0"
}

resource "azurerm_cognitive_deployment" "ca_openai_deployment" {
  name                 = "${var.org}-cd-openai-${var.env}"
  cognitive_account_id = azurerm_cognitive_account.ca_openai.id
  model {
    format  = "OpenAI"
    name    = "gpt-4o-realtime-preview"
    version = "2024-10-01"
  }

  sku {
    name = "GlobalStandard"
  }
}