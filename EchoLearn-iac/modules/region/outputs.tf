output "region" {
  value       = local.regions[local.region_slug]
  description = "Azure region in standard format"
}

output "region_short" {
  value       = local.short_names[local.region_slug]
  description = "Azure region in short format for resource naming purpose"
}
