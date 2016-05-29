package lookupd

import (
	"errors"
	"sync"

	"golang.org/x/net/context"

	pb "github.com/lintflow/core/proto"
)

func New() pb.LookupdServiceServer {
	return &lookupd{}
}

type lookupd struct{}

type kv struct {
	lock  sync.RWMutex
	store map[string]*pb.Service
}

func (k *kv) Get(key string) *pb.Service {
	var service *pb.Service
	k.lock.Lock()
	service = k.store[key]
	k.lock.Unlock()
	return service
}

func (k *kv) Set(key string, service *pb.Service) {
	if service != nil {
		k.lock.Lock()
		k.store[key] = service
		k.lock.Unlock()
	}
}

func (k *kv) List(filter *pb.ListRequest) []*pb.Service {
	services := []*pb.Service{}
	for _, service := range k.store {
		if filter.Type == pb.ListRequest_ANY {
			services = append(services, service)
		} else if service.Type.String() == filter.Type.String() {
			services = append(services, service)
		}
	}
	return services
}

var store = &kv{store: make(map[string]*pb.Service, 0)}

func (l *lookupd) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	service := req.GetService()
	if service == nil {
		return nil, errors.New(`bad value`)
	}

	store.Set(service.Id, service)
	return &pb.RegisterResponse{Ok: true}, nil
}

func (l *lookupd) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	services := store.List(req)
	return &pb.ListResponse{Services: services}, nil
}
