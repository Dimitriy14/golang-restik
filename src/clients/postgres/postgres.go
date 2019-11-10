package postgres

import (
	"fmt"

	"github.com/Dimitriy14/golang-restik/src/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const dbInfo = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

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

	_ = db
	return nil
}
