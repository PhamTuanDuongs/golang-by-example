package main

import (
	chat "chatbidirection/genproto"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish a gRPC connection to the server at the specified address ":9999".
	// Use insecure transport credentials (for testing purposes only).
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Can not connect to server %v", err)
	}

	//  creat a channel to mark done of program
	done := make(chan bool)

	// Create a gRPC client for the chat service using the established connection.
	client := chat.NewServiceChatClient(conn)

	// Open a bidirectional streaming RPC for the "ChatA" operation.
	// The stream variable will be used to send and receive messages.
	stream, err := client.ChatC(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Declare tag, mess
	var tag, mess string
	fmt.Println("Who do you want send to?")
	_, err = fmt.Scanln(&tag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Write your message")
	_, err = fmt.Scanln(&mess)
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()
	//Send data to Server
	err = stream.Send(&chat.Request{
		Tag:  tag,
		Mess: mess,
	})
	//Check processing sending data to Server
	if err != nil {
		log.Println(err)
		return
	} else {
		elapsedTime := time.Since(startTime)
		log.Printf("RPC call completed in %s", elapsedTime)
	}

	// creat a goroutine
	go func() {
		// receive continuously data on a stream
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			value := resp.Mess
			log.Printf("Response from server C: %v", value)
		}
	}()

	<-done
	log.Printf("Finished")
}
