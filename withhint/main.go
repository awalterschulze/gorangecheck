//go:noinline

package main

import "math/rand"

func encodeFixed64(dAtA []byte, offset int, v uint64) int {
	_ = dAtA[offset+7]
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}

func main() {
	buf := make([]byte, 8)
	u := rand.Uint64()
	encodeFixed64(buf, 0, u)
}
