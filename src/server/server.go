package main

import (
	"context"
	"gRPC/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{Result: in.A + in.B}, nil
}

func (s server) Divide(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	return &pb.CalculationResponse{Result: in.A / in.B}, nil
}

func (s server) Sum(ctx context.Context, in *pb.NumbersRequest) (*pb.CalculationResponse, error) {
	var sum int64
	for _, num := range in.Numbers {
		sum += num
	}
	return &pb.CalculationResponse{Result: sum}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}
	s := grpc.NewServer()
	// this allows gRpcUi to get the proto types.
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed", err)
	}
}
