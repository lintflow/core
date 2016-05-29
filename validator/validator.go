package validator

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"errors"
	"fmt"
	"github.com/golang/lint"
	pb "github.com/lintflow/core/proto"
	"google.golang.org/grpc"
	"io"
	"sync"
)

func New(s *pb.Service, lookuper pb.LookupdServiceClient) pb.ValidatorServiceServer {
	v := &validator{
		s:      s,
		lookup: lookuper,
		c:      make(map[string]*grpc.ClientConn, 0),
	}

	_, err := v.lookup.Register(context.Background(), &pb.RegisterRequest{Service: s})
	if err != nil {
		grpclog.Fatalf(`can't register service: %s`, err.Error())
	}
	return v
}

type validator struct {
	s      *pb.Service
	lookup pb.LookupdServiceClient
	//addrs
	sync.Mutex
	//clients connections
	c map[string]*grpc.ClientConn
}

func (v *validator) GetConnection(addr string) (conn *grpc.ClientConn, err error) {

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

func (v *validator) Resourcer(addr string) (pb.ResourcerServiceClient, error) {
	conn, err := v.GetConnection(addr)
	if err != nil {
		return nil, err
	}
	return pb.NewResourcerServiceClient(conn), nil
}

func (v *validator) Reporter(addr string) (pb.ReporterServiceClient, error) {
	conn, err := v.GetConnection(addr)
	if err != nil {
		return nil, err
	}
	return pb.NewReporterServiceClient(conn), nil
}

func (v *validator) Validate(t *pb.ValidationTask, s pb.ValidatorService_ValidateServer) error {

	resourser := t.GetResourcer()
	rs, err := v.Resourcer(resourser.GetService().Address)
	if err != nil {
		return err
	}
	rp, err := v.Reporter(t.Reporter().GetService().Address)
	if err != nil {
		return err
	}

	// открывает стрим на получение данных
	streamOfResourse, err := rs.Get(s.Context(), resourser.Config)
	if err != nil {
		return err
	}

	linter := new(lint.Linter)

	for {
		file, err := streamOfResourse.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return err
		}

		problems, err := linter.Lint(file.Header, file.Body)
		if err != nil {
			return err
		}

		rp
	}
	return nil
}
