package logic

import (
	"errors"
	"go.uber.org/zap"
	"stu_srv/dao/dao_mysql"
	"stu_srv/proto_stu/student"
)

func StuRegister(in *student.StuRegisterRequest) (*student.StuRegisterResponse, error) {
	name, err := dao_mysql.FindUserByName(in.Username)
	if err != nil {
		return nil, err
	}
	if name.ID != 0 {
		zap.L().Info("用户已存在！")
		return nil, errors.New("用户已存在！")
	}
	register, err := dao_mysql.StuRegister(in)
	if err != nil {
		return nil, err
	}
	if register.ID == 0 {
		zap.L().Info("用户注册失败！")
		return nil, errors.New("用户注册失败！")
	}
	return &student.StuRegisterResponse{
		StuID:   int64(register.ID),
		StuName: register.Username,
		Gender:  register.Gender,
	}, nil
}
