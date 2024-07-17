// config/database.go
package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=temp1 port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	}
	fmt.Println("Database connection successfully opened")
}


func Migrate(models ...interface{}) {
	for _, model := range models {
		DB.AutoMigrate(model)
	}
}