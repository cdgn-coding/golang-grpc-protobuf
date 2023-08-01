package main

import (
	"github.com/cdgn-coding/golang-grpc-protobuf/internal/platform/grpc"
	"go.uber.org/zap"
)

type container struct {
	logger  *zap.Logger
	service grpc.StudentService
}

func bootstrap() *container {
	logger, err := zap.NewDevelopment()
	mustSuccess(err, "Cannot create logger")

	studentService := grpc.NewStudentService()
	dependencies := &container{
		logger:  logger,
		service: studentService,
	}

	return dependencies
}
