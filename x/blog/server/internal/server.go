package internal

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type storer interface {
	Has(key []byte) bool
	Set(key, value []byte)
}

type Server struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
}

func NewServer(cdc codec.Codec, storeKey sdk.StoreKey) Server {
	s := Server{cdc: cdc, storeKey: storeKey}

	return s
}
