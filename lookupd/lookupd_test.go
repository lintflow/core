package lookupd

import (
	"testing"

	"golang.org/x/net/context"

	pb "github.com/lintflow/core/proto"
)

func TestLookupd_Register(t *testing.T) {
	srv := &pb.Service{
		Id:          `some-1`,
		Name:        `name`,
		Description: `description`,
		Address:     `localhost:1234`,
		Type:        pb.Service_LINTER,
		Tags:        []string{`test`, `start`},
		TaskConfig:  []byte(`{"config":"json"}`),
	}
	service := New()
	resp, err := service.Register(context.TODO(), &pb.RegisterRequest{srv})
	if err != nil {
		t.Errorf(`expected err == nil, but - %s`, err)
	}
	if !resp.Ok {
		t.Errorf(`expected resp OK = true`)
	}
	if len(store.store) != 1 {
		t.Errorf(`expected len store after register == 1, but - %d`, len(store.store))
	}

	// we can get service data
	list, err := service.List(context.TODO(), new(pb.ListRequest))
	if err != nil {
		t.Errorf(`expected err == nil, but - %s`, err)
	}
	if len(list.Services) != 1 {
		t.Errorf(`expected len list services after register == 1, but - %d`, len(list.Services))
	}

}
