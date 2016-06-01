package reporter

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"bufio"
	pb "github.com/lintflow/core/proto"
	"github.com/pborman/uuid"
	"io"
	"os"
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
	link, writer, closer, err := NewFileReport()
	if err != nil {
		return err
	}
	defer closer()

	for {
		problem, err := s.Recv()
		if err == io.EOF {
			err = writer.Flush()
			if err != nil {
				return err
			}
			return s.SendAndClose(&pb.ReportSummary{
				Link:  link,
				Total: int64(total),
			})
		}
		if err != nil {
			return err
		}

		if problem == nil {
			continue
		}

		total++

		for _, detail := range problem.Details {
			_, err = writer.WriteString(string(detail.Fragment) + ` - ` + detail.Description + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewFileReport() (string, *bufio.Writer, func() error, error) {
	filename := "/tmp/report-" + uuid.New() + `.txt`
	f, err := os.Create(filename)
	return filename, bufio.NewWriter(f), f.Close, err
}
