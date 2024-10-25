package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "simple-grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthUserServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}

	return &pb.HelloReply{Message: "Hello, " + in.Name}, nil
}

func (s *server) CheckToken(ctx context.Context, in *pb.TokenRequest) (*pb.TokenReply, error) {
	if in.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Token is required")
	}

	const dummyToken = "beneboba"
	if in.Token != dummyToken {
		return &pb.TokenReply{Status: false, Message: "Token is invalid"}, nil
	}

	return &pb.TokenReply{Status: true, Message: "Token valid success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthUserServer(s, &server{})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
