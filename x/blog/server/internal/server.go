package internal

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// storer is a test hook
type storer interface {
	Has(key []byte) bool
	Set(key, value []byte)
}

type Server struct {
	cdc          codec.Codec
	storeKey     sdk.StoreKey
	storeFactory func(ctx context.Context, prefix []byte) storer // factory
}

func NewServer(cdc codec.Codec, storeKey sdk.StoreKey) Server {
	s := Server{
		cdc:      cdc,
		storeKey: storeKey,
		storeFactory: func(ctx context.Context, pref []byte) storer {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			return prefix.NewStore(sdkCtx.KVStore(storeKey), pref)
		},
	}

	return s
}
