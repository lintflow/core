package main

import (
	"flag"
	"net"

	"github.com/lintflow/core/lookupd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/lintflow/core/proto"
)

var (
	addr = flag.String(`addr`, `localhost:4567`, `address for listen service`)
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterLookupdServiceServer(grpcServer, lookupd.New())
	grpcServer.Serve(lis)
}
