package libsignal

/*
#cgo CFLAGS: -W
#cgo LDFLAGS: -L. ./lib/libsignal_ffi.a -lpthread -ldl -lm
#include "lib/libsignal_ffi.h"

SignalPublicKey* uglycast(void* value) { return (SignalPublicKey*)value; }
*/
import (
	"C"
)
import (
	"unsafe"

	log "github.com/sirupsen/logrus"
)

// PreKeyBundle contains the data required to initialize a sender session.
// It is constructed from PreKeys registered by the peer.
type PreKeyBundle struct {
	RegistrationID uint32
	DeviceID       uint32

	PreKeyID     uint32
	PreKeyPublic *ECPublicKey

	SignedPreKeyID        int32
	SignedPreKeyPublic    *ECPublicKey
	SignedPreKeySignature [64]byte

	IdentityKey  *IdentityKey
	PreKeyBundle *C.SignalPreKeyBundle
}

// IdentityKey represents a Curve25519 public key used as a public identity.
type IdentityKey struct {
	ECPublicKey
}

// ECPublicKey represents a 256 bit Curve25519 public key.
type ECPublicKey struct {
	Key [32]byte
}

func NewPrekeyBundle(bundle PreKeyBundle) ([]byte, error) {
	log.Println("NewPrekeybundle")
	out := (*C.SignalPreKeyBundle)(unsafe.Pointer(&C.SignalPreKeyBundle{}))
	err := C.signal_pre_key_bundle_new(
		(**C.SignalPreKeyBundle)(unsafe.Pointer(out)),
		cUint32(bundle.RegistrationID),                    //registrationId
		cUint32(bundle.DeviceID),                          //deviceId
		cUint32(bundle.PreKeyID),                          //preKeyId
		cSignalPublicKey(bundle.PreKeyPublic),             //preKey
		cUint32(uint32(bundle.SignedPreKeyID)),            //signedPreKeyId
		cSignalPublicKey(bundle.SignedPreKeyPublic),       //signedPreKey
		cBytes(bundle.SignedPreKeySignature[:]),           //signedPreKeySignature
		cLen(bundle.SignedPreKeyPublic.Key[:]),            //signedPreKeyLength
		cSignalPublicKey(&bundle.IdentityKey.ECPublicKey), //identityKey
	)
	bundle.PreKeyBundle = (*C.SignalPreKeyBundle)(unsafe.Pointer(bundle.PreKeyBundle))
	log.Debugln("NewPrekeybundle err", err, bundle.PreKeyBundle)
	if err != nil {
		return nil, errFromCode(err)
	}
	return nil, nil
}

func cSignalPublicKey(key *ECPublicKey) *C.SignalPublicKey {
	return (*C.SignalPublicKey)(unsafe.Pointer(key))
}
func cSignalIdentityKey(key *ECPublicKey) *C.SignalIdentityKey {
	return (*C.SignalIdentityKey)(unsafe.Pointer(&C.SignalIdentityKey{
		public_key: C.SignalPublicKey{key: cBytes(key.Key[:])}}))
}

func NewPreKey() {

}
