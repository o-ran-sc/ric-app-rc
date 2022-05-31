#./grpcurl -plaintext -d "{ \"e2NodeID\": \"36000000\", \"plmnID\": \"111\", \"ranName\": \"gnb_131_133_36000000\", \"RICE2APHeaderData\": { \"RanFuncId\": 300, \"RICRequestorID\": 2 }, \"RICControlHeaderData\": { \"ControlStyle\": 3, \"ControlActionId\": 1, \"UEID\": \"00006\" }, \"RICControlMessageData\": { \"RICControlCellTypeVal\": 4, \"TargetCellID\": \"1113\" }, \"RICControlAckReqVal\": 0 }"  10.96.106.3:7777 rc.MsgComm.SendRICControlReqServiceGrpc

#10.96.106.3:7777  is the grpc server ip and port
#Values of other parameters can be provided as shown above

./grpcurl -plaintext -d "{ \"e2NodeID\": \"gnb_208_092_303030\", \"plmnID\": \"111\", \"ranName\": \"gnb_208_092_303030\", \"RICE2APHeaderData\": { \"RanFuncId\": 300, \"RICRequestorID\": 2 }, \"RICControlHeaderData\": { \"ControlStyle\": 3, \"ControlActionId\": 1, \"UEID\": {\"GnbUEID\": { \"amfUENGAPID\": 10 , \"guami\": {\"pLMNIdentity\": \"111111\",\"aMFRegionID\": \"00001001\", \"aMFSetID\": \"000000 \", \"aMFPointer\": \" 000\"  },\"gNBCUUEF1APID\": 2, \"gNBCUCPUEE1APID\": 1 }} }, \"RICControlMessageData\": { \"RICControlCellTypeVal\": 4, \"TargetCellID\": \"1113\" }, \"RICControlAckReqVal\": 1 }" 10.99.225.43:7777 rc.MsgComm.SendRICControlReqServiceGrpc

