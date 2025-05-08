package handler

import (
	"Api/client"
	"context"
	"stu_srv/proto_stu/student"
)

func StuRegister(ctx context.Context, i *student.StuRegisterRequest) (*student.StuRegisterResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in student.StudentClient) (interface{}, error) {
		register, err := in.StuRegister(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*student.StuRegisterResponse), nil
}
