package control

/*
#include <e2sm/wrapper.h>
#cgo LDFLAGS: -le2smwrapper -lm
#cgo CFLAGS: -I/usr/local/include/e2sm
*/
import "C"

import (
	"errors"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	"unsafe"
)

type E2sm struct {
}

func (c *E2sm) SetRicControlHeader(buffer []byte, ueIDbuf []byte, ricControlStyleType int64, ricControlActionID int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	cptr_ueIDbuf := unsafe.Pointer(&ueIDbuf[0])
	size := C.e2sm_encode_ric_control_header(cptr, C.size_t(len(buffer)), cptr_ueIDbuf, C.size_t(len(ueIDbuf)),
		C.long(ricControlStyleType), C.long(ricControlActionID))

	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	xapp.Logger.Info("E2sm SetRicControlHeader is success")
	return
}

func (c *E2sm) SetRicControlMessage(buffer []byte, targetPrimaryCell int64, targetCell int64, nrCGIOrECGI int64, nrOrEUtraCell int64, ranParameterValue []byte) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	cptrRanParameterValue := unsafe.Pointer(&ranParameterValue[0])

	size := C.e2sm_encode_ric_control_message(cptr, C.size_t(len(buffer)), C.long(targetPrimaryCell),
		C.long(targetCell), C.long(nrOrEUtraCell), C.long(nrCGIOrECGI), cptrRanParameterValue, C.size_t(len(ranParameterValue)))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set RicControlMessage due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8) //TOCHECK: if C.int(size) is returning bits not bytes to get correct size of the buffer.
	xapp.Logger.Info("E2sm SetRicControlMessage is success")
	return
}
