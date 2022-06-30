package control

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"gerrit.o-ran-sc.org/r/ric-app/rc/protocol/grpc/ricmsgcommrpc/rc"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

var (
	gControlData               *Control
	gChan_RicControlReq_handle = make(chan *rc.RicControlGrpcReq, 2000) //Make it configurable
)

func NewControl() Control {

	file := "/opt/rc.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	logLevel := xapp.Config.GetInt("controls.logLevel")
	xapp.Logger.SetLevel(logLevel)

	xapp.Logger.Debug("GRPC Server Port = %v ", xapp.Config.GetString("controls.ricHOControlgRpcServerPort"))
	xapp.Logger.Debug("Log Level = %d ", xapp.Config.GetInt("controls.logLevel"))

	go StartgRPCRCControlCommServerRoutine()
	xapp.Logger.Debug("StartgRPCRCControlCommServerRoutine done")

	//To Handle RIC Control Message
	go StartHandleControlReqRoutine()

	return Control{5,
		make(chan *xapp.RMRParams, 1000), //Make it configurable
		make(map[int]bool),
		&sync.Mutex{},
		0}
}

func ReadyCB(i interface{}) {
	gControlData = i.(*Control)
	go controlLoop()
}

func (aControlData *Control) Run() {
	xapp.SetReadyCB(ReadyCB, aControlData)
	xapp.Run(aControlData)
}

func (aControlData *Control) Consume(rp *xapp.RMRParams) (err error) {
	gControlData.rcChan <- rp
	return
}

func (aControlData *Control) rmrSend(params *xapp.RMRParams) (err error) {
	if !xapp.Rmr.Send(params, false) {
		err = errors.New("rmr.Send() failed")
		xapp.Logger.Info("Failed to rmrSend to %v", err)
		log.Printf("Failed to rmrSend to %v", err)
	}
	return
}

func (aControlData *Control) rmrReplyToSender(params *xapp.RMRParams) (err error) {
	if !xapp.Rmr.Send(params, true) {
		err = errors.New("rmr.Send() failed")
		xapp.Logger.Error("Failed to rmrReplyToSender to %v", err)
		log.Printf("Failed to rmrReplyToSender to %v", err)
	}
	return
}

func controlLoop() {
	for {
		msg := <-gControlData.rcChan
		xapp.Logger.Debug("Received message type: %d", msg.Mtype)
		log.Printf("Received message type: %d", msg.Mtype)
		switch msg.Mtype {
		case 12041:
			go HandleControlResponse(msg)
		case 12042:
			go HandleControlFailure(msg)
		default:
			err := errors.New("Message Type " + strconv.Itoa(msg.Mtype) + " is discarded")
			xapp.Logger.Error("Unknown message type: %v", err)
			log.Printf("Unknown message type: %v", err)
		}
	}
}

func StartHandleControlReqRoutine() {

	log.Printf("Starting Go Routine for Handling GRPC RIC Control msg ")
	xapp.Logger.Info("Starting Go Routine for Handling GRPC RIC Control msg ")
	for {
		HandlegRPCRICControlMsgReq(<-gChan_RicControlReq_handle)
	}
	xapp.Logger.Debug("StartHandleControlReqRoutine Done")
}

func HandlegRPCRICControlMsgReq(aPtrRicControlGrpcReq *rc.RicControlGrpcReq) {

	xapp.Logger.Debug("HandlegRPCRICControlMsgReq :%v", *aPtrRicControlGrpcReq)
	lRicHoControlMsg := RicHoControlMsg{}
	lRicHoControlMsg.RicControlGrpcReqPtr = aPtrRicControlGrpcReq

	lUEID := lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID
	xapp.Logger.Debug("HandlegRPCRICControlMsgReq UEID = %v ", lUEID)
	//Mandatory parameters validation
	if lRicHoControlMsg.RicControlGrpcReqPtr.E2NodeID == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.TargetCellID == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.AmfUENGAPID < 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.Guami.PLMNIdentity == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.Guami.AMFRegionID == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.Guami.AMFSetID == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.Guami.AMFPointer == "" ||
		len(lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.GNBCUUEF1APID) == 0 ||
		len(lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID.GnbUEID.GNBCUCPUEE1APID) == 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.PlmnID == "" ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.ControlActionId < 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.ControlStyle < 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.RICControlCellTypeVal < 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICE2APHeaderData.RICRequestorID < 0 ||
		lRicHoControlMsg.RicControlGrpcReqPtr.RICE2APHeaderData.RanFuncId < 0 {
		xapp.Logger.Error("Mandaroty parameters missing, dont send control request ")
		return
	}

	lRicHoControlMsg.GetSequenceNumber()

	go lRicHoControlMsg.SendRicControlRequest(lRicHoControlMsg.GetSequenceNumber())

	return
}

func (aRicHoControlMsg *RicHoControlMsg) GetSequenceNumber() int {

	//Incrementing the RIC Requestor Instance Id to make the request unique and traceable.
	gControlData.eventRicControlReqExpiredMu.Lock()
	gControlData.ricRequestInstanceID = gControlData.ricRequestInstanceID + 1
	gControlData.eventRicControlReqExpiredMu.Unlock()

	return gControlData.ricRequestInstanceID
}

func (aRicHoControlMsg *RicHoControlMsg) setEventRicControlCreateExpiredTimer(aSeqNum int) {

	gControlData.eventRicControlReqExpiredMu.Lock()
	gControlData.eventRicControlReqExpiredMap[aSeqNum] = false
	gControlData.eventRicControlReqExpiredMu.Unlock()
	controlAckTimer := xapp.Config.GetInt("controls.controlAckTimer")
	xapp.Logger.Debug("configured controlAckTimer = %d and controlAckTimer = %d ", xapp.Config.GetInt("controls.controlAckTimer"),controlAckTimer)

	timer := time.NewTimer(time.Duration(controlAckTimer) * time.Second)
	go func(t *time.Timer) {
		defer t.Stop()
		xapp.Logger.Debug("RIC_CONTROL_REQ[%s]: Waiting for RIC_CONTROL_RESP...", aSeqNum)
		log.Printf("RIC_CONTROL_REQ[%s]: Waiting for RIC_CONTROL_RESP...", aSeqNum)
		for {
			select {
			case <-t.C:
				gControlData.eventRicControlReqExpiredMu.Lock()
				isResponsed := gControlData.eventRicControlReqExpiredMap[aSeqNum]
				delete(gControlData.eventRicControlReqExpiredMap, aSeqNum)
				gControlData.eventRicControlReqExpiredMu.Unlock()
				if !isResponsed {
					xapp.Logger.Debug("RIC_CONTROL_REQ[%s]: RIC Event Create Timer experied!", aSeqNum)
					log.Printf("RIC_CONTROL_REQ[%s]: RIC Event Create Timer experied!", aSeqNum)
					//Send ErrorIndication message on Timeout
					return
				}
			default:
				gControlData.eventRicControlReqExpiredMu.Lock()
				flag := gControlData.eventRicControlReqExpiredMap[aSeqNum]
				if flag {
					delete(gControlData.eventRicControlReqExpiredMap, aSeqNum)
					gControlData.eventRicControlReqExpiredMu.Unlock()
					xapp.Logger.Debug("RIC_CONTROL_REQ[%s]: RIC Event Create Timer canceled!", aSeqNum)
					log.Printf("RIC_CONTROL_REQ[%s]: RIC Event Create Timer canceled!", aSeqNum)
					return
				} else {
					gControlData.eventRicControlReqExpiredMu.Unlock()
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(timer)
}
func (aRicHoControlMsg *RicHoControlMsg) SendRicControlRequest(aRequestSN int) (err error) {
	var e2ap *E2ap
	var e2sm *E2sm

	xapp.Logger.Info("SendRicControlRequest Enter for RanName = %s", aRicHoControlMsg.RicControlGrpcReqPtr.RanName)

	if aRicHoControlMsg.RicControlGrpcReqPtr == nil {
		return err
	}

	var lRicControlStyleType int64 = aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.ControlStyle
	var lRicControlActionID int64 = aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.ControlActionId
	lUEID := aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID

	var ueId_data UEid

	xapp.Logger.Debug("UEID:%v, ueId_data strct :%v", aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID,
		ueId_data)
	ueId_data.amf_UE_NGAP_Id = lUEID.GnbUEID.AmfUENGAPID
	ueId_data.pLMNIdentitybuf = lUEID.GnbUEID.Guami.PLMNIdentity
        ueId_data.aMFRegionIDbuf = lUEID.GnbUEID.Guami.AMFRegionID
        ueId_data.aMFSetIDbuf = lUEID.GnbUEID.Guami.AMFSetID
        ueId_data.aMFPointerbuf = lUEID.GnbUEID.Guami.AMFPointer
        ueId_data.F1AP_id = lUEID.GnbUEID.GNBCUUEF1APID
        ueId_data.E1AP_id = lUEID.GnbUEID.GNBCUCPUEE1APID

	var lRicControlHeader []byte = make([]byte, 1024) //Check the Size
	lRicControlHeaderEncoded, err := e2sm.SetRicControlHeader(lRicControlHeader, &ueId_data, lRicControlStyleType, lRicControlActionID)
	if err != nil {
		xapp.Logger.Error("SetRicControlHeader Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		log.Printf("SetRicControlHeader Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		return err
	} else {
		xapp.Logger.Info("SetRicControlHeader is success: %x", lRicControlHeaderEncoded)
		fmt.Fprintf(os.Stderr, "Encoded RIC Control Header PDU:\n")
		for i := 0; i < len(lRicControlHeaderEncoded); i++ {
			fmt.Fprintf(os.Stderr, " %02x", lRicControlHeaderEncoded[i])
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	var lTargetPrimaryCell int64 = RIC_CONTROL_TARGET_PRIMARY_CELL
	var lTargetCell int64 = RIC_CONTROL_TARGET_CELL
	var lNrCGIOrECGI int64 = RIC_CONTROL_CGI_TYPE

	lNrOrEUtraCellType := aRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.RICControlCellTypeVal
	lTargetCellVal := aRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.TargetCellID
	lTargetCellValBuf := []byte(lTargetCellVal)
	//lNRPlmnId := []byte(aRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.TargetCellID.PlmnID)
        //lNRCellId := aRicHoControlMsg.RicControlGrpcReqPtr.RICControlMessageData.TargetCellID.NRCellID


	var lRicControlMessage []byte = make([]byte, 1024)
	lRicControlMessageEncoded, err := e2sm.SetRicControlMessage(lRicControlMessage, lTargetPrimaryCell, lTargetCell, lNrCGIOrECGI, int64(lNrOrEUtraCellType), lTargetCellValBuf)
	if err != nil {
		xapp.Logger.Error("SetRicControlMessage Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		log.Printf("SetRicControlMessage Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		return err
	} else {
		xapp.Logger.Debug("SetRicControlMessage is success: %x", lRicControlMessageEncoded)
		fmt.Fprintf(os.Stderr, "Encoded RIC Control Message PDU:\n")
		for i := 0; i < len(lRicControlMessageEncoded); i++ {
			fmt.Fprintf(os.Stderr, " %02x", lRicControlMessageEncoded[i])
		}
		fmt.Fprintf(os.Stderr, "\n")
	}

	lParams := &xapp.RMRParams{}
	lParams.Mtype = 12040 //RIC_CONTROL_REQ
	lParams.SubId = -1

	var lRequestorId uint16 = uint16(aRicHoControlMsg.RicControlGrpcReqPtr.RICE2APHeaderData.RICRequestorID)
	var lFuncId uint16 = uint16(aRicHoControlMsg.RicControlGrpcReqPtr.RICE2APHeaderData.RanFuncId)

	lParams.Payload = make([]byte, 2048)
	lParams.Payload, err = e2ap.SetRicControlRequestPayload(lParams.Payload, lRequestorId, uint16(aRequestSN), lFuncId,
		lRicControlHeaderEncoded, lRicControlMessageEncoded)
	if err != nil {
		xapp.Logger.Error("SetRicControlRequestPayload Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		log.Printf("SetRicControlRequestPayload Failed: %v, UEID:%v", err, aRicHoControlMsg.RicControlGrpcReqPtr.RICControlHeaderData.UEID)
		return err
	} else {

		xapp.Logger.Debug("Encoding RicControlRequestPayload is success. UEID: %v, Payload: %x", lUEID, lParams.Payload)
		fmt.Fprintf(os.Stderr, "Encoded RIC Control Req PDU:\n")
		for i := 0; i < len(lParams.Payload); i++ {
			fmt.Fprintf(os.Stderr, " %02x", lParams.Payload[i])
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	lParams.PayloadLen = len(lParams.Payload)
	valEnbId := aRicHoControlMsg.RicControlGrpcReqPtr.E2NodeID
	valRanName := aRicHoControlMsg.RicControlGrpcReqPtr.RanName
	valPlmnId := aRicHoControlMsg.RicControlGrpcReqPtr.PlmnID
	lParams.Meid = &xapp.RMRMeid{PlmnID: valPlmnId, EnbID: valEnbId, RanName: valRanName}

	xapp.Logger.Debug("The RIC Control RMR message to be sent is with MsgType:%d  SubId=%d, lParams.Meid: %v, UEID: %v", lParams.Mtype, lParams.SubId, lParams.Meid, lUEID)

	xapp.Logger.Debug("Sending RIC Control message to RanName: %s, UEID: %v  ", aRicHoControlMsg.RicControlGrpcReqPtr.RanName, lUEID)
	
	xapp.Logger.Info("lParams %v ",lParams)
	err = gControlData.rmrSend(lParams)
	if err != nil {
		xapp.Logger.Error("Failed to send RIC_CONTROL_REQ: %v", err)
		log.Printf("Failed to send RIC_CONTROL_REQ: %v", err)
		xapp.Logger.Info("Failed to send RIC_CONTROL_REQ: %v", err)
		return err
	}

	xapp.Logger.Info("Sending RIC Control message to RanName: %s, UEID: %v  Success", aRicHoControlMsg.RicControlGrpcReqPtr.RanName, lUEID)
	//Start Timer Operation.
	aRicHoControlMsg.setEventRicControlCreateExpiredTimer(aRequestSN) //TODO check if this is required as we are not expecting Control ACK

	return nil
}

func HandleControlResponse(params *xapp.RMRParams) (err error) {

	var e2ap *E2ap
        var e2sm *E2sm
	xapp.Logger.Debug("The SubId in RIC_CONTROL_RESP is %d", params.SubId)
	log.Printf("The SubId in RIC_CONTROL_RESP is %d", params.SubId)



	//Decode the RIC Control Ack message
	controlAck, err := e2ap.GetControlAckMsg(params.Payload)
	if err != nil {
		xapp.Logger.Error("Failed to decode RIC Control Ack: %v", err)
                log.Println("Failed to decode RIC Control Ack: %v", err)
                return
	}
	log.Println("E2ap RIC RIC Control Ack message decoded \n")
	xapp.Logger.Info("E2ap RIC RIC Control Ack message decoded \n")
	gControlData.eventRicControlReqExpiredMu.Lock()
        _,ok := gControlData.eventRicControlReqExpiredMap[int(controlAck.InstanceId)]
        if !ok {
                gControlData.eventRicControlReqExpiredMu.Unlock()
                xapp.Logger.Debug("RIC_CONTROL_REQ has been deleted!")
                log.Printf("RIC_CONTROL_REQ has been deleted!")
                return nil
        } else {
                gControlData.eventRicControlReqExpiredMap[int(controlAck.InstanceId)] = true
		gControlData.eventRicControlReqExpiredMu.Unlock()
        }

	controlOutcome, err := e2sm.GetControlOutcome(controlAck.ControlOutcome)
        if err != nil {
		xapp.Logger.Error("Failed to decode RIC Control Outcome: %v", err)
                log.Println("Failed to decode RIC Control Outcome: %v", err)
                return
        }
        log.Println("E2SM-RC RIC Control Outcome  decoded \n",controlOutcome)
        xapp.Logger.Info("E2SM-RC RIC Control Outcome  decoded \n",controlOutcome)



	return nil
}

func HandleControlFailure(params *xapp.RMRParams) (err error) {

	var e2ap *E2ap

	xapp.Logger.Debug("The SubId in RIC_CONTROL_FAILURE is %d", params.SubId)
	log.Printf("The SubId in RIC_CONTROL_FAILURE is %d", params.SubId)

	controlAck, err := e2ap.GetControlFailureMsg(params.Payload)
	if err != nil {
                xapp.Logger.Debug("Failed to decode RIC Control message: %v", err)
                log.Println("Failed to decode RIC Control Ack: %v", err)
                return
        }
        log.Println("E2ap RIC  Control Ack message decoded \n")
        xapp.Logger.Debug("E2ap RIC Control Ack message decoded \n")

        gControlData.eventRicControlReqExpiredMu.Lock()
        _, ok := gControlData.eventRicControlReqExpiredMap[int(controlAck.InstanceId)]
        if !ok {
                gControlData.eventRicControlReqExpiredMu.Unlock()
                xapp.Logger.Debug("RIC_CONTROL_REQ has been deleted!")
                log.Printf("RIC_CONTROL_REQ has been deleted!")
                return nil
        } else {
                gControlData.eventRicControlReqExpiredMap[int(controlAck.InstanceId)] = true
                gControlData.eventRicControlReqExpiredMu.Unlock()
        }
	return nil
}
