package client

import (
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
)

//************************************辅助函数************************************

// SwitchChunkServer client 与 chunkServer 建立连接,并且执行逻辑
func SwitchChunkServer(chunkServerSocket *string, command *string, chunkHandle *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)

	// 1. 建立连接，端口是服务端开放的端口 没有证书会报错
	conn, err := grpc.Dial(*chunkServerSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Client connected chunkServer at " + *chunkServerSocket)
	// 退出时关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close the connection" + err.Error())
		}
	}(conn)

	// 2. 调用 Product.pb.go 中的 NewProdServiceClient 方法
	clientForCS := pb.NewChunkServerToClientClient(conn)
	clientForCSCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 3. 调用方法
	switch *command {
	case "create":
		chunkServerReply, _ := clientForCS.Create(clientForCSCtx, &pb.Request{SendMessage: *chunkHandle})
		// 根据 chunkServer 的返回码来输出信息
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
	case "list":

	}

}

//************************************业务函数************************************

// CreateFile 客户端创建文件
func CreateFile(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).CreateFile(*clientForMSCtx, &pb.Request{SendMessage: *filePath})
	// 根据 masterServer 的返回码来输出信息
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message("Response from masterServer:" + masterServerReply.ReplyMessage)
		result := strings.Split(masterServerReply.ReplyMessage, "|")
		chunkHandle := result[0]
		for i := 1; i < len(result); i++ {
			chunkServerSocket := "127.0.0.1:" + result[i]
			command := "create"
			// 转向与 ChunkServer 交互
			SwitchChunkServer(&chunkServerSocket, &command, &chunkHandle)
		}
	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}

// ListFiles 客户端展示文件
func ListFiles(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).ListFiles(*clientForMSCtx, &pb.Request{SendMessage: *filePath})
	// 根据 masterServer 的返回码来输出信息
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message("Response from masterServer: " + masterServerReply.ReplyMessage)
	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}

// WriteFile 客户端展示文件
func WriteFile(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string, data *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).WriteFile(*clientForMSCtx, &pb.Request{SendMessage: *filePath + *data})
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message(masterServerReply.ReplyMessage)
	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}
