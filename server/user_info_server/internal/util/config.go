package util

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

}
