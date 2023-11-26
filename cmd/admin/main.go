package main

import (
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/admin/applicator"
	"github.com/agadilkhan/pickup-point-service/internal/admin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	//nolint:all
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "admin-service"))

	cfg, err := loadConfig("config/admin")
	if err != nil {
		l.Fatalf("failed to load config: %v", err)
	}

	app := applicator.NewApplicator(l, cfg)
	app.Run()
}

func loadConfig(path string) (config *config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %v", err)
	}

	return config, err
}
