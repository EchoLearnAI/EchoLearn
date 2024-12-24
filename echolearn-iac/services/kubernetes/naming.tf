module "naming" {
  source = "../../modules/naming"
  location = var.location
  project = var.project
  environment = var.environment
}