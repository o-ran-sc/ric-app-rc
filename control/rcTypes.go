package control

import (
	"sync"

	rc "gerrit.o-ran-sc.org/r/ric-app/rc/protocol/grpc/ricmsgcommrpc/rc"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

const MAX_CONTROL_REQ_ATTEMPTS = 2
const RIC_CONTROL_STYLE_TYPE = 3
const RIC_CONTROL_ACTION_ID = 1

const TIME_NANOSEC = 1000000000

const RIC_CONTROL_TARGET_PRIMARY_CELL = 1
const RIC_CONTROL_TARGET_CELL = 2
const RIC_CONTROL_CELL_TYPE = 3
const RIC_CONTROL_CGI_TYPE = 4

//GRPC rsp codes
const GRPC_SUCCESS = 0
const GRPC_ERROR = -1

type Control struct {
	eventRicControlReqTimePeriod int32                //maximum time for the RIC Control Request creation procedure in the E2 Node
	rcChan                       chan *xapp.RMRParams //channel for receiving rmr message
	eventRicControlReqExpiredMap map[int]bool         //map for recording the RIC Control Request event creation procedure is expired or not
	eventRicControlReqExpiredMu  *sync.Mutex          //mutex for eventCreateExpiredMap
	ricRequestInstanceID         int
}

type RicHoControlMsg struct {
	RicControlGrpcReqPtr *rc.RicControlGrpcReq
}


type UEid struct {
        amf_UE_NGAP_Id   int64
        pLMNIdentitybuf  string
        aMFRegionIDbuf   string
        aMFSetIDbuf      string
        aMFPointerbuf    string
        F1AP_id         []int64
        E1AP_id         []int64
}

type ValueInt int64

type valueReal float64

type Null int32

type PrintableString OctetString

type OctetString struct {
        Buf  []byte
        Size int
}

type BitString struct {
        Buf        []byte
        Size       int
        BitsUnused int
}

type ControlAckMsg struct {
        RequestID             int32
        InstanceId            int32
        FuncID                int32
        CallProcessID         []byte
        CallProcessIDLength   int32
        ControlOutcome        []byte
        ControlOutcomeLength  int32
}

type ControlOutcomeFormat1ItemType struct {
	RanParameterID     int64
	RANParameterValueType int
	RANParameterValue interface{}
}

type ControlOutcomeFormat1 struct {
	ControlOutcomeFormat1Item      []ControlOutcomeFormat1ItemType
        ControlOutcomeFormat1ItemCount int
}

type ControlOutcomeMsg struct {
        ControlOutcomeType  int32
        ControlOutcome      interface{}
}
