package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/evidence/exported"
	"github.com/okxtuta/go-anatha/x/evidence/internal/types"
)

func testHandler(sdk.Context, exported.Evidence) error { return nil }

func TestRouterSeal(t *testing.T) {
	r := types.NewRouter()
	r.Seal()
	require.Panics(t, func() { r.AddRoute("test", nil) })
	require.Panics(t, func() { r.Seal() })
}

func TestRouter(t *testing.T) {
	r := types.NewRouter()
	r.AddRoute("test", testHandler)
	require.True(t, r.HasRoute("test"))
	require.Panics(t, func() { r.AddRoute("test", testHandler) })
	require.Panics(t, func() { r.AddRoute("    ", testHandler) })
}
