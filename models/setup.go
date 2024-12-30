package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Adicione a senha correta na string DSN
	dsn := "host=localhost user=postgres password=go_blog dbname=go_blog port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migração automática para criar a tabela com base no modelo
	database.AutoMigrate(&Post{})

	DB = database
}
