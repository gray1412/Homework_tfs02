package storage

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Storage interface {
	ConnectDatabase() (db *gorm.DB)
	NewFileLogger(filepath string) (*zap.Logger, error)
}
