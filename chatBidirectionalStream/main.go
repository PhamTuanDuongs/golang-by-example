package main

import (
	chat "chatbidirection/genproto"
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
	msg, err := ch.Recv()

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Message is received from A " + msg.Mess)
	c.mut.Lock()

	// add an new element to map
	c.CacheMessage[msg.GetTag()] = msg.Mess
	c.mut.Unlock()

	value, ok := c.CacheMessage["A"]
	if ok && value != "" {
		startTime := time.Now()
		errSend := ch.Send(&chat.Response{
			Mess: "Hi client " + msg.Tag + ": " + " I am server A",
		})
		if errSend != nil {
			fmt.Println(errSend)
		} else {
			elapsedTime := time.Since(startTime)
			log.Printf("Request processed in %s", elapsedTime)
			fmt.Println("send to A done")
			c.mut.Lock()
			c.CacheMessage["A"] = ""
			c.mut.Unlock()
		}
	}

	return nil
}



func (c *chatService) ChatB(ch chat.ServiceChat_ChatBServer) error {
	fmt.Println("Have connection from client B")
	msg, err := ch.Recv()

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Message is received from B " + msg.Mess)
	c.mut.Lock()

	// add an new element to map
	c.CacheMessage[msg.GetTag()] = msg.Mess
	c.mut.Unlock()

	value, ok := c.CacheMessage["B"]
	if ok && value != "" {
		startTime := time.Now()
		errSend := ch.Send(&chat.Response{
			Mess: "Hi client " + msg.Tag + ": " + " I am server B",
		})
		if errSend != nil {
			fmt.Println(errSend)
		} else {
			elapsedTime := time.Since(startTime)
			log.Printf("Request processed in %s", elapsedTime)
			fmt.Println("send to B done")
			c.mut.Lock()
			c.CacheMessage["B"] = ""
			c.mut.Unlock()
		}
	}

	return nil
}

func (c *chatService) ChatC(ch chat.ServiceChat_ChatCServer) error {
	fmt.Println("Have connection from client C")
	msg, err := ch.Recv()

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Message is received from C: " + msg.Mess)
	c.mut.Lock()

	// add an new element to map
	c.CacheMessage[msg.GetTag()] = msg.Mess
	c.mut.Unlock()

	value, ok := c.CacheMessage["C"]
	if ok && value != "" {
		startTime := time.Now()
		errSend := ch.Send(&chat.Response{
			Mess: "Hi client " + msg.Tag + ": " + " I am server C",
		})
		if errSend != nil {
			fmt.Println(errSend)
		} else {
			elapsedTime := time.Since(startTime)
			log.Printf("Request processed in %s", elapsedTime)
			fmt.Println("send to C done")
			c.mut.Lock()
			c.CacheMessage["C"] = ""
			c.mut.Unlock()
		}
	}

	return nil
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