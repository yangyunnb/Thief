package initserver

import (
	"github.com/jinzhu/configor"
)

var Conf config

type config struct {
	Server          ServerConfig    `yaml:"server"`
	Mysql           MysqlConfig     `yaml:"mysql"`
	WebSocketConfig WebSocketConfig `yaml:"WebSocketConfig"`
	RedisConfig     RedisConfig     `yaml:"redis"`
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

type WebSocketConfig struct {
	ReadBufferSize   int   `yaml:"readBufferSize"`
	WriteBufferSize  int   `yaml:"writeBufferSize"`
	HandshakeTimeout int64 `json:"handshakeTimeout"`
	WriteTimeout     int64 `json:"writeTimeout"`
	ReadTimeout      int64 `json:"readTimeout"`
}

type RedisConfig struct {
	Host         string `yaml:"host"`
	PoolSize     int    `yaml:"poolSize"`
	MaxConn      int    `yaml:"maxConn"`
	ConnTimeout  string `yaml:"connTimeout"`
	ReadTimeout  string `yaml:"readTimeout"`
	WriteTimeout string `yaml:"writeTimeout"`
}

func InitConfig() {
	const configFile = "config/dev.yaml"
	if err := configor.Load(&Conf, configFile); err != nil {
		panic(err)
	}
}
