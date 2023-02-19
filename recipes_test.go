package recipes

import (
	"testing"
)

func TestSwapEndianUint32(t *testing.T) {
	var b uint32 = 0xaabbccdd
	var expected uint32 = 0xddccbbaa

	swap := SwapEndianUint32(b)
	if swap != expected {
		t.Errorf("Error while swaping endian. Expected: 0x%x, result: 0x%x", expected, swap)
	}
}

func TestConcatUint8(t *testing.T) {
	var x uint8 = 0xf2
	var y uint8 = 0x75
	var expected uint16 = 0xf275

	res := ConcatUint8(x, y)

	if res != expected {
		t.Errorf("Wrong uint8 concatenation, expected 0x%x, received: 0x%x", expected, res)
	}
}

func TestConcatUint16(t *testing.T) {
	var x uint16 = 0xf1f2
	var y uint16 = 0xf3f4
	var expected uint32 = 0xf1f2f3f4

	res := ConcatUint16(x, y)

	if res != expected {
		t.Errorf("Wrong uint16 concatenation, expected 0x%x, received: 0x%x", expected, res)
	}
}

func TestConcatUint32(t *testing.T) {
	var x uint32 = 0xf1f2f3f4
	var y uint32 = 0xb1b2b3b4
	var expected uint64 = 0xf1f2f3f4b1b2b3b4

	res := ConcatUint32(x, y)

	if res != expected {
		t.Errorf("Wrong uint32 concatenation, expected 0x%x, received: 0x%x", expected, res)
	}
}

func TestReverse(t *testing.T) {
	var val uint16 = 0xaabb

	reversed := reverse(val)

	if reversed != 0xbbaa {
		t.Error("Wrong reverse")
	}
}
