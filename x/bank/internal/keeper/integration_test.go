package keeper_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/okxtuta/go-anatha/simapp"
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/auth"
)

func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})

	app.AccountKeeper.SetParams(ctx, auth.DefaultParams())
	app.BankKeeper.SetSendEnabled(ctx, true)

	return app, ctx
}
