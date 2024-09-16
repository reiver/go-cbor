package uint16s

import (
	"github.com/reiver/go-cbor/internal/initialbyte"
	"github.com/reiver/go-cbor/internal/uint8s"
)

func Marshal(value uint16) ([]byte, error) {

	switch {
	case value < 256:
		return uint8s.Marshal(uint8(value))
	}

	return []byte{initialbyte.Uint16,
		byte( (value & 0xff_00) >> (8*1) ),
		byte( (value & 0x00_ff)          ),
	}, nil
}
