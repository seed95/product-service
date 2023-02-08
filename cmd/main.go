package main

import (
	"fmt"
	nativeLog "log"
	"net"
)

func main() {

	factory, err := NewServerFactory()
	if err != nil {
		nativeLog.Fatal(err)
	}
	config := factory.Config
	grpcServer := factory.GRPRServer

	// Running gRPC Server
	listener, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		nativeLog.Fatal("cannot create grpc server: ", err)
	}

	fmt.Printf("Running gRPC server on port %s\n", config.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		nativeLog.Fatalf("failed ro bind gRPC server on port %s, error: %s", config.GRPCPort, err.Error())
	}

}
