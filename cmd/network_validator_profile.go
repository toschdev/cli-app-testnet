package cmd

import (
	"github.com/spf13/cobra"
)

// NewNetworkValidatorProfile creates a new chain command that holds some other
// sub commands related to launching a testnet for a chain.
func NewNetworkValidatorProfile() *cobra.Command {
	c := &cobra.Command{
		Use:   "profile",
		Short: "Manage validator profile",
		Long: `Set and display information in the validator's profile.
`,
	}

	c.AddCommand(
		NewNetworkValidatorProfileSet(),
		NewNetworkValidatorProfileShow(),
	)

	return c
}
