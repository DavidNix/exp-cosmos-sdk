package blog

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateComment_ValidateBasic(t *testing.T) {
	author, err := bech32.ConvertAndEncode("regen", []byte("test"))
	require.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		comment := &MsgCreateComment{
			PostSlug: "slug",
			Author:   author,
			Body:     "body",
		}
		require.NoError(t, comment.ValidateBasic())
	})

	t.Run("missing data", func(t *testing.T) {
		// TODO - trim whitespace tests, e.g. MsgCreateComment{Body: "   "}

		for _, tt := range []struct {
			MsgCreateComment
			WantErr string
		}{
			{MsgCreateComment{}, "no author: invalid request"},
			{MsgCreateComment{Author: author}, "no body: invalid request"},
			{MsgCreateComment{Author: author, Body: "body"}, "no post slug: invalid request"},
		} {
			err := tt.MsgCreateComment.ValidateBasic()
			require.Error(t, err, tt)
			require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest, tt)
			require.EqualError(t, err, tt.WantErr, tt)
		}
	})

	t.Run("invalid bech32 author address", func(t *testing.T) {
		t.Skip("TODO - validate address is a public key (if possible)")
	})
}
