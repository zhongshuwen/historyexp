# 中数文版权保护联盟链历史
[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/zhongshuwen/dfuse-eosio)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Getting started

If it's the first time you boot a `nodeos` node, please review
https://developers.eos.io/welcome/latest/tutorials/bios-boot-sequence
and make sure you get a grasp of what this blockchain node is capable.

The default settings of `zswlishi` allow you to quickly bootstrap a working
development chain by also managing the block producing node for you.

## Requirements

### Operating System
* Linux or macOS (no Windows support for now, should work through WSL2 (**untested**))

### dfuse Instrumented nodeos (deep-mind)
* See [DEPENDENCIES.md](DEPENDENCIES.md) for instructions on how to get an instrumented `nodeos` binary.

## Installing

### From a pre-built release

* Download a tarball from the [GitHub Releases Tab](https://github.com/zhongshuwen/historyexp/releases).
* Put the binary `zswlishi` in your `PATH`.

### From source

Build requirements:
* `Git`
* `Go` 1.14 or higher ([installation](https://golang.org/doc/install#install))
* `yarn` 1.15 or higher ([installation](https://classic.yarnpkg.com/en/docs/install))

```bash
./scripts/build.sh
```

This will install the binary in your `$GOPATH/bin` folder (normally `$HOME/go/bin`). Make sure this folder is in your `PATH` env variable. If it's missing, take a look at [TROUBLESHOOTING.md](TROUBLESHOOTING.md#gopathbin-folder-missing-from-path-env-variable).

_**Note** -- If you're getting yarn dependency warnings while running the `yarn install && yarn build` commands above, you can normally safely ignore those and move forward with the installation. If you're getting an error while installing and/or compiling, see [TROUBLESHOOTING.md](./TROUBLESHOOTING.md#installing--compiling-error)._

## Creating a new local chain with `zswlishi`

### 1. Initialize

Initialize a few configuration files in your working directory (`dfuse.yaml`, `mindreader/config.ini`, ...)

```
zswlishi init
```

Answer `y` (yes) when being prompted for `Do you want ZSWLiShi to run a producing node for you?`. If you intend is to sync an existing chain, follow [Syncing an existing chain with `zswlishi`](#syncing-an-existing-chain-with-zswlishi) instead.

### 2. Boot

Optionally, you can also copy over a boot sequence to have dfuse bootstraps your chain with accounts + system contracts to have a chain ready for development in a matter of seconds:

```
wget -O bootseq.yaml https://raw.githubusercontent.com/zhongshuwen/historyexp/develop/devel/standard/bootseq.yaml
```

When you're ready, boot your instance with:

```
zswlishi start
```

A successful start will list the launching applications as well as the graphical interfaces with their relevant links:

```
Dashboard:        http://localhost:8081

Explorer & APIs:  http://localhost:8080
GraphiQL:         http://localhost:8080/graphiql
```

In this mode, two nodeos instances will now be running on your machine, a block producer node and a mindreader node, and the dfuse services should be ready in a couple seconds.

## Syncing an existing chain with `zswlishi`

If you chose to sync to an existing chain, only the mindreader node will launch. It may take a while for the initial sync depending on the size of the chain and the services may generate various error logs until it catches up (more options for quickly syncing with an existing chain will be proposed in upcoming releases).

* See [Syncing a chain partially](./PARTIAL_SYNC.md)
* See the following issue about the complexity of [syncing a large chain](https://github.com/zhongshuwen/historyexp/issues/26)

You should also take a look at our Docs:
* [System Admin Guide](https://docs.dfuse.io/eosio/admin-guide/)
* [Large Chains Preparation](https://docs.dfuse.io/eosio/admin-guide/large-chains-preparation/)

## Filtering

* See [Filtering](https://docs.dfuse.io/eosio/admin-guide/filtering/)

## Overview - Repository Map

The glue:
* The [zswlishi](./cmd/zswlishi) binary.
* The [launcher](./launcher) which starts all the internal services

The EOSIO-specific services:
* [abicodec](./abicodec): ABI encoding and decoding service
* [statedb](./statedb): The **dfuse State** database for EOSIO, with all tables at any block height
* [kvdb-loader](./kvdb-loader): Service that loads data into the `kvdb` storage
* [dashboard](./dashboard): Server and UI for the **ZSWLiShi** dashboard.
* [eosq](./eosq): The famous https://eosq.app block explorer
* [eosws](./eosws): The REST, Websocket service, push guarantee, chain pass-through service.

dfuse Products's EOSIO-specific hooks and plugins:
* [search plugin](./search), object mappers, EOSIO-specific indexer, results mapper (along with the [search client](./search-client).
* [dgraphql resolvers](./dgraphql), with all data schemas for EOSIO
* [blockmeta plugin](./blockmeta), for EOS-specific `kvdb` bridge.

## Logging

See [Logging](./LOGGING.md)

## Troubleshooting

See [Troubleshooting](./TROUBLESHOOTING.md)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our Code of Conduct & processes for submitting pull requests, and [CONVENTIONS.md](CONVENTIONS.md) for our coding conventions.

## License

[Apache 2.0](LICENSE)

## References

- [dfuse Docs](https://docs.dfuse.io)
- [dfuse on Telegram](https://t.me/dfuseAPI) - Community & Team Support
