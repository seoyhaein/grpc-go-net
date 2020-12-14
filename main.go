package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", port) // func Listen(net, laddr string) (Listener, error): 프로토콜, IP 주소, 포트 번호를 설정하여 네트워크 연결을 대기합니다.

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //NewServer는 등록 된 서비스가없고 아직 요청 수락을 시작하지 않은 gRPC 서버를 만듭니다.

	if err := s.Serve(lis); err != nil { // https://github.com/grpc/grpc-go/blob/81b95b1854d7caf3cc21aed316fc222e1749cf31/server.go#L692
		log.Fatalf("failed to serve: %v", err)
	}
}
