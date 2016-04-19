package inspector

import (
	pb "github.com/lintflow/core/proto"
)

type inspector struct{}

func (i *inspector) Inspect(task *pb.Task, stream pb.InspectorService_InspectServer) error {
	return nil
}
