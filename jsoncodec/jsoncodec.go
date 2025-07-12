package jsoncodec

import (
	"encoding/json"

	"google.golang.org/grpc/encoding"
)

type JsonCodec struct{}

func (j *JsonCodec) Name() string {
	return "json"
}

func (j *JsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *JsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Init() {
	encoding.RegisterCodec(&JsonCodec{})
}
