package config

import (
	"errors"

	"github.com/spf13/viper"
)

type GRPCServer struct {
	Host string
	Port string
}

type Database struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

type Configuration struct {
	GRPCServer GRPCServer
	Database   Database
}

var (
	ErrFilePathEmpty = errors.New("file path is empty")
	ErrReadFile      = errors.New("can't read file")
)

func Init(path string) (Configuration, error) {
	var configuration Configuration

	if path == "" {
		return configuration, ErrFilePathEmpty
	}

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return configuration, ErrReadFile
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return configuration, ErrReadFile
	}

	return configuration, nil
}
