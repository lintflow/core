package validator

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	//"errors"
	//"fmt"
	"github.com/golang/lint"
	pb "github.com/lintflow/core/proto"
	"github.com/pborman/uuid"
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
	rp, err := v.Reporter(t.GetReporter().GetService().Address)
	if err != nil {
		return err
	}

	// открывает стрим на получение данных
	streamOfResourse, err := rs.Get(context.Background(), &pb.ConfigRequest{Config: resourser.Config})
	if err != nil {
		return err
	}

	writer, err := rp.Record(context.Background())
	if err != nil {
		return err
	}

	linter := new(lint.Linter)

	current := 0
	countProblems := 0
	for {
		file, err := streamOfResourse.Recv()
		current++
		if err == io.EOF {
			// закрываем ресурсер
			streamOfResourse.CloseSend()
			//if err != nil {
			//	println(`streamOfResourse.CloseSend:`, err.Error())
			//	return err
			//}

			// закрываем репортер и получаем ссылку на отчет
			summary, err := writer.CloseAndRecv()
			if err != nil {
				return err
			}
			var total int64 = 123
			link := `some link`
			if summary != nil {
				total = summary.Total
				link = summary.Link
			}
			// пишем последний прогресс
			err = s.Send(&pb.ValidateProgress{
				Reporter: &pb.ValidateProgress_Progress{
					Id:      uuid.New(),
					Total:   int64(current),
					Current: total,
				},
				Resourser: &pb.ValidateProgress_Progress{
					Id:      uuid.New(),
					Total:   int64(current),
					Current: int64(current),
				},
				LinkToReport: link,
			})
			if err != nil {
				println(`error Send`, err.Error())
			}
			return nil
		}
		if err != nil {
			return err
		}

		problems, err := linter.Lint(file.Header, file.Body)
		if err != nil {
			return err
		}

		report := &pb.Problem{
			Id:       uuid.New() + `/` + file.Header,
			Original: file.Body,
		}

		for _, problem := range problems {
			report.Details = append(report.Details, &pb.Problem_Detail{
				Id:          uuid.New(),
				Fragment:    []byte(problem.Position.String()),
				Description: problem.Category + `:` + problem.String(),
			})
			countProblems++
		}

		err = writer.Send(report)
		if err != nil {
			return err
		}

		err = s.Send(&pb.ValidateProgress{
			Reporter: &pb.ValidateProgress_Progress{
				Id:      uuid.New(),
				Total:   file.Total,
				Current: int64(countProblems),
			},
			Resourser: &pb.ValidateProgress_Progress{
				Id:      uuid.New(),
				Total:   file.Total,
				Current: int64(current),
			},
			LinkToReport: "",
		})

		if err != nil {
			return err
		}
	}
	return nil
}
