package masterserver

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"strconv"
	"strings"
)

type MasterServer struct {
	pb.UnimplementedMasterServerToClientServer
	port     string
	fileList []string
	metadata MetaData
}

//************************************辅助函数************************************

// NewMasterServer 构造函数
func NewMasterServer(port *string, locations []string) *MasterServer {
	return &MasterServer{port: *port, metadata: *NewMetaData(locations)}
}

//*********************************** 业务函数 ************************************

// CreateFile 创建文件
func (ms *MasterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command CreateFile " + filePath)
	// 创建文件
	var chunkHandle string
	var locations []string
	var statusCode cm.StatusCode
	ms.createFile(&filePath, &chunkHandle, &locations, &statusCode)
	// 打印状态
	switch statusCode.Value {
	case 0:
		logger.Message(statusCode.Exception)
	default:
		logger.Warn(statusCode.Exception)
	}
	// 返回信息给客户端
	if statusCode.Value != 0 {
		return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: strconv.Itoa(statusCode.Value)}, nil
	}
	replyMessage := ""
	for i := 0; i < len(locations); i++ {
		replyMessage = replyMessage + "|" + locations[i]
	}
	return &pb.Reply{ReplyMessage: replyMessage, StatusCode: strconv.Itoa(statusCode.Value)}, nil
}

func (ms *MasterServer) createFile(filePath *string, chunkHandle *string, locations *[]string, statusCode *cm.StatusCode) {
	//logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	*chunkHandle = cm.GenerateChunkHandle()
	ms.metadata.CreateNewFile(filePath, chunkHandle, statusCode)
	if statusCode.Value != 0 {
		return
	}
	files := ms.metadata.GetFiles()
	file, ok := (*files)[*filePath]
	if ok {
		chunk := (*file.GetChunks())[*chunkHandle]
		*locations = chunk.locations
	}
}

// ListFiles 展示文件列表
func (ms *MasterServer) ListFiles(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command ListFiles " + filePath)
	// 存放文件
	var files []string
	ms.listFiles(&filePath, &files)
	// 返回信息
	var statusCode cm.StatusCode
	replyMessage := ""
	if len(files) == 0 {
		statusCode.Value = cm.FilePathNotExistsValue
		replyMessage = filePath + ":" + cm.FilePathNotExistsException
		statusCode.Exception = replyMessage
		return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: strconv.Itoa(statusCode.Value)}, nil
	}
	for _, file := range files {
		replyMessage = replyMessage + "|" + file
	}
	return &pb.Reply{ReplyMessage: replyMessage, StatusCode: "0"}, nil
}

func (ms *MasterServer) listFiles(filePath *string, files *[]string) {
	for k, _ := range ms.metadata.files {
		// 长度太小，直接下次循环
		if len(k) < len(*filePath) {
			continue
		}
		// 截取子串,符合条件的加入切片中,切片是头闭尾开
		subK := k[0:len(*filePath)]
		if strings.Compare(subK, *filePath) == 0 {
			*files = append(*files, k)
		}
	}
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
