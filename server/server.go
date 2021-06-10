package main

import (
	"context"
	"fmt"
	"io"
	"net"

	pb "github.com/Markogoodman/grpctest/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world " + r.Name}, nil
}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for i := 0; i < 5; i++ {
		stream.Send(&pb.HelloReply{Message: fmt.Sprintf("%d hello %s", i, r.Name)})
	}
	return nil
}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "Server receive all"})
		}
		if err != nil {
			return err
		}
		fmt.Println("Server receive", req.Name)
	}
}
func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8001")
	server.Serve(lis)
}
