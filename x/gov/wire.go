package gov

import (
	"github.com/kingblockio/kingblock/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {

	cdc.RegisterConcrete(MsgSubmitProposal{}, "kingblock/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "kingblock/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "kingblock/MsgVote", nil)

	cdc.RegisterInterface((*Proposal)(nil), nil)
	cdc.RegisterConcrete(&TextProposal{}, "gov/TextProposal", nil)
}

var msgCdc = wire.NewCodec()
