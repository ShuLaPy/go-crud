package main

import (
	"github.com/ShuLaPy/go-crud/initializers"
	"github.com/ShuLaPy/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
