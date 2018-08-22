package cli

import (
	"github.com/spf13/cobra"

	"github.com/kingblockio/kingblock/client/context"
	sdk "github.com/kingblockio/kingblock/types"
	"github.com/kingblockio/kingblock/wire"
	authcmd "github.com/kingblockio/kingblock/x/auth/client/cli"
	"github.com/kingblockio/kingblock/x/slashing"
)

// create unrevoke command
func GetCmdUnrevoke(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unrevoke",
		Args:  cobra.ExactArgs(1),
		Short: "unrevoke validator previously revoked for downtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))

			validatorAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := slashing.NewMsgUnrevoke(validatorAddr)

			// build and sign the transaction, then broadcast to Tendermint
			err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg}, cdc)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
