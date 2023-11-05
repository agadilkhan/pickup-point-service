package postgres

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/user/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		return nil, fmt.Errorf("connection open err: %v", err)
	}

	return &Db{DB: db}, nil
}

func (d *Db) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to returning *sql.DB err: %v", err)
	}

	return db.Close()
}
