package cmd

import (
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/ignite/cli/ignite/pkg/cosmosutil"
	"github.com/spf13/cobra"
	launchtypes "github.com/tendermint/spn/x/launch/types"

	"github.com/toschdev/ignite-testnet/network"
	"github.com/toschdev/ignite-testnet/network/networktypes"
)

// NewNetworkGenesisEditRemoveValidator creates a new command to send remove validator request.
func NewNetworkCoordinatorGenesisEditRemoveValidator() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove-validator [launch-id] [address]",
		Short: "Remove a validator from the genesis file",
		RunE:  networkRequestRemoveValidatorHandler,
		Args:  cobra.ExactArgs(2),
	}

	flagSetClearCache(c)
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetHome())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())
	return c
}

func networkRequestRemoveValidatorHandler(cmd *cobra.Command, args []string) error {
	session := cliui.New(cliui.StartSpinner())
	defer session.End()

	cacheStorage, err := newCache(cmd)
	if err != nil {
		return err
	}

	nb, err := newNetworkBuilder(cmd, CollectEvents(session.EventBus()))
	if err != nil {
		return err
	}

	// parse launch ID
	launchID, err := network.ParseID(args[0])
	if err != nil {
		return err
	}

	// get the address for the account and change the prefix for Ignite Chain
	address, err := cosmosutil.ChangeAddressPrefix(args[1], networktypes.SPN)
	if err != nil {
		return err
	}

	n, err := nb.Network()
	if err != nil {
		return err
	}

	// create the remove validator request
	removeValidatorRequest := launchtypes.NewValidatorRemoval(
		address,
	)

	// simulate the remove validator request
	if err := verifyRequestsFromRequestContents(
		cmd.Context(),
		cacheStorage,
		nb,
		launchID,
		removeValidatorRequest,
	); err != nil {
		return err
	}

	// send the request
	return n.SendRequest(cmd.Context(), launchID, removeValidatorRequest)
}
