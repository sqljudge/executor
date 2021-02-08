package server

import (
	"context"

	pb "github.com/sqljudge/executor/internal/executor"
)

// ExecutorServer implements gRPC server for the executor service
type ExecutorServer struct {
	pb.UnimplementedExecutorServer
}

// Execute executes given SQL query on the corresponding database
func (s *ExecutorServer) Execute(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
