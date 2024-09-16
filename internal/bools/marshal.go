package bools

import (
	"github.com/reiver/go-cbor/internal/initialbyte"
)

func Marshal(value bool) ([]byte, error) {
	switch value {
	case true:
		return []byte{initialbyte.True}, nil
	default:
		return []byte{initialbyte.False}, nil
	}

	return []byte{initialbyte.Null}, nil
}
