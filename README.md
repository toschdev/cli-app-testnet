"Make it as easy as possible to start a distributed testnet."

# ignite testnet

`ignite testnet` is an app developed for [Ignite CLI](https://github.com/ignite/cli).

It is designed to facilitate the launching and management of new Cosmos testnet blockchains. This application streamlines the process of coordinating with validators and managing testnet deployments by leveraging the robust features of the Ignite Chain. With its specialized ignite testnet commands, users can efficiently initiate, configure, and control various aspects of their Cosmos testnets.

## Install

1. **Install Ignite:**
```bash
curl https://get.ignite.com/cli! | bash
```


2. **Install Testnet App:**
```bash
git clone https://github.com/toschdev/ignite-testnet && cd ignite-testnet
ignite app install -g $(pwd)
```

## Launching a testnet

For a complete list of the workflow, check [workflow.md](./workflow.md).

### **Workflow for Publishing a Chain on Ignite Testnet**

Learn about [coordinator](./coordinator.md) and [validator](./validator.md) roles.

1. **Coordinator: Publish Chain Information**
    - Command: **`ignite testnet coordinator manage start <repository URL>  --chain-id example-1`**
    - Purpose: This step involves the coordinator publishing details about the chain to be launched on the Ignite blockchain. The URL points to a repository with a Cosmos SDK chain.
    - Output: A launch ID
1. **Validators: Setup Nodes and Request to Join**
   
   Do these steps on a dedicated validator node.
    - Command: **`ignite testnet node setup <launch ID>`**
    - Purpose: Fechtes and installs the chains binary from GitHub, create a gentx for a published chain ID. This command will prompt for values like self-delegation and commission. These values will be used in the validator's gentx.
    - Command: **`ignite testnet validator join <launch ID> --amount <stake amount>`**
    - Purpose: Validators specify the amount of stake they are committing.
2. **Coordinator: List Validator Genesis Edits**
    - Command: **`ignite testnet coordinator genesis review-requests <launch ID>`**
    - Purpose: The coordinator reviews genesis edit requests from validators. These edits may include proposals for initial token allocations, validator details, and other genesis parameters. 
3. **Coordinator: Approve or Reject Validator Genesis Edits**   
    - Command: **`ignite testnet coordinator genesis approve <launch ID> <request IDs>`**
    - Command: **`ignite testnet coordinator genesis reject <launch ID> <request IDs>`**
    - Purpose: The coordinator decides on the validator requests, shaping the genesis file's final form. Approved requests are incorporated into the genesis file.
4. **Coordinator: Announce Chain Launch Readiness**  
    - Command: **`ignite testnet coordinator manage announce-ready <launch ID>`**
    - Purpose: This marks the end of the genesis file's editing phase. The coordinator declares the testnet ready for launch, signaling that no further genesis edits will be accepted.
5. **Validators: Prepare Nodes for Launch**
    - Command: **`ignite testnet coordinator manage prepare-launch <launch ID>`**
    - Purpose: Validators finalize their node configurations based on the approved genesis file. This step ensures all nodes are synchronized with the agreed-upon genesis state.
6. **Validators: Launch Nodes**
    - Action: Validators use the provided command (e.g., **`exampled --home ~/.example`**) to launch their nodes.
    - Purpose: Once a sufficient number of validators have activated their nodes, the testnet goes live with the blockchain operating as per the agreed genesis state.

## Developer instruction

Get some tokens from the spn faucet:

```bash
curl -i -X POST -d  '{"address": "spn1kwlf..."}' https://faucet.devnet.ignite.com
```

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