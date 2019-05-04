package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
)

func main() {
	fmt.Println("Client is starting...")
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	//doSum(c)
	//doNumberDecompoisition(c)
	doComputeAverage(c)
}

func doSum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start to do a sum request...")
	req := &calculatorpb.SumRequest{
		Numbers: &calculatorpb.Numbers{
			FirstNumber: 5,
			LastNumber: 6,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call sum: %v", err)
	}
	fmt.Printf("Response from server: %v \n", res)
}

func doNumberDecompoisition(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start to do a number decomposition request...")
	req := &calculatorpb.NumberDecompositionRequest{
		Numbers: &calculatorpb.Numbers{
			FirstNumber: 2,
			LastNumber: 120,
		},
	}
	resStream, err := c.DecomposeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling DecomposeNumber: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading message: %v", err)
		}
		log.Printf("Response from DecomposeNumber: %v", msg)
	}
}

func doComputeAverage(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start to do a ComputeAverage request...")
	requests := []*calculatorpb.ComputeAverageRequest{
		{ Number: 1 },
		{ Number: 2 },
		{ Number: 3 },
		{ Number: 4 },
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverage function: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send req: %v\n", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response from ComputeAverage: %v", err)
	}
	fmt.Printf("Response from ComputeAverage: %v\n", res)
}