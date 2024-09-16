package nulls

import (
	"github.com/reiver/go-cbor/internal/initialbyte"
)

func Marshal() ([]byte, error) {
	return []byte{initialbyte.Null}, nil
}
