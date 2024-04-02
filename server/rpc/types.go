package rpc

import (
	pb "echo/proto"

	"go.uber.org/zap"
)

type Server struct {
	pb.EchoServer
	Log *zap.Logger
}
