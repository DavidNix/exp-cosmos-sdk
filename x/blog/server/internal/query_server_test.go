package internal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/bec/x/blog"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
)

type mockIterator struct {
	MaxItems int
	idx      int

	DidClose bool
}

func (m *mockIterator) Domain() (start []byte, end []byte) {
	panic("implement me")
}

func (m *mockIterator) Valid() bool {
	return m.idx < m.MaxItems
}

func (m *mockIterator) Next() {
	m.idx++
}

func (m *mockIterator) Key() (key []byte) {
	panic("implement me")
}

func (m *mockIterator) Value() (value []byte) {
	return []byte(strconv.Itoa(m.idx))
}

func (m *mockIterator) Error() error {
	panic("implement me")
}

func (m *mockIterator) Close() error {
	m.DidClose = true
	return nil
}

type mockUnmarshaller struct {
}

func (m *mockUnmarshaller) Unmarshal(bz []byte, proto codec.ProtoMarshaler) error {
	comment := proto.(*blog.Comment)
	comment.PostSlug = fmt.Sprintf("slug%s", bz)
	comment.Author = fmt.Sprintf("author%s", bz)
	comment.Body = fmt.Sprintf("body%s", bz)
	return nil
}

func TestFetchAllComments_AllComments(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		iter := mockIterator{MaxItems: 3}
		var cdc mockUnmarshaller
		fetch := NewFetchAllComments(&cdc, func(ctx context.Context, prefix []byte) dbm.Iterator {
			require.NotNil(t, ctx)
			require.Equal(t, "comment", string(prefix))
			return &iter
		})

		req := &blog.QueryAllCommentsRequest{PostSlug: "slug1"}
		resp, err := fetch(ctx, req)

		require.NoError(t, err)
		require.Len(t, resp.Comments, 1)
	})

	t.Run("zero state - post exists but no comments", func(t *testing.T) {
		iter := mockIterator{MaxItems: 3}
		var cdc mockUnmarshaller
		fetch := NewFetchAllComments(&cdc, func(ctx context.Context, prefix []byte) dbm.Iterator {
			return &iter
		})

		req := &blog.QueryAllCommentsRequest{PostSlug: "won't find me"}
		resp, err := fetch(ctx, req)

		require.NoError(t, err)
		require.Empty(t, resp.Comments)
	})

	t.Run("post does not exist", func(t *testing.T) {
		t.Skip("TODO - will need to inject Store interface similar to CreateComment")
	})

	t.Run("error", func(t *testing.T) {
		t.Skip("TODO - only error possible seems to be unmarshalling")
	})
}
