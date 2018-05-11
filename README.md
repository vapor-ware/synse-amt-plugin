# Synse AMT Plugin
A [Synse Server][synse-server] plugin for communicating with Intel AMT enabled machines.

This plugin supports two types of devices:
- boot_target
- power

They provide functionality for:

**Reads**:
- Power status *(on/off)*

**Writes**:
- Power control *(on/off/cycle)*
- Boot target selection *(pxe/hd/cd)*

## Getting Started

### Getting the Plugin
You can get the Synse AMT plugin either by cloning this repo and running one of:
```console
# Build the AMT plugin binary locally
$ make setup build

# Build the AMT plugin Docker image locally
$ make docker
```

You can also use a pre-built docker image from [DockerHub][plugin-dockerhub]
```console
$ docker pull vaporio/amt-plugin
```

Or a pre-built binary from the latest [release][plugin-release].

### Running the Plugin
If you are using the plugin binary:
```console
# The name of the plugin binary may differ depending on whether it is built
# locally or a pre-built binary is used.
$ ./plugin
```

If you are using the Docker image:
```console
$ docker run vaporio/amt-plugin
```

In either case, the plugin should run, but you should not see any devices configured,
and you should see errors in the logs saying that various configurations were not found.
See the next section for how to configure your plugin.

For information and examples on how to deploy a plugin with Synse Server, see the
[Plugin SDK Documentation][sdk-documentation]

## Configuring the Plugin for your deployment
Plugins have three different types of configurations - these are all described in detail
in the [SDK Configuration Documentation][sdk-config-docs].

For your deployment, you will need to provide your own device instance config, examples of
which can be found in [config/device](config/device).

### instance config
The instance configuration for the AMT plugin is fairly standard. It requires devices to
have:
- `ip`: The IP address/hostname for the AMT-enabled machine. (port not included)
- `password`: The password for AMT.

An example of a power device config for an AMT-enabled machine at 10.1.2.3 with password
ADMIN is as follows:

```yaml
version: 1.0
locations:
  r1vec:
    rack: rack-1
    board: vec
devices:
  - type: power
    model: amt-power
    instances:
      - ip: "10.1.2.3"
        password: "ADMIN"
        location: r1vec
        info: vec_power
```

Note that `location` and `info` are reserved for internal usage. The `location` is required,
but `info` is optional.

Once you have your own config, you can either mount it into the container to the default location
at `/etc/synse/config/device/<filename>.yml`,  or mount it anywhere in the container, e.g.
`/tmp/cfg/<filename>.yml`, and specify that path in the device instance config override environment
variable, `PLUGIN_DEVICE_PATH=/tmp/cfg`.

## Feedback
Feedback for this plugin, or any component of the Synse ecosystem, is greatly appreciated!
If you experience any issues, find the documentation unclear, have requests for features,
or just have questions about it, we'd love to know. Feel free to open an issue for any
feedback you may have.

## Contributing
We welcome contributions to the project. The project maintainers actively manage the issues
and pull requests. If you choose to contribute, we ask that you either comment on an existing

The Synse AMT Plugin, and all other components of the Synse ecosystem, is released under the
[GPL-3.0](LICENSE) license.


[synse-server]: https://github.com/vapor-ware/synse-server
[plugin-dockerhub]: https://hub.docker.com/r/vaporio/amt-plugin
[plugin-release]: https://github.com/vapor-ware/synse-amt-plugin/releases
[sdk-config-docs]: http://synse-sdk.readthedocs.io/en/latest/user/configuration.html
[sdk-documentation]:http://synse-sdk.readthedocs.io/en/latest/user/tutorial.html#build-and-run-the-plugin
