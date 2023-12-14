package cmd

import "github.com/spf13/cobra"

// NewNetworkValidator creates a new validator command
// it contains sub commands to manage validator profile.
func NewNetworkValidator() *cobra.Command {
	c := &cobra.Command{
		Use:   "validator",
		Short: "Validator commands for profile, system setup and joining a testnet",
	}
	c.AddCommand(
		NewNetworkValidatorProfile(),
		NewNetworkValidatorJoin(),
		NewNetworkValidatorSetupNode(),
	)
	return c
}
