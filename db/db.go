package db

import (
	"fmt"
	"os"

	"github.com/sobhan/tod/entitys"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func NewDatabase() (*Database, error) {
	password := os.Getenv("PGPASSWORD")
	if password == "" {
		return nil, fmt.Errorf("PGPASSWORD must be set")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envOrDefault("PGHOST", "localhost"),
		envOrDefault("PGUSER", "postgres"),
		password,
		envOrDefault("PGDATABASE", "todo"),
		envOrDefault("PGPORT", "5432"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	err = db.AutoMigrate(&entitys.TodoLists{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}


	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
