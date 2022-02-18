package types

import (
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/supply/exported"
)

// SupplyKeeper defines the expected supply Keeper (noalias)
type SupplyKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetModuleAccount(ctx sdk.Context, moduleName string) exported.ModuleAccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}
