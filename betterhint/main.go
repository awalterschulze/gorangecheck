//go:noinline

package main

import "math/rand"

func encodeFixed64(dAtA []byte, offset int, v uint64) int {
	buf := dAtA[offset:]
	_ = buf[7]
	buf[0] = uint8(v)
	buf[1] = uint8(v >> 8)
	buf[2] = uint8(v >> 16)
	buf[3] = uint8(v >> 24)
	buf[4] = uint8(v >> 32)
	buf[5] = uint8(v >> 40)
	buf[6] = uint8(v >> 48)
	buf[7] = uint8(v >> 56)
	return offset + 8
}

func main() {
	buf := make([]byte, 8)
	u := rand.Uint64()
	encodeFixed64(buf, 0, u)
}
