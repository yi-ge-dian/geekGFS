package masterserver

import (
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
)

type MasterServer struct {
	pb.UnimplementedMasterServerToClientServer
	port     string
	fileList []string
	metadata MetaData
}

func (ms *MasterServer) MasterService(port string, locations []string) {
	ms.port = port
	ms.metadata.SetLocations(locations)
}

// ListFiles 展示文件列表
func (ms *MasterServer) ListFiles(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command ListFiles" + filePath)

	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// CreateFile todo
func (ms *MasterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// AppendFile todo
func (ms *MasterServer) AppendFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// CreateChunk todo
func (ms *MasterServer) CreateChunk(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// ReadFile todo
func (ms *MasterServer) ReadFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// WriteFile todo
func (ms *MasterServer) WriteFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// DeleteFile todo
func (ms *MasterServer) DeleteFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}
