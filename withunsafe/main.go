//go:noinline

package main

import (
	"math/rand"
	"unsafe"
)

func encodeFixed64(dAtA []byte, offset int, v uint64) int {
	*(*uint64)(unsafe.Pointer(&dAtA[offset])) = v
	return offset + 8
}

func main() {
	buf := make([]byte, 8)
	u := rand.Uint64()
	encodeFixed64(buf, 0, u)
}
