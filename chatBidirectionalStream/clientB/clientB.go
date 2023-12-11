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
	conn, err := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	stuck := make(chan bool)
	// create stream
	client := chat.NewServiceChatClient(conn)
	chReceiVal := make(chan string)
	clientB, err := client.ChatB(context.Background())
	go func() {
		cout := 0
		// Dat vong for o day khong bi hieu nang, boi vi chi khi nhan duoc messgage thi no moi run tiep code phia sau clientB.Recv()
		for {
			cout++
			if err != nil {
				log.Fatalf("openn stream error %v", err)
			}
			res, _ := clientB.Recv()
			if res != nil {
				chReceiVal <- res.Mess
			} else {
				stuck <- true
				chReceiVal <- clientB.Context().Err().Error()
			}
			fmt.Println(cout)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-chReceiVal:
				fmt.Println(msg)
			case <-clientB.Context().Done():
				fmt.Println("Stream is cancled")
				return
			}
		}
	}()
	<-stuck
}

// go func() {
// 	stream, err := client.ChatB(context.Background())
// 	if err != nil {
// 		log.Fatalf("openn stream error %v", err)
// 	}

// 	var w1 string
// 	n, err := fmt.Scanln(&w1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("number of items read: %d\n", n)
// 	err = stream.Send(&chat.Request{
// 		Tag:  w1,
// 		Mess: "Xin chao client " + w1,
// 	})

// 	if err != nil {
// 		log.Println("Send error")
// 		return
// 	}
// }()
