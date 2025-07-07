package main

import (
	"fmt"
	"log"

	"kumparan-backend-position-interview/bin/config"
	articles "kumparan-backend-position-interview/bin/modules/articles/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := createConnection()
	if err != nil {
		panic("Failed to connect database postgre")
	}

	log.Println("Migration: START")
	if err := db.AutoMigrate(ModelTables...); err != nil {
		panic("Migration: " + err.Error())
	}
	log.Println("Migration: SUCCESS")
}

var ModelTables []interface{} = []interface{}{
	&articles.Articles{},
}

func createConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgreSQL.Host,
		config.GlobalEnv.PostgreSQL.User,
		config.GlobalEnv.PostgreSQL.Password,
		config.GlobalEnv.PostgreSQL.DBName,
		config.GlobalEnv.PostgreSQL.Port,
		config.GlobalEnv.PostgreSQL.SSLMode,
	)
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}
