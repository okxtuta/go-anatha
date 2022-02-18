package types

import "github.com/okxtuta/go-anatha/codec"

// Register the sdk message type
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}
