package codec

import (
	"github.com/lesismal/arpc/codec"
	"google.golang.org/protobuf/proto"
)

type GRPCCodec struct {
	codec.Codec
}

func (g *GRPCCodec) Marshal(v any) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (g *GRPCCodec) Unmarshal(data []byte, v any) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
