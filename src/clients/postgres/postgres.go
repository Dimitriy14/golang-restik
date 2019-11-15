package postgres

import (
	"fmt"

	"github.com/Dimitriy14/golang-restik/src/config"
	"github.com/Dimitriy14/golang-restik/src/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const dbInfo = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

var Client PGClient

type PGClient struct {
	Session *gorm.DB
}

func Load() error {
	db, err := gorm.Open("postgres",
		fmt.Sprintf(
			dbInfo,
			config.Conf.Postgres.Host,
			config.Conf.Postgres.Port,
			config.Conf.Postgres.User,
			config.Conf.Postgres.Password,
			config.Conf.Postgres.DBName,
		))
	if err != nil {
		return err
	}

	Client = PGClient{Session: db}
	db.SetLogger(logger.NewGormLogger(logger.Log))
	db.LogMode(true)
	return nil
}
