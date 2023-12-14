# testnet-app

`testnet-app` is an app developed for [Ignite CLI](https://github.com/ignite/cli).

The Ignite App adds `ignite testnet` commands that allow launching new Cosmos testnet blockchains by interacting with the Ignite Chain to coordinate with validators.

## Install

1. **Install Ignite:**
```bash
curl https://get.ignite.com/cli! | bash
```


2. **Install Testnet App:**
```bash
git clone https://github.com/toschdev/testnet-app && cd testnet-app
ignite app install -g $(pwd)
```

## Launching a testnet

### **Workflow for Publishing a Chain on Ignite Testnet**

1. **Coordinator: Publish Chain Information**
    - Command: **`ignite testnet manage start <repository URL>`**
    - Purpose: This step involves the coordinator publishing details about the chain to be launched on the Ignite blockchain. The URL points to a repository with a Cosmos SDK chain.
    - Output: A launch ID
1. **Validators: Setup Nodes and Request to Join**
   Do these steps on a dedicated validator node.
    - Command: **`ignite testnet manage setup <launch ID>`**
    - Purpose: Validators setup their nodes and request to join the testnet as validators.
    - Command: **`ignite testnet manage join <launch ID> --amount <stake amount>`**
    - Purpose: Validators specify the amount of stake they are committing.
2. **Coordinator: List Validator Requests**
    - Command: **`ignite testnet request list <launch ID>`**
    - Purpose: The coordinator lists all validator requests to review and approve.
3. **Coordinator: Approve Validator Requests**
    - Command: **`ignite testnet request approve <launch ID> <request IDs>`**
    - Purpose: The coordinator approves the validator requests essential for the validator set.
4. **Coordinator: Announce Chain Launch Readiness**
    - Command: **`ignite testnet manage announce-ready <launch ID>`**
    - Purpose: Once the necessary validators are approved, this command signals that the chain is ready for launch.
5. **Validators: Prepare Nodes for Launch**
    - Command: **`ignite testnet manage prepare-launch <launch ID>`**
    - Purpose: Validators prepare their nodes for the launch, following the instructions provided by the output of this command.
6. **Validators: Launch Nodes**
    - Action: Validators use the provided command (e.g., **`exampled --home ~/.example`**) to launch their nodes.
    - Purpose: When enough validators have launched their nodes, the blockchain becomes live.

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