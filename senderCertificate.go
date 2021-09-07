package libsignal

/*
#cgo CFLAGS: -W
#cgo LDFLAGS: -L. ./lib/libsignal_ffi.a -lpthread -ldl -lm
#include "lib/libsignal_ffi.h"
int size(unsigned char *ptr)
{
    //variable used to access the subsequent array elements.
    int offset = 0;
    //variable that counts the number of elements in your array
    int count = 0;

    //While loop that tests whether the end of the array has been reached
    while (*(ptr + offset) != '\0')
    {
        //increment the count variable
        ++count;
        //advance to the next element of the array
        ++offset;
    }
    //return the size of the array
    return count;
}

typedef struct Out_signal_sender_certificate {
	SignalSenderCertificate *Certificate;
	SignalFfiError *Error;
} Out_signal_sender_certificate;
Out_signal_sender_certificate call_signal_sender_certificate_deserialize(SignalSenderCertificate *out1, const unsigned char *data,
                                                      size_t data_len){
	SignalSenderCertificate** out = malloc(sizeof(out));
	SignalFfiError* err = signal_sender_certificate_deserialize(out, data, data_len);
	Out_signal_sender_certificate out_struct;
	out_struct.Certificate = *out;
	out_struct.Error = err;
	return out_struct;
}

*/
import (
	"C"
)
import (
	"fmt"
	"unsafe"
)

type SenderCertificate struct {
	certificate []byte
	signature   []byte
}

func SenderCertificateDeserialize(certificate []byte) (*SenderCertificate, error) {
	out := SenderCertificate{}
	out2 := C.call_signal_sender_certificate_deserialize(
		(*C.SignalSenderCertificate)(unsafe.Pointer(&out)),
		cBytes(certificate), cLen(certificate))
	if out2.Error != nil {
		return nil, errFromCode(out2.Error)
	}
	fmt.Printf("hey %+v\n", out2.Certificate)
	out.certificate = bytesFromCBytes(out2.Certificate.certificate, C.size(out2.Certificate.certificate))
	out.signature = bytesFromCBytes(out2.Certificate.signature, C.size(out2.Certificate.signature))
	return &out, nil
}

func (s *SenderCertificate) GetCertificate() []byte {
	return s.certificate
}
