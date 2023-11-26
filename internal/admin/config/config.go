package config

import "time"

type Config struct {
	Auth       `yaml:"Auth"`
	HTTPServer `yaml:"HTTPServer"`
	Transport  `yaml:"Transport"`
}

type Auth struct {
	JWTSecretKey string `yaml:"JWTSecretKey"`
}

type HTTPServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type Transport struct {
	UserGrpcTransport `yaml:"UserGrpc"`
}

type UserGrpcTransport struct {
	Host string `yaml:"Host"`
}
