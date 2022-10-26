package main

import (
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	// 日志库，六个级别，Log、Message、Info、Warn、Debug、Error
	// 服务器运行情况:Info、Error
	// 服务器交流信息：Message、Warn
	// 服务器查找Bug：Debug
	// 打印消息：Log
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// 1. 建立连接，端口是服务端开放的30001端口 没有证书会报错
	masterServerTarget := "127.0.0.1:30001"
	conn, err := grpc.Dial("127.0.0.1:30001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("client connected masterServer at " + masterServerTarget)
	// 退出时关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close the connection" + err.Error())
		}
	}(conn)

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := pb.NewMasterServerToClientClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. todo：截取命令行参数，现在只是书写测试即可
	resp, _ := productServiceClient.CreateFile(ctx, &pb.Request{SendMessage: "/home/1.txt"})
	switch resp.StatusCode {
	case "0":
		logger.Message(resp.ReplyMessage)
	default:
		logger.Warn(resp.ReplyMessage)
	}
}
