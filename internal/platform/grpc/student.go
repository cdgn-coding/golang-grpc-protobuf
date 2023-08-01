package grpc

import (
	context "context"
	"github.com/cdgn-coding/golang-grpc-protobuf/proto"
)

type StudentService struct {
	proto.UnimplementedStudentServiceServer
}

func NewStudentService() StudentService {
	return StudentService{}
}

func (s StudentService) GetStudent(ctx context.Context, request *proto.GetStudentRequest) (*proto.Student, error) {
	student := &proto.Student{
		Id:   "Fake",
		Name: "Fake",
		Age:  30,
	}
	return student, nil
}

func (s StudentService) SetStudent(ctx context.Context, student *proto.Student) (*proto.SetStudentResponse, error) {
	//TODO implement me
	panic("implement me")
}
