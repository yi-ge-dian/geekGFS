package masterserver

import "GeekGFS/src/pb"
import "context"

type MasterServer struct {
	pb.UnimplementedMasterServerToClientServer
}

func (s *MasterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}
