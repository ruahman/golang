/***
 * protoc --go_out=. *.proto
 */

package main

import (
	"fmt"
	"go_grpc/chat"

	"google.golang.org/grpc"

	// "log"
	"net"
)

func main() {

	// test := chat.Person{}
	// test.Age = 32
	// test.Name = "diego"

	fmt.Println("start grpc")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		// log.Fatalf("Faild to listen on port 9000: %v", err)
		fmt.Println("server failed to running on port 9000")
	}

	fmt.Println("port is okay")

	s := chat.Server{}

	grpcServer := grpc.NewServer()

  chat.RegisterChatServiceServer(grpcServer, &s)

	fmt.Println("can create grpc server")

	if err := grpcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to server grpc server on port 9000: %v", err)
		fmt.Println("server failed")
	}

}
