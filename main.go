package main

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var log hclog.Logger

const (
	port = ":50051"
)

type server struct{} // 여기에다가 추가해줘야 함.

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
