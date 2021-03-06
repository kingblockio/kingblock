package slashing

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/kingblockio/kingblock/types"
)

func TestMsgUnrevokeGetSignBytes(t *testing.T) {
	addr := sdk.AccAddress("abcd")
	msg := NewMsgUnrevoke(addr)
	bytes := msg.GetSignBytes()
	require.Equal(t, string(bytes), `{"address":"kingblockioaccaddr1v93xxeqhyqz5v"}`)
}
