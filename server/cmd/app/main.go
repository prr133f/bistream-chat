package main

import (
	"echo/pkg/logging"
	"echo/server/rpc"
	"fmt"
	"net"

	pb "echo/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := logging.InitLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		logger.Fatal("error while listening tcp", zap.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterEchoServer(s, &rpc.Server{Log: logger})

	logger.Info("Serving app")
	if err := s.Serve(lis); err != nil {
		logger.Fatal("error while serving server", zap.Error(err))
	}
}
