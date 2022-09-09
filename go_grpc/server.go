package main

import (
	"fmt"
	"google.golang.org/grpc"
	// "log"
	"net"
)

func main() {

	fmt.Println("start grpc")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		// log.Fatalf("Faild to listen on port 9000: %v", err)
		fmt.Println("server failed to running on port 9000")
	}

	fmt.Println("port is okay")

	grpcServer := grpc.NewServer()

	fmt.Println("can create grpc server")

	if err := grpcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to server grpc server on port 9000: %v", err)
		fmt.Println("server failed")
	}

}
