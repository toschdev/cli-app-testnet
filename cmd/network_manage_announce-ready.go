package cmd

import (
	"time"

	timeparser "github.com/aws/smithy-go/time"
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/spf13/cobra"

	"github.com/toschdev/testnet-app/network"
)

const (
	flagLauchTime = "launch-time"
)

// NewNetworkManageAnnounceReady creates a announce-ready message that the chain is awaiting to be launched
func NewNetworkManageAnnounceReady() *cobra.Command {
	c := &cobra.Command{
		Use:   "announce-ready [launch-id]",
		Short: "Trigger the announcement of a chain, no further changes accepted",
		Long: `The announce-ready command communicates to the world that the chain is ready to be
launched.

Only the coordinator of the chain can execute the announce-ready command.

	ignite testnet manage announce-ready 42

After the announce-ready command is executed no changes to the genesis are accepted. For
example, validators will no longer be able to successfully execute the "ignite
testnet manage join" command to apply as a validator.

The announce-ready command sets the date and time after which the chain will start. By
default, the current time is set. To give validators more time to prepare for
the announce-ready, set the time with the "--launch-time" flag:

	ignite testnet manage announce-ready 42 --launch-time 2023-01-01T00:00:00Z

After the announce-ready command is executed, validators can generate the finalized
genesis and prepare their nodes for the launch. For example, validators can run
"ignite testnet manage prepare" to generate the genesis and populate the peer
list.

If you want to change the launch time or open up the genesis file for changes
you can use "ignite testnet manage revert-launch" to make it possible, for
example, to accept new validators and add accounts.
`,
		Args: cobra.ExactArgs(1),
		RunE: networkChainLaunchHandler,
	}

	c.Flags().String(
		flagLauchTime,
		"",
		"timestamp the chain is effectively launched (example \"2022-01-01T00:00:00Z\")",
	)
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())

	return c
}

func networkChainLaunchHandler(cmd *cobra.Command, args []string) error {
	session := cliui.New(cliui.StartSpinner())
	defer session.End()

	nb, err := newNetworkBuilder(cmd, CollectEvents(session.EventBus()))
	if err != nil {
		return err
	}

	// parse launch ID
	launchID, err := network.ParseID(args[0])
	if err != nil {
		return err
	}

	// parse launch time
	var launchTime time.Time
	launchTimeStr, _ := cmd.Flags().GetString(flagLauchTime)

	if launchTimeStr != "" {
		launchTime, err = timeparser.ParseDateTime(launchTimeStr)
		if err != nil {
			return err
		}
	}

	n, err := nb.Network()
	if err != nil {
		return err
	}

	return n.TriggerLaunch(cmd.Context(), launchID, launchTime)
}
