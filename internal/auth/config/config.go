package config

import "time"

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Database   `yaml:"Database"`
	Auth       `yaml:"Auth"`
	Transport  `yaml:"Transport"`
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
	Password string `env-required:"true" env:"DB_PASSWORD"`
	Name     string `yaml:"Name"`
}

type Auth struct {
	JWTSecretKey      string `yaml:"JWTSecretKey"`
	PasswordSecretKey string `yaml:"PasswordSecretKey"`
}

type Transport struct {
	User UserTransport `yaml:"User"`
}

type UserTransport struct {
	Host    string        `yaml:"Host"`
	Timeout time.Duration `yaml:"Timeout"`
}
