package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	appConf *AppConfiguration
)

// AppConfiguration ...
type AppConfiguration struct {
	AppName string   `json:"appName"`
	Mode    string   `json:"mode"` //debug, release, test
	Port    string   `json:"port"`
	DB      DBConfig `json:"db"`
}

// DBConfig ...
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Encoding string `json:"encoding"`
	PoolSize int    `json:"poolsize"`
}

func AppConfig(path string) *AppConfiguration {
	if appConf == nil {
		c := &AppConfiguration{}
		loadConfig(path, c)
		appConf = c
	}

	return appConf
}

func loadConfig(path string, c *AppConfiguration) {
	rawData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rawData, c)
	if err != nil {
		panic(err)
	}
}
