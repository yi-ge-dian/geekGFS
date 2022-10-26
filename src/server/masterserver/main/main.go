package main

import (
	cm "GeekGFS/src/common"
	"GeekGFS/src/pb"
	ms "GeekGFS/src/server/masterserver"
	"github.com/sadlil/gologger"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 日志库，六个级别，Log、Message、Info、Warn、Debug、Error
	// 服务器运行情况:Info、Error
	// 服务器交流信息：Message、Warn
	// 服务器查找Bug：Debug
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// masterServer配置
	masterAddressPort := "127.0.0.1:30001"
	gfsConfig := cm.NewGFSConfig(cm.GFSChunkSize, cm.GFSChunkServerLocations, cm.GFSChunkServerRoot)
	masterServer := ms.NewMasterServer(&masterAddressPort, gfsConfig.ChunkServerLocations())
	// grpc服务器
	s := grpc.NewServer()
	// 注册服务器端服务
	pb.RegisterMasterServerToClientServer(s, masterServer)
	listener, err := net.Listen("tcp", masterAddressPort)
	if err != nil {
		logger.Error("Failed to listen to the port" + err.Error())
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			logger.Error("Failed to close the port")
		}
	}(listener)
	// 启动监听
	logger.Info("masterServer listening at " + listener.Addr().String())
	_ = s.Serve(listener)
}
