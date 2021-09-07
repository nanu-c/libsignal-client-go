package libsignal

/*
#cgo CFLAGS: -W
#cgo LDFLAGS: -L. ./lib/libsignal_ffi.a -lpthread -ldl -lm
#include "lib/signal_ffi.h"
*/
import "C"
import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ErrVerificationFailed if verification failed.
var ErrVerificationFailed = errors.New("verification failed")

// ErrInternal if internal error.
var ErrInternal = errors.New("internal error")

// ErrInvalidInput if invalid input.
var ErrInvalidInput = errors.New("invalid input")

// Error is a generic error with code.
type Error struct {
	Code int
}

func errFromCode(code *C.SignalFfiError) error {
	log.Infoln("erFromCode:", C.signal_error_get_type(code))
	switch C.signal_error_get_type(code) {
	case C.SignalErrorCode_UnknownError:
		return errors.New("UnknownError")
	case C.SignalErrorCode_InvalidState:
		return errors.New("InvalidState")
	case C.SignalErrorCode_InternalError:
		return errors.New("InternalError")
	case C.SignalErrorCode_NullParameter:
		return errors.New("NullParameter")
	case C.SignalErrorCode_InvalidArgument:
		return errors.New("InvalidArgument")
	case C.SignalErrorCode_InvalidType:
		return errors.New("InvalidType")
	case C.SignalErrorCode_InvalidUtf8String:
		return errors.New("InvalidUtf8String")
	case C.SignalErrorCode_InsufficientOutputSize:
		return errors.New("InsufficientOutputSize")
	case C.SignalErrorCode_ProtobufError:
		return errors.New("ProtobufError")
	case C.SignalErrorCode_InvalidCiphertext:
		return errors.New("InvalidCiphertext")
	case C.SignalErrorCode_LegacyCiphertextVersion:
		return errors.New("LegacyCiphertextVersion")
	case C.SignalErrorCode_UnknownCiphertextVersion:
		return errors.New("UnknownCiphertextVersion")
	case C.SignalErrorCode_UnrecognizedMessageVersion:
		return errors.New("UnrecognizedMessageVersion")
	case C.SignalErrorCode_InvalidMessage:
		return errors.New("InvalidMessage")
	case C.SignalErrorCode_SealedSenderSelfSend:
		return errors.New("SealedSenderSelfSend")
	case C.SignalErrorCode_InvalidKey:
		return errors.New("InvalidKey")
	case C.SignalErrorCode_InvalidSignature:
		return errors.New("InvalidSignature")
	case C.SignalErrorCode_FingerprintIdentifierMismatch:
		return errors.New("FingerprintIdentifierMismatch")
	case C.SignalErrorCode_FingerprintVersionMismatch:
		return errors.New("FingerprintVersionMismatch")
	case C.SignalErrorCode_FingerprintParsingError:
		return errors.New("FingerprintParsingError")
	case C.SignalErrorCode_UntrustedIdentity:
		return errors.New("UntrustedIdentity")
	case C.SignalErrorCode_InvalidKeyIdentifier:
		return errors.New("InvalidKeyIdentifier")
	case C.SignalErrorCode_SessionNotFound:
		return errors.New("SessionNotFound")
	case C.SignalErrorCode_DuplicatedMessage:
		return errors.New("DuplicatedMessage")
	case C.SignalErrorCode_CallbackError:
		return errors.New("CallbackError")
	default:
		return Error{Code: int(C.signal_error_get_type(code))}
	}
	b := C.signal_error_get_type(code)
	fmt.Println("%d", b)
	return nil
}

func (e Error) Error() string {
	return fmt.Sprintf("zkgroup error %d", e.Code)
}
