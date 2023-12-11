package main

import (
	chat "chatbidirection/genproto"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// Create a chatService struct
type chatService struct {
	CacheMessage map[string]string
	chat.UnimplementedServiceChatServer
	mut sync.Mutex
}

// func to receive and send message to clientA
func (c *chatService) ChatA(ch chat.ServiceChat_ChatAServer) error {
	fmt.Println("Have connection from client A")
	startTime := time.Now()

	msg, err := ch.Recv()
	if err != nil {
		log.Println(err)
		log.Println("Co loi ro")
		return err
	}
	fmt.Println("Message is received from A " + msg.Mess)
	c.mut.Lock()
	// write receiver message
	c.CacheMessage[msg.GetTag()] = msg.Mess
	c.mut.Unlock()
	value, ok := c.CacheMessage["A"]
	if ok && value != "" {
		errSend := ch.Send(&chat.Response{
			Mess: "Xin chao client " + msg.Tag + ": " + " toi la server A day",
		})
		if errSend != nil {
			fmt.Println(errSend)
		} else {
			fmt.Println("send to A done")
			c.mut.Lock()
			c.CacheMessage["A"] = ""
			c.mut.Unlock()
		}
	}
	elapsedTime := time.Since(startTime)
	log.Printf("Request processed in %s", elapsedTime)
	return nil
}

// func to receive and send message to clientB
func (c *chatService) ChatB(ch chat.ServiceChat_ChatBServer) error {
	fmt.Println("Have connection from client B")
	startTime := time.Now()

	//	Create a channel to store values be sent from the client
	storeValue := make(chan string)

	// goroutine to handle sending messages to the client
	go func(receiveOnlyValue chan string, ctx context.Context) {
		for {
			cout := 1
			select {
			case msg := <-receiveOnlyValue:
				err := ch.Send(&chat.Response{
					Mess: msg,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("send to B done")
				}
				elapsedTime := time.Since(startTime)
				log.Printf("Request processed in %s", elapsedTime)
				fmt.Println("Receive from B")
			case <-ctx.Done():
				fmt.Println("Ngat ket noi nhe")
				return
			}
			cout++
			fmt.Println(cout)
		}
	}(storeValue, ch.Context())

	for {
		c.mut.Lock()
		value, ok := c.CacheMessage["B"]
		c.mut.Unlock()
		if len(value) > 0 && ok {
			storeValue <- value
			c.mut.Lock()
			c.CacheMessage["B"] = ""
			c.mut.Unlock()
		}
	}

}

func (c *chatService) ChatC(ch chat.ServiceChat_ChatCServer) error {
	fmt.Println("Have connection from Client C")
	startTime := time.Now()

	//	Create a channel to store values be sent from the client
	storeValue := make(chan string)

	// goroutine to handle sending messages to the client
	go func(receiveOnlyValue chan string, ctx context.Context) {
		for {
			cout := 1
			select {
			case msg := <-receiveOnlyValue:
				err := ch.Send(&chat.Response{
					Mess: msg,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("send to C done")
				}
				elapsedTime := time.Since(startTime)
				log.Printf("Request processed in %s", elapsedTime)
			case <-ctx.Done():
				return
			}
			cout++
			fmt.Println(cout)
		}
	}(storeValue, ch.Context())

	for {
		c.mut.Lock()
		value, ok := c.CacheMessage["C"]
		c.mut.Unlock()
		if len(value) > 0 && ok {
			storeValue <- value
			c.mut.Lock()
			c.CacheMessage["C"] = ""
			c.mut.Unlock()
		}
	}

}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Create a new server from grpc
	server := grpc.NewServer()

	// Call Register ChatServer from protobuf
	chat.RegisterServiceChatServer(server, &chatService{
		CacheMessage: map[string]string{},
		mut:          sync.Mutex{},
	})

	fmt.Println("Server is listening on port 9999")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
