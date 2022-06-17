package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/regen-network/bec/x/blog/server/internal"

	"github.com/regen-network/bec/x/blog"
)

func RegisterServices(cdc codec.Codec, storeKey sdk.StoreKey, configurator module.Configurator) {
	impl := internal.NewServer(cdc, storeKey)
	blog.RegisterMsgServer(configurator.MsgServer(), impl)
	blog.RegisterQueryServer(configurator.QueryServer(), impl)
}
