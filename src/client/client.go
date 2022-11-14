package client

import (
	"GeekGFS/src/pb"
	"github.com/sadlil/gologger"
)

func RunClient(command string, filePath string, args string) {

	// 4. 调用对应的函数
	var resp *pb.Reply
	switch command {
	case "create":
		resp, _ = productServiceClient.CreateFile(ctx, &pb.Request{SendMessage: filePath})
	case "list":
		resp, _ = productServiceClient.ListFiles(ctx, &pb.Request{SendMessage: filePath})
	case "write":
		if len(args) == 0 {
			resp.ReplyMessage = "No input data given to write"
			resp.StatusCode = "-1"
		} else {
			resp, _ = productServiceClient.WriteFile(ctx, &pb.Request{SendMessage: filePath + args})
		}
	default:
		PrintUsage()
	}
	// 5. 打印对应的信息
	switch resp.StatusCode {
	case "0":
		logger.Message(resp.ReplyMessage)
	default:
		logger.Warn(resp.ReplyMessage)
	}
}

func PrintUsage() {
	logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	logger.Info("Usage:")
	logger.Info("<command> " + " <filePath> " + "<args>(optional) ")
	logger.Info("create filePath")
	logger.Info("list filePath")
	logger.Info("write filePath offset data")
}
