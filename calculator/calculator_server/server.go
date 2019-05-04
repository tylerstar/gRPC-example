package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"net"
	"time"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with: %v\n", req)
	firstNumber := req.GetNumbers().GetFirstNumber()
	lastNumber := req.GetNumbers().GetLastNumber()
	result := firstNumber + lastNumber
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func (*server) DecomposeNumber(req *calculatorpb.NumberDecompositionRequest, stream calculatorpb.CalculatorService_DecomposeNumberServer) error {
	fmt.Printf("DecomposeNumber function was invoked with: %v\n", req)
	firstNumber := req.GetNumbers().GetFirstNumber()
	lastNumber := req.GetNumbers().GetLastNumber()

	for {
		if lastNumber <= 1 {
			break
		}
		if lastNumber % firstNumber == 0 {
			result := firstNumber
			res := &calculatorpb.NumberDecompositionResponse{
				Result: result,
			}
			if err := stream.Send(res); err != nil {
				log.Fatalf("Failed to send: %v, error: %v", res, err)
			}
			lastNumber /= firstNumber
		} else {
			firstNumber += 1
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Start to read stream from ComputeAverage...")

	var sum, count float64 = 0, 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: sum / count,
			})
		}
		if err != nil {
			log.Fatalf("Failed to read the message from the stream: %v", err)
		}
		number := msg.GetNumber()
		sum += float64(number)
		count += 1
	}
}

func (*server) FindMaxNumber(stream calculatorpb.CalculatorService_FindMaxNumberServer) error {
	fmt.Println("FindMaxNumber function was invoked...")
	var maxNumber int32 = 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to reading streaming: %v", err)
			return err
		}
		number := res.GetNumber()
		if number > maxNumber {
			maxNumber = number
			err = stream.Send(&calculatorpb.FindMaxNumberResponse{
				Result: maxNumber,
			})
			if err != nil {
				log.Fatalf("Failed to send message to the clinet: %v", err)
				return err
			}
		}
	}
}

func main() {
	fmt.Println("Calculator Server is starting...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen to %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
