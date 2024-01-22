package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel   string `json:"LogLevel"`
	AppID      uint64 `json:"AppID"`
	Token      string `json:"Token"`
	ServerPath string `json:"ServerPath"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(confName string) error {
	filePath := "./" + confName
	f, err := os.Open(filePath)
	if err != nil {
		return FileNotExist
	}
	defer func() {
		_ = f.Close()
	}()
	c := new(Config)
	d := json.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return err
	}
	CONF = c
	return nil
}

var DefaultConfig = &Config{
	LogLevel:   "Info",
	ServerPath: "./bedrock_server.exe",
}
