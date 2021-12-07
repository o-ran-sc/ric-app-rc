package control

import (
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	"gerrit.o-ran-sc.org/r/ric-app/rc/protocol/grpc/ricmsgcommrpc/rc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

// RCControlServer represents the grpc server
type RCControlServer struct {
}

//Start gRPC Server For receiving Control   Messages from dataingestion
func StartgRPCRCControlCommServerRoutine() error {

	//log.Println("Starting Go Routine for Handling gRPC Server for handling gRPC RCControl Handling")
	xapp.Logger.Info("Starting Go Routine for Handling gRPC Server for handling gRPC RCControl Handling ")
	port := xapp.Config.GetString("controls.ricHOControlgRpcServerPort")
	xapp.Logger.Info("GrpcServer started listening on Port port =  %v", port)
	lConnStr := ":" + port

	// create a listener on TCP port Configured
	lLis, lErr := net.Listen("tcp", lConnStr)
	if lErr != nil {
		//log.Printf("GrpcServer Listen failed with error", lErr)
		xapp.Logger.Info("GrpcServer Listen failed with error = %v", lErr)
		return lErr
	}

	xapp.Logger.Info("GrpcServer started listening on Connection: %v", lConnStr)
	// create a server instance
	lRCControlSrv := RCControlServer{}

	// create a grpc server object
	lgRPCServer := grpc.NewServer()

	// attach the Ping service to the server
	rc.RegisterMsgCommServer(lgRPCServer, &lRCControlSrv)
	reflection.Register(lgRPCServer)

	//Register to health service
	grpc_health_v1.RegisterHealthServer(lgRPCServer, health.NewServer())

	xapp.Logger.Debug("GrpcServer Serve start with port = %v and lConnStr = %v and lgRPCServer = %v ", lLis, lConnStr, lgRPCServer)

	// start the server
	if lErr := lgRPCServer.Serve(lLis); lErr != nil {
		xapp.Logger.Info("GrpcServer Serve failed with error", lErr)
		return lErr
	}

	return nil
}
