# testnet-app

`testnet-app` is an app developed for [Ignite CLI](https://github.com/ignite/cli).

The Ignite App adds `ignite testnet` commands that allow launching new Cosmos testnet blockchains by interacting with the Ignite Chain to coordinate with validators.

[**Check out our documentation for launching chains with the commands**](https://docs.ignite.com/nightly/network/introduction)

## Developer instruction

- Clone this repo locally
- Run `ignite app install -g $(pwd)` to add the app to global config
- The `ignite testnet` command is now available with the local version of the app

Then repeat the following loop:

- Hack on the plugin code
- Rerun `ignite testnet` to automatically recompile the app and test

If something corrupted your plugins, remove them from:

```bash
nano ~/.ignite/apps/igniteapps.yml # remove line with the app
```