package inspector

import (
	"golang.org/x/net/context"

	pb "github.com/lintflow/core/proto"
	"google.golang.org/grpc"
	"io"
	"sync"
)

func New(lookuper pb.LookupdServiceClient) pb.InspectorServiceServer {
	return &inspector{
		lookup: lookuper,
		c:      make(map[string]*grpc.ClientConn, 0),
	}
}

func (v *inspector) GetConnection(addr string) (conn *grpc.ClientConn, err error) {

	var ok bool
	v.Lock()
	conn, ok = v.c[addr]
	v.Unlock()
	if ok {
		return conn, nil
	}

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	v.Lock()
	v.c[addr] = conn
	v.Unlock()

	return conn, nil
}

func (v *inspector) Validator(addr string) (pb.ValidatorServiceClient, error) {
	conn, err := v.GetConnection(addr)
	if err != nil {
		return nil, err
	}
	return pb.NewValidatorServiceClient(conn), nil
}

type inspector struct {
	lookup pb.LookupdServiceClient

	sync.Mutex
	c map[string]*grpc.ClientConn
}

func (i *inspector) Inspect(task *pb.Task, stream pb.InspectorService_InspectServer) error {
	lintOptions := task.GetValidators()

	validator, err := i.Validator(lintOptions.GetService().Address)
	if err != nil {
		return err
	}
	progressor, err := validator.Validate(stream.Context(), &pb.ValidationTask{
		Config:    lintOptions.Config,
		Reporter:  task.GetReporters(),
		Resourcer: task.GetResourcers(),
	})
	if err != nil {
		return err
	}
	for {
		progress, err := progressor.Recv()
		if err == io.EOF {
			progressor.CloseSend()
		}
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Progress{
			Id:       progress.GetResourser().Id,
			Total:    progress.GetResourser().Total,
			Current:  progress.GetResourser().Current,
			Link:     progress.LinkToReport,
			Problems: progress.GetReporter().Current,
		})
		if err != nil {
			return err
		}
	}
	return nil

}

func (i *inspector) Services(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return i.lookup.List(ctx, req)
}
