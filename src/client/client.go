package client

import (
	"GeekGFS/src/pb"
	"context"
	"github.com/sadlil/gologger"
)

// CreateFile 客户端创建文件
func CreateFile(productServiceClient *pb.MasterServerToClientClient, ctx *context.Context, filePath *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	var resp *pb.Reply
	resp, _ = (*productServiceClient).CreateFile(*ctx, &pb.Request{SendMessage: *filePath})
	switch resp.StatusCode {
	case "0":
		logger.Message(resp.ReplyMessage)
	default:
		logger.Warn(resp.ReplyMessage)
	}
}

// ListFiles 客户端展示文件
func ListFiles(productServiceClient *pb.MasterServerToClientClient, ctx *context.Context, filePath *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	var resp *pb.Reply
	resp, _ = (*productServiceClient).ListFiles(*ctx, &pb.Request{SendMessage: *filePath})
	switch resp.StatusCode {
	case "0":
		logger.Message(resp.ReplyMessage)
	default:
		logger.Warn(resp.ReplyMessage)
	}
}

// WriteFile 客户端展示文件
func WriteFile(productServiceClient *pb.MasterServerToClientClient, ctx *context.Context, filePath *string, data *string) {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	var resp *pb.Reply
	resp, _ = (*productServiceClient).ListFiles(*ctx, &pb.Request{SendMessage: *filePath + *data})
	switch resp.StatusCode {
	case "0":
		logger.Message(resp.ReplyMessage)
	default:
		logger.Warn(resp.ReplyMessage)
	}
}
