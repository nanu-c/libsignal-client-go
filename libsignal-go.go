package libsignal

/*
#cgo CFLAGS: -W
#cgo LDFLAGS: -L. ./lib/libsignal_ffi.a -lpthread -ldl -lm
#include "lib/signal_ffi.h"
//returns the size of a character array using a pointer to the first element of the character array


*/
import "C"
import (
	"crypto/rand"
	"unsafe"
)

func cBytes(b []byte) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&b[0]))
}

func cLen(b []byte) C.ulong {
	return C.ulong(len(b))
}
func cUint32(i uint32) C.uint32_t {
	return C.uint32_t(i)
}

func randBytes(length int) []byte {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return buf
}
func bytesFromCBytes(cBytes *C.uchar, length C.int) []byte {
	return C.GoBytes(unsafe.Pointer(cBytes), length)
}

// func NewPrekeyBundle() {

// 	out := &C.SignalPreKeyBundle{}
// 	// if res := C.signal_pre_key_bundle_new(&out); res != 0 {
// 	// 	return nil, errFromCode(res)
// 	// }

// }
