package main

import (
	"context"
	pb "grpc_greet/proto"
	"log"
	"time"
)

func callSayHelloClientStreaming(names *pb.NameList, client pb.GreetServiceClient) error {
	log.Println("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err = stream.Send(req); err != nil {
			log.Printf("Erro whil sending %v", err)
		}
		log.Printf("Sent the request name with %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Println("stream finished")
	if err != nil {
		return err
	}
	log.Printf("%v", res.Messages)
	return nil
}
