package main

import (
	cl "GeekGFS/src/client"
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func printUsage() {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	logger.Info("Usage:")
	logger.Info("<command> " + " <filePath> " + "<args>(optional) ")
	logger.Info("create filePath")
	logger.Info("list filePath")
	logger.Info("write filePath offset data")
}

func main() {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	printUsage()
	if len(os.Args) < 3 {
		logger.Warn("输入参数过少，至少为3，请参考Usage")
		return
	}
	// 日志库，六个级别，Log、Message、Info、Warn、Debug、Error
	// 服务器运行情况: Info、Error()
	// 服务器交流信息: Message、Warn
	// 服务器查找Bug: Debug
	// 1. 建立连接，端口是服务端开放的30001端口 没有证书会报错
	masterServerSocket := "127.0.0.1:30001"
	conn, err := grpc.Dial(masterServerSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("client connected masterServer at " + masterServerSocket)
	// 退出时关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close the connection" + err.Error())
		}
	}(conn)

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := pb.NewMasterServerToClientClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 3. 根据 command 调用方法
	command := os.Args[1]
	switch command {
	case "create":

	}

	var args string
	for i := 3; i < len(os.Args); i++ {
		args = args + "|" + os.Args[i]
	}
	// command filePath
	cl.RunClient(os.Args[1], os.Args[2], args)
}
