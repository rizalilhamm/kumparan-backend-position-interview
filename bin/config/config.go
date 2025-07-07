package config

import (
	"crypto/rsa"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	RootApp    string
	HTTPPort   uint16
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey


	PostgreSQL struct {
		Host         string
		User         string
		Password     string
		DBName       string
		Port         uint16
		SSLMode      string
		MaxIdleConns int
		MaxOpenConns int
		MaxLifeTime  int
	}
}

// GlobalEnv global environment
var GlobalEnv Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	rootApp := strings.TrimSuffix(path, "/bin/config")
	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp
	GlobalEnv.HTTPPort = 8080

	// loadPostgreSQL()
}
