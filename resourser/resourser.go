package resourser

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"encoding/json"
	"errors"
	pb "github.com/lintflow/core/proto"

	"github.com/pborman/uuid"
	"os/exec"
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

func (r *resourser) Get(req *pb.ConfigRequest, stream pb.ResourcerService_GetServer) error {
	if len(req.Config) {
		return errors.New(`bad config`)
	}
	git := &struct {
		URL string `json:"url"`
	}{}
	err := json.Unmarshal(req.Config, git)
	if err != nil {
		return err
	}

	iter := r.newFsIter(git.URL)
	total, err := iter.Init()
	if err != nil {
		return err
	}

	for iter.Next() {
		filename, blob := iter.File()
		err = stream.Send(&pb.Resource{
			Header: filename,
			Body:   blob,
			Total:  total,
		})
		if err != nil {
			return err
		}
	}
	return iter.Err()
}

type FsIter interface {
	Next() bool
	File() (string, []byte)
	Init() (int64, error)
	Err() error
}

func (r *resourser) newFsIter(uri string) FsIter {
	return &fsiter{uuid.New(), uri, nil}
}

type fsiter struct {
	id  string
	uri string
	err error
}

func (f *fsiter) Next() bool {
	return false
}
func (f *fsiter) File() (string, []byte) {
	return "", []byte(``)
}

func (f *fsiter) Err() error {
	return f.err
}

func (f *fsiter) Init() (int64, error) {
	err := exec.Command(`git`, `clone`, f.uri, `/tmp/`+f.id).Run()
	if err != nil {
		return 0, err
	}

}
