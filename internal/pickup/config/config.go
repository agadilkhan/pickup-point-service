package config

import "time"

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Database   `yaml:"Database"`
	Auth       `yaml:"Auth"`
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
	JWTSecretKey string `yaml:"JWTSecretKey"`
}
