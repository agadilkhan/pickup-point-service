package postgres

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Db struct {
	*gorm.DB
}

type Config config.DBNode

func (c Config) dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
	)
}

func New(cfg config.DBNode) (*Db, error) {
	conf := Config(cfg)

	db, err := gorm.Open(postgres.Open(conf.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Db{DB: db}, nil
}

func (d *Db) Close() {
	db, err := d.DB.DB()
	if err != nil {
		log.Printf("error closing database: %s", err)
	}

	db.Close()
}
