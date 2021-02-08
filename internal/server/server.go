package server

import (
	"context"

	pb "github.com/sqljudge/executor/internal/executor"
)

type ExecutorServer struct {
	pb.UnimplementedExecutorServer
}

func (s *ExecutorServer) Execute(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
