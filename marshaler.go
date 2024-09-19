package cbor

type Marshaler interface {
	MarshalCBOR() ([]byte, error)
}
