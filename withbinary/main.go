//go:noinline

package main

import "math/rand"
import "encoding/binary"

func encodeFixed64(dAtA []byte, offset int, v uint64) int {
	binary.LittleEndian.PutUint64(dAtA[offset:], v)
	return offset + 8
}

func main() {
	buf := make([]byte, 8)
	u := rand.Uint64()
	encodeFixed64(buf, 0, u)
}
