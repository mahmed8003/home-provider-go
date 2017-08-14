package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
Server :
*/
type Server struct {
	Port       string `mapstructure:"port"`
	EnableLogs bool   `mapstructure:"enable_logs"`
}

/*
Jwt :
*/
type Jwt struct {
	Algorithm string `mapstructure:"algorithm"`
	Secret    string `mapstructure:"secret"`
	Duration  int    `mapstructure:"duration"`
}

/*
Database :
*/
type Database struct {
	Uri      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

/*
Redis :
*/
type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

/*
AppConfig :
*/
type AppConfig struct {
	LogLevel string   `mapstructure:"log_level"`
	Server   Server   `mapstructure:"server"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
}

var config AppConfig

/*
LoadConfig : Load configuration based on envviourment
*/
func LoadConfig(env string) error {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName(env)
	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	return v.Unmarshal(&config)
}

/*
GetConfig : It will return
*/
func GetConfig() AppConfig {
	return config
}
