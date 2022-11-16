package chunkserver

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"io"
	"os"
	"strings"
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

func CheckFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// ************************************业务函数************************************

// Create 创建
func (cs *ChunkServer) Create(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	chunkHandle := req.SendMessage
	logger.Message(cs.port + "Create Chunk " + chunkHandle)
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
		return
	}
	// 关闭文件
	err = file.Close()
	if err != nil {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: " + err.Error()
		return
	}
	statusCode.Value = "0"
	statusCode.Exception = "SUCCESS: Chunk Created"
}

// Write 写
func (cs *ChunkServer) Write(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	slice := strings.Split(req.SendMessage, "|")
	chunkHandle := slice[0]
	data := slice[1]
	logger.Message(cs.port + "Write Chunk" + data + "to " + chunkHandle)
	// 定义变量，传进去
	var statusCode cm.StatusCode
	// 核心逻辑
	cs.write(&chunkHandle, &data, &statusCode)
	// 返回信息给客户端
	return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: statusCode.Value}, nil
}

// write 核心逻辑
func (cs *ChunkServer) write(chunkHandle *string, data *string, statusCode *cm.StatusCode) {
	filePath := cs.root + "/" + *chunkHandle
	if !CheckFileExist(filePath) {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: File not exists, please create one"
		return
	}
	err := os.WriteFile(filePath, []byte(*data), 0777)
	if err != nil {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: " + err.Error()
		return
	}
	statusCode.Value = "0"
	statusCode.Exception = "SUCCESS: " + cs.port + " write data into " + *chunkHandle
}

// Read 读
func (cs *ChunkServer) Read(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	chunkHandle := req.SendMessage
	// 定义变量传进去
	data := ""
	var statusCode cm.StatusCode
	// 核心逻辑
	cs.read(&chunkHandle, &data, &statusCode)
	// 返回逻辑
	if statusCode.Value != "0" {
		return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: statusCode.Value}, nil
	}
	return &pb.Reply{ReplyMessage: data, StatusCode: statusCode.Value}, nil
}

// read 核心逻辑
func (cs *ChunkServer) read(chunkHandle *string, data *string, statusCode *cm.StatusCode) {
	filePath := cs.root + "/" + *chunkHandle
	if !CheckFileExist(filePath) {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: File not exists, please create one"
		return
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: " + err.Error()
		return
	}
	defer func(file *os.File) {
		err_ := file.Close()
		if err_ != nil {
			statusCode.Value = "-1"
			statusCode.Exception = "ERROR: " + err.Error()
			return
		}
	}(file)
	// 按字节读取文件
	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err_ := file.Read(buf)
		if err_ != nil && err_ != io.EOF {
			panic(err_)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	*data = string(chunks)
	statusCode.Value = "0"
	statusCode.Exception = "SUCCESS: " + cs.port + " read data from " + *chunkHandle
}
