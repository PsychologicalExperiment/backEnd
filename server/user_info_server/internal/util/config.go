package util

import (
	"fmt"
	"github.com/PsychologicalExperiment/backEnd/util/plugins"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SqlConfig       MySqlConf            `yaml:"MySqlConf"`
	TokenSecretKey  string               `yaml:"tokenSecretKey"`
	TokenExpireHour int                  `yaml:"tokenExpireHour"`
	LoggerConfig    plugins.LoggerConfig `yaml:"LoggerConfig"`
}

type MySqlConf struct {
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	TableName string `yaml:"tableName"`
}

var GConfig Config

func InitConfig() {
	config, err := ioutil.ReadFile("./configs/dev/config.yaml")
	if err != nil {
		fmt.Printf("read config failed, error: %+v", err)
		return
	}
	err = yaml.Unmarshal(config, &GConfig)
	if err != nil {
		fmt.Printf("decode yaml failed, error: %+v", err)
		return
	}
}
