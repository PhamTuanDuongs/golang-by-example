package main

import (
	chat "chatbidirection/genproto"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
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

// func to receive and send message to clientB
func (c *chatService) ChatB(ch chat.ServiceChat_ChatBServer) error {
	fmt.Println("Have connection from client B")

	//	Create a channel to store values be sent from the client
	storeValue := make(chan string)

	// goroutine to handle sending messages to the client
	go func(receiveOnlyValue chan string, ctx context.Context) {
		for {
			select {
			case msg := <-receiveOnlyValue:
				startTime := time.Now()
				err := ch.Send(&chat.Response{
					Mess: msg,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("send to B done")
					elapsedTime := time.Since(startTime)
					log.Printf("Request processed in %s", elapsedTime)
				}
			case <-ctx.Done():
				fmt.Println(ctx.Err().Error() + " B")
				return
			}
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

func performDangerousOperation(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Performing dangerous operation")

	// Simulating a panic condition
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("Recovered from panic:", r)
	// 	}
	// }()
	// listen, err := net.Listen("tcp", ":9999")
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }
	// // Create a new server from grpc
	// server := grpc.NewServer()

	// // Call Register ChatServer from protobuf
	// chat.RegisterChatServer(server, &chatService{
	// 	CacheMessage: map[string]string{},
	// 	mut:          sync.Mutex{},
	// })

	// fmt.Println("Server is listening on port 9999")
	// if err := server.Serve(listen); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
	// Example function that may panic
	// performDangerousOperation()
	arr := [3]int{1, 2, 3}

	// Truy cập phần tử nằm ngoài giới hạn của mảng (index 3)
	// Điều này sẽ gây panic
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}

	// This code won't be executed because of the panic
	fmt.Println("This line will not be reached")
}

type Person struct {
	Id int32
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go performDangerousOperation(&wg)
	// fmt.Println("This line will not be reached")
	// fmt.Println("Program continues execution after recovering from panic")
	// wg.Wait()

	// Truy cập thông tin từ con trỏ nil
	value := Person{2}
	value = Person{}

	// Dòng này sẽ không được thực hiện vì đã có panic trước đó
	fmt.Println(value)

}

/*
https://stackoverflow.com/questions/25356602/golang-panic-crash-prevention
*/
