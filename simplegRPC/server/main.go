package main

import (
	"context"
	"fmt"
	"log"
	"net"
	calculator "simplegrpc/proto"

	// Assume this is the package generated from the protobuf file

	"google.golang.org/grpc"
)

type calculatorService struct {
	calculator.UnimplementedCalculatorServer
}

func (s *calculatorService) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.Operand1 + req.Operand2

	return &calculator.AddResponse{
		Result: result,
		Status: "OK",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	calculator.RegisterCalculatorServer(server, &calculatorService{})

	fmt.Println("Server is listening on port 9999")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
