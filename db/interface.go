package db

import (
	"errors"
	"fmt"
	"github.com/arung-agamani/mitsukeru-go/config"
	"github.com/arung-agamani/mitsukeru-go/models"
	"github.com/arung-agamani/mitsukeru-go/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	switch config.AppConfig.DbConfig.DatabaseType {
	case "sqlite":
		gormDb, err := gorm.Open(sqlite.Open(config.AppConfig.DbConfig.DatabaseLink), &gorm.Config{})
		if err != nil {
			logger.Errorf("Error when initializing database: %v", err)
			panic(err)
		}
		db = gormDb
		break
	case "postgres":
		dbconfig := config.AppConfig.DbConfig
		//connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		//	dbconfig.DatabaseUser,
		//	dbconfig.DatabasePass,
		//	dbconfig.DatabaseLink,
		//	dbconfig.DatabasePort,
		//	dbconfig.DatabaseName)
		gormDb, err := gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s",
				dbconfig.DatabaseUser,
				dbconfig.DatabasePass,
				dbconfig.DatabaseName,
				dbconfig.DatabaseLink,
				dbconfig.DatabasePort,
			),
		}), &gorm.Config{})
		if err != nil {
			logger.Errorf("Error when initializing database: %v", err)
			panic(err)
		}
		db = gormDb
		break
	default:
		logger.Errorf("Unsupported database type. Only accepts \"sqlite\" for now.")
		panic(0)
	}
	AutoMigrate(db)
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.LostItem{}, &models.FoundItem{}, &models.Event{})
	if err != nil {
		logger.Errorf("Error when migrating db model")
		logger.Error(err)
	}
}

func GetDB() *gorm.DB { return db }

func HandleError(err error) (bool, string) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return false, "Record not found"
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return false, "Duplicate key"
	default:
		return false, fmt.Sprintf("Unhandled error: %s", err.Error())
	}
}
