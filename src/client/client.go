package client

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	"context"
	"fmt"
	"github.com/sadlil/gologger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"strings"
	"time"
)

//************************************辅助函数************************************

// SwitchChunkServer client 与 chunkServer 建立连接,并且执行逻辑
func SwitchChunkServer(chunkServerSocket *string, command *string, sendData *string) string {
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
		chunkServerReply, _ := clientForCS.Create(clientForCSCtx, &pb.Request{SendMessage: *sendData})
		// 根据 chunkServer 的返回码来输出信息
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
		return ""
	case "write":
		chunkServerReply, _ := clientForCS.Write(clientForCSCtx, &pb.Request{SendMessage: *sendData})
		// 根据 chunkServer 的返回码来输出信息
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
		return ""
	case "read":
		chunkServerReply, _ := clientForCS.Read(clientForCSCtx, &pb.Request{SendMessage: *sendData})
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
		return chunkServerReply.ReplyMessage
	case "getChunkSpace":
		chunkServerReply, _ := clientForCS.GetChunkSpace(clientForCSCtx, &pb.Request{SendMessage: *sendData})
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
		return chunkServerReply.ReplyMessage
	case "append":
		chunkServerReply, _ := clientForCS.Append(clientForCSCtx, &pb.Request{SendMessage: *sendData})
		switch chunkServerReply.StatusCode {
		case "0":
			logger.Message("Response from chunkServer: " + chunkServerReply.ReplyMessage)
		default:
			logger.Warn(chunkServerReply.ReplyMessage)
		}
		return ""
	default:
		logger.Warn("No this command")
		return ""
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

// WriteFile 客户端写文件
func WriteFile(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string, data *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).WriteFile(*clientForMSCtx, &pb.Request{SendMessage: *filePath + "|" + *data})
	// 根据 masterServer 的返回码来输出信息
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message(masterServerReply.ReplyMessage)
		result := strings.Split(masterServerReply.ReplyMessage, "|")
		// 切片第一个是空，直接移除
		result = result[1:]
		size := len(result)
		// 定义变量存放切片中的数据
		var ports []string
		chunkHandle := ""
		for i := 0; i < size; i++ {
			if i%4 == 0 {
				dataStart := i / 4 * 64
				dataSize := 64
				chunkHandle = result[i]
				ports = append(ports, result[i+1], result[i+2], result[i+3])
				for portId := 0; portId < len(ports); portId++ {
					chunkServerSocket := "127.0.0.1:" + ports[portId]
					command := "write"
					sendData := ""
					if dataStart < len(*data) {
						if dataStart+dataSize < len(*data) {
							sendData = chunkHandle + "|" + (*data)[dataStart:dataStart+dataSize]
						} else {
							sendData = chunkHandle + "|" + (*data)[dataStart:len(*data)]
						}
					}
					// 转向与 ChunkServer 交互
					SwitchChunkServer(&chunkServerSocket, &command, &sendData)
				}
				// 优雅的清空切片
				ports = ports[0:0]
			}
		}
	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}

// ReadFile 客户端读文件
func ReadFile(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).ReadFile(*clientForMSCtx, &pb.Request{SendMessage: *filePath})
	// 根据 masterServer 的返回码来输出信息
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message("Response from masterServer: " + masterServerReply.ReplyMessage)
		result := strings.Split(masterServerReply.ReplyMessage, "|")
		// 切片第一个是空，直接移除
		result = result[1:]
		size := len(result)
		data := ""
		for i := 0; i < size; i = i + 2 {
			command := "read"
			chunkHandle := result[i]
			location := result[i+1]
			chunkServerSocket := "127.0.0.1:" + location
			messageFromChunkServer := SwitchChunkServer(&chunkServerSocket, &command, &chunkHandle)
			data = data + messageFromChunkServer
			logger.Message("Response from chunkServer, all data = " + data)
		}
	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}

// AppendFile 客户端追加文件
func AppendFile(clientForMS *pb.MasterServerToClientClient, clientForMSCtx *context.Context, filePath *string, data *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	masterServerReply, _ := (*clientForMS).WriteFile(*clientForMSCtx, &pb.Request{SendMessage: *filePath + "|" + *data})
	// 根据 masterServer 的返回码来输出信息
	switch masterServerReply.StatusCode {
	case "0":
		logger.Message("Response from masterServer: " + masterServerReply.ReplyMessage)
		result := strings.Split(masterServerReply.ReplyMessage, "|")
		// 切片第一个是空，直接移除
		result = result[1:]
		latestChunkHandle := result[0]
		// 向 chunkServer 询问 这个chunk 还有多少空间
		chunkServerSocket := "127.0.0.1:" + result[1]
		command := "getChunkSpace"
		existSizeString := SwitchChunkServer(&chunkServerSocket, &command, &latestChunkHandle)
		existSize, err := strconv.Atoi(existSizeString)
		if err != nil {
			return
		}
		fmt.Println(existSize)
		availableSize := cm.GFSChunkSize - existSize
		fmt.Println(availableSize)
		// 看看我要追加的数据有多大
		dataSize := len(*data)
		// 如果我的数据比可用的小，那么我直接追加就可以了
		if dataSize <= availableSize {
			for i := 1; i < len(result); i++ {
				chunkServerSocket_ := "127.0.0.1:" + result[i]
				command_ := "append"
				sendData := latestChunkHandle + "|" + *data
				SwitchChunkServer(&chunkServerSocket_, &command_, &sendData)
			}
		} else {
			// todo：***
			// 如果我的数据比可用的大，我追加完成之后需要再创建新的，直至追加完毕
		}

	default:
		logger.Warn(masterServerReply.ReplyMessage)
	}
}
