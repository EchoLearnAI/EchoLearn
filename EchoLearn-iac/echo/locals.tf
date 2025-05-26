locals {
  # List of regions available per landing zone
  available_regions = {
    echo = ["eu", "us"]
  }
  # Return a list of available landing zone regions
  landing_zone_regions = toset(flatten([
    for zone, regions in local.available_regions : [for region in regions : "${zone}-${region}"]
  ]))
  tags = merge(var.tags, data.azurerm_subscription.current.tags)
}