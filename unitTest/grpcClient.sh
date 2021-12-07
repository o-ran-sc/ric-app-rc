./grpcurl -plaintext -d "{ \"e2NodeID\": \"36000000\", \"plmnID\": \"111\", \"ranName\": \"gnb_131_133_36000000\", \"RICE2APHeaderData\": { \"RanFuncId\": 300, \"RICRequestorID\": 2 }, \"RICControlHeaderData\": { \"ControlStyle\": 3, \"ControlActionId\": 1, \"UEID\": \"00006\" }, \"RICControlMessageData\": { \"RICControlCellTypeVal\": 4, \"TargetCellID\": \"1113\" }, \"RICControlAckReqVal\": 0 }"  10.96.106.3:7777 rc.MsgComm.SendRICControlReqServiceGrpc

#10.96.106.3:7777  is the grpc server ip and port
#Values of other parameters can be provided as shown above

