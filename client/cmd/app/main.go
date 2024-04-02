package main

import (
	"context"
	"echo/pkg/logging"
	pb "echo/proto"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logging.InitLogger()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	stream, err := client.EventStream(context.Background())
	if err != nil {
		logger.Fatal(err.Error())
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				logger.Error(err.Error())
				return
			}
			fmt.Printf("Accepted: %s\n", in.GetText())
		}
	}()

	fmt.Println("Waiting for messages...")
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "" {
			break
		}
		if err := stream.Send(&pb.Events{
			Text: msg,
		}); err != nil {
			logger.Fatal(err.Error())
		}
	}
	stream.CloseSend()
	<-waitc
}
