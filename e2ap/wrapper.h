#ifndef	_WRAPPER_H_
#define	_WRAPPER_H_

#include "RICcontrolRequest.h"
#include "E2AP-PDU.h"
#include "InitiatingMessage.h"
#include "SuccessfulOutcome.h"
#include "UnsuccessfulOutcome.h"
#include "ProtocolIE-Container.h"
#include "ProtocolIE-Field.h"
#include "RICactionDefinition.h"
#include "RICsubsequentAction.h"
//#include "CauseRIC.h"

typedef struct RICControlAcknowledgeMsg {
        long requestorID;
        long instanceID;
        long ranfunctionID;
        uint8_t *callProcessID;
        size_t callProcessIDSize;
        uint8_t *ricControlOutCome;
        size_t ricControlOutComeSize;
}RICControlAcknowledge;

typedef struct RICControlFailureMsg {
        long requestorID;
        long instanceID;
        long ranfunctionID;
        uint8_t *callProcessID;
        size_t callProcessIDSize;
        uint8_t *ricControlOutCome;
        size_t ricControlOutComeSize;
        int     causeType;
        long    causeValue;

}RICControlFailure;

size_t encode_E2AP_PDU(E2AP_PDU_t* pdu, void* buffer, size_t buf_size);
E2AP_PDU_t* decode_E2AP_PDU(const void* buffer, size_t buf_size);

/* RICcontrol */
ssize_t e2ap_encode_ric_control_request_message(void *buffer, size_t buf_size, long ricRequestorID, long ricRequestSequenceNumber, long ranFunctionID, void *ricControlHdr, size_t ricControlHdrSize, void *ricControlMsg, size_t ricControlMsgSize);

RICControlAcknowledge* e2ap_decode_ric_control_acknowledge_message(void *buffer, size_t buf_size);
void e2ap_free_decoded_ric_control_ack(RICControlAcknowledge* msg);

RICControlFailure* e2ap_decode_ric_control_failure_message(void *buffer, size_t buf_size);
void e2ap_free_decoded_ric_control_failure(RICControlFailure* msg);
#endif /* _WRAPPER_H_ */
