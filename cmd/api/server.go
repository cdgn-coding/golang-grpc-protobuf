package main

import (
	"github.com/cdgn-coding/golang-grpc-protobuf/cmd/api/middleware"
	"github.com/cdgn-coding/golang-grpc-protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func start(container *container) {
	listener, err := net.Listen("tcp", ":50051")
	mustSuccess(err, "Cannot create listener")

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.NewLoggerInterceptor(container.logger),
			middleware.NewRecoveryInterceptor(),
		),
	)

	server.RegisterService(&proto.StudentService_ServiceDesc, container.service)
	reflection.Register(server)

	err = server.Serve(listener)
	mustSuccess(err, "Cannot start server")
}
