package services

import (
	"gorm.io/gorm"
)

type UserInfoServerImpl struct {
	sqlConn *gorm.DB
}

type userInfo struct {
	gorm.Model         // ID,CreateAt,UpdateAt,DeletedAt
	Email       string `gorm:"type:varchar(100);unique_index"` //唯一索引
	PhoneNumber string `gorm:"type:varchar(20);unique"`
	Uin         string `gorm:"unique;not null"` //唯一并且不为空
	UserName    string `gorm:"type:varchar(10)"`
	Gender      uint32 `gorm:"type:tinyint(3);not null"`
}
