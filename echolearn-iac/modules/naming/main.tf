locals {
  projects = {
    el = "Echo Learn"
  }

  project = lookup(zipmap(values(local.projects), keys(local.projects)), var.project)
}

module "region" {
  source = "../region"
  count  = var.location != "" ? 1 : 0

  name = var.location
}

module "az_naming" {
  source  = "Azure/naming/azurerm"
  version = "0.4.2"
  suffix = compact([
    local.segment,
    local.project,
    var.product,
    var.name,
    var.environment,
    var.location != "" ? module.region[0].region_short : "",
    var.increment != null ? format("%02d", var.increment) : ""
  ])
}

# for resources that are limited to 24 characters
module "az_naming_limited" {
  source  = "Azure/naming/azurerm"
  version = "0.4.2"
  suffix = [join("", compact([
    substr(local.segment, 0, 2),
    substr(local.project, 0, 2),
    var.product != "" ? substr(var.product, 0, 3) : "",
    substr(var.name, 0, 19 - (var.product != "" ? 3 : 0) - (var.location != "" ? 2 : 0) - (var.increment != null ? 2 : 0)),
    substr(var.environment, 0, 1),
    var.location != "" ? substr(module.region[0].region_short, 0, 2) : "",
    var.increment != null ? format("%02d", var.increment) : ""
  ]))]
}

locals {
  spn_suffix = compact([
    local.segment,
    local.project,
    var.product,
    var.environment,
    var.name,
    var.increment != null ? format("%02d", var.increment) : ""
  ])
  spn = "spn-${join("-", local.spn_suffix)}"
}
