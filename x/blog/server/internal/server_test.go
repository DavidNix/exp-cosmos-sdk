package internal

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

const testStoreKey mockStoreKey = "test-store-key"

type mockStoreKey string

func (m mockStoreKey) Name() string   { return string(m) }
func (m mockStoreKey) String() string { return string(m) }

type mockCodec struct {
	GotProtoMarshaler codec.ProtoMarshaler
	MarshalBytes      []byte
	Err               error
}

func (m *mockCodec) Marshal(o codec.ProtoMarshaler) ([]byte, error) {
	m.GotProtoMarshaler = o
	return m.MarshalBytes, m.Err
}

func (m *mockCodec) MustMarshal(o codec.ProtoMarshaler) []byte {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MarshalLengthPrefixed(o codec.ProtoMarshaler) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MustMarshalLengthPrefixed(o codec.ProtoMarshaler) []byte {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) Unmarshal(bz []byte, ptr codec.ProtoMarshaler) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MustUnmarshal(bz []byte, ptr codec.ProtoMarshaler) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) UnmarshalLengthPrefixed(bz []byte, ptr codec.ProtoMarshaler) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr codec.ProtoMarshaler) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) UnpackAny(any *types.Any, iface interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MarshalJSON(o proto.Message) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MustMarshalJSON(o proto.Message) []byte {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MarshalInterfaceJSON(i proto.Message) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) UnmarshalInterfaceJSON(bz []byte, ptr interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) UnmarshalJSON(bz []byte, ptr proto.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockCodec) MustUnmarshalJSON(bz []byte, ptr proto.Message) {
	//TODO implement me
	panic("implement me")
}
