package cmd

import "github.com/spf13/cobra"

// NewNetworkCoordinator creates a new coordinator command
// it contains sub commands to manage coordinator profile.
func NewNetworkCoordinatorProfile() *cobra.Command {
	c := &cobra.Command{
		Use:   "profile",
		Short: "Show and update a coordinator profile",
	}
	c.AddCommand(
		NewNetworkCoordinatorShow(),
		NewNetworkCoordinatorSet(),
	)
	return c
}
