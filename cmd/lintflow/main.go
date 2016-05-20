package main

import (
	"flag"

	pb "github.com/lintflow/core/proto"

	"code.google.com/p/go.net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var DefaultInspector string = `localhost:4568`

func main() {
	flag.Parse()
	command := flag.Arg(0)

	// Set up a connection to the lookupd services
	conn, err := grpc.Dial(DefaultInspector, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to listen inspector services: %v", err)
	}
	defer conn.Close()

	inspector := pb.NewInspectorServiceClient(conn)

	switch command {
	case `services`:
		resp, err := inspector.Services(context.Background(), new(pb.ListRequest))
		if err != nil {
			grpclog.Fatalf("failed get services: %v", err)
		} else {
			grpclog.Println(`services:`)
			for _, service := range resp.GetServices() {
				grpclog.Println("\t" + service.String())
			}
		}
	}
}
