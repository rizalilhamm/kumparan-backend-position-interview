package databases

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../kumparan_articles.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to SQLite database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database handle from GORM")
	}

	// Optional: Limit connections (not critical for SQLite)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
