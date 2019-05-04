package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"time"
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
	//doComputeAverage(c)
	//doFindMaxNumber(c)
	doErrorUnary(c)
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

func doFindMaxNumber(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start to send a FindMaxNumberRequest...")

	requests := []*calculatorpb.FindMaxNumberRequest{
		{ Number: 1 },
		{ Number: 5 },
		{ Number: 3 },
		{ Number: 6 },
		{ Number: 2 },
		{ Number: 20 },
	}

	waitc := make(chan struct{})
	// Create a stream
	stream, err := c.FindMaxNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to invoke function FindMaxNumber: %v", err)
	}

	// Send messages to the client
	go func() {
		for _, req := range requests {
			fmt.Printf("Send request: %v\n", req)
			if err = stream.Send(req); err != nil {
				log.Fatalf("Failed to send request to the server.\n")
				close(waitc)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("Failed to close the stream: %v\n", err)
			close(waitc)
		}
	}()

	// Receive messages from the client
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive message from the server: %v\n", err)
				break
			}
			fmt.Printf("Receive: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// Block the flow
	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start to do a SquareRoot Unary RPC...")

	// Correct call
	doErrorCall(c, 10)

	// Error call
	doErrorCall(c, -2)
}


func doErrorCall(c calculatorpb.CalculatorServiceClient, number int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: number})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v\n", err)
			return
		}
	}
	fmt.Printf("Result of sqaure root if %v: %v\n", number, res.GetNumberRoot())
}