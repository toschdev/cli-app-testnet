package cmd

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/ignite/cli/ignite/pkg/cosmosutil"
	"github.com/spf13/cobra"
	launchtypes "github.com/tendermint/spn/x/launch/types"

	"github.com/toschdev/ignite-testnet/network"
	"github.com/toschdev/ignite-testnet/network/networkchain"
	"github.com/toschdev/ignite-testnet/network/networktypes"
)

// NewNetworkCoordinatorGenesisEditAddAccount creates a new command to send add account request.
func NewNetworkCoordinatorGenesisEditAddAccount() *cobra.Command {
	c := &cobra.Command{
		Use:   "add-account [launch-id] [address] [coins]",
		Short: "Initiate an addition of a new account to the genesis file",
		Long: `The "add-account" command is designed to facilitate the creation of a request for adding a new account to the genesis file of a blockchain chain.

		Key Features:
		
		- Address and Balance Specification: Users can specify the account's address and its initial coin balance.
		- Duplication Check: The command ensures no existing genesis or vesting account shares the same address in the launch information. If such an account exists, the request is automatically rejected.
		- Uniform Balance Option: In scenarios where the coordinator mandates uniform balances across all genesis accounts (a common practice in testnets), this command requires only the address. Providing a token balance in such cases will lead to an error.
		- Usage Context: Ideal for coordinators who need to manage the initial state of accounts in the genesis file, particularly useful in setting up test environments or specific distribution scenarios.
`,
		RunE: networkRequestAddAccountHandler,
		Args: cobra.RangeArgs(2, 3),
	}

	flagSetClearCache(c)
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetHome())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())
	return c
}

func networkRequestAddAccountHandler(cmd *cobra.Command, args []string) error {
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

	chainLaunch, err := n.ChainLaunch(cmd.Context(), launchID)
	if err != nil {
		return err
	}

	c, err := nb.Chain(networkchain.SourceLaunch(chainLaunch))
	if err != nil {
		return err
	}

	var balance sdk.Coins
	if c.IsAccountBalanceFixed() {
		balance = c.AccountBalance()
		if len(args) == 3 {
			return fmt.Errorf(
				"balance can't be provided, balance has been set by coordinator to %s",
				balance.String(),
			)
		}
	} else {
		if len(args) < 3 {
			return errors.New("account balance expected")
		}
		balanceStr := args[2]
		balance, err = sdk.ParseCoinsNormalized(balanceStr)
		if err != nil {
			return err
		}
	}

	// create the add account request
	addAccountRequest := launchtypes.NewGenesisAccount(
		launchID,
		address,
		balance,
	)

	// simulate the add account request
	if err := verifyRequestsFromRequestContents(
		cmd.Context(),
		cacheStorage,
		nb,
		launchID,
		addAccountRequest,
	); err != nil {
		return err
	}

	// send the request
	return n.SendRequest(cmd.Context(), launchID, addAccountRequest)
}
