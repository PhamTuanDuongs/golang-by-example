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
	startTime := time.Now()
	conn, err := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	//  creat a channel
	done := make(chan bool)

	// create stream
	client := chat.NewChatClient(conn)
	stream, err := client.ChatA(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	var tag string
	var mess string

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

	err = stream.Send(&chat.Request{
		Tag:  tag,
		Mess: mess,
	})
	if err != nil {
		log.Println("Send error")
		return
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
			log.Printf("Response from server A: %v", value)
		}
	}()

	<-done
	elapsedTime := time.Since(startTime)
	log.Printf("RPC call completed in %s", elapsedTime)
	log.Printf("Finished")
}
