package main

import (
	pb "grpc_greet/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request names : %v\n", req)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		log.Printf("sent message to - %v\n", name)
		time.Sleep(2 * time.Second)
	}
	return nil
}
