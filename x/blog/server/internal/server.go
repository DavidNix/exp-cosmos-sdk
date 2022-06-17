package internal

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tm-db"
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

	// An example how to break up grpc interface methods into smaller functions or service objects that have 1 responsibility.
	// This prevents one object like Server from becoming large with many methods and struct fields.
	// Instead, we embed the grpc service methods into Server, thus composing the Server from smaller parts.
	// The compiler guarantees the grpc service methods are implemented. Therefore, we just have to make sure
	// the constructor below is updated such that embedded types are non-nil.
	FetchAllComments
}

func NewServer(cdc codec.Codec, storeKey sdk.StoreKey) Server {
	s := Server{
		cdc:      cdc,
		storeKey: storeKey,
		storeFactory: func(ctx context.Context, pref []byte) storer {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			return prefix.NewStore(sdkCtx.KVStore(storeKey), pref)
		},

		FetchAllComments: NewFetchAllComments(cdc, func(ctx context.Context, prefix []byte) dbm.Iterator {
			// TODO: The below boilerplate would have better cohesion if close to Query implementation vs. here in the constructor.
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			store := sdkCtx.KVStore(storeKey) // Capturing via closure, may want to add storeKey to test coverage.
			return sdk.KVStorePrefixIterator(store, prefix)
		}),
	}

	return s
}
