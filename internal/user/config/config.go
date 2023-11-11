package config

import "time"

type Config struct {
	HttpServer `yaml:"HttpServer"`
	GrpcServer `yaml:"GrpcServer"`
	Database   `yaml:"Database"`
}

type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type GrpcServer struct {
	Port string `ymal:"Port"`
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
