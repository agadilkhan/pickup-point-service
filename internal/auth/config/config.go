package config

import (
	"time"
)

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Database   `yaml:"Database"`
	Auth       `yaml:"Auth"`
	Transport  `yaml:"Transport"`
	Kafka      `yaml:"Kafka"`
	Redis      `yaml:"Redis"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type Database struct {
	Main    DBNode `yaml:"Main"`
	Replica DBNode `yaml:"Replica"`
}

type DBNode struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}

type Auth struct {
	JWTSecretKey      string `yaml:"JWTSecretKey"`
	PasswordSecretKey string `yaml:"PasswordSecretKey"`
}

type Transport struct {
	UserGrpc UserGrpcTransport `yaml:"UserGrpc"`
}

type UserGrpcTransport struct {
	Host string `yaml:"Host"`
}

type Kafka struct {
	Brokers  []string `yaml:"Brokers"`
	Producer Producer `yaml:"Producer"`
	Consumer Consumer `yaml:"Consumer"`
}

type Producer struct {
	Topic string `yaml:"topic"`
}

type Consumer struct {
	Topics []string `yaml:"topics"`
}

type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}
