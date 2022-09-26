package internal

import (
	"context"
	"encoding/hex"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/bec/x/blog"
	"github.com/stretchr/testify/require"
)

type mockStore struct {
	GotHasKey string
	StubHas   bool

	Saved map[string]string
}

func (m *mockStore) Has(key []byte) bool {
	m.GotHasKey = string(key)
	return m.StubHas
}

func (m *mockStore) Set(key, value []byte) {
	if m.Saved == nil {
		m.Saved = make(map[string]string)
	}
	m.Saved[hex.EncodeToString(key)] = string(value)
}

func TestServer_CreateComment(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		cdc := mockCodec{MarshalBytes: []byte(`stub marshalled`)}
		srv := NewServer(&cdc, testStoreKey)
		store := mockStore{StubHas: true}
		var factoryCalls int
		srv.storeFactory = func(ctx context.Context, prefix []byte) storer {
			if ctx == nil {
				panic("nil context")
			}
			switch factoryCalls {
			case 0:
				require.Equal(t, "post", string(prefix))
			case 1:
				require.Equal(t, "comment", string(prefix))
			default:
				panic("expected only 2 calls to storeFactory")
			}
			factoryCalls++
			return &store
		}
		req := &blog.MsgCreateComment{
			PostSlug: "post1",
			Author:   "John",
			Body:     "Post!",
		}
		resp, err := srv.CreateComment(ctx, req)
		require.NoError(t, err)

		gotComment := cdc.GotProtoMarshaler.(*blog.Comment)
		require.Equal(t, "post1", gotComment.PostSlug)
		require.Equal(t, "John", gotComment.Author)
		require.Equal(t, "Post!", gotComment.Body)

		require.Len(t, store.Saved, 1)

		// echo -n "post1JohnPost\!" | sha256sum
		const wantKey = "80eb43643aeecf58a9652b5fdb7d82efb7e79af12e13a4218145cd8926e1b899"
		require.Equal(t, "stub marshalled", store.Saved[wantKey])
		require.Equal(t, wantKey, resp.Hash)
	})

	t.Run("associated post does not exist", func(t *testing.T) {
		var cdc mockCodec
		srv := NewServer(&cdc, testStoreKey)
		var store mockStore
		srv.storeFactory = func(ctx context.Context, prefix []byte) storer {
			require.Equal(t, "post", string(prefix))
			return &store
		}
		req := &blog.MsgCreateComment{
			PostSlug: "post1",
			Author:   "John",
			Body:     "Post!",
		}
		_, err := srv.CreateComment(ctx, req)
		require.Error(t, err)
		require.ErrorIs(t, err, sdkerrors.ErrNotFound)
		require.EqualError(t, err, "post slug post1: not found")

		require.Nil(t, cdc.GotProtoMarshaler)
		require.Empty(t, store.Saved)
	})

	t.Run("author account does not exist", func(t *testing.T) {
		t.Skip("TODO - validate author has an account on chain to prevent comment spam")
	})
}
