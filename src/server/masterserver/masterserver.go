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
		statusCode.Value = "-1"
		statusCode.Exception = "ERROR: file doesn't exist in " + *filePath
	} else {
		statusCode.Value = "0"
		statusCode.Exception = "SUCCESS: file exists in " + *filePath
	}
}

//*********************************** 业务函数 ************************************

// CreateFile 创建文件
func (ms *MasterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command CreateFile " + filePath)
	// 定义变量，传进去
	var chunkHandle string
	var locations []string
	var statusCode cm.StatusCode
	// 核心逻辑
	ms.createFile(&filePath, &chunkHandle, &locations, &statusCode)
	// 打印状态
	switch statusCode.Value {
	case "0":
		logger.Message(statusCode.Exception)
	default:
		logger.Warn(statusCode.Exception)
	}
	// 返回信息给客户端
	if statusCode.Value != "0" {
		return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: statusCode.Value}, nil
	}
	replyMessage := ""
	for i := 0; i < len(locations); i++ {
		replyMessage = replyMessage + "|" + locations[i]
	}
	return &pb.Reply{ReplyMessage: replyMessage, StatusCode: statusCode.Value}, nil
}

// createFile 核心逻辑
func (ms *MasterServer) createFile(filePath *string, chunkHandle *string, locations *[]string, statusCode *cm.StatusCode) {
	// 元数据信息需要修改
	ms.metadata.CreateNewFile(filePath, chunkHandle, statusCode)
	// 返回码错误
	if statusCode.Value != "0" {
		return
	}
	// 返回码正确
	files := ms.metadata.GetFiles()
	if file, ok := (*files)[*filePath]; ok {
		chunks := file.GetChunks()
		chunk := (*chunks)[*chunkHandle]
		*locations = chunk.locations
	}
}

// ListFiles 展示文件列表
func (ms *MasterServer) ListFiles(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	filePath := req.SendMessage
	logger.Message("Command ListFiles " + filePath)
	// 定义变量，传进去
	var filePaths []string
	// 核心逻辑
	ms.listFiles(&filePath, &filePaths)
	// 返回信息给客户端
	var statusCode cm.StatusCode
	replyMessage := ""
	if len(filePaths) == 0 {
		statusCode.Value = "-1"
		statusCode.Exception = filePath + ":" + "is not exist"
		return &pb.Reply{ReplyMessage: statusCode.Exception, StatusCode: statusCode.Value}, nil
	}
	for _, file := range filePaths {
		replyMessage = replyMessage + "|" + file
	}
	return &pb.Reply{ReplyMessage: replyMessage, StatusCode: "0"}, nil
}

// listFiles 核心逻辑
func (ms *MasterServer) listFiles(filePath *string, filePaths *[]string) {
	for k, _ := range ms.metadata.files {
		// 长度太小，直接下次循环
		if len(k) < len(*filePath) {
			continue
		}
		// 截取子串,符合条件的加入切片中,切片是头闭尾开
		subK := k[0:len(*filePath)]
		if strings.Compare(subK, *filePath) == 0 {
			*filePaths = append(*filePaths, k)
		}
	}
}

// WriteFile 写文件
func (ms *MasterServer) WriteFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// 分割串
	slice := strings.Split(req.SendMessage, "|")
	filePath := slice[0]
	offset := slice[1]
	data := ""
	for i := 2; i < len(slice); i++ {
		data = data + slice[i]
	}
	logger.Message("Command WriteFile " + data + " to " + filePath + "offset" + offset)
	// 定义变量，传进去
	var statusCode cm.StatusCode
	var chunksLocations string
	// 核心逻辑
	ms.writeFile(&filePath, &offset, &data, &chunksLocations, &statusCode)

	// todo 处理返回逻辑
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

// writeFile 核心逻辑
func (ms *MasterServer) writeFile(filePath *string, offset *string, data *string, chunksLocations *string, statusCode *cm.StatusCode) {
	ms.checkValidFile(filePath, statusCode)
	if statusCode.Value != "0" {
		return
	}
	// 先拿到文件
	files := ms.metadata.GetFiles()
	file := (*files)[*filePath]
	// 获得原始的 chunk 数量
	chunkOriginNum := len(file.chunkHandleSet)
	// 转成 int
	offsetInt, _ := strconv.Atoi(*offset)
	dataInt, _ := strconv.Atoi(*data)
	// 根据偏移量获得在第几个chunk中，chunk内的偏移量是多少
	chunkId := offsetInt / cm.GFSChunkSize
	chunkOffset := offsetInt % cm.GFSChunkSize
	// 偏移量开始的位置在已有的中间
	if chunkId >= 0 && chunkId <= chunkOriginNum {
		// 加上偏移量算一次
		offsetData := offsetInt + dataInt
		newChunkId := offsetData / cm.GFSChunkSize
		newChunkOffset := offsetData % cm.GFSChunkSize
		// 加上偏移量也没有超出

	}

	var chunkNum int
	if len(*data)%cm.GFSChunkSize != 0 {
		chunkNum = len(*data)/cm.GFSChunkSize + 1
	} else {
		chunkNum = len(*data) / cm.GFSChunkSize
	}
	//chunks := file.GetChunks()
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
