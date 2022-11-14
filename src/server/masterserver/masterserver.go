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

// checkValidFile 检查文件是否有效
func (ms *MasterServer) checkValidFile(filePath *string, statusCode *cm.StatusCode) {
	files := ms.metadata.GetFiles()
	// 如果该文件不存在
	if _, ok := (*files)[*filePath]; !ok {
		statusCode.Value = cm.FileNotExistsValue
		statusCode.Exception = cm.FileNotExistsException + "in " + *filePath
	} else {
		statusCode.Value = 0
		statusCode.Exception = "SUCCESS: file exists in " + *filePath
	}
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

// WriteFile 写文件
func (ms *MasterServer) WriteFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// 分割串
	slice := strings.Split(req.SendMessage, "|")
	filePath := slice[0]
	var data string
	for i := 1; i < len(slice); i++ {
		data = data + slice[i]
	}
	logger.Message("Command WriteFile " + data + " to " + filePath)
	var statusCode cm.StatusCode
	var chunksLocations string
	ms.writeFile(&filePath, &data, &chunksLocations, &statusCode)

	// todo 处理返回逻辑
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

func (ms *MasterServer) writeFile(filePath *string, data *string, chunksLocations *string, statusCode *cm.StatusCode) {
	ms.checkValidFile(filePath, statusCode)
	if statusCode.Value != 0 {
		return
	}
	var chunkNum int
	if len(*data)%cm.GFSChunkSize != 0 {
		chunkNum = len(*data)/cm.GFSChunkSize + 1
	} else {
		chunkNum = len(*data) / cm.GFSChunkSize
	}
	// 从 master 这里只能获取到 chunkServer 的位置，写数据去chunkServer
	files := ms.metadata.GetFiles()
	file := (*files)[*filePath]
	chunks := file.GetChunks()
	// 如果想要写的 chunk 数目比现在的小，直接给你我以前的数据，可能会出现跨页的情况，这种情况去 chunkServer处理，我反正只给你位置信息
	if chunkNum <= len(file.chunkHandleSet) {

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

// DeleteFile todo
func (ms *MasterServer) DeleteFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}
