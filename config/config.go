package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreGorm(logger *logrus.Logger, config Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s %s ",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("not connect gorm postgre", err.Error())
		return nil
	}
	return db
}
