package network

import (
	"context"
	"errors"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	rewardtypes "github.com/tendermint/spn/x/reward/types"

	"github.com/toschdev/cli-plugin-testnet/network/networktypes"
	"github.com/toschdev/cli-plugin-testnet/network/testutil"
)

func TestSetReward(t *testing.T) {
	t.Run("successfully set reward", func(t *testing.T) {
		var (
			account          = testutil.NewTestAccount(t, testutil.TestAccountName)
			suite, network   = newSuite(account)
			coins            = sdk.NewCoins(sdk.NewCoin(TestDenom, sdkmath.NewInt(TestAmountInt)))
			lastRewardHeight = int64(10)
		)

		addr, err := account.Address(networktypes.SPN)
		require.NoError(t, err)

		suite.CosmosClientMock.
			On(
				"BroadcastTx",
				context.Background(),
				account,
				&rewardtypes.MsgSetRewards{
					Provider:         addr,
					LaunchID:         testutil.LaunchID,
					Coins:            coins,
					LastRewardHeight: lastRewardHeight,
				},
			).
			Return(testutil.NewResponse(&rewardtypes.MsgSetRewardsResponse{
				PreviousCoins:            nil,
				PreviousLastRewardHeight: lastRewardHeight - 1,
				NewCoins:                 coins,
				NewLastRewardHeight:      lastRewardHeight,
			}), nil).
			Once()

		setRewardError := network.SetReward(context.Background(), testutil.LaunchID, lastRewardHeight, coins)
		require.NoError(t, setRewardError)
		suite.AssertAllMocks(t)
	})
	t.Run("failed to set reward, failed to broadcast set reward tx", func(t *testing.T) {
		var (
			account          = testutil.NewTestAccount(t, testutil.TestAccountName)
			suite, network   = newSuite(account)
			coins            = sdk.NewCoins(sdk.NewCoin(TestDenom, sdkmath.NewInt(TestAmountInt)))
			lastRewardHeight = int64(10)
			expectedErr      = errors.New("failed to set reward")
		)

		addr, err := account.Address(networktypes.SPN)
		require.NoError(t, err)

		suite.CosmosClientMock.
			On(
				"BroadcastTx",
				context.Background(),
				account,
				&rewardtypes.MsgSetRewards{
					Provider:         addr,
					LaunchID:         testutil.LaunchID,
					Coins:            coins,
					LastRewardHeight: lastRewardHeight,
				},
			).
			Return(testutil.NewResponse(&rewardtypes.MsgSetRewardsResponse{}), expectedErr).
			Once()
		setRewardError := network.SetReward(context.Background(), testutil.LaunchID, lastRewardHeight, coins)
		require.Error(t, setRewardError)
		require.Equal(t, expectedErr, setRewardError)
		suite.AssertAllMocks(t)
	})
}
