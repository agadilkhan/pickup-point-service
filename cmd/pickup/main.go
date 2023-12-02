package main

import (
	"fmt"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/applicator"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	//nolint:all
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With("app", "pickup-service")

	cfg, err := loadConfig("config/pickup")
	if err != nil {
		l.Fatalf("failed to load config: %v", err)
	}

	app := applicator.NewApplicator(cfg, l)
	app.Run()
}

func loadConfig(path string) (cfg *config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("ReadInConfig err: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to Unmarshal config err: %v", err)
	}

	return cfg, err
}
