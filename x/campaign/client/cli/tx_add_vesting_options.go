package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
)

func CmdAddVestingOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-vesting-options [campaign-id] [address] [total-shares] [vesting-shares] [vesting-end-time]",
		Short: "Add a mainnet vesting account",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			campaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			address := args[1]

			totalShares, err := types.NewShares(args[2])
			if err != nil {
				return err
			}

			vestingShares, err := types.NewShares(args[3])
			if err != nil {
				return err
			}
			endTime, _ := strconv.ParseUint(args[4], 10, 64)
			delayedVesting := *types.NewShareDelayedVesting(totalShares, vestingShares, int64(endTime))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddVestingOptions(
				campaignID,
				clientCtx.GetFromAddress().String(),
				address,
				delayedVesting,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
