package masterserver

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sadlil/gologger"
)

//type chunkHandle string

type MasterServer struct {
	pb.UnimplementedMasterServerToClientServer
	port     string
	fileList []string
	metadata MetaData
}

//************************************辅助函数************************************

// StartMasterService 启动 Master 服务
func (ms *MasterServer) StartMasterService(port string, locations []string) {
	ms.port = port
	ms.metadata.SetLocations(locations)
}

// GetChunkHandle 得到 chunkHandle，形式：5b912ae9-71c1-464d-8e32-712b4b506430
func (ms *MasterServer) GetChunkHandle(chunkHandle *string) {
	uid, _ := uuid.NewV4()
	fmt.Println(uid)
	*chunkHandle = uid.String()
}

//*********************************** 业务函数************************************

// CreateFile 创建文件
func (ms *MasterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command CreateFile " + filePath)
	//创建文件
	var chunkHandle string
	var locations []string
	var s cm.StatusCode
	ms.createFile(&filePath, &chunkHandle, locations, &s)

	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

func (ms *MasterServer) createFile(filePath *string, chunkHandle *string, files []string, locations *cm.StatusCode) {
	ms.GetChunkHandle(chunkHandle)
}

// ListFiles 展示文件列表
func (ms *MasterServer) ListFiles(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command ListFiles " + filePath)

	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

func (ms *MasterServer) listFiles(filePath *string, files []string) {

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
