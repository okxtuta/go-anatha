package bank

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/okxtuta/go-anatha/types"
	sdkerrors "github.com/okxtuta/go-anatha/types/errors"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res, err := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)

	_, _, log := sdkerrors.ABCIInfo(err, false)
	require.True(t, strings.Contains(log, "unrecognized bank message type"))
}
