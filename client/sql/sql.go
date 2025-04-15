package sql

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error

	sqlite := sqlite.New(sqlite.Config{
		DSN: "file:db-app.db",
	})

	db, err = gorm.Open(sqlite, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *gorm.DB {
	return db
}
