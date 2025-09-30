package dao_mysql

import "Common/models/mysql"

func FindUserByUserId(userid int) (m *mysql.Student, err error) {
	m = &mysql.Student{}
	err = m.FindUserById(userid)
	if err != nil {
		return nil, err
	}
	return m, nil
}
