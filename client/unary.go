package main

import (
	"context"
	pb "grpc_greet/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not get greet  %v", err)
	}
	log.Printf("%s\n", resp.Message)
}
