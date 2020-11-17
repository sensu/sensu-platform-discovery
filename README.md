[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/sensu-platform-discovery)
![Go Test](https://github.com/sensu/sensu-platform-discovery/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/sensu/sensu-platform-discovery/workflows/goreleaser/badge.svg)

# Sensu Platform Discovery

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

Discover system platform information and output a list of agent
subscriptions. This plugin can be used in combination with the [Sensu
Entity Manager handler](https://github.com/sensu/sensu-entity-manager)
to automate Sensu Go agent subscription management.

## Usage examples

```
$ sensu-platform-discovery -h
Discover system platform information and output a list of agent subscriptions.

Usage:
  sensu-platform-discovery [flags]
  sensu-platform-discovery [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --get-cloud-provider   Determine the cloud provider and include it in subscriptions.
  -h, --help                 help for sensu-platform-discovery

Use "sensu-platform-discovery [command] --help" for more information about a command.
```

```
$ sensu-platform-discovery
linux
debian
```

```
$ sensu-platform-discovery --get-cloud-provider
aws
linux
debian
```

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add sensu/sensu-platform-discovery
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/sensu/sensu-platform-discovery).

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-platform-discovery
  namespace: default
spec:
  command: sensu-platform-discovery --get-cloud-provider
  subscriptions:
  - discovery
  runtime_assets:
  - sensu/sensu-platform-discovery
  interval: 60
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-platform-discovery repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
