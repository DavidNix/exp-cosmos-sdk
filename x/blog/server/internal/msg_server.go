package internal

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/multierr"

	"github.com/regen-network/bec/x/blog"
)

var _ blog.MsgServer = Server{}

func (s Server) CreatePost(goCtx context.Context, request *blog.MsgCreatePost) (*blog.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	key := []byte(request.Slug)
	if store.Has(key) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate slug %s found", request.Slug)
	}

	post := blog.Post{
		Author: request.Author,
		Slug:   request.Slug,
		Title:  request.Title,
		Body:   request.Body,
	}

	bz, err := s.cdc.Marshal(&post)
	if err != nil {
		return nil, err
	}

	store.Set(key, bz)

	return &blog.MsgCreatePostResponse{}, nil
}

func (s Server) CreateComment(ctx context.Context, req *blog.MsgCreateComment) (*blog.MsgCreateCommentResponse, error) {
	if err := s.validateComment(req); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	blogStore := s.storeFactory(ctx, blog.KeyPrefix(blog.PostKey))
	if !blogStore.Has([]byte(req.PostSlug)) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "post slug %s", req.PostSlug)
	}

	store := s.storeFactory(ctx, blog.KeyPrefix(blog.CommentKey))

	key := sha256.Sum256([]byte(req.PostSlug + req.Author + req.Body))
	post := blog.Comment{
		PostSlug: req.PostSlug,
		Author:   req.Author,
		Body:     req.Body,
	}

	bz, err := s.cdc.Marshal(&post)
	if err != nil {
		return nil, err
	}
	store.Set(key[:], bz)

	return &blog.MsgCreateCommentResponse{Hash: hex.EncodeToString(key[:])}, nil
}

func (s Server) validateComment(comment *blog.MsgCreateComment) error {
	var err error
	// TODO - trim strings?
	if comment.PostSlug == "" {
		multierr.AppendInto(&err, errors.New("post slug missing"))
	}
	if comment.Author == "" {
		multierr.AppendInto(&err, errors.New("author missing"))
	}
	if comment.Body == "" {
		multierr.AppendInto(&err, errors.New("body missing"))
	}
	return err
}
