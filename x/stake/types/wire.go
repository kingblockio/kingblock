package types

import (
	"github.com/kingblockio/kingblock/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "kingblock/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "kingblock/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "kingblock/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgBeginUnbonding{}, "kingblock/BeginUnbonding", nil)
	cdc.RegisterConcrete(MsgCompleteUnbonding{}, "kingblock/CompleteUnbonding", nil)
	cdc.RegisterConcrete(MsgBeginRedelegate{}, "kingblock/BeginRedelegate", nil)
	cdc.RegisterConcrete(MsgCompleteRedelegate{}, "kingblock/CompleteRedelegate", nil)
}

// generic sealed codec to be used throughout sdk
var MsgCdc *wire.Codec

func init() {
	cdc := wire.NewCodec()
	RegisterWire(cdc)
	wire.RegisterCrypto(cdc)
	MsgCdc = cdc
	//MsgCdc = cdc.Seal() //TODO use when upgraded to go-amino 0.9.10
}
