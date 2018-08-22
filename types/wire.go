package types

import wire "github.com/kingblockio/kingblock/wire"

// Register the sdk message type
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}
