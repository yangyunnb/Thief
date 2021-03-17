package initserver

import (
	"github.com/jinzhu/configor"
)

var Conf config

type config struct {
	Server ServerConfig `yaml:"server"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

type ServerConfig struct {
	Port                              int    `yaml:"port"`
	Charset                           string `yaml:"charset"`
	DisablePathCorrection             bool   `yaml:"disablePathCorrection"`
	EnablePathEscape                  bool   `yaml:"enablePathEscape"`
	FireMethodNotAllowed              bool   `yaml:"fireMethodNotAllowed"`
	DisableBodyConsumptionOnUnmarshal bool   `yaml:"disableBodyConsumptionOnUnmarshal"`
}

type MysqlConfig struct {
	DSN         string `yaml:"dsn"`
	MaxIDConn   int    `yaml:"maxIdConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

func InitConfig() {
	const configFile = "config/dev.yaml"
	if err := configor.Load(&Conf, configFile); err != nil {
		panic(err)
	}
}
