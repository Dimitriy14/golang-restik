package postgres

import (
	"fmt"

	"github.com/Dimitriy14/golang-restik/pkg/config"
	"github.com/Dimitriy14/golang-restik/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const dbInfo = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

var Client PGClient

type PGClient struct {
	Session *gorm.DB
}

func Load() error {
	url := config.Conf.HerokuPg
	if url == "" {
		url = fmt.Sprintf(
			dbInfo,
			config.Conf.Postgres.Host,
			config.Conf.Postgres.Port,
			config.Conf.Postgres.User,
			config.Conf.Postgres.Password,
			config.Conf.Postgres.DBName,
		)
	}

	db, err := gorm.Open("postgres", url)
	if err != nil {
		return err
	}

	Client = PGClient{Session: db}
	db.SetLogger(logger.NewGormLogger(logger.Log))
	db.LogMode(true)
	return nil
}
