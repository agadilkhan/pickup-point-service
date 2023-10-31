package main

import (
	"github.com/agadilkhan/pickup-point-service/internal/auth/applicator"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "auth-service"))

	cfg, err := loadConfig("config/auth")
	if err != nil {
		l.Fatalf("failed to load config: %v", err)
	}

	app := applicator.NewApplicator(l, &cfg)
	app.Run()
}

func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	dbPass := viper.GetString("DB_PASSWORD")

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	config.Database.Main.Password = dbPass
	config.Database.Replica.Password = dbPass

	return config, err
}
