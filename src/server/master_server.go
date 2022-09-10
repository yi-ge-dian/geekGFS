package main

import (
	"GeekGFS/src/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type masterServer struct {
	pb.UnimplementedMasterServerToClientServer
}

func (s *masterServer) CreateFile(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{ReplyMessage: "1", StatusCode: "1"}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterMasterServerToClientServer(s, &masterServer{})
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
