package cbor

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-rfc8949/types/bools"
	"github.com/reiver/go-rfc8949/types/nils"
	"github.com/reiver/go-rfc8949/types/int8s"
	"github.com/reiver/go-rfc8949/types/int16s"
	"github.com/reiver/go-rfc8949/types/int32s"
	"github.com/reiver/go-rfc8949/types/int64s"
	"github.com/reiver/go-rfc8949/types/uint8s"
	"github.com/reiver/go-rfc8949/types/uint16s"
	"github.com/reiver/go-rfc8949/types/uint32s"
	"github.com/reiver/go-rfc8949/types/uint64s"
)

func Marshal(value any) ([]byte, error) {

	if nil == value {
		return nils.Marshal()
	}

	switch casted := value.(type) {
	case Marshaler:
		return casted.MarshalCBOR()
	case bool:
		return bools.Marshal(casted)
	case int:
		return Marshal(int64(casted))
	case int8:
		return int8s.Marshal(casted)
	case int16:
		return int16s.Marshal(casted)
	case int32: // rune
		return int32s.Marshal(casted)
	case int64:
		return int64s.Marshal(casted)
	case uint:
		return Marshal(uint64(casted))
	case uint8: // byte
		return uint8s.Marshal(casted)
	case uint16:
		return uint16s.Marshal(casted)
	case uint32:
		return uint32s.Marshal(casted)
	case uint64:
		return uint64s.Marshal(casted)
	default:
		return nil, erorr.Errorf("cbor: cannot marshal value of type %T", value)
	}
}
