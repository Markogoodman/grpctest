package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/Markogoodman/grpctest/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8001", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	//  SayHello(client)
	// SayList(client)
	SayRecord(client)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "marko"})
	fmt.Println(resp.Message)
	return nil
}

func SayList(client pb.GreeterClient) error {
	stream, _ := client.SayList(context.Background(), &pb.HelloRequest{Name: "marko"})
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(resp.Message)
	}

	return nil
}

func SayRecord(client pb.GreeterClient) error {
	stream, _ := client.SayRecord(context.Background())
	for i := 0; i < 5; i++ {
		stream.Send(&pb.HelloRequest{Name: fmt.Sprint("marko", i)})
	}
	resp, _ := stream.CloseAndRecv()
	fmt.Println("Client receive", resp)
	return nil
}
