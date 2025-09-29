package dao_mysql

import (
	"Common/models/mysql"
	"stu_srv/proto_stu/student"
)

func StuRegister(in *student.StuRegisterRequest) (s *mysql.Student, err error) {
	s = &mysql.Student{
		Username: in.Username,
		Password: in.Password,
		Gender:   in.Gender,
		Phone:    in.Phone,
		Email:    in.Email,
		Status:   1,
	}
	err = s.CreateUser()
	if err != nil {

	}
	return nil, nil
}
