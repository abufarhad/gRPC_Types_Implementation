package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc_greet/proto"
	"log"
)

const address = ":8085"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	//unary
	//callSayHello(client)

	names := &pb.NameList{
		Names: []string{"Faisal vai", "Shafin", "Zihan"},
	}
	
	//ServerStreaming
	//callSayHelloServerStreaming(names, client)
	
	//ClientStreaming
	//callSayHelloClientStreaming(names, client)
	
	//BidirectionalStreaming
	callSayHelloBiDirStreaming(names, client)
}
