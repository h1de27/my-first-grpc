package main

import (
	"context"

	pb "github.com/h1de27/my-first-grpc/echo/proto"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context,
	req *pb.EchoRequest) (*pb.EchoResponse, error) {
		return &pb.EchoResponse{
			Message: req.GetMessage()}, nil
}
