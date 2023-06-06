package main

import (
	"context"
	pb "grpc_greet/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirStreaming(names *pb.NameList, client pb.GreetServiceClient) error {
	log.Println("Bidirectional streaming started")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}
	waitc := make(chan struct{})
	go func() {
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
		close(waitc)
	}()

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
	if err := stream.CloseSend(); err != nil {
		return err
	}
	<-waitc
	log.Println("Bidirectional stream finished")
	return nil
}
