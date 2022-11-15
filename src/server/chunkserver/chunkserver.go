package chunkserver

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"os"
)

type ChunkServer struct {
	pb.UnimplementedChunkServerToClientServer
	port string
	root string
}

// ************************************辅助函数************************************'

func NewChunkServer(port *string, root string) *ChunkServer {
	var cs ChunkServer
	cs.port = *port
	cs.root = "./" + root + "/" + *port
	err := os.MkdirAll(cs.root, os.ModePerm)
	if err != nil {
		return nil
	}
	return &cs
}

// ************************************业务函数************************************

// Create 创建
func (cs *ChunkServer) Create(ctx context.Context, request *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	chunkHandle := request.SendMessage
	logger.Message(cs.port + "Create Chunk" + chunkHandle)
	// 定义变量，传进去
	var statusCode cm.StatusCode
	// 核心逻辑
	cs.create(&chunkHandle, &statusCode)
	// 返回信息给客户端
	return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: statusCode.Value}, nil
}

// create 核心逻辑
func (cs *ChunkServer) create(chunkHandle *string, statusCode *cm.StatusCode) {
	filePath := cs.root + "/" + *chunkHandle
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: " + err.Error()
	}
	// 关闭文件
	err = file.Close()
	if err != nil {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: " + err.Error()
	}
	statusCode.Value = "0"
	statusCode.Exception = "SUCCESS: Chunk Created "
}
