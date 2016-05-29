package reporter

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	pb "github.com/lintflow/core/proto"
	"io"
)

type reporter struct {
	s      *pb.Service
	lookup pb.LookupdServiceClient
}

func New(s *pb.Service, lookuper pb.LookupdServiceClient) pb.ReporterServiceServer {
	v := &reporter{
		s:      s,
		lookup: lookuper,
	}

	_, err := v.lookup.Register(context.Background(), &pb.RegisterRequest{Service: s})
	if err != nil {
		grpclog.Fatalf(`can't register service: %s`, err.Error())
	}

	return v
}

func (r *reporter) Record(s pb.ReporterService_RecordServer) error {
	total := 0
	link, writer, err := NewFileReport()
	if err != nil {
		return err
	}
	for {
		problem, err := s.Recv()
		if err == io.EOF {
			s.SendAndClose(&pb.ReportSummary{
				Link:  link,
				Total: total,
			})
		}
		total++
		_, err = writer.Write(problemToBytes(problem))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewFileReport() (string, io.Writer, error) {

}

func problemToBytes(p *pb.Problem) []byte {
	return []byte(p.String())
}
