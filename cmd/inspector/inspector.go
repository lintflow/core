package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"fmt"
	"github.com/lintflow/core/inspector"
	pb "github.com/lintflow/core/proto"
	"net/http"
	_ "net/http/pprof"
)

var (
	addr    = flag.String(`addr`, `localhost:4568`, `address for listen service`)
	lookupd = flag.String(`lookupd`, `localhost:4567`, `address for listen lookupd`)
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Set up a connection to the lookupd services
	conn, err := grpc.Dial(*lookupd, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to listen lookupd: %v", err)
	}

	defer conn.Close()

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	go http.ListenAndServe(fmt.Sprintf(":%d", 36663), nil)
	pb.RegisterInspectorServiceServer(grpcServer, inspector.New(pb.NewLookupdServiceClient(conn)))
	grpcServer.Serve(lis)
}
