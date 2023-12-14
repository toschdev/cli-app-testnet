### The Validator

**Overview:**
Validators play a key role in maintaining the blockchain's integrity and security. They are responsible for validating transactions, creating new blocks, and participating in consensus mechanisms. Setting up their nodes correctly and aligning with the network's parameters are critical for the smooth operation of the blockchain.

**(optional) Set up your Profile**

Validators can use the `ignite testnet validator profile` command to manage their public profiles. This can be important for network transparency and for coordinators to identify and assess validators. A well-managed profile includes accurate information about the validator's node, its capabilities, and the validator's contact information, contributing to a robust and trustworthy network ecosystem.

**Key Commands:**

1. **Join a Testnet**
    - **`ignite testnet validator join`**
    - Purpose: Send a request to join a testnet as a validator.
2. **Manage Validator Profile**
    - **`ignite testnet validator profile`**
    - Purpose: Manage personal profile information as a validator.

### Example Workflow

1. **Coordinator Starts a Chain:**
    - Executes **`ignite testnet coordinator manage start`** with the necessary repository URL to initiate a new blockchain.
2. **Validators Set Up and Apply:**
    - Run **`ignite testnet validator setup-node`** which fechtes and installs the chains binary from GitHub, creates a gentx for a published chain ID. This command will prompt for values like self-delegation and commission.
    - Run **`ignite testnet validator join`** with their node information and stake amount to apply as a validator.
3. **Coordinator Manages Genesis Requests:**
    - Uses **`ignite testnet coordinator genesis`** to review and approve or reject validator requests.
4. **Coordinator Announces Readiness:**
    - Once satisfied with the validator set, the coordinator runs **`ignite testnet coordinator manage announce-ready`**.
5. **Validators Prepare Nodes:**
    - After the readiness announcement, validators execute **`ignite testnet validator prepare-launch`** to finalize their node setup.
6. **Launch and Monitoring:**
    - The coordinator monitors the network, and validators ensure their nodes are running smoothly.