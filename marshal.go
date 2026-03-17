package cbor

import (
	"net/url"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-rfc8949/types/bools"
	"github.com/reiver/go-rfc8949/types/bytestrings"
	"github.com/reiver/go-rfc8949/types/nils"
	"github.com/reiver/go-rfc8949/types/int8s"
	"github.com/reiver/go-rfc8949/types/int16s"
	"github.com/reiver/go-rfc8949/types/int32s"
	"github.com/reiver/go-rfc8949/types/int64s"
	"github.com/reiver/go-rfc8949/types/uint8s"
	"github.com/reiver/go-rfc8949/types/uint16s"
	"github.com/reiver/go-rfc8949/types/uint32s"
	"github.com/reiver/go-rfc8949/types/tags/uris"
	"github.com/reiver/go-rfc8949/types/textstrings"
	"github.com/reiver/go-rfc8949/types/uint64s"
)

func Marshal(value any) ([]byte, error) {

	if nil == value {
		return nils.Marshal()
	}

	switch casted := value.(type) {
	case Marshaler: // This must come first.
		return casted.MarshalCBOR()
	case []byte:
		return bytestrings.Marshal(casted)
	case bool:
		return bools.Marshal(casted)
	case int:
		return int64s.Marshal(int64(casted))
	case int8:
		return int8s.Marshal(casted)
	case int16:
		return int16s.Marshal(casted)
	case int32: // rune
		return int32s.Marshal(casted)
	case int64:
		return int64s.Marshal(casted)
	case uint:
		return uint64s.Marshal(uint64(casted))
	case uint8: // byte
		return uint8s.Marshal(casted)
	case uint16:
		return uint16s.Marshal(casted)
	case uint32:
		return uint32s.Marshal(casted)
	case uint64:
		return uint64s.Marshal(casted)
	case string:
		return textstrings.Marshal(casted)
	case url.URL:
		return uris.Marshal(&casted)
	case *url.URL:
		return uris.Marshal(casted)
	default:
		return nil, erorr.Errorf("cbor: cannot marshal value of type %T", value)
	}
}
