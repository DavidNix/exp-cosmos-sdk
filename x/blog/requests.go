package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreatePost{}
	_ sdk.Msg = &MsgCreateComment{}
)

func (m *MsgCreatePost) ValidateBasic() error {
	if m.Author == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no author")
	}
	if m.Body == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no body")
	}
	if m.Slug == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no slug")
	}
	if m.Title == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no title")
	}

	return nil
}

func (m *MsgCreatePost) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Author)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCreateComment) ValidateBasic() error {
	if m.Author == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no author")
	}
	if m.Body == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no body")
	}
	if m.PostSlug == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no post slug")
	}
	return nil
}

func (m *MsgCreateComment) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Author)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
