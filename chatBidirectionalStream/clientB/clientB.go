package main

import (
	chat "chatbidirection/genproto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish a gRPC connection to the server at the specified address ":9999".
	// Use insecure transport credentials (for testing purposes only).

	conn, err := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can not connect to server %v", err)
	}

	//  creat a channel to mark done of program
	stuck := make(chan bool)
	// create stream
	client := chat.NewServiceChatClient(conn)
	chReceiVal := make(chan string)
	chCheckContext := make(chan string)

	clientB, err := client.ChatB(context.Background())
	go func() {
		// Dat vong for o day khong bi hieu nang, boi vi chi khi nhan duoc messgage thi no moi run tiep code phia sau clientB.Recv()
		for {
			if err != nil {
				log.Fatalln(err)
			}
			res, _ := clientB.Recv()
			if res != nil {
				chReceiVal <- res.Mess
			} else {
				chCheckContext <- clientB.Context().Err().Error()
			}
		}
	}()

	// create a goroutine to dis continuously data
	// Start a new goroutine to run the following function concurrently.
	go func() {
		// Infinite loop to continuously listen for messages on channels.
		for {
			select {
			// Case: Receive a message from the chReceiVal channel.
			case msg := <-chReceiVal:
				fmt.Println(msg)

			// Case: The clientB context is done (indicating termination or cancellation).
			case <-clientB.Context().Done():
				// Signal that the goroutine is no longer stuck by sending false to the 'stuck' channel.
				stuck <- false

				// Print the result received on the 'chCheckContext' channel.
				fmt.Println(<-chCheckContext)

				// Exit the goroutine.
				return
			}
		}
	}()

	<-stuck
}
