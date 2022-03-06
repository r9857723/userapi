package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port        int    `json:"port"`
	HostName    string `json:"hostname"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	MaxLifeTime int    `json:"maxlifetime"`
	MaxOpenConn int    `json:"maxopenconn"`
	MaxIdleConn int    `json:"maxidleconn"`
	DataBase    string `json:"database"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = loadConfig()
	}
	return config
}

func loadConfig() *Config {
	config := &Config{}
	viper := viper.New()
	viper.SetConfigFile(".yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper ReadConfig err : %s", err)
	}

	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("viper unmarshal err : %s", err)
	}
	return config
}
