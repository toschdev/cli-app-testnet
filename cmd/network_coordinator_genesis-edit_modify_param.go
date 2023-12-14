package cmd

import (
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/spf13/cobra"
	launchtypes "github.com/tendermint/spn/x/launch/types"

	"github.com/toschdev/ignite-testnet/network"
	"github.com/toschdev/ignite-testnet/network/networkchain"
)

// NewNetworkCoordinatorGenesisModifyParam creates a new command to send param change request.
func NewNetworkCoordinatorGenesisModifyParam() *cobra.Command {
	c := &cobra.Command{
		Use:   "modify-param [launch-id] [module-name] [param-name] [value (json, string, number)]",
		Short: "Initiate modification of module parameters in the genesis file.",
		Long: `The modify-param command is aimed at facilitating coordinators to propose changes to specific module parameters in the genesis file of a blockchain.
		
		Functionality:

		- Parameter Customization: Allows coordinators to adjust module parameters to tailor the blockchain's behavior according to specific needs or objectives.
		- Genesis File Integration: Ensures that the proposed changes are seamlessly integrated into the genesis file, pending approval.
		- Flexibility: Provides the flexibility to alter various aspects of blockchain modules, such as consensus parameters, staking rules, or governance policies.

		Usage Context:

		This command is particularly useful for coordinators who are fine-tuning the blockchain settings before launch, ensuring that all module parameters are correctly set to align with the intended operational framework of the chain.


		Example Command:
		ignite testnet coordinator genesis modify-param [module_name] [parameter_name] [new_value]

		This command structure allows coordinators to precisely target and modify specific parameters within a module, enhancing the customization capabilities of the blockchain setup process.
		`,
		RunE: networkRequestChangeParamHandler,
		Args: cobra.ExactArgs(4),
	}

	flagSetClearCache(c)
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetHome())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())
	return c
}

func networkRequestChangeParamHandler(cmd *cobra.Command, args []string) error {
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

	module := args[1]
	param := args[2]
	value := []byte(args[3])

	n, err := nb.Network()
	if err != nil {
		return err
	}

	// fetch chain information
	chainLaunch, err := n.ChainLaunch(cmd.Context(), launchID)
	if err != nil {
		return err
	}

	c, err := nb.Chain(networkchain.SourceLaunch(chainLaunch))
	if err != nil {
		return err
	}

	// check validity of request
	err = c.CheckRequestChangeParam(
		cmd.Context(),
		module,
		param,
		value,
	)
	if err != nil {
		return err
	}

	// create the param change request
	paramChangeRequest := launchtypes.NewParamChange(
		launchID,
		module,
		param,
		value,
	)

	// simulate the param change request
	if err := verifyRequestsFromRequestContents(
		cmd.Context(),
		cacheStorage,
		nb,
		launchID,
		paramChangeRequest,
	); err != nil {
		return err
	}

	// send the request
	return n.SendRequest(cmd.Context(), launchID, paramChangeRequest)
}
