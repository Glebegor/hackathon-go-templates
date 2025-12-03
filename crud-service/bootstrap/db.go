package bootstrap

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(config DbConfig, logger *slog.Logger) (*gorm.DB, error) {
	if logger == nil {
		logger = slog.Default()
	}

	db, err := gorm.Open(postgres.Open(config.DbUrl))
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to the database successfully")
	return db, nil
}
