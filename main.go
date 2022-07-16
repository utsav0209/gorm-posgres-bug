package main

import (
	"fmt"
	"github.com/utsav0209/gorm-postgres-bug/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Demonstrate Save method bug in gorm

	postgresHost := "localhost"
	postgresPort := "5433"
	postgresDB := "postgres_bug_check"

	postgresDSN := fmt.Sprintf("host=%v port=%v user=postgres dbname=%v password=postgres sslmode=disable", postgresHost, postgresPort, postgresDB)

	DBSession, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Error while connecting to database")
	}

	// Auto Migrate User table
	_ = DBSession.AutoMigrate(&models.User{})

	// Add users
	tx := DBSession.Begin()
	tx = tx.Save(&models.User{UserId: 1, Name: "User 1"})
	tx.Commit()

	// Add another user with same ID different name
	tx = DBSession.Begin()
	tx = tx.Save(&models.User{UserId: 1, Name: "User 2"})
	tx.Commit()

}
