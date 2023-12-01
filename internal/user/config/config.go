package config

type Config struct {
	GrpcServer `yaml:"GrpcServer"`
	Database   `yaml:"Database"`
}

type GrpcServer struct {
	Port string `yaml:"Port"`
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
