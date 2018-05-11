# Synse AMT Plugin
A plugin for [Synse Server][synse-server] to communicate with Intel AMT enabled machines.

### Devices
This plugin currently supports 2 types of AMT functionality exposed as devices:
  - AMT Power
    - Read enabled
    - Write enabled
  - AMT Boot Target
    - Read disabled
    - Write enabled

#### Write Action Payloads
##### AMT Power:
    {'action': 'state', 'raw': <power_state>}

Supported power states are `on`, `off`, `cycle`

##### AMT Boot Target
    {'action': 'target', 'raw': <boot_target>}

Supported boot targets are `pxe`, `hd`, `cd`

#### Device Instance Configuration
The following fields are required for both AMT Power and AMT Boot Target device instance configurations
```yml
- ip: IP Address for contacting AMT machine, port not included
  password: AMT password
```
#### Building and Running
The plugin can be run independently as a go package via:
`$ go build -o plugin`

For information and examples on how to deploy a plugin with Synse Server, see the [Plugin SDK Documentation][sdk-documentation]

[synse-server]: https://github.com/vapor-ware/synse-server
[sdk-documentation]:http://synse-sdk.readthedocs.io/en/latest/user/tutorial.html#build-and-run-the-plugin
