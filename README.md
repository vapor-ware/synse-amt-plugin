[![CircleCI](https://circleci.com/gh/vapor-ware/synse-amt-plugin.svg?style=shield)](https://circleci.com/gh/vapor-ware/synse-amt-plugin)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin?ref=badge_shield)

# Synse AMT Plugin
A plugin for [Synse Server][synse-server] used to communicate with Intel AMT enabled machines.

## Plugin Support
### Outputs
Outputs should be referenced by name. A single device can have more than one instance
of an output type. A value of `-` in the table below indicates that there is no value
set for that field.


| Name | Description | Unit | Precision | Scaling Factor |
| ---- | ----------- | ---- | --------- | -------------- |
| power.state | An output type for power state readings (on/off). | - | - | - |
| boot.target | An output type for boot target settings. | - | - | - |


### Device Handlers
Device Handlers should be referenced by name.

| Name | Description | Read | Write | Bulk Read |
| ---- | ----------- | ---- | ----- | --------- |
| boot_target | A handler for setting an AMT device's boot target. | ✗ | ✓ | ✗ |
| power | A handler for power control of an AMT device. | ✓ | ✓ | ✗ |


### Write Values
This plugin supports the following values when writing to a device via a handler.

| Handler | Write Action | Write Data |
| ------- | ------------ | ---------- |
| boot_target | `target` | `pxe`, `hd`, `cd` |
| chassis.power | `state` | `on`, `off`, `cycle` |


## Getting Started

### Getting the Plugin
You can get the Synse AMT plugin either by cloning this repo, setting up the project dependencies,
and building the binary or docker image:
```bash
# Setup the project
$ make setup

# Build the binary
$ make build

# Build the docker image
$ make docker
```

You can also use a pre-built docker image from [DockerHub][plugin-dockerhub]
```bash
$ docker pull vaporio/amt-plugin
```

Or a pre-built binary from the latest [release][plugin-release].

### Running the Plugin
If you are using the plugin binary:
```bash
# The name of the plugin binary may differ depending on whether it is built
# locally or a pre-built binary is used.
$ ./plugin
```

If you are using the Docker image:
```bash
$ docker run vaporio/amt-plugin
```

In either case, the plugin should run, but you should not see any devices configured,
and you should see errors in the logs saying that various configurations were not found.
See the next section for how to configure your plugin.

For information and examples on how to deploy a plugin with Synse Server, see the
[Plugin SDK Documentation][sdk-documentation]

## Configuring the Plugin for your deployment
Plugin and device configuration are described in detail in the [SDK Configuration Documentation][sdk-config-docs].

For your deployment, you will need to provide your own device config, examples of
which can be found in [config/device](config/device).

### device config
The device configuration for the AMT plugin is fairly standard. It requires devices to
have:
- `ip`: The IP address/hostname for the AMT-enabled machine. (port not included)
- `password`: The password for AMT.

An example of a power device config for an AMT-enabled machine at 10.1.2.3 with password
ADMIN is as follows:

```yaml
version: 1.0
locations:
  - name: r1vec
    rack: 
      name: rack-1
    board:
      name: vec
devices:
  - name: power
    outputs:
      - type: power
    instances:
      - info: System Power
        location: r1vec
        data:
          ip: 10.1.2.3
          password: ADMIN
```

Once you have your own config, you can either mount it into the container to the default location
at `/etc/synse/plugin/config/device`,  or mount it anywhere in the container, e.g.
`/tmp/cfg/<filename>.yml`, and specify that path in the device instance config override environment
variable, `PLUGIN_DEVICE_CONFIG=/tmp/cfg`.

## Feedback
Feedback for this plugin, or any component of the Synse ecosystem, is greatly appreciated!
If you experience any issues, find the documentation unclear, have requests for features,
or just have questions about it, we'd love to know. Feel free to open an issue for any
feedback you may have.

## Contributing
We welcome contributions to the project. The project maintainers actively manage the issues
and pull requests. If you choose to contribute, we ask that you either comment on an existing
issue or open a new one.

The Synse AMT Plugin, and all other components of the Synse ecosystem, is released under the
[GPL-3.0](LICENSE) license.


[synse-server]: https://github.com/vapor-ware/synse-server
[plugin-dockerhub]: https://hub.docker.com/r/vaporio/amt-plugin
[plugin-release]: https://github.com/vapor-ware/synse-amt-plugin/releases
[sdk-config-docs]: http://synse-sdk.readthedocs.io/en/latest/user/configuration.html
[sdk-documentation]:http://synse-sdk.readthedocs.io/en/latest/user/tutorial.html#build-and-run-the-plugin


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin?ref=badge_large)