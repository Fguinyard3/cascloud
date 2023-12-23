package db

import (
	"fmt"

	"cascloud/config"
	model "cascloud/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB(cfg *config.Config) (*gorm.DB, error) {
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := gorm.Open(postgres.Open(info), &gorm.Config{})

	if cfg.Environment == "dev" {
		return db, nil
	}
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database")
		return nil, err
	}

	// we need to do auto migration for our models
	migrateErr := db.AutoMigrate(
		&model.User{},
		&model.Workspace{},
		&model.Collaborations{},
		&model.Role{},
		&model.File{},
		&model.Folder{},
	)
	if migrateErr != nil {
		log.Error().Err(migrateErr).Msg("Error migrating models")
		return nil, migrateErr
	}

	return db, nil
}
