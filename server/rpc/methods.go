package rpc

import (
	pb "echo/proto"
	"io"

	"go.uber.org/zap"
)

func (s *Server) EventStream(stream pb.Echo_EventStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.Log.Error(err.Error())
			return err
		}
		s.Log.Info("Accepted message: ", zap.String("msg", in.GetText()))

		if err := stream.Send(&pb.Events{
			Text: in.GetText(),
		}); err != nil {
			s.Log.Error(err.Error())
			return err
		}
		s.Log.Info("Sended message: ", zap.String("msg", in.GetText()))
	}
}
