package db

import (
	"fmt"

	"github.com/sobhan/tod/entitys"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

const (
	user     = "postgres"
	password = "Ala.13495782"
	dbname   = "todo"
	port     = "5432"
	host     = "localhost"
)

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

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
