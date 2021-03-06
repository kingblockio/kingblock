package cli

import (

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kingblockio/kingblock/client"
	"github.com/kingblockio/kingblock/client/context"

	sdk "github.com/kingblockio/kingblock/types"
	wire "github.com/kingblockio/kingblock/wire"

	authcmd "github.com/kingblockio/kingblock/x/auth/client/cli"
	"github.com/kingblockio/kingblock/x/ibc"
)

const (
	flagTo     = "to"
	flagAmount = "amount"
	flagChain  = "chain"
)

// IBC transfer command
func IBCTransferCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))

			// get the from address
			from, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}

			// build the message
			msg, err := buildMsg(from)
			if err != nil {
				return err
			}

			// get password
			err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg}, cdc)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().String(flagTo, "", "Address to send coins")
	cmd.Flags().String(flagAmount, "", "Amount of coins to send")
	cmd.Flags().String(flagChain, "", "Destination chain to send coins")
	return cmd
}

func buildMsg(from sdk.AccAddress) (sdk.Msg, error) {
	amount := viper.GetString(flagAmount)
	coins, err := sdk.ParseCoins(amount)
	if err != nil {
		return nil, err
	}

                        toStr := viper.GetString(flagTo)
			to, err := sdk.AccAddressFromBech32(toStr)
			if err != nil {
				return nil, err
			}

	packet := ibc.NewIBCPacket(from, to, coins, viper.GetString(client.FlagChainID),
		viper.GetString(flagChain))

	msg := ibc.IBCTransferMsg{
		IBCPacket: packet,
	}

	return msg, nil
}
