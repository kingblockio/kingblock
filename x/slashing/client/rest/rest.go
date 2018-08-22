package rest

import (
	"github.com/gorilla/mux"

	"github.com/kingblockio/kingblock/client/context"
	"github.com/kingblockio/kingblock/crypto/keys"
	"github.com/kingblockio/kingblock/wire"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec, kb keys.Keybase) {
	registerQueryRoutes(ctx, r, cdc)
	registerTxRoutes(ctx, r, cdc, kb)
}
