package resourser

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"encoding/json"
	"errors"
	pb "github.com/lintflow/core/proto"

	"github.com/pborman/uuid"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	if len(req.Config) == 0 {
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
			Total:  int64(total),
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
	Init() (int, error)
	Err() error
}

func (r *resourser) newFsIter(uri string) FsIter {
	return &fsiter{
		id:  uuid.New(),
		uri: uri,
	}
}

type fsiter struct {
	id      string
	uri     string
	err     error
	files   []string
	current int
}

func (f *fsiter) Next() bool {
	length := len(f.files)
	if length == 0 {
		f.err = errors.New(`dir not have golang project or file *.go patterns`)
		return false
	}
	f.current++
	// end of array
	if f.current-1 == length {
		return false
	}
	return true
}
func (f *fsiter) File() (string, []byte) {
	filename := f.files[f.current-1]
	blob, _ := ioutil.ReadFile(filename)
	return filename, blob
}

func (f *fsiter) Err() error {
	return f.err
}

func (f *fsiter) walk(path string, info os.FileInfo, err error) error {
	if info == nil {
		return nil
	}
	if info.IsDir() && info.Name() == `.git` {
		return filepath.SkipDir
	}
	if info.IsDir() {
		return nil
	}

	if !strings.Contains(path, `.go`) {
		return nil
	}

	f.files = append(f.files, path)
	return nil
}

func (f *fsiter) Init() (int, error) {
	err := exec.Command(`git`, `clone`, f.uri, `/tmp/`+f.id).Run()
	if err != nil && err.Error() != `exit status 128` {
		return 0, err
	}

	err = filepath.Walk(`/tmp/`+f.id, f.walk)
	if err != nil {
		return 0, err
	}
	return len(f.files), nil
}
