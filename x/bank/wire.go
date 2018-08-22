package bank

import (
	"github.com/kingblockio/kingblock/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "kingblock/Send", nil)
	cdc.RegisterConcrete(MsgIssue{}, "kingblock/Issue", nil)
}

var msgCdc = wire.NewCodec()

func init() {
	RegisterWire(msgCdc)
}
