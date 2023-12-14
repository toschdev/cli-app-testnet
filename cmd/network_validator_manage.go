package cmd

import (
	"github.com/spf13/cobra"
)

// NewNetworkValidatorManage creates a new chain command that holds some other
// sub commands related to launching a testnet for a chain.
func NewNetworkValidatorManage() *cobra.Command {
	c := &cobra.Command{
		Use:   "manage",
		Short: "Manage a validator node and join a testnet",
		Long: `Setup your system for a validator node and request to join the testnet
`,
	}

	c.AddCommand(
		NewNetworkValidatorManageJoin(),
		NewNetworkValidatorManageSetup(),
	)

	return c
}
