package main

import (
	"context"
	"fmt"
	"log"

	calculator "simplegrpc/proto" // Assume this is the package generated from the protobuf file

	"google.golang.org/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	req := &calculator.AddRequest{
		Operand1: 0,
		Operand2: 0,
		Status:   "1",
	}
	size := proto.Size(req)
	fmt.Println(size)

	response, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}

	fmt.Printf("Result: %d\n", response.Result)
}
