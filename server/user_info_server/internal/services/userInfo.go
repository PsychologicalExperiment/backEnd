package services

import (
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
)

type UserInfoServerImpl struct {
	sqlConn  *gorm.DB
	redisCli *redis.Client
}

type userInfo struct {
	gorm.Model         // ID,CreateAt,UpdateAt,DeletedAt
	Email       string `gorm:"type:varchar(100);unique_index"` //唯一索引
	PhoneNumber string `gorm:"type:varchar(20);unique"`
	UserName    string `gorm:"type:varchar(10)"`
	Gender      uint32 `gorm:"type:tinyint(3);not null"`
}

func (u *UserInfoServerImpl) insertUserInfo(
	user *userInfo,
) error {
	res := u.sqlConn.Create(user)
	if res.Error != nil {
		grpclog.Errorf("insert into db failed, error: %+v, userInfo: %+v", res.Error, user)
		return res.Error
	}
	return nil
}

func (u *UserInfoServerImpl) isEmailUsed(
	email string,
) (bool, error) {
	users := []userInfo{}
	res := u.sqlConn.Where("email = ?", email).Find(&users)
	if res.Error != nil {
		grpclog.Errorf("read db failed, error: %+v", res.Error)
		return false, res.Error
	}
	return len(users) > 0, nil
}
