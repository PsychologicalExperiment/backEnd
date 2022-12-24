package util

import (
	"fmt"
	"github.com/PsychologicalExperiment/backEnd/util/plugins"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	SqlConfig       MySqlConf            `yaml:"MySqlConf"`
	TokenSecretKey  string               `yaml:"tokenSecretKey"`
	TokenExpireHour int                  `yaml:"tokenExpireHour"`
	LoggerConfig    plugins.LoggerConfig `yaml:"LoggerConfig"`
}

type MySqlConf struct {
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
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
