package initserver

import (
	"github.com/jinzhu/configor"
	"github.com/kataras/iris"
)

var Conf config

type config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port                              int    `yaml:"port"`
	Charset                           string `yaml:"charset"`
	DisablePathCorrection             bool   `yaml:"disablePathCorrection"`
	EnablePathEscape                  bool   `yaml:"enablePathEscape"`
	FireMethodNotAllowed              bool   `yaml:"fireMethodNotAllowed"`
	DisableBodyConsumptionOnUnmarshal bool   `yaml:"disableBodyConsumptionOnUnmarshal"`
}

func InitConfig() {
	const configFile = "config/dev.yaml"
	iris.YAML(configFile)
	if err := configor.Load(&Conf, configFile); err != nil {
		panic(err)
	}
}
