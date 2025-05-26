terraform {
  required_version = ">= 1.10.3"

  required_providers {
    azapi = {
      source  = "azure/azapi"
      version = "2.2.0"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.15.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.3"
    }
  }

  backend "azurerm" {}
}

provider "azapi" {}

provider "azurerm" {
  storage_use_azuread = true
  subscription_id     = var.subscription_id
  features {}
}
