package undefineds

import (
	"github.com/reiver/go-cbor/internal/initialbyte"
)

func Marshal() ([]byte, error) {
	return []byte{initialbyte.Undefined}, nil
}
