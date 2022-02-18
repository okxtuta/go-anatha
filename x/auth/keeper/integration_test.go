package keeper_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/okxtuta/go-anatha/simapp"
	sdk "github.com/okxtuta/go-anatha/types"
	authtypes "github.com/okxtuta/go-anatha/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
