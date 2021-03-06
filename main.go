package main

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"

	pb "github.com/seoyhaein/grpc-go-net/protos"
)

var log hclog.Logger

const (
	port = ":50051"
)

type server struct{ pb.UnimplementedGreeterServer } // 여기에다가 추가해줘야 함.

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Info("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	log = hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()
	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(gs)

	s := &server{}
	// grpc 서비스를 등록해야함.
	pb.RegisterGreeterServer(gs, s)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("연결안됨", "error", err)
		os.Exit(1)
	}

	if err := gs.Serve(l); err != nil {
		log.Error("grpc 서버 실패", "error", err)
	}

	log.Info("서비스 시작 localhost:50051")
}
