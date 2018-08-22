package ibc

import (
	"github.com/kingblockio/kingblock/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(IBCTransferMsg{}, "kingblock/IBCTransferMsg", nil)
	cdc.RegisterConcrete(IBCReceiveMsg{}, "kingblock/IBCReceiveMsg", nil)
}
