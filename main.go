package main

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/routers"
)

func main() {
	config.Connect()
	config.ConnectToRedis()
	config.Migrate(&models.Log{})

	r := routers.SetupRouter()
	r.Run(":8080")
}
