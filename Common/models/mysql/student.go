package mysql

import (
	"Common/global"
	"gorm.io/gorm"
)

// TODO:注册学员
type Student struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);not null;comment:'学员昵称'"`
	Password string `gorm:"type:varchar(30);not null;comment:'学员密码'"`
	Gender   string `gorm:"type:varchar(10);not null;comment:'学员性别'"`
	Phone    string `gorm:"type:varchar(11);not null;comment:'学员电话'"`
	Email    string `gorm:"type:varchar(30);not null;comment:'学员邮箱'"`
	Status   int    `gorm:"type:int;not null;comment:'学员状态'"`
}

func (s *Student) CreateUser() error {
	return global.DB.Create(&s).Error
}

func (s *Student) FindUserByName(username string) error {
	return global.DB.Where("username = ?", username).Limit(1).Find(&s).Error
}
