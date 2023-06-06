package main

import (
	"context"
	pb "grpc_greet/proto"
	"io"
	"log"
)

func callSayHelloServerStreaming(names *pb.NameList, client pb.GreetServiceClient) {
	log.Printf("streaming started...")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not sends names : %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Println("streaming finished")
}
