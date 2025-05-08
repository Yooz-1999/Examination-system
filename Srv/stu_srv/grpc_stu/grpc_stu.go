package grpc_stu

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func RegisterStuGrpc(call func(grpc *grpc.Server)) {
	// 1.监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 9991))
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	// 2.实例化gRPC
	s := grpc.NewServer()
	call(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	// 3.在gRPC上注册微服务
	//pb.RegisterUserInfoServiceServer(s, &u)
	// 4.启动服务端
	log.Printf("server listening at %v", listener.Addr())
	if err = s.Serve(listener); err != nil {
		log.Fatalf("server failed to listening at %v", err)
	}

}
