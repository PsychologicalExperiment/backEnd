package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"

	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
)

const (
	searchKeyEmail       string = "email"
	searchKeyPhoneNumber string = "phone_number"
)

type UserInfoServerImpl struct {
	sqlConn  *gorm.DB
	redisCli *redis.Client
}

func NewUserInfoServerImpl(db *gorm.DB) *UserInfoServerImpl {
	//return &UserInfoServerImpl{db.Table("user_info"), nil}
	return &UserInfoServerImpl{db, nil}
}

type userInfo struct {
	gorm.Model         // ID,CreateAt,UpdateAt,DeletedAt
	Email       string `gorm:"type:varchar(100);unique_index"` //唯一索引
	PhoneNumber string `gorm:"type:varchar(20);unique"`
	UserName    string `gorm:"type:varchar(10)"`
	Gender      uint32 `gorm:"type:tinyint(3);not null"`
	Password    string `gorm:"type:varchar(20)"`
	UserType    uint32 `gorm:"type:tinyint(3);index"`
	Extra       string `gorm:"type:text"`
}

func (u *UserInfoServerImpl) insertUserInfo(
	user *userInfo,
) error {
	res := u.sqlConn.Debug().Create(user)
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

func (u *UserInfoServerImpl) getUserInfosByKey(
	queryKey, queryVal string,
) ([]userInfo, error) {
	users := []userInfo{}
	res := u.sqlConn.Table("user_info").Where(fmt.Sprintf("%s = ?", queryKey), queryVal).Debug().Find(&users)
	//res := u.sqlConn.Where(fmt.Sprintf("%s = ?", queryKey), queryVal).Debug().Find(&users)
	//_ = u.sqlConn.Where(fmt.Sprintf("%s = ?", queryKey), queryVal).Debug().Find(&users)
	if res.Error != nil {
		grpclog.Errorf("read db failed, error: %+v", res.Error)
		return nil, serverErr.New(serverErr.ErrMySqlError)
	}
	return users, nil
}

func (u *UserInfoServerImpl) setTokenForUser(
	ctx context.Context,
	email, token string,
) error {
	err := u.redisCli.Set(ctx, email, token, 0).Err()
	if err != nil {
		grpclog.Errorf("set redis failed, error: %+v", err)
		return serverErr.New(serverErr.ErrSetRedisFailed)
	}
	return nil
}
