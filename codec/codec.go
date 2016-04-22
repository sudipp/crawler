package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Codec interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

var (
	JSON = jsonCodec{}
	Gob  = gobCodec{}
)

type jsonCodec struct{}

func (_ jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (_ jsonCodec) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

type gobCodec struct{}

func (c gobCodec) Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (c gobCodec) Unmarshal(b []byte, v interface{}) error {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	return dec.Decode(v)
}
