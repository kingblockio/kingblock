package slashing

import (
	"github.com/kingblockio/kingblock/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(MsgUnrevoke{}, "kingblock/MsgUnrevoke", nil)
}

var cdcEmpty = wire.NewCodec()
