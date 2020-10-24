package main

import (
	"flag"
	"fmt"
	"golang-grpc-examples/messages/messenger"
	"golang-grpc-examples/server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", ":5000", "set service port")
}

func main() {
	flag.Parse()

	log.Printf("Service listening on port %v", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	grpcServer := grpc.NewServer()

	messengerSvc := service.NewMessengerService()

	messenger.RegisterMessengerServiceServer(grpcServer, messengerSvc)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
