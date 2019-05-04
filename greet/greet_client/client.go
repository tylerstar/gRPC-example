package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"
	"time"
)

func main() {

	fmt.Println("Hello, I'm a client.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDirectionalStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName: "Smith",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v\n", err)
	}
	log.Printf("Response from Greet: %v", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start to do a server streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName: "Smith",
		},
	}
	resStreams, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v\n", err)
	}
	for {
		msg, err := resStreams.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming rpc...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Jane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling the LongGreet: %v\n", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send message: %v\n", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongFreet: %v\n", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBiDirectionalStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start to do a Bi Directional streaming RPC...")

	// Create a stream by invoking the cline
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Jane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})
	// Send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send message to the server: %v\n", err)
				close(waitc)
			}
			time.Sleep(time.Millisecond * 1000)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Failed to close the stream: %v\n", err)
			close(waitc)
		}
	}()

	// Receive a bunch of messages from the cline (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// Block until everything is done
	<-waitc
}


