[![Build Status](https://build.vio.sh/buildStatus/icon?job=vapor-ware/synse-amt-plugin/master)](https://build.vio.sh/blue/organizations/jenkins/vapor-ware%2Fsynse-amt-plugin/activity)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin?ref=badge_shield)

# Synse AMT Plugin

A plugin for [Synse Server][synse-server] used to communicate with Intel AMT enabled machines.

## Getting Started

### Getting

You can install the AMT plugin via a [release](https://github.com/vapor-ware/synse-amt-plugin/releases)
binary or via Docker image

```
docker pull vaporio/amt-plugin
```

If you wish to use a development build, fork and clone the repo and build the plugin
from source.

### Running

The AMT plugin requires device configurations for the AMT-enabled servers that it will
communicate with in order for it to run. As such, running the plugin without additional configuration
will cause it to fail.

A simple example of what device configurations may look like can be found in the
[config/device](config/device) directory. Once you have your plugin configurations defined,
you can update the [compose file](docker-compose.yaml) to mount them into the plugin container
and run it with:

```bash
docker-compose up -d
```

## Plugin Configuration

Plugin and device configuration are described in detail in the [SDK Documentation][sdk-docs].

There is an additional config scheme specific to this plugin for the contents of a configured
device's `data` field. Device `data` may be specified in two places (the prototype config and
the instance config sections). The data scheme describes the resulting unified config from
both sources.

An example:

```yaml
devices:
- type: boot_target
  instances:
  - info: Server Boot Target
    data:
      ip: "127.0.0.1"
      password: "guest"
```

| Field      | Required | Type   | Description                                              |
| ---------- | -------- | ------ | -------------------------------------------------------- |
| `ip`       | yes      | string | The hostname/ip of the AMT-enabled server to connect to. |
| `password` | yes      | string | The AMT password for the server.                         |

### Reading Outputs

Outputs are referenced by name. A single device may have more than one instance
of an output type. A value of `-` in the table below indicates that there is no value
set for that field. The *built-in* section describes outputs this plugin uses which
are [built-in to the SDK](https://synse.readthedocs.io/en/latest/sdk/concepts/reading_outputs/#built-ins).

**Built-in**

| Name  | Description                       | Unit  | Type    | Precision |
| ----- | --------------------------------- | :---: | ------- | :-------: |
| state | The power state of an AMT device. | -     | `state` | -         |

### Device Handlers

Device Handlers are referenced by name.

| Name        | Description                                | Outputs | Read  | Write | Bulk Read | Listen |
| ----------- | ------------------------------------------ | ------- | :---: | :---: | :-------: | :----: |
| boot_target | A handler for setting server boot target.  | `-`     | ✗     | ✓     | ✗         | ✗      |
| power       | A handler for managing server power state. | `state` | ✓     | ✓     | ✗         | ✗      |

### Write Values

This plugin supports the following values when writing to a device via a handler.

| Handler     | Write Action  | Write Data           | Description |
| ----------- | :-----------: | :------------------: | ----------- |
| boot_target | `target`      | `pxe`, `hd`, `cd`    | The boot target to set. |
| power       | `state`       | `on`, `off`, `cycle` | The minimum bound for readings to be generated within. |

## Compatibility

Below is a table describing the compatibility of plugin versions with Synse platform versions.

|             | Synse v2 | Synse v3 |
| ----------- | -------- | -------- |
| plugin v1.x | ✓        | ✗        |
| plugin v2.x | ✗        | ✓        |

## Troubleshooting

### Debugging

The plugin can be run in debug mode for additional logging. This is done by:

- Setting the `debug` option  to `true` in the plugin configuration YAML ([config.yml](config.yml))

  ```yaml
  debug: true
  ```

- Passing the `--debug` flag when running the binary/image

  ```
  docker run vaporio/amt-plugin --debug
  ```

- Running the image with the `PLUGIN_DEBUG` environment variable set to `true`

  ```
  docker run -e PLUGIN_DEBUG=true vaporio/amt-plugin
  ```

## Contributing / Reporting

If you experience a bug, would like to ask a question, or request a feature, open a
[new issue](https://github.com/vapor-ware/synse-amt-plugin/issues) and provide as much
context as possible. All contributions, questions, and feedback are welcomed and appreciated.

## License

The Synse AMT Plugin is licensed under GPLv3. See [LICENSE](LICENSE) for more info.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-amt-plugin?ref=badge_large)

[synse-server]: https://github.com/vapor-ware/synse-server
[sdk-docs]: https://synse.readthedocs.io/en/latest/sdk/intro/
