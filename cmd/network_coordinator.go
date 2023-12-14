package cmd

import "github.com/spf13/cobra"

// NewNetworkCoordinator creates a new coordinator command
// it contains sub commands to manage coordinator profile.
func NewNetworkCoordinator() *cobra.Command {
	c := &cobra.Command{
		Use:   "coordinator",
		Short: "Use the coordinator to manage setting up your testnet",
	}
	c.AddCommand(
		NewNetworkCoordinatorProfile(),
		NewNetworkCoordinatorManage(),

		NewNetworkCoordinatorGenesisEdit(),
	)
	return c
}
