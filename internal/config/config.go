package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	GrpcPort     int
	AMQPHost     string
	AMQPPort     int
	AMQPLogin    string
	AMQPPassword string
	AMQPVhost    string
	DBHost       string
	DBPort       int
	DBName       string
	DBUser       string
	DBPassword   string
}

var (
	ErrReadFile      = errors.New("can't read file")
	ErrWithValue      = errors.New("error reading value")
)

func Init() (*Configuration, error) {
	err := godotenv.Overload()
	if err != nil {
		return nil, err
	}

	grpcPort, err1 := strconv.Atoi(os.Getenv("GRPC_PORT"))
	amqpPort, err2 := strconv.Atoi(os.Getenv("AMQP_PORT"))
	dbPort, err3 := strconv.Atoi(os.Getenv("DB_PORT"))

	if err1 != nil || err2 != nil || err3 != nil {
		return nil, ErrWithValue
	}

	return &Configuration{
		GrpcPort:     grpcPort,
		AMQPHost:     os.Getenv("AMQP_HOST"),
		AMQPPort:     amqpPort,
		AMQPLogin:    os.Getenv("AMQP_LOGIN"),
		AMQPPassword: os.Getenv("AMQP_PASSWORD"),
		AMQPVhost:    os.Getenv("AMQP_VHOST"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       dbPort,
		DBName:       os.Getenv("DB_DATABASE"),
		DBUser:       os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
	}, nil
}
