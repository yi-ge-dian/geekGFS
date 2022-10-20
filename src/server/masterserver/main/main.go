package main

import (
	"GeekGFS/src/pb"
	ms "GeekGFS/src/server/masterserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterMasterServerToClientServer(s, &ms.MasterServer{})
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal("Failed to listen to the port", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("Failed to close the port")
		}
	}(listener)
	log.Printf("server listening at %v", listener.Addr())
	_ = s.Serve(listener)
}
