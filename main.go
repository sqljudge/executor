package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/sqljudge/executor/internal/executor"
	"github.com/sqljudge/executor/internal/server"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 5555, "gRPC server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExecutorServer(grpcServer, &server.ExecutorServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
