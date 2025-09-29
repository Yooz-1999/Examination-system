package server

import (
	"context"
	"stu_srv/internal/logic"
	"stu_srv/proto_stu/student"
)

type ServerStudent struct {
	student.UnimplementedStudentServer
}

func (s ServerStudent) StuRegister(ctx context.Context, in *student.StuRegisterRequest) (*student.StuRegisterResponse, error) {
	register, err := logic.StuRegister(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}
