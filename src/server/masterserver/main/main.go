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
	// 服务器运行情况:INFO、ERROR
	// 交流信息：Message
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// grpc服务器
	s := grpc.NewServer()
	masterAddressPort := "127.0.0.1:30001"
	gfsConfig := new(cm.GFSConfig)
	gfsConfig.Start()
	masterServer := new(ms.MasterServer)
	masterServer.MasterService(masterAddressPort, gfsConfig.GetChunkServerLocations())
	// 注册服务器段服务
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
	logger.Info("server listening at " + listener.Addr().String())
	_ = s.Serve(listener)
}
