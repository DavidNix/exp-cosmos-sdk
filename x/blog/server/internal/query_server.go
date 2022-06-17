package internal

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/bec/x/blog"
	tmdb "github.com/tendermint/tm-db"
)

var _ blog.QueryServer = Server{}

func (s Server) AllPosts(goCtx context.Context, request *blog.QueryAllPostsRequest) (*blog.QueryAllPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, blog.KeyPrefix(blog.PostKey))

	defer iterator.Close()

	var posts []*blog.Post
	for ; iterator.Valid(); iterator.Next() {
		var msg blog.Post
		err := s.cdc.Unmarshal(iterator.Value(), &msg)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &msg)
	}

	return &blog.QueryAllPostsResponse{
		Posts: posts,
	}, nil
}

// FetchAllComments is a function meant to be embedded into Server.
type FetchAllComments func(ctx context.Context, request *blog.QueryAllCommentsRequest) (*blog.QueryAllCommentsResponse, error)

type unmarshaller interface {
	Unmarshal(bz []byte, ptr codec.ProtoMarshaler) error
}

func NewFetchAllComments(cdc unmarshaller, iteratorFactory func(ctx context.Context, prefix []byte) tmdb.Iterator) FetchAllComments {
	return func(ctx context.Context, request *blog.QueryAllCommentsRequest) (*blog.QueryAllCommentsResponse, error) {
		iterator := iteratorFactory(ctx, blog.KeyPrefix(blog.CommentKey))
		defer iterator.Close()

		var comments []*blog.Comment
		// TODO: There's got to be a more efficient way to search by post_slug vs. iterating through everything.
		for ; iterator.Valid(); iterator.Next() {
			var comment blog.Comment
			if err := cdc.Unmarshal(iterator.Value(), &comment); err != nil {
				return nil, err
			}
			if comment.PostSlug != request.PostSlug {
				continue
			}
			comments = append(comments, &comment)
		}
		return &blog.QueryAllCommentsResponse{Comments: comments}, nil
	}
}

// AllComments implements blog.QueryServer and delegates to the receiver.
func (fn FetchAllComments) AllComments(ctx context.Context, request *blog.QueryAllCommentsRequest) (*blog.QueryAllCommentsResponse, error) {
	return fn(ctx, request)
}
