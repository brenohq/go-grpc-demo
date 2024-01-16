package main

import (
	"log"

	pb "github.com/brenohq/go-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)
}
