package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBMYSQL *gorm.DB

func ConnectMySQL() {
	log.Println("Connecting to MySQL")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, errPing := db.DB()
			if errPing == nil && sqlDB.Ping() == nil {
				log.Println("✅ Connected to MySQL")
				DBMYSQL = db
				return
			}
		}

		log.Printf("⏳ Attempt %d: Waiting for MySQL... (%v)", i, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ Failed to connect to MySQL after multiple attempts: %v", err)
}
