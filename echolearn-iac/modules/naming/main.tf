# Combine variables to form a standardized name
locals {
  resource_name = "${var.location}-${var.project}-${var.environment}"
}
