package client

import (
	"google.golang.org/grpc/grpclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlInitOpt struct {
	Dsn string
}

func InitSqlClient(opt SqlInitOpt) *gorm.DB {
	db, err := gorm.Open(mysql.Open(opt.Dsn), &gorm.Config{})
	if err != nil {
		grpclog.Fatalf("init sql client failed|error: %+v, dsn: %s", err, opt.Dsn)
		return nil
	}
	return db
}
