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
	_ "log"
	_"encoding/binary"
	"encoding/hex"
	"strings"
	"strconv"
)

type E2sm struct {
}

func (c *E2sm) SetRicControlHeader(buffer []byte, ueIdData *UEid, ricControlStyleType int64, ricControlActionID int64) (newBuffer []byte, err error) {

	xapp.Logger.Info("SetRicControlHeader Enter \n")
	cptr := unsafe.Pointer(&buffer[0])
	//cptr_ueIDbuf := unsafe.Pointer(&ueIDbuf[0])

	lplmnIdBuf := strings.Join(strings.Fields(ueIdData.pLMNIdentitybuf), "")
	lIntvar, _ := strconv.Atoi(lplmnIdBuf)
	xapp.Logger.Info("lIntvar = %d\n", lIntvar)

	lIntegerByte := ConverIntegerToByte(lIntvar)
	//ml.MavLog(ml.INFO, lTransId, " lIntegerByte = %v\n", lIntegerByte)
	xapp.Logger.Info(" lIntegerByte = %v\n", lIntegerByte)
        lOutByte := get_bytepack_plmnId(lIntegerByte)
        xapp.Logger.Info("lOutByte Len:%d\n", len(lOutByte))
        xapp.Logger.Info("lOutByte = %02X\n", lOutByte)
	cptrRanParameterValue := unsafe.Pointer(&lOutByte[0])
	

		pLMNIdentity := C.CString(ueIdData.pLMNIdentitybuf)
	defer C.free(unsafe.Pointer(pLMNIdentity))
        	pLMNIdentity_size := C.size_t(len(ueIdData.pLMNIdentitybuf))
        	aMFRegionID := C.CString(ueIdData.aMFRegionIDbuf)
	defer C.free(unsafe.Pointer(aMFRegionID))
        	aMFSetID := C.CString(ueIdData.aMFSetIDbuf)
	defer C.free(unsafe.Pointer(aMFSetID))
        	aMFPointer := C.CString(ueIdData.aMFPointerbuf)
	defer C.free(unsafe.Pointer(aMFPointer))

	uEID_C := C.struct_uEID{
                amf_UE_NGAP_Id: C.long(ueIdData.amf_UE_NGAP_Id),
                pLMNIdentity : pLMNIdentity,
                pLMNIdentity_size : pLMNIdentity_size,
                aMFRegionID : aMFRegionID,
                aMFRegionID_size : 1, //AMFRegionID ::= BIT STRING (SIZE(8))
                aMFSetID : aMFSetID,
                aMFSetID_size: 2, //BIT STRING (SIZE(10))
                aMFPointer : aMFPointer,
                aMFPointer_size : 1,  //BIT STRING (SIZE(6))
        }

	xapp.Logger.Debug("uEID_C amf_UE_NGAP_Id  = %l",uEID_C.amf_UE_NGAP_Id)
	xapp.Logger.Debug("uEID_C pLMNIdentity = %s",uEID_C.pLMNIdentity)
	xapp.Logger.Debug("uEID_C aMFRegionID %s = ",uEID_C.aMFRegionID)
	xapp.Logger.Debug("uEID_C aMFSetID  %s = ",uEID_C.aMFSetID)
	xapp.Logger.Debug("uEID_C aMFPointer %s = ",uEID_C.aMFPointer)

	size := C.e2sm_encode_ric_control_header(cptr, C.size_t(len(buffer)),&uEID_C ,(*C.long)(&ueIdData.F1AP_id[0]),C.size_t(len(ueIdData.F1AP_id)),(*C.long)(&ueIdData.E1AP_id[0]),C.size_t(len(ueIdData.E1AP_id)),(C.long)(ricControlStyleType), (C.long)(ricControlActionID), cptrRanParameterValue,C.size_t(len(lOutByte)))

	//size := C.e2sm_encode_ric_control_header(cptr, C.size_t(len(buffer)),&uEID_C ,(*C.long)(&ueIdData.F1AP_id[0]),(*C.long)(&ueIdData.E1AP_id[0]),(C.long)(ricControlStyleType), (C.long)(ricControlActionID))
	
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	xapp.Logger.Info("E2sm SetRicControlHeader is success")
	return
}

func (c *E2sm) SetRicControlMessage(buffer []byte, targetPrimaryCell int64, targetCell int64, nrCGIOrECGI int64, nrOrEUtraCell int64, ranParameterValue []byte) (newBuffer []byte, err error) {
	
	xapp.Logger.Info("SetRicControlMessagei Enter  ")

	cptr := unsafe.Pointer(&buffer[0])
	cptrRanParameterValue := unsafe.Pointer(&ranParameterValue[0])

	/*
	lIntvar, _ := strconv.Atoi(string(ranParameterValue))
	xapp.Logger.Info("lIntvar %d\n", lIntvar)
	lIntegerByte := ConverIntegerToByte(lIntvar)
	xapp.Logger.Info("lIntegerByte %v\n", lIntegerByte)
	lOutByte := get_bytepack_plmnId(lIntegerByte)
	xapp.Logger.Info("lOutByte Len:%d\n", len(lOutByte))
	xapp.Logger.Info("lOutByte %02X\n", lOutByte)
	cptrRanParameterValue := unsafe.Pointer(&lOutByte[0])
	xapp.Logger.Info("cptrRanParameterValue = %v\n ", cptrRanParameterValue)
	*/

	size := C.e2sm_encode_ric_control_message(cptr, C.size_t(len(buffer)), C.long(targetPrimaryCell),
		C.long(targetCell), C.long(nrOrEUtraCell), C.long(nrCGIOrECGI), cptrRanParameterValue, C.size_t(len(ranParameterValue)))

	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set RicControlMessage due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8) //TOCHECK: if C.int(size) is returning bits not bytes to get correct size of the buffer.
	xapp.Logger.Info("E2sm SetRicControlMessage is success")
	return
}

func (c *E2sm) GetControlOutcome(buffer []byte) (controlOutcome *ControlOutcomeMsg, err error) {
        cptr := unsafe.Pointer(&buffer[0])
        controlOutcome = &ControlOutcomeMsg{}
	//lTransId := gTransIdDefVal
        //ml.MavLog(ml.INFO, lTransId, "Decode Control Outcome: ")
	xapp.Logger.Info("Decode Control Outcome: ")
        decodedMsg := C.e2sm_decode_ric_call_process_outcome(cptr, C.size_t(len(buffer)))
        if decodedMsg == nil {
                return controlOutcome, errors.New("e2sm wrapper is unable to get ControlOutcome due to wrong or invalid input")
        }
        defer C.e2sm_free_ric_call_process_outcome(decodedMsg)
        //ml.MavLog(ml.INFO, lTransId, "Decoded Control Outcome: ")
	controlOutcome.ControlOutcomeType = int32(decodedMsg.ric_controlOutcome_formats.present)
	if controlOutcome.ControlOutcomeType == 1 {
                controlOutcome := &ControlOutcomeFormat1{}
                controlOutcomeFormat1_C := *(**C.E2SM_RC_ControlOutcome_Format1_t)(unsafe.Pointer(&decodedMsg.ric_controlOutcome_formats.choice[0]))
		controlOutcome.ControlOutcomeFormat1ItemCount = int(controlOutcomeFormat1_C.ranP_List.list.count)
		controlOutcome.ControlOutcomeFormat1Item = make([]ControlOutcomeFormat1ItemType, controlOutcome.ControlOutcomeFormat1ItemCount +1)
                //ml.MavLog(ml.INFO, lTransId, "RIC Control Outcome  Format1 Item Count = %d ", controlOutcome.ControlOutcomeFormat1ItemCount)
                for i := 0; i < controlOutcome.ControlOutcomeFormat1ItemCount; i++ {
			ControlOutcomeFormat1Item :=  &controlOutcome.ControlOutcomeFormat1Item[i]
			var sizeof_E2SM_RC_ControlOutcome_Format1_Item_t *C.E2SM_RC_ControlOutcome_Format1_Item_t
			ControlOutcomeFormat1Item_C :=  *(*(**C.E2SM_RC_ControlOutcome_Format1_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(controlOutcomeFormat1_C.ranP_List.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_E2SM_RC_ControlOutcome_Format1_Item_t))))
			ControlOutcomeFormat1Item.RanParameterID = int64(ControlOutcomeFormat1Item_C.ranParameter_ID)
			ControlOutcomeFormat1Item.RANParameterValueType = int(ControlOutcomeFormat1Item_C.ranParameter_value.present)
		        ranParameterValue_C := (C.RANParameter_Value_t)(ControlOutcomeFormat1Item_C.ranParameter_value)
			if ControlOutcomeFormat1Item.RANParameterValueType == 5 {
				//ml.MavLog(ml.INFO, lTransId, "RANParameterValueType  is RANParameter_Value_PR_valueOctS")
				ranParameterValue := &OctetString{}
				ranParameter_C := (*C.OCTET_STRING_t)(unsafe.Pointer(&ranParameterValue_C.choice[0]))
                                ranParameterValue.Size = int(ranParameter_C.size)
                                ranParameterValue.Buf = C.GoBytes(unsafe.Pointer(ranParameter_C.buf), C.int(ranParameter_C.size))
                                ControlOutcomeFormat1Item.RANParameterValue = ranParameterValue
				//ml.MavLog(ml.INFO, lTransId, "RANParameterValue = ", ControlOutcomeFormat1Item.RANParameterValue)
			}
		}
	} else {
		return controlOutcome, errors.New("Unknown RIC Control Outcome type")
	}
	return controlOutcome , err
}

func (c *E2sm) get_plmn_from_mcc_mnc(plmnIdBuf []byte)(bytePackedPlmnId []uint8) {
	//lTransId := gTransIdDefVal
	hexResp, _ := hex.DecodeString(string(plmnIdBuf))
	//ml.MavLog(ml.INFO, lTransId, "get_plmn_from_mcc_mnc plmnIdBuf and hexResp = ", plmnIdBuf,hexResp)
	var mcc []uint8 = make([]uint8, 3)
        var mnc []uint8 = make([]uint8, 3)

        mcc[0] = hexResp[0] >> 4
        mcc[1] = hexResp[0] & 0xf
        mcc[2] = hexResp[1] >> 4
        mnc[0] = hexResp[1] & 0xf
        mnc[1] = hexResp[2] >> 4
        mnc[2] = hexResp[2] & 0xf

        var tmpPlmnId []uint8 = make([]uint8, 6)
        tmpPlmnId[0] = mcc[0]
        tmpPlmnId[1] =  mcc[1]
        tmpPlmnId[2] =  mcc[2]
      if len(mcc) == 2 {
                tmpPlmnId[3] = 0x0f
                tmpPlmnId[4] = mnc[0]
                tmpPlmnId[5] = mnc[1]
        } else {
                tmpPlmnId[3] = mnc[0]
                tmpPlmnId[4] = mnc[1]
                tmpPlmnId[5] = mnc[2]
        }
        //ml.MavLog(ml.INFO, lTransId, "\n tmpPlmnId " , tmpPlmnId)
        bytePackedPlmnId = make([]uint8,3)
        bytePackedPlmnId[0] = ((tmpPlmnId[1] & 0x0F) << 4) | (tmpPlmnId[0] & 0x0F)
        bytePackedPlmnId[1] = ((tmpPlmnId[3] & 0x0F) << 4) | (tmpPlmnId[2] & 0x0F)
        bytePackedPlmnId[2] = ((tmpPlmnId[5] & 0x0F) << 4) | (tmpPlmnId[4] & 0x0F)
        //ml.MavLog(ml.INFO, lTransId, "\n bytePackedPlmnId " , bytePackedPlmnId)
        return
}

