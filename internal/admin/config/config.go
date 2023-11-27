package config

import "time"

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Auth       `yaml:"Auth"`
	Transport  `yaml:"Transport"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
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
