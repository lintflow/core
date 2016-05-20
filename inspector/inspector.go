package inspector

import (
	"golang.org/x/net/context"

	pb "github.com/lintflow/core/proto"
)

func New(lookuper pb.LookupdServiceClient) pb.InspectorServiceServer {
	return &inspector{
		lookup: lookuper,
	}
}

type inspector struct {
	lookup pb.LookupdServiceClient
}

func (i *inspector) Inspect(task *pb.Task, stream pb.InspectorService_InspectServer) error {
	return nil
}

func (i *inspector) Services(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return i.lookup.List(ctx, req)
}
