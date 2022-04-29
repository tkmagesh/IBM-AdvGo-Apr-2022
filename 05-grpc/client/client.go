package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	/*
		certFile := "ssl-2/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalln(sslErr)
		}
		options := grpc.WithTransportCredentials(creds)
	*/
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	service := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	doRequestResponse(ctx, service)
	//doRequestResponseWithInterrupt(ctx, service)
	//doServerStreaming(ctx, service)
	//doClientStreaming(ctx, service)
	//doBiDiStreaming(ctx, service)
	//doRequestResponseWithTimeout(ctx, service)
}

func doRequestResponse(ctx context.Context, service proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := service.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Operation : Result =", addResponse.GetResult())
}

func doRequestResponseWithInterrupt(ctx context.Context, service proto.AppServiceClient) {
	reqCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		fmt.Println("Hit ENTER to cancel....")
		var input string
		fmt.Scanln(&input)
		cancel()
	}()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := service.Add(reqCtx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Operation : Result =", addResponse.GetResult())
}

func doRequestResponseWithTimeout(ctx context.Context, service proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeOutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	res, err := service.Add(timeOutCtx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout error")
			} else {
				log.Fatalln(err)
			}
		}
		log.Fatalln(err)
	}
	fmt.Println("Result = ", res.GetResult())
}

func doServerStreaming(ctx context.Context, service proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := service.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No :=%d\n", res.GetPrimeNo())
	}
	fmt.Println("Done")
}

func doClientStreaming(ctx context.Context, service proto.AppServiceClient) {
	nos := []int32{3, 1, 4, 2, 5, 9, 6, 8, 7}
	clientStream, err := service.CalculateAverage(ctx)
	if err != nil {
		log.Println("Connection error")
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Sending no : %d\n", no)
		req := &proto.AverageRequest{
			No: no,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average = %d\n", res.GetResult())
}

func doBiDiStreaming(ctx context.Context, service proto.AppServiceClient) {
	personNames := []proto.PersonName{
		proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		proto.PersonName{FirstName: "Ganesh", LastName: "Easwaran"},
		proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := service.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	//handle responses
	done := func() <-chan struct{} {
		doneCh := make(chan struct{})
		go func() {
			for {
				res, err := clientStream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
				msg := res.GetGreetMessage()
				fmt.Println(msg)
			}
			close(doneCh)
		}()
		return doneCh
	}()

	for _, personName := range personNames {
		fmt.Printf("Sending Person %v\n", personName)
		time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	clientStream.CloseSend()
	<-done
}
