package cmd

import (
	"github.com/spf13/cobra"
)

// NewNetworkCoordinatorManage creates a new chain command that holds some other
// sub commands related to launching a testnet for a chain.
func NewNetworkCoordinatorManage() *cobra.Command {
	c := &cobra.Command{
		Use:   "manage",
		Short: "Manage a testnet as coordinator",
		Long: `The "manage" namespace features the most commonly used commands for launching
blockchains with Ignite.

As a coordinator you start your blockchain with Ignite. When enough validators
are approved for the genesis and no changes are excepted to be made to the
genesis, a coordinator announces that the chain is ready for launch with the
"announce-ready" command. In the case of an unsuccessful launch, the coordinator can revert it
using the "revert-launch" command.

As a validator, you setup your node and apply to become a validator for a
blockchain with the "join" command. After the launch of the chain is announced,
validators can generate the finalized genesis and download the list of peers with the
"prepare-launch" command.

The "install" command can be used to download, compile the source code and
install the chain's binary locally. The binary can be used, for example, to
initialize a validator node or to interact with the chain after it has been
launched.

All chains published to Ignite can be listed by using the "list" command.
`,
	}

	c.AddCommand(
		NewNetworkCoordinatorManageStart(),
		NewNetworkCoordinatorManagePrepareLaunch(),
		NewNetworkCoordinatorManageAnnounceReady(),
		NewNetworkCoordinatorManageRevertLaunch(),
		NewNetworkCoordinatorManageInstallBinaries(),
	)

	return c
}
