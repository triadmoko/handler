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
