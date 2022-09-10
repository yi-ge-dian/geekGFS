package main

import (
	"GeekGFS/src/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8002端口
	// 没有证书会报错
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	// 退出时关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Failed to close the connection", err)
		}
	}(conn)
	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := pb.NewMasterServerToClientClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := productServiceClient.ListFiles(ctx, &pb.Request{SendMessage: "111"})
	if err != nil {
		log.Fatal("Error in calling gRPC method: ", err)
	}
	fmt.Println("Success in calling gRPC method:", resp.ReplyMessage)
}
