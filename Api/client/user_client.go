package client

import (
	"context"
	"google.golang.org/grpc"
	"stu_srv/proto_stu/student"
)

type HandlerStudent func(ctx context.Context, in student.StudentClient) (interface{}, error)

func UserClient(ctx context.Context, handlerStudent HandlerStudent) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:9991", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := student.NewStudentClient(dial)
	res, err := handlerStudent(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
