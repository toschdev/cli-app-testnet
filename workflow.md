### **Revised Namespace and Workflow for Ignite Testnet Setup**

### For Coordinators:

1. **Profile Management**
    - **Command:** **`ignite testnet coordinator profile`**
    - **Subcommands:**
        - **`set`**: Set information in the coordinator's profile.
        - **`show`**: Display the coordinator's profile.
2. **Testnet Management**
    - **Command:** **`ignite testnet coordinator manage`**
    - **Subcommands:**
        - **`start`**: Initialize a new testnet.
        - **`list`**: List all published chains.
        - **`show`**: Show details of a specific chain.
        - **`announce-ready`**: Announce the readiness of the chain, locking further changes.
        - **`prepare-launch`**: Prepare the final genesis file and peers for the validator node.
        - **`revert-launch`**: Revert the launch process if needed.
3. **Genesis File Management**
    - **Command:** **`ignite testnet coordinator genesis`**
    - **Subcommands:**
        - **`review-requests`**: List all requests for genesis file modifications.
        - **`approve`**: Approve pending genesis edit requests.
        - **`reject`**: Reject genesis edit requests.
        - **`add-account`**: Request to add an account.
        - **`remove-account`**: Request to remove an account.
        - **`modify-param`**: Request changes to module parameters.
        - **`remove-validator`**: Request to remove a validator.
        - **`simulate`**: Verify and simulate the chain genesis from requests.

### For Validators:

1. **Profile Management**
    - **Command:** **`ignite testnet validator profile`**
    - **Subcommands:**
        - **`set-profile`**: Set information in the validator's profile.
        - **`show`**: Display the validator's profile.
2. **Validator Setup and Join**
    - **Command:** **`ignite testnet validator manage`**
    - **Subcommands:**
        - **`setup`**: Set up a chain from a published chain ID.
        - **`join`**: Send a request to join a testnet as a validator.

### General Commands:

1. **Tool Access**
    - **Command:** **`ignite testnet tools`**
    - **Subcommands:** Various commands to run subsidiary tools.
2. **Version Information**
    - **Command:** **`ignite testnet version`**
    - **Purpose:** Check the version of the Ignite CLI.

### **Usage**

- Coordinators use **`ignite testnet coordinator ...`** for managing the testnet and handling all aspects of genesis file modifications.
- Validators use **`ignite testnet validator ...`** for setting up their nodes and joining the testnet.
- General tools and version information can be accessed with **`ignite testnet tools`** and **`ignite testnet version`**.