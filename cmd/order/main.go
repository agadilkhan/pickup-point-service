package main

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/order/applicator"
	"github.com/agadilkhan/pickup-point-service/internal/order/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With("app", "order-service")

	cfg, err := loadConfig("config/order")
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

	dbPass := viper.GetString("DB_PASSWORD")

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("ReadInConfig err: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to Unmarshal config err: %v", err)
	}

	cfg.Main.Password = dbPass
	cfg.Replica.Password = dbPass

	return cfg, err
}
