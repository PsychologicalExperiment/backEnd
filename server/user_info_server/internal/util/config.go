package util

import (
	"io/ioutil"

	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SqlConfig MySqlConf `yaml:"MySqlConf"`
}

type MySqlConf struct {
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var GConfig Config

func InitConfig() {
	config, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		grpclog.Errorf("read config failed, error: %+v", err)
		return
	}
	err = yaml.Unmarshal(config, &GConfig)
	if err != nil {
		grpclog.Errorf("decode yaml failed, error: %+v", err)
		return
	}
}
