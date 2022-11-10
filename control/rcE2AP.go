package control

/*
#include <stdlib.h>
#include <e2ap/wrapper.h>
#cgo LDFLAGS: -le2apwrapper
#cgo CFLAGS: -I/usr/local/include/e2ap
*/
import "C"

import (
	"errors"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	"unsafe"
)

type E2ap struct {
}

func (c *E2ap) SetRicControlRequestPayload(payload []byte, ricRequestorID uint16, ricRequestSequenceNumber uint16, ranFunctionID uint16,
	ricControlHdr []byte, ricControlMsg []byte) (newPayload []byte, err error) {

	cptr := unsafe.Pointer(&payload[0])
	cptr_ricControlHdr := unsafe.Pointer(&ricControlHdr[0])
	cptr_ricControlMsg := unsafe.Pointer(&ricControlMsg[0])

	xapp.Logger.Debug("\n")
	xapp.Logger.Debug("ricControlHdr\n", ricControlHdr)
	xapp.Logger.Debug("\n")
	xapp.Logger.Debug("ricControlMsg\n", ricControlMsg)
	xapp.Logger.Debug("\n")

	size := C.e2ap_encode_ric_control_request_message(cptr, C.size_t(len(payload)), C.long(ricRequestorID), C.long(ricRequestSequenceNumber),
		C.long(ranFunctionID), cptr_ricControlHdr, C.size_t(len(ricControlHdr)), cptr_ricControlMsg, C.size_t(len(ricControlMsg)))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Request Payload due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	xapp.Logger.Info("SetRicControlHeader is success")
	return
}


func (c *E2ap) GetControlAckMsg(payload []byte) (decodedMsg *ControlAckMsg, err error) {
        cptr := unsafe.Pointer(&payload[0])
        decodedMsg = &ControlAckMsg{}
        decodedCMsg := C.e2ap_decode_ric_control_acknowledge_message(cptr, C.size_t(len(payload)))
        defer C.e2ap_free_decoded_ric_control_ack(decodedCMsg)
        if decodedCMsg == nil {
                return decodedMsg, errors.New("e2ap wrapper is unable to decode indication message due to wrong or invalid payload")
        }

        decodedMsg.RequestID = int32(decodedCMsg.requestorID)
        decodedMsg.InstanceId = int32(decodedCMsg.instanceID)
        decodedMsg.FuncID = int32(decodedCMsg.ranfunctionID)
        callproc := unsafe.Pointer(decodedCMsg.callProcessID)
        decodedMsg.CallProcessID = C.GoBytes(callproc, C.int(decodedCMsg.callProcessIDSize))
        decodedMsg.CallProcessIDLength = int32(decodedCMsg.callProcessIDSize)
        controlOutcome := unsafe.Pointer(decodedCMsg.ricControlOutCome)
        decodedMsg.ControlOutcome = C.GoBytes(controlOutcome, C.int(decodedCMsg.ricControlOutComeSize))
        decodedMsg.ControlOutcomeLength = int32(decodedCMsg.ricControlOutComeSize)
        return
}

func (c *E2ap) GetControlFailureMsg(payload []byte) (decodedMsg *ControlFailureMsg, err error) {
        cptr := unsafe.Pointer(&payload[0])
        decodedMsg = &ControlFailureMsg{}
        decodedCMsg := C.e2ap_decode_ric_control_failure_message(cptr, C.size_t(len(payload)))
        defer C.e2ap_free_decoded_ric_control_failure(decodedCMsg)
        if decodedCMsg == nil {
                return decodedMsg, errors.New("e2ap wrapper is unable to decode control failure message due to wrong or invalid payload")
        }

        decodedMsg.RequestID = int32(decodedCMsg.requestorID)
        decodedMsg.InstanceId = int32(decodedCMsg.instanceID)
        decodedMsg.FuncID = int32(decodedCMsg.ranfunctionID)
        callproc := unsafe.Pointer(decodedCMsg.callProcessID)
        decodedMsg.CallProcessID = C.GoBytes(callproc, C.int(decodedCMsg.callProcessIDSize))
        decodedMsg.CallProcessIDLength = int32(decodedCMsg.callProcessIDSize)
        decodedMsg.CauseType = int32(decodedCMsg.causeType)
        decodedMsg.CauseValue = int64(decodedCMsg.causeValue)
        controlOutcome := unsafe.Pointer(decodedCMsg.ricControlOutCome)
        decodedMsg.ControlOutcome = C.GoBytes(controlOutcome, C.int(decodedCMsg.ricControlOutComeSize))
        decodedMsg.ControlOutcomeLength = int32(decodedCMsg.ricControlOutComeSize)
        return
}
