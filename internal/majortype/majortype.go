package majortype

func MajorTypeFromInitialByte(b byte) byte {
	return (b & 0b111_00000)
}

const (
	MajorType0 byte = (0 << 5)
	MajorType1 byte = (1 << 5)
	MajorType2 byte = (2 << 5)
	MajorType3 byte = (3 << 5)
	MajorType4 byte = (4 << 5)
	MajorType5 byte = (5 << 5)
	MajorType6 byte = (6 << 5)
	MajorType7 byte = (7 << 5)
)

const (
	UnsignedInteger byte = MajorType0
	NegativeInteger byte = MajorType1
	ByteString      byte = MajorType2
	TextString      byte = MajorType3
	Array           byte = MajorType4
	Map             byte = MajorType5
	Tagged          byte = MajorType6
)

const (
	FloatingPoint byte = MajorType7
	SimpleValue   byte = MajorType7
	BreakStopCode byte = MajorType7
)

func IsUnsignedInteger(b byte) bool {
	return b == UnsignedInteger
}

func IsNegativeInteger(b byte) bool {
	return b == NegativeInteger
}

func IsByteString(b byte) bool {
	return b == ByteString
}

func IsTextString(b byte) bool {
	return b == TextString
}

func IsArray(b byte) bool {
	return b == Array
}

func IsMap(b byte) bool {
	return b == Map
}

func IsTagged(b byte) bool {
	return b == Tagged
}

func CouldBeFloatingPoint(b byte) bool {
	return b == MajorType7
}

func CouldBeSimpleValue(b byte) bool {
	return b == MajorType7
}

func CouldBeBreakStopCode(b byte) bool {
	return b == MajorType7
}
