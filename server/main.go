package main

import (
	"google.golang.org/grpc"
	pb "grpc_greet/proto"
	"log"
	"net"
)

const port = ":8085"

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server listening at %v\n", lis.Addr())
	if sErr := grpcServer.Serve(lis); sErr != nil {
		log.Fatalf("Failed to serve : %v", sErr)
	}
}
