package services

import (
	"fmt"

	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
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
		return serverErr.New(serverErr.ErrMySqlError)
	}
	return nil
}

func (u *UserInfoServerImpl) isUinqueKeyUsed(
	uniqueKey, queryKey string,
) (bool, error) {
	users := []userInfo{}
	res := u.sqlConn.Where(fmt.Sprintf("%s = ?", queryKey), uniqueKey).Find(&users)
	if res.Error != nil {
		grpclog.Errorf("read db failed, error: %+v", res.Error)
		return false, serverErr.New(serverErr.ErrMySqlError)
	}
	return len(users) > 0, nil
}
