package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (res *proto.AddResponse, er error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add Operation invoked with x=%d and y=%d\n", x, y)
	timeOut := time.After(10 * time.Second)

LOOP:
	for {
		select {
		case <-timeOut:
			result := x + y
			res = &proto.AddResponse{
				Result: result,
			}
			break LOOP
		case <-ctx.Done():
			fmt.Println("Cancel instruction received")
			er = errors.New("interrupt received")
			break LOOP
		}
	}
	return
}

func (asi *appServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Generated prime no : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			serverStream.Send(res)
		}
	}
	return nil
}

func (asi *appServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	fmt.Println("Calculate Average")
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		no := req.GetNo()
		fmt.Printf("Received : %d\n", no)
		sum += no
		count++
		time.Sleep(1 * time.Second)
	}
	avg := sum / count
	res := &proto.AverageResponse{
		Result: avg,
	}
	serverStream.SendAndClose(res)
	return nil
}

func (asi *appServiceImpl) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		personName := req.GetPerson()
		msg := fmt.Sprintf("Hi %s, %s!", personName.GetFirstName(), personName.GetLastName())
		res := &proto.GreetResponse{
			GreetMessage: msg,
		}
		er := stream.Send(res)
		if er != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	asi := &appServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		certFile := "ssl-2/server.crt"
		keyFile := "ssl-2/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalln(sslErr)
		}
		opts := grpc.Creds(creds)
		grpcServer := grpc.NewServer(opts)
	*/
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
