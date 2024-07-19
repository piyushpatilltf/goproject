package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	errEnv:=godotenv.Load();
	if errEnv != nil {
        log.Fatalf("Error loading .env file")
    }
	dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
	dbName:=os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")
	dbPort:=os.Getenv("DB_PORT")
	dbSsl:=os.Getenv("DB_SSL")
	
	dsn := "host="+dbHost+" user="+dbUser+" password="+dbPassword+" dbname="+dbName+" port="+dbPort+" sslmode="+dbSsl
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