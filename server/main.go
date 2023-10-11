package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	fmt.Println("Starting...")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(lis, grpcServer

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port", port)


}
