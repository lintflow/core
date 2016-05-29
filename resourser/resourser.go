package resourser

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	pb "github.com/lintflow/core/proto"
)

type resourser struct {
	s      *pb.Service
	lookup pb.LookupdServiceClient
}

func New(s *pb.Service, lookuper pb.LookupdServiceClient) pb.ResourcerServiceServer {
	v := &resourser{
		s:      s,
		lookup: lookuper,
	}

	_, err := v.lookup.Register(context.Background(), &pb.RegisterRequest{Service: s})
	if err != nil {
		grpclog.Fatalf(`can't register service: %s`, err.Error())
	}

	return v
}
