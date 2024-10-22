terraform {
  backend "azurerm" {
    storage_account_name = "elsadjz1o7"
    container_name       = "tfstate"
    key                  = "echotf/terraform.tfstate"
    use_azuread_auth     = true
    client_id            = "client_id"
    client_secret        = "client_secret"
    tenant_id            = "tenant_id"
    subscription_id      = "subscription_id"
  }
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>4.6.0"
    }
  }
}

provider "azurerm" {
  features {} 
  storage_use_azuread = true
  subscription_id     = "subscription_id"
}

resource "azurerm_resource_group" "rl_rg" {
  name     = "${var.org}-rg-${var.env}"
  location = var.location
}

resource "azurerm_cognitive_account" "example" {
  name                = "${var.org}-ca-${var.env}"
  location            = azurerm_resource_group.rl_rg.location
  resource_group_name = azurerm_resource_group.rl_rg.name
  kind                = "SpeechServices"

  sku_name = "S0"
}