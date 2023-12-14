package cmd

import (
	"time"

	timeparser "github.com/aws/smithy-go/time"
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/spf13/cobra"

	"github.com/toschdev/ignite-testnet/network"
)

const (
	flagLauchTime = "launch-time"
)

// NewNetworkCoordinatorManageAnnounceReady creates a announce-ready message that the chain is awaiting to be launched
func NewNetworkCoordinatorManageAnnounceReady() *cobra.Command {
	c := &cobra.Command{
		Use:   "announce-ready [launch-id]",
		Short: "Trigger the announcement of a chain, no further genesis changes accepted",
		Long: `# Overview

		The "announce-ready" command is a pivotal operation within the Ignite testnet environment, signifying the readiness of a blockchain chain for launch. This command is exclusive to the coordinator's role, marking a critical transition in the chain's setup process.
		
		## Usage
		
		Command "ignite testnet coordinator manage announce-ready <launch ID>"
		
		- Example: "ignite testnet coordinator manage announce-ready 42"
		
		## Description
		
		The execution of "announce-ready" communicates globally that the chain is prepared for launch. It is a crucial signal indicating the chain's transition from the setup phase to operational status.
		
		Key Points:
		
		1. Coordinator Exclusive: Only the designated coordinator of the chain has the authority to execute this command.
		2. Locks Genesis Changes: Once executed, the genesis file becomes immutable - further modifications or additions of validators (via "ignite testnet validator join") are no longer possible.
		3. Launch Time Setting: The command defaults to the current time for the chain's launch. However, coordinators can schedule a future launch using the "-launch-time" flag.
			- Example: "ignite testnet coordinator manage announce-ready 42 --launch-time 2023-01-01T00:00:00Z"
		
		### Post-Execution Actions
		
		After successfully running "announce-ready", validators are prompted to finalize the genesis file and configure their nodes for the impending launch. This preparation includes generating the final genesis file and establishing the peer list.
		
		- Validator Command: Validators should use "ignite testnet manage prepare for these final preparatory steps.
		
		### Reverting the Launch
		
		In cases where a revision of the launch time or reopening of the genesis file for updates is necessary, coordinators can utilize the "revert-launch" command.
		
		- Usage: ignite testnet coordinator manage revert-launch
		- Purpose: This allows for modifications such as incorporating new validators or adding accounts, effectively resetting the chain's status to pre-"announce-ready" state.
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
