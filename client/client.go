package main

import (
	"context"
	"fmt"

	pb "github.com/Markogoodman/grpctest/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

func main() {
	conn, _ := grpc.Dial(":8001", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "marko"})
	fmt.Println(resp.Message)
	return nil
}
