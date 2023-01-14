package recipes

// SwapEndianUint32 swaps endianess order
// From big endian to little endian or vice versa
func SwapEndianUint32(b uint32) uint32 {
	b0 := b & 0x000000ff << 24
	b1 := b & 0x0000ff00 << 8
	b2 := b & 0x00ff0000 >> 8
	b3 := b & 0xff000000 >> 24

	swap := b0 | b1 | b2 | b3
	return swap
}

// ConcatUint8 concatenate two uint8 into uint16
func ConcatUint8(x, y uint8) uint16 {
	return uint16(x)<<8 | uint16(y)
}

// ConcatUint16 concatenate two uint16 into uint
func ConcatUint16(x, y uint16) uint32 {
	return uint32(x)<<16 | uint32(y)
}

// ConcatUint16 concatenate two uint16 into uint
func ConcatUint32(x, y uint32) uint64 {
	return uint64(x)<<32 | uint64(y)
}
