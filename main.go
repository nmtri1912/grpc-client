package main

import (
	"context"
	"grpc-client/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())

	if err != nil {
		log.Printf("Cannot connect grpc server")
	}

	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	message := chat.RequestMessage{
		Body: "hello from client",
	}

	response, err := client.SayHello(context.Background(), &message)
	if err != nil {
		log.Println("error when process method SayHello: %s", err)
	}

	log.Printf("Response from grpc server: %s", response.Body)
}
