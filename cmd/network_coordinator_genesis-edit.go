package cmd

import "github.com/spf13/cobra"

// NewNetworkCoordinatorGenesisEdit creates a new approval request command that holds some other
// sub commands related to handle request for a chain.
func NewNetworkCoordinatorGenesisEdit() *cobra.Command {
	c := &cobra.Command{
		Use:   "genesis",
		Short: "Create, show, reject and approve requests for the genesis file",
		Long: `The "genesis" namespace contains commands for creating, showing, approving, and
rejecting genesis change requests.

A genesis request is a mechanism in Ignite that allows changes made to the genesis
file like adding accounts with token balances and validators. Anyone can submit
a request, but only the coordinator of a chain can approve or reject a genesis request.

Each request has a status:

* Pending: waiting for the approval of the coordinator
* Approved: approved by the coordinator, its content has been applied to the
  launch information
* Rejected: rejected by the coordinator or the request creator
`,
	}

	c.AddCommand(
		NewNetworkCoordinatorGenesisEditShow(),
		NewNetworkCoordinatorGenesisEditReviewRequests(),
		NewNetworkCoordinatorGenesisEditApprove(),
		NewNetworkCoordinatorGenesisEditReject(),
		NewNetworkCoordinatorGenesisEditSimulate(),
		NewNetworkCoordinatorGenesisEditAddAccount(),
		NewNetworkCoordinatorGenesisEditRemoveAccount(),
		NewNetworkCoordinatorGenesisEditRemoveValidator(),
		NewNetworkCoordinatorGenesisModifyParam(),
	)

	return c
}
