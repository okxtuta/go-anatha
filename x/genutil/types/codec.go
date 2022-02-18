package types

import (
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	authtypes "github.com/okxtuta/go-anatha/x/auth/types"
	stakingtypes "github.com/okxtuta/go-anatha/x/staking/types"
)

// ModuleCdc defines a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

// TODO: abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	ModuleCdc = codec.New()
	stakingtypes.RegisterCodec(ModuleCdc)
	authtypes.RegisterCodec(ModuleCdc)
	sdk.RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
