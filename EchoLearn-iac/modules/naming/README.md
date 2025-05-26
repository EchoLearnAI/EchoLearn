# Naming module

This module intend to provide a consistent naming convention across Azure
resources.

## Getting started

```hcl
module "naming" {
  source      = "../modules/naming"
  name        = "ef"
  environment = "nonprd"
  increment   = 1
  location    = "eastus"
}

output "spn" {
  value = module.naming.service_principal
}

output "k8s" {
  value = module.naming.kubernetes_cluster
}
```

<!-- BEGIN_TF_DOCS -->
## Inputs

| Name | Description | Type |
|------|-------------|------|
| <a name="input_environment"></a> [environment](#input\_environment) | The environment in which the resource is deployed, e.g. nonprd, prd. | `string` |
| <a name="input_name"></a> [name](#input\_name) | Name of the application or service the resource is deployed for. | `string` |
| <a name="input_increment"></a> [increment](#input\_increment) | (Optional) A 2-digit increment number | `number` |
| <a name="input_location"></a> [location](#input\_location) | (Optional) The region in which the resource is deployed. | `string` |
| <a name="input_product"></a> [product](#input\_product) | (Optional) The product to which this resource belongs to. | `string` |
| <a name="input_project"></a> [project](#input\_project) | (Optional) The project to which this resource belongs to. | `string` |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_analysis_services_server"></a> [analysis\_services\_server](#output\_analysis\_services\_server) | n/a |
| <a name="output_api_management"></a> [api\_management](#output\_api\_management) | n/a |
| <a name="output_app_configuration"></a> [app\_configuration](#output\_app\_configuration) | n/a |
| <a name="output_app_service"></a> [app\_service](#output\_app\_service) | n/a |
| <a name="output_app_service_environment"></a> [app\_service\_environment](#output\_app\_service\_environment) | n/a |
| <a name="output_app_service_plan"></a> [app\_service\_plan](#output\_app\_service\_plan) | n/a |
| <a name="output_application_gateway"></a> [application\_gateway](#output\_application\_gateway) | n/a |
| <a name="output_application_insights"></a> [application\_insights](#output\_application\_insights) | n/a |
| <a name="output_application_security_group"></a> [application\_security\_group](#output\_application\_security\_group) | n/a |
| <a name="output_automation_account"></a> [automation\_account](#output\_automation\_account) | n/a |
| <a name="output_automation_certificate"></a> [automation\_certificate](#output\_automation\_certificate) | n/a |
| <a name="output_automation_credential"></a> [automation\_credential](#output\_automation\_credential) | n/a |
| <a name="output_automation_runbook"></a> [automation\_runbook](#output\_automation\_runbook) | n/a |
| <a name="output_automation_schedule"></a> [automation\_schedule](#output\_automation\_schedule) | n/a |
| <a name="output_automation_variable"></a> [automation\_variable](#output\_automation\_variable) | n/a |
| <a name="output_availability_set"></a> [availability\_set](#output\_availability\_set) | n/a |
| <a name="output_bastion_host"></a> [bastion\_host](#output\_bastion\_host) | n/a |
| <a name="output_batch_account"></a> [batch\_account](#output\_batch\_account) | n/a |
| <a name="output_batch_application"></a> [batch\_application](#output\_batch\_application) | n/a |
| <a name="output_batch_certificate"></a> [batch\_certificate](#output\_batch\_certificate) | n/a |
| <a name="output_batch_pool"></a> [batch\_pool](#output\_batch\_pool) | n/a |
| <a name="output_bot_channel_directline"></a> [bot\_channel\_directline](#output\_bot\_channel\_directline) | n/a |
| <a name="output_bot_channel_email"></a> [bot\_channel\_email](#output\_bot\_channel\_email) | n/a |
| <a name="output_bot_channel_ms_teams"></a> [bot\_channel\_ms\_teams](#output\_bot\_channel\_ms\_teams) | n/a |
| <a name="output_bot_channel_slack"></a> [bot\_channel\_slack](#output\_bot\_channel\_slack) | n/a |
| <a name="output_bot_channels_registration"></a> [bot\_channels\_registration](#output\_bot\_channels\_registration) | n/a |
| <a name="output_bot_connection"></a> [bot\_connection](#output\_bot\_connection) | n/a |
| <a name="output_bot_web_app"></a> [bot\_web\_app](#output\_bot\_web\_app) | n/a |
| <a name="output_cdn_endpoint"></a> [cdn\_endpoint](#output\_cdn\_endpoint) | n/a |
| <a name="output_cdn_profile"></a> [cdn\_profile](#output\_cdn\_profile) | n/a |
| <a name="output_cognitive_account"></a> [cognitive\_account](#output\_cognitive\_account) | n/a |
| <a name="output_container_app"></a> [container\_app](#output\_container\_app) | n/a |
| <a name="output_container_app_environment"></a> [container\_app\_environment](#output\_container\_app\_environment) | n/a |
| <a name="output_container_group"></a> [container\_group](#output\_container\_group) | n/a |
| <a name="output_container_registry"></a> [container\_registry](#output\_container\_registry) | n/a |
| <a name="output_container_registry_webhook"></a> [container\_registry\_webhook](#output\_container\_registry\_webhook) | n/a |
| <a name="output_cosmosdb_account"></a> [cosmosdb\_account](#output\_cosmosdb\_account) | n/a |
| <a name="output_cosmosdb_cassandra_cluster"></a> [cosmosdb\_cassandra\_cluster](#output\_cosmosdb\_cassandra\_cluster) | n/a |
| <a name="output_cosmosdb_cassandra_datacenter"></a> [cosmosdb\_cassandra\_datacenter](#output\_cosmosdb\_cassandra\_datacenter) | n/a |
| <a name="output_cosmosdb_postgres"></a> [cosmosdb\_postgres](#output\_cosmosdb\_postgres) | n/a |
| <a name="output_custom_provider"></a> [custom\_provider](#output\_custom\_provider) | n/a |
| <a name="output_dashboard"></a> [dashboard](#output\_dashboard) | n/a |
| <a name="output_data_factory"></a> [data\_factory](#output\_data\_factory) | n/a |
| <a name="output_data_factory_dataset_mysql"></a> [data\_factory\_dataset\_mysql](#output\_data\_factory\_dataset\_mysql) | n/a |
| <a name="output_data_factory_dataset_postgresql"></a> [data\_factory\_dataset\_postgresql](#output\_data\_factory\_dataset\_postgresql) | n/a |
| <a name="output_data_factory_dataset_sql_server_table"></a> [data\_factory\_dataset\_sql\_server\_table](#output\_data\_factory\_dataset\_sql\_server\_table) | n/a |
| <a name="output_data_factory_integration_runtime_managed"></a> [data\_factory\_integration\_runtime\_managed](#output\_data\_factory\_integration\_runtime\_managed) | n/a |
| <a name="output_data_factory_linked_service_data_lake_storage_gen2"></a> [data\_factory\_linked\_service\_data\_lake\_storage\_gen2](#output\_data\_factory\_linked\_service\_data\_lake\_storage\_gen2) | n/a |
| <a name="output_data_factory_linked_service_key_vault"></a> [data\_factory\_linked\_service\_key\_vault](#output\_data\_factory\_linked\_service\_key\_vault) | n/a |
| <a name="output_data_factory_linked_service_mysql"></a> [data\_factory\_linked\_service\_mysql](#output\_data\_factory\_linked\_service\_mysql) | n/a |
| <a name="output_data_factory_linked_service_postgresql"></a> [data\_factory\_linked\_service\_postgresql](#output\_data\_factory\_linked\_service\_postgresql) | n/a |
| <a name="output_data_factory_linked_service_sql_server"></a> [data\_factory\_linked\_service\_sql\_server](#output\_data\_factory\_linked\_service\_sql\_server) | n/a |
| <a name="output_data_factory_pipeline"></a> [data\_factory\_pipeline](#output\_data\_factory\_pipeline) | n/a |
| <a name="output_data_factory_trigger_schedule"></a> [data\_factory\_trigger\_schedule](#output\_data\_factory\_trigger\_schedule) | n/a |
| <a name="output_data_lake_analytics_account"></a> [data\_lake\_analytics\_account](#output\_data\_lake\_analytics\_account) | n/a |
| <a name="output_data_lake_analytics_firewall_rule"></a> [data\_lake\_analytics\_firewall\_rule](#output\_data\_lake\_analytics\_firewall\_rule) | n/a |
| <a name="output_data_lake_store"></a> [data\_lake\_store](#output\_data\_lake\_store) | n/a |
| <a name="output_data_lake_store_firewall_rule"></a> [data\_lake\_store\_firewall\_rule](#output\_data\_lake\_store\_firewall\_rule) | n/a |
| <a name="output_database_migration_project"></a> [database\_migration\_project](#output\_database\_migration\_project) | n/a |
| <a name="output_database_migration_service"></a> [database\_migration\_service](#output\_database\_migration\_service) | n/a |
| <a name="output_databricks_cluster"></a> [databricks\_cluster](#output\_databricks\_cluster) | n/a |
| <a name="output_databricks_high_concurrency_cluster"></a> [databricks\_high\_concurrency\_cluster](#output\_databricks\_high\_concurrency\_cluster) | n/a |
| <a name="output_databricks_standard_cluster"></a> [databricks\_standard\_cluster](#output\_databricks\_standard\_cluster) | n/a |
| <a name="output_databricks_workspace"></a> [databricks\_workspace](#output\_databricks\_workspace) | n/a |
| <a name="output_dev_test_lab"></a> [dev\_test\_lab](#output\_dev\_test\_lab) | n/a |
| <a name="output_dev_test_linux_virtual_machine"></a> [dev\_test\_linux\_virtual\_machine](#output\_dev\_test\_linux\_virtual\_machine) | n/a |
| <a name="output_dev_test_windows_virtual_machine"></a> [dev\_test\_windows\_virtual\_machine](#output\_dev\_test\_windows\_virtual\_machine) | n/a |
| <a name="output_disk_encryption_set"></a> [disk\_encryption\_set](#output\_disk\_encryption\_set) | n/a |
| <a name="output_dns_a_record"></a> [dns\_a\_record](#output\_dns\_a\_record) | n/a |
| <a name="output_dns_aaaa_record"></a> [dns\_aaaa\_record](#output\_dns\_aaaa\_record) | n/a |
| <a name="output_dns_caa_record"></a> [dns\_caa\_record](#output\_dns\_caa\_record) | n/a |
| <a name="output_dns_cname_record"></a> [dns\_cname\_record](#output\_dns\_cname\_record) | n/a |
| <a name="output_dns_mx_record"></a> [dns\_mx\_record](#output\_dns\_mx\_record) | n/a |
| <a name="output_dns_ns_record"></a> [dns\_ns\_record](#output\_dns\_ns\_record) | n/a |
| <a name="output_dns_ptr_record"></a> [dns\_ptr\_record](#output\_dns\_ptr\_record) | n/a |
| <a name="output_dns_txt_record"></a> [dns\_txt\_record](#output\_dns\_txt\_record) | n/a |
| <a name="output_dns_zone"></a> [dns\_zone](#output\_dns\_zone) | n/a |
| <a name="output_eventgrid_domain"></a> [eventgrid\_domain](#output\_eventgrid\_domain) | n/a |
| <a name="output_eventgrid_domain_topic"></a> [eventgrid\_domain\_topic](#output\_eventgrid\_domain\_topic) | n/a |
| <a name="output_eventgrid_event_subscription"></a> [eventgrid\_event\_subscription](#output\_eventgrid\_event\_subscription) | n/a |
| <a name="output_eventgrid_topic"></a> [eventgrid\_topic](#output\_eventgrid\_topic) | n/a |
| <a name="output_eventhub"></a> [eventhub](#output\_eventhub) | n/a |
| <a name="output_eventhub_authorization_rule"></a> [eventhub\_authorization\_rule](#output\_eventhub\_authorization\_rule) | n/a |
| <a name="output_eventhub_consumer_group"></a> [eventhub\_consumer\_group](#output\_eventhub\_consumer\_group) | n/a |
| <a name="output_eventhub_namespace"></a> [eventhub\_namespace](#output\_eventhub\_namespace) | n/a |
| <a name="output_eventhub_namespace_authorization_rule"></a> [eventhub\_namespace\_authorization\_rule](#output\_eventhub\_namespace\_authorization\_rule) | n/a |
| <a name="output_eventhub_namespace_disaster_recovery_config"></a> [eventhub\_namespace\_disaster\_recovery\_config](#output\_eventhub\_namespace\_disaster\_recovery\_config) | n/a |
| <a name="output_express_route_circuit"></a> [express\_route\_circuit](#output\_express\_route\_circuit) | n/a |
| <a name="output_express_route_gateway"></a> [express\_route\_gateway](#output\_express\_route\_gateway) | n/a |
| <a name="output_firewall"></a> [firewall](#output\_firewall) | n/a |
| <a name="output_firewall_application_rule_collection"></a> [firewall\_application\_rule\_collection](#output\_firewall\_application\_rule\_collection) | n/a |
| <a name="output_firewall_ip_configuration"></a> [firewall\_ip\_configuration](#output\_firewall\_ip\_configuration) | n/a |
| <a name="output_firewall_nat_rule_collection"></a> [firewall\_nat\_rule\_collection](#output\_firewall\_nat\_rule\_collection) | n/a |
| <a name="output_firewall_network_rule_collection"></a> [firewall\_network\_rule\_collection](#output\_firewall\_network\_rule\_collection) | n/a |
| <a name="output_firewall_policy"></a> [firewall\_policy](#output\_firewall\_policy) | n/a |
| <a name="output_firewall_policy_rule_collection_group"></a> [firewall\_policy\_rule\_collection\_group](#output\_firewall\_policy\_rule\_collection\_group) | n/a |
| <a name="output_frontdoor"></a> [frontdoor](#output\_frontdoor) | n/a |
| <a name="output_frontdoor_firewall_policy"></a> [frontdoor\_firewall\_policy](#output\_frontdoor\_firewall\_policy) | n/a |
| <a name="output_function_app"></a> [function\_app](#output\_function\_app) | n/a |
| <a name="output_hdinsight_hadoop_cluster"></a> [hdinsight\_hadoop\_cluster](#output\_hdinsight\_hadoop\_cluster) | n/a |
| <a name="output_hdinsight_hbase_cluster"></a> [hdinsight\_hbase\_cluster](#output\_hdinsight\_hbase\_cluster) | n/a |
| <a name="output_hdinsight_interactive_query_cluster"></a> [hdinsight\_interactive\_query\_cluster](#output\_hdinsight\_interactive\_query\_cluster) | n/a |
| <a name="output_hdinsight_kafka_cluster"></a> [hdinsight\_kafka\_cluster](#output\_hdinsight\_kafka\_cluster) | n/a |
| <a name="output_hdinsight_ml_services_cluster"></a> [hdinsight\_ml\_services\_cluster](#output\_hdinsight\_ml\_services\_cluster) | n/a |
| <a name="output_hdinsight_rserver_cluster"></a> [hdinsight\_rserver\_cluster](#output\_hdinsight\_rserver\_cluster) | n/a |
| <a name="output_hdinsight_spark_cluster"></a> [hdinsight\_spark\_cluster](#output\_hdinsight\_spark\_cluster) | n/a |
| <a name="output_hdinsight_storm_cluster"></a> [hdinsight\_storm\_cluster](#output\_hdinsight\_storm\_cluster) | n/a |
| <a name="output_image"></a> [image](#output\_image) | n/a |
| <a name="output_iotcentral_application"></a> [iotcentral\_application](#output\_iotcentral\_application) | n/a |
| <a name="output_iothub"></a> [iothub](#output\_iothub) | n/a |
| <a name="output_iothub_consumer_group"></a> [iothub\_consumer\_group](#output\_iothub\_consumer\_group) | n/a |
| <a name="output_iothub_dps"></a> [iothub\_dps](#output\_iothub\_dps) | n/a |
| <a name="output_iothub_dps_certificate"></a> [iothub\_dps\_certificate](#output\_iothub\_dps\_certificate) | n/a |
| <a name="output_key_vault"></a> [key\_vault](#output\_key\_vault) | n/a |
| <a name="output_key_vault_certificate"></a> [key\_vault\_certificate](#output\_key\_vault\_certificate) | n/a |
| <a name="output_key_vault_key"></a> [key\_vault\_key](#output\_key\_vault\_key) | n/a |
| <a name="output_key_vault_secret"></a> [key\_vault\_secret](#output\_key\_vault\_secret) | n/a |
| <a name="output_kubernetes_cluster"></a> [kubernetes\_cluster](#output\_kubernetes\_cluster) | n/a |
| <a name="output_kusto_cluster"></a> [kusto\_cluster](#output\_kusto\_cluster) | n/a |
| <a name="output_kusto_database"></a> [kusto\_database](#output\_kusto\_database) | n/a |
| <a name="output_kusto_eventhub_data_connection"></a> [kusto\_eventhub\_data\_connection](#output\_kusto\_eventhub\_data\_connection) | n/a |
| <a name="output_lb"></a> [lb](#output\_lb) | n/a |
| <a name="output_lb_nat_rule"></a> [lb\_nat\_rule](#output\_lb\_nat\_rule) | n/a |
| <a name="output_linux_virtual_machine"></a> [linux\_virtual\_machine](#output\_linux\_virtual\_machine) | n/a |
| <a name="output_linux_virtual_machine_scale_set"></a> [linux\_virtual\_machine\_scale\_set](#output\_linux\_virtual\_machine\_scale\_set) | n/a |
| <a name="output_load_test"></a> [load\_test](#output\_load\_test) | n/a |
| <a name="output_local_network_gateway"></a> [local\_network\_gateway](#output\_local\_network\_gateway) | n/a |
| <a name="output_log_analytics_workspace"></a> [log\_analytics\_workspace](#output\_log\_analytics\_workspace) | n/a |
| <a name="output_logic_app_workflow"></a> [logic\_app\_workflow](#output\_logic\_app\_workflow) | n/a |
| <a name="output_machine_learning_workspace"></a> [machine\_learning\_workspace](#output\_machine\_learning\_workspace) | n/a |
| <a name="output_managed_disk"></a> [managed\_disk](#output\_managed\_disk) | n/a |
| <a name="output_maps_account"></a> [maps\_account](#output\_maps\_account) | n/a |
| <a name="output_mariadb_database"></a> [mariadb\_database](#output\_mariadb\_database) | n/a |
| <a name="output_mariadb_firewall_rule"></a> [mariadb\_firewall\_rule](#output\_mariadb\_firewall\_rule) | n/a |
| <a name="output_mariadb_server"></a> [mariadb\_server](#output\_mariadb\_server) | n/a |
| <a name="output_mariadb_virtual_network_rule"></a> [mariadb\_virtual\_network\_rule](#output\_mariadb\_virtual\_network\_rule) | n/a |
| <a name="output_monitor_action_group"></a> [monitor\_action\_group](#output\_monitor\_action\_group) | n/a |
| <a name="output_monitor_autoscale_setting"></a> [monitor\_autoscale\_setting](#output\_monitor\_autoscale\_setting) | n/a |
| <a name="output_monitor_diagnostic_setting"></a> [monitor\_diagnostic\_setting](#output\_monitor\_diagnostic\_setting) | n/a |
| <a name="output_monitor_scheduled_query_rules_alert"></a> [monitor\_scheduled\_query\_rules\_alert](#output\_monitor\_scheduled\_query\_rules\_alert) | n/a |
| <a name="output_mssql_database"></a> [mssql\_database](#output\_mssql\_database) | n/a |
| <a name="output_mssql_elasticpool"></a> [mssql\_elasticpool](#output\_mssql\_elasticpool) | n/a |
| <a name="output_mssql_managed_instance"></a> [mssql\_managed\_instance](#output\_mssql\_managed\_instance) | n/a |
| <a name="output_mssql_server"></a> [mssql\_server](#output\_mssql\_server) | n/a |
| <a name="output_mysql_database"></a> [mysql\_database](#output\_mysql\_database) | n/a |
| <a name="output_mysql_firewall_rule"></a> [mysql\_firewall\_rule](#output\_mysql\_firewall\_rule) | n/a |
| <a name="output_mysql_server"></a> [mysql\_server](#output\_mysql\_server) | n/a |
| <a name="output_mysql_virtual_network_rule"></a> [mysql\_virtual\_network\_rule](#output\_mysql\_virtual\_network\_rule) | n/a |
| <a name="output_nat_gateway"></a> [nat\_gateway](#output\_nat\_gateway) | n/a |
| <a name="output_network_ddos_protection_plan"></a> [network\_ddos\_protection\_plan](#output\_network\_ddos\_protection\_plan) | n/a |
| <a name="output_network_interface"></a> [network\_interface](#output\_network\_interface) | n/a |
| <a name="output_network_security_group"></a> [network\_security\_group](#output\_network\_security\_group) | n/a |
| <a name="output_network_security_group_rule"></a> [network\_security\_group\_rule](#output\_network\_security\_group\_rule) | n/a |
| <a name="output_network_security_rule"></a> [network\_security\_rule](#output\_network\_security\_rule) | n/a |
| <a name="output_network_watcher"></a> [network\_watcher](#output\_network\_watcher) | n/a |
| <a name="output_notification_hub"></a> [notification\_hub](#output\_notification\_hub) | n/a |
| <a name="output_notification_hub_authorization_rule"></a> [notification\_hub\_authorization\_rule](#output\_notification\_hub\_authorization\_rule) | n/a |
| <a name="output_notification_hub_namespace"></a> [notification\_hub\_namespace](#output\_notification\_hub\_namespace) | n/a |
| <a name="output_point_to_site_vpn_gateway"></a> [point\_to\_site\_vpn\_gateway](#output\_point\_to\_site\_vpn\_gateway) | n/a |
| <a name="output_postgresql_database"></a> [postgresql\_database](#output\_postgresql\_database) | n/a |
| <a name="output_postgresql_firewall_rule"></a> [postgresql\_firewall\_rule](#output\_postgresql\_firewall\_rule) | n/a |
| <a name="output_postgresql_server"></a> [postgresql\_server](#output\_postgresql\_server) | n/a |
| <a name="output_postgresql_virtual_network_rule"></a> [postgresql\_virtual\_network\_rule](#output\_postgresql\_virtual\_network\_rule) | n/a |
| <a name="output_powerbi_embedded"></a> [powerbi\_embedded](#output\_powerbi\_embedded) | n/a |
| <a name="output_private_dns_a_record"></a> [private\_dns\_a\_record](#output\_private\_dns\_a\_record) | n/a |
| <a name="output_private_dns_aaaa_record"></a> [private\_dns\_aaaa\_record](#output\_private\_dns\_aaaa\_record) | n/a |
| <a name="output_private_dns_cname_record"></a> [private\_dns\_cname\_record](#output\_private\_dns\_cname\_record) | n/a |
| <a name="output_private_dns_mx_record"></a> [private\_dns\_mx\_record](#output\_private\_dns\_mx\_record) | n/a |
| <a name="output_private_dns_ptr_record"></a> [private\_dns\_ptr\_record](#output\_private\_dns\_ptr\_record) | n/a |
| <a name="output_private_dns_srv_record"></a> [private\_dns\_srv\_record](#output\_private\_dns\_srv\_record) | n/a |
| <a name="output_private_dns_txt_record"></a> [private\_dns\_txt\_record](#output\_private\_dns\_txt\_record) | n/a |
| <a name="output_private_dns_zone"></a> [private\_dns\_zone](#output\_private\_dns\_zone) | n/a |
| <a name="output_private_dns_zone_group"></a> [private\_dns\_zone\_group](#output\_private\_dns\_zone\_group) | n/a |
| <a name="output_private_endpoint"></a> [private\_endpoint](#output\_private\_endpoint) | n/a |
| <a name="output_private_link_service"></a> [private\_link\_service](#output\_private\_link\_service) | n/a |
| <a name="output_private_service_connection"></a> [private\_service\_connection](#output\_private\_service\_connection) | n/a |
| <a name="output_proximity_placement_group"></a> [proximity\_placement\_group](#output\_proximity\_placement\_group) | n/a |
| <a name="output_public_ip"></a> [public\_ip](#output\_public\_ip) | n/a |
| <a name="output_public_ip_prefix"></a> [public\_ip\_prefix](#output\_public\_ip\_prefix) | n/a |
| <a name="output_recovery_services_vault"></a> [recovery\_services\_vault](#output\_recovery\_services\_vault) | n/a |
| <a name="output_redis_cache"></a> [redis\_cache](#output\_redis\_cache) | n/a |
| <a name="output_redis_firewall_rule"></a> [redis\_firewall\_rule](#output\_redis\_firewall\_rule) | n/a |
| <a name="output_relay_hybrid_connection"></a> [relay\_hybrid\_connection](#output\_relay\_hybrid\_connection) | n/a |
| <a name="output_relay_namespace"></a> [relay\_namespace](#output\_relay\_namespace) | n/a |
| <a name="output_resource_group"></a> [resource\_group](#output\_resource\_group) | n/a |
| <a name="output_role_assignment"></a> [role\_assignment](#output\_role\_assignment) | n/a |
| <a name="output_role_definition"></a> [role\_definition](#output\_role\_definition) | n/a |
| <a name="output_route"></a> [route](#output\_route) | n/a |
| <a name="output_route_table"></a> [route\_table](#output\_route\_table) | n/a |
| <a name="output_search_service"></a> [search\_service](#output\_search\_service) | n/a |
| <a name="output_service_fabric_cluster"></a> [service\_fabric\_cluster](#output\_service\_fabric\_cluster) | n/a |
| <a name="output_service_principal"></a> [service\_principal](#output\_service\_principal) | n/a |
| <a name="output_servicebus_namespace"></a> [servicebus\_namespace](#output\_servicebus\_namespace) | n/a |
| <a name="output_servicebus_namespace_authorization_rule"></a> [servicebus\_namespace\_authorization\_rule](#output\_servicebus\_namespace\_authorization\_rule) | n/a |
| <a name="output_servicebus_queue"></a> [servicebus\_queue](#output\_servicebus\_queue) | n/a |
| <a name="output_servicebus_queue_authorization_rule"></a> [servicebus\_queue\_authorization\_rule](#output\_servicebus\_queue\_authorization\_rule) | n/a |
| <a name="output_servicebus_subscription"></a> [servicebus\_subscription](#output\_servicebus\_subscription) | n/a |
| <a name="output_servicebus_subscription_rule"></a> [servicebus\_subscription\_rule](#output\_servicebus\_subscription\_rule) | n/a |
| <a name="output_servicebus_topic"></a> [servicebus\_topic](#output\_servicebus\_topic) | n/a |
| <a name="output_servicebus_topic_authorization_rule"></a> [servicebus\_topic\_authorization\_rule](#output\_servicebus\_topic\_authorization\_rule) | n/a |
| <a name="output_shared_image"></a> [shared\_image](#output\_shared\_image) | n/a |
| <a name="output_shared_image_gallery"></a> [shared\_image\_gallery](#output\_shared\_image\_gallery) | n/a |
| <a name="output_signalr_service"></a> [signalr\_service](#output\_signalr\_service) | n/a |
| <a name="output_snapshots"></a> [snapshots](#output\_snapshots) | n/a |
| <a name="output_sql_elasticpool"></a> [sql\_elasticpool](#output\_sql\_elasticpool) | n/a |
| <a name="output_sql_failover_group"></a> [sql\_failover\_group](#output\_sql\_failover\_group) | n/a |
| <a name="output_sql_firewall_rule"></a> [sql\_firewall\_rule](#output\_sql\_firewall\_rule) | n/a |
| <a name="output_sql_server"></a> [sql\_server](#output\_sql\_server) | n/a |
| <a name="output_static_web_app"></a> [static\_web\_app](#output\_static\_web\_app) | n/a |
| <a name="output_storage_account"></a> [storage\_account](#output\_storage\_account) | n/a |
| <a name="output_storage_blob"></a> [storage\_blob](#output\_storage\_blob) | n/a |
| <a name="output_storage_container"></a> [storage\_container](#output\_storage\_container) | n/a |
| <a name="output_storage_data_lake_gen2_filesystem"></a> [storage\_data\_lake\_gen2\_filesystem](#output\_storage\_data\_lake\_gen2\_filesystem) | n/a |
| <a name="output_storage_queue"></a> [storage\_queue](#output\_storage\_queue) | n/a |
| <a name="output_storage_share"></a> [storage\_share](#output\_storage\_share) | n/a |
| <a name="output_storage_share_directory"></a> [storage\_share\_directory](#output\_storage\_share\_directory) | n/a |
| <a name="output_storage_table"></a> [storage\_table](#output\_storage\_table) | n/a |
| <a name="output_stream_analytics_function_javascript_udf"></a> [stream\_analytics\_function\_javascript\_udf](#output\_stream\_analytics\_function\_javascript\_udf) | n/a |
| <a name="output_stream_analytics_job"></a> [stream\_analytics\_job](#output\_stream\_analytics\_job) | n/a |
| <a name="output_stream_analytics_output_blob"></a> [stream\_analytics\_output\_blob](#output\_stream\_analytics\_output\_blob) | n/a |
| <a name="output_stream_analytics_output_eventhub"></a> [stream\_analytics\_output\_eventhub](#output\_stream\_analytics\_output\_eventhub) | n/a |
| <a name="output_stream_analytics_output_mssql"></a> [stream\_analytics\_output\_mssql](#output\_stream\_analytics\_output\_mssql) | n/a |
| <a name="output_stream_analytics_output_servicebus_queue"></a> [stream\_analytics\_output\_servicebus\_queue](#output\_stream\_analytics\_output\_servicebus\_queue) | n/a |
| <a name="output_stream_analytics_output_servicebus_topic"></a> [stream\_analytics\_output\_servicebus\_topic](#output\_stream\_analytics\_output\_servicebus\_topic) | n/a |
| <a name="output_stream_analytics_reference_input_blob"></a> [stream\_analytics\_reference\_input\_blob](#output\_stream\_analytics\_reference\_input\_blob) | n/a |
| <a name="output_stream_analytics_stream_input_blob"></a> [stream\_analytics\_stream\_input\_blob](#output\_stream\_analytics\_stream\_input\_blob) | n/a |
| <a name="output_stream_analytics_stream_input_eventhub"></a> [stream\_analytics\_stream\_input\_eventhub](#output\_stream\_analytics\_stream\_input\_eventhub) | n/a |
| <a name="output_stream_analytics_stream_input_iothub"></a> [stream\_analytics\_stream\_input\_iothub](#output\_stream\_analytics\_stream\_input\_iothub) | n/a |
| <a name="output_subnet"></a> [subnet](#output\_subnet) | n/a |
| <a name="output_template_deployment"></a> [template\_deployment](#output\_template\_deployment) | n/a |
| <a name="output_traffic_manager_profile"></a> [traffic\_manager\_profile](#output\_traffic\_manager\_profile) | n/a |
| <a name="output_user_assigned_identity"></a> [user\_assigned\_identity](#output\_user\_assigned\_identity) | n/a |
| <a name="output_virtual_machine"></a> [virtual\_machine](#output\_virtual\_machine) | n/a |
| <a name="output_virtual_machine_extension"></a> [virtual\_machine\_extension](#output\_virtual\_machine\_extension) | n/a |
| <a name="output_virtual_machine_scale_set"></a> [virtual\_machine\_scale\_set](#output\_virtual\_machine\_scale\_set) | n/a |
| <a name="output_virtual_machine_scale_set_extension"></a> [virtual\_machine\_scale\_set\_extension](#output\_virtual\_machine\_scale\_set\_extension) | n/a |
| <a name="output_virtual_network"></a> [virtual\_network](#output\_virtual\_network) | n/a |
| <a name="output_virtual_network_gateway"></a> [virtual\_network\_gateway](#output\_virtual\_network\_gateway) | n/a |
| <a name="output_virtual_network_gateway_connection"></a> [virtual\_network\_gateway\_connection](#output\_virtual\_network\_gateway\_connection) | n/a |
| <a name="output_virtual_network_peering"></a> [virtual\_network\_peering](#output\_virtual\_network\_peering) | n/a |
| <a name="output_virtual_wan"></a> [virtual\_wan](#output\_virtual\_wan) | n/a |
| <a name="output_windows_virtual_machine"></a> [windows\_virtual\_machine](#output\_windows\_virtual\_machine) | n/a |
| <a name="output_windows_virtual_machine_scale_set"></a> [windows\_virtual\_machine\_scale\_set](#output\_windows\_virtual\_machine\_scale\_set) | n/a |
<!-- END_TF_DOCS -->
