package dao_mysql

import (
	"Common/models/mysql"
)

func FindUserByName(username string) (s *mysql.Student, err error) {
	s = &mysql.Student{}
	err = s.FindUserByName(username)
	if err != nil {
		return nil, err
	}
	return s, nil
}
