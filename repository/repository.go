package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
}
type repository struct {
	loggger *logrus.Logger
	dbGorm  *gorm.DB
}

func NewRepository(loggger *logrus.Logger, dbGorm *gorm.DB) *repository {
	return &repository{loggger: loggger, dbGorm: dbGorm}
}

func (r *repository) Ping() (string, error) {
	// Get generic database object sql.DB to use its functions
	sqlDB, err := r.dbGorm.DB()
	if err != nil {
		return "", err
	}
	return "connected", sqlDB.Ping()
}
