package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/lintflow/core/reporter"
	_ "net/http/pprof"

	"fmt"
	pb "github.com/lintflow/core/proto"
	"net/http"
)

var (
	addr        = flag.String(`addr`, `localhost:4525`, `address for listen service`)
	lookupd     = flag.String(`lookupd`, `localhost:4567`, `address for listen lookupd`)
	id          = flag.String(`id`, `reporter-1`, `id of your reporter`)
	name        = flag.String(`name`, `reporter`, `name of your reporter`)
	description = flag.String(`description`, `core test reporter`, `description of your reporter`)
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
	srv := &pb.Service{
		Id:          *id,
		Name:        *name,
		Description: *description,
		Address:     *addr,
		Type:        pb.Service_REPORTER,
		Tags:        []string{`test`, `start`},
		TaskConfig:  []byte(`{"config":"json"}`),
	}
	grpclog.Println(srv.String())
	go http.ListenAndServe(fmt.Sprintf(":%d", 36668), nil)
	pb.RegisterReporterServiceServer(grpcServer, reporter.New(srv, pb.NewLookupdServiceClient(conn)))
	grpcServer.Serve(lis)
}
