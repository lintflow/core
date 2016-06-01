package main

import (
	"flag"

	pb "github.com/lintflow/core/proto"

	"golang.org/x/net/context"

	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"io"
	"strconv"
)

var DefaultInspector string = `localhost:4568`

var (
	uri = flag.String(`repo`, `https://github.com/lintflow/golang-test-project.git`, `your repo with go project for validate here`)
)

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
			fmt.Println(`services:`)
			for _, service := range resp.GetServices() {
				fmt.Println("\t" + service.String())
			}
		}
	case `inspect`:
		resp, err := inspector.Services(context.Background(), new(pb.ListRequest))
		task := &pb.Task{}
		if err != nil {
			grpclog.Fatalf("failed get services: %v", err)
		} else {
			for _, service := range resp.GetServices() {
				switch service.Type {
				case pb.Service_LINTER:
					task.Validators = &pb.Task_Args{Service: service}
				case pb.Service_REPORTER:
					task.Reporters = &pb.Task_Args{Service: service}
				case pb.Service_RESOURCER:
					task.Resourcers = &pb.Task_Args{
						Service: service,
						Config:  []byte(`{"url":"` + *uri + `"}`),
					}
				}
			}
		}

		progress, err := inspector.Inspect(context.Background(), task)
		if err != nil {
			grpclog.Fatalf("failed start incpect task %s: %v", task, err)
		}

		for {
			info, err := progress.Recv()
			if err == io.EOF || info == nil {
				progress.CloseSend()
				println(`Finish!`)
				return
			}
			if err != nil {
				grpclog.Fatalf("failed recive data: %v", err)
			}
			println("\t", info.Current, `/`, info.Total)
			if info.Link != "" {
				println(`see your report here - ` + info.Link)
				println(`was finded ` + strconv.Itoa(int(info.Problems)) + ` problems in ` + strconv.Itoa(int(info.Total)) + ` files.`)
			}
		}
	}

}
