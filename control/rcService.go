package control

import (
	"log"
	"runtime/debug"
	_ "strings"
	"time"

	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	"gerrit.o-ran-sc.org/r/ric-app/rc/protocol/grpc/ricmsgcommrpc/rc"
	"golang.org/x/net/context"
)

func (aPtrSendMsgSrv *RCControlServer) SendRICControlReqServiceGrpc(aCtx context.Context, aPtrRicControlGrpcReq *rc.RicControlGrpcReq) (*rc.RicControlGrpcRsp, error) {

	xapp.Logger.Info("SendRICControlReqServiceGrpc Enter ")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Encountered panic:", r)
			xapp.Logger.Error("Encountered panic:", r)
			debug.SetTraceback("all")
			time.Sleep(100 * time.Millisecond)
			panic(r)
		}
	}()

	var lRicControlGrpcRsp rc.RicControlGrpcRsp
	lRicControlGrpcRsp.RspCode = GRPC_ERROR 
	lRicControlGrpcRsp.Description = "The client specified an invalid argument. or some parameters are missing "
	if aPtrRicControlGrpcReq == nil {
		xapp.Logger.Error("Received nil data from Send error rsp")
	} else {
		if len(aPtrRicControlGrpcReq.E2NodeID) == 0 || len(aPtrRicControlGrpcReq.RICControlMessageData.TargetCellID) == 0 || len(aPtrRicControlGrpcReq.RICControlHeaderData.UEID) == 0  || len(aPtrRicControlGrpcReq.RanName) == 0  || len(aPtrRicControlGrpcReq.PlmnID) == 0{
			xapp.Logger.Error("Mandatory parameters are not received send Error rsp to client,no control Request will be initiated ")
		} else {
			xapp.Logger.Info("GRPC Control request validated, initiate Control Request to RAN")
			gChan_RicControlReq_handle <- aPtrRicControlGrpcReq
			xapp.Logger.Debug("Received RIC Control Request with RicRequestorID: %d", aPtrRicControlGrpcReq.RICE2APHeaderData.RICRequestorID)
			xapp.Logger.Debug("Received RIC Control Request with RicActionId: %d", aPtrRicControlGrpcReq.RICE2APHeaderData.RanFuncId)

			xapp.Logger.Debug("Received RIC Control Request with UEID: %s", aPtrRicControlGrpcReq.RICControlHeaderData.UEID)
			xapp.Logger.Debug("Received RIC Control Request with ControlStyle: %d", aPtrRicControlGrpcReq.RICControlHeaderData.ControlStyle)
			xapp.Logger.Debug("Received RIC Control Request with ControlActionId: %d", aPtrRicControlGrpcReq.RICControlHeaderData.ControlActionId)

			xapp.Logger.Debug("Received RIC Control Request with CellType: %d", aPtrRicControlGrpcReq.RICControlMessageData.RICControlCellTypeVal)
			xapp.Logger.Debug("Received RIC Control Request with TargetCellID: %s", aPtrRicControlGrpcReq.RICControlMessageData.TargetCellID)
			xapp.Logger.Debug("Received RIC Control Request with RanName: %s", aPtrRicControlGrpcReq.RanName)
			lRicControlGrpcRsp.RspCode = GRPC_SUCCESS //int32(lGrpcRspStatusCode)
			lRicControlGrpcRsp.Description = "Success Response"
			xapp.Logger.Debug("RICHOCONTROL_EVENT: Success Response sent with  RspCode:%d, Description:%s", lRicControlGrpcRsp.RspCode, lRicControlGrpcRsp.Description)
		}
	}

	xapp.Logger.Info("RICHOCONTROL_EVENT: Response sent with  RspCode:%d, Description:%s", lRicControlGrpcRsp.RspCode, lRicControlGrpcRsp.Description)

	log.Printf(" Response sent to , rsp: {Code: ", lRicControlGrpcRsp.RspCode, ", Description: ", lRicControlGrpcRsp.Description, "}")
	return &lRicControlGrpcRsp, nil
}
