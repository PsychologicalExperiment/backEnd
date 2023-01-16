package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	Config GlobalConfig
)

type ServerConfig struct {
	ServerName string `yaml:"server_name"`
	Port       int32  `yaml:"port"`
	Log        string `yaml:"log"`
}

type PrometheusConfig struct {
	IP   string `yaml:"ip"`
	Port int32  `yaml:"port"`
}

type DbConfig struct {
	Master MysqlConfig `yaml:"master"`
	Slave  MysqlConfig `yaml:"slave"`
}

type MysqlConfig struct {
	IP     string `yaml:"ip"`
	Port   int32  `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

type EtcdConfig struct {
	IP   string `yaml:"ip"`
	Port int32  `yaml:"port"`
}

type GlobalConfig struct {
	Server       ServerConfig     `yaml:"server"`
	Monitor      PrometheusConfig `yaml:"prometheus"`
	Db           DbConfig         `yaml:"mysql"`
	NamingServer EtcdConfig       `yaml:"etcd"`
}

func init() {
	config, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Printf("read config failed, error: %+v", err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(config, &Config); err != nil {
		fmt.Printf("config.yaml error: %+v", err)
		os.Exit(1)
	}
}
