resource "azurerm_public_ip" "main" {
  count = var.deploy_virtual_machine && var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-pip-${var.env}"
  resource_group_name = azurerm_resource_group.el_rg_net[0].name
  location            = azurerm_resource_group.el_rg_net[0].location
  allocation_method   = "Static"
}

resource "azurerm_network_interface" "main" {
  count = var.deploy_virtual_machine && var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-nic-${var.env}"
  location            = azurerm_resource_group.el_rg_net[0].location
  resource_group_name = azurerm_resource_group.el_rg_net[0].name

  ip_configuration {
    name                          = "ipconfig1"
    subnet_id                     = azurerm_subnet.admin[0].id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.main[0].id
  }
}

resource "azurerm_windows_virtual_machine" "main" {
  count = var.deploy_virtual_machine && var.public_network_access_enabled ? 0 : 1

  name                = "${var.org}-vm-${var.env}"
  location            = azurerm_resource_group.el_rg.location
  resource_group_name = azurerm_resource_group.el_rg.name
  size                = "Standard_F2"
  admin_username      = "adminuser"
  admin_password      = "P@$$w0rd1234!"
  network_interface_ids = [
    azurerm_network_interface.main[0].id,
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2016-Datacenter"
    version   = "latest"
  }
}

resource "azurerm_dev_test_global_vm_shutdown_schedule" "main" {
  count = var.deploy_virtual_machine && var.public_network_access_enabled == false && var.enable_automatic_vm_shutdown ? 1 : 0

  virtual_machine_id = azurerm_windows_virtual_machine.main[0].id
  location           = azurerm_resource_group.el_rg.location
  enabled            = true

  daily_recurrence_time = "2200"
  timezone              = "Eastern Standard Time"

  notification_settings {
    enabled = false
  }
}