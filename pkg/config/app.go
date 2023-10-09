package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {

	absPath, _ := filepath.Abs("./credentials.env")
	err := godotenv.Load(absPath)
	if err != nil {
		panic(err.Error())
	}

	dsn := os.Getenv("DB_DSN")

	d, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
