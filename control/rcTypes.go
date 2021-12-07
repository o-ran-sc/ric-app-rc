package control

import (
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	rc "gerrit.o-ran-sc.org/r/ric-app/rc/protocol/grpc/ricmsgcommrpc/rc"
	"sync"
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
	eventRicControlReqTimePeriod int32                //maximum time for the RIC Subscription Request event creation procedure in the E2 Node
	rcChan                       chan *xapp.RMRParams //channel for receiving rmr message
	eventRicControlReqExpiredMap map[int]bool //map for recording the RIC Subscription Request event creation procedure is expired or not
	eventRicControlReqExpiredMu  *sync.Mutex  //mutex for eventCreateExpiredMap
	ricRequestInstanceID         int
}

type RicHoControlMsg struct {
	RicControlGrpcReqPtr *rc.RicControlGrpcReq
}

