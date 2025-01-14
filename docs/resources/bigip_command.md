---
layout: "bigip"
page_title: "BIG-IP: bigip_command"
sidebar_current: "docs-bigip-resource-device-x"
description: |-
    Provides details about bigip device
---

# bigip_command

`bigip_command` Run TMSH commands on F5 devices

This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device

## Example Usage

```hcl
resource "bigip_command" "test-command" {
  commands = ["show sys version"]
}

# Create ltm node
resource "bigip_command" "test-command" {
  commands = ["create ltm node 10.10.10.70"]
}

# Destroy ltm node
resource "bigip_command" "test-command" {
  when     = "destroy"
  commands = ["delete ltm node 10.10.10.70"]
}
```

It is also possible to send Bash commands however care is needed with quoting:

```hcl
resource "bigip_command" "byol-license" {
  commands = [
    "bash -c \"cp /config/bigip.license /config/bigip.license.bak.$(date +%s)\"",
    "bash -c \"echo ${var.license_file_contents_base64} | base64 --decode > /config/bigip.license\"",
    "bash -c \"reloadlic\""
  ]
}
```

Note that use of single quotes is not supported, thus this will not work:

```hcl
resource "bigip_command" "hello-world" {
  commands = ["bash -c 'echo hello world'"]
}
```

## Argument Reference

* `commands` - (Required) The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
* `when` - (Optional, possible values: `apply` or `destroy`) default value will be `apply`,can be set to `destroy` for terraform destroy call.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `command_result` - The resulting output from the `commands` executed.