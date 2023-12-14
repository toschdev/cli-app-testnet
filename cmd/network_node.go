package cmd

import "github.com/spf13/cobra"

// NewNetworkNode creates a new coordinator command
// it contains sub commands to manage coordinator profile.
func NewNetworkNode() *cobra.Command {
	c := &cobra.Command{
		Use:   "node",
		Short: "The node command is used to install or setup a blockchain environment on a node",
	}
	c.AddCommand(
		NewNetworkNodeInstall(),
		NewNetworkNodeSetup(),
	)
	return c
}
