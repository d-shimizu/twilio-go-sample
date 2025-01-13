package main

import (
	"fmt"
	"log"
	"os"

	"localhost/twilio-go-sample/infra/database/migrate"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// データベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーションの実行
	if err := migrate.Migrate(db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migration completed successfully")

}
