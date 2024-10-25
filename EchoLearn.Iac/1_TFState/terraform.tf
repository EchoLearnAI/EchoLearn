terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>4.6.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "~>3.0.2"
    }
  }
}