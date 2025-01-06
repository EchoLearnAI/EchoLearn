module "naming" {
  source      = "../../modules/naming"
  name        = var.name
  environment = var.environment
  increment   = var.increment
  location    = var.location
}