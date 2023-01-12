package main

import (
	"github.com/ShuLaPy/go-crud/controllers"
	"github.com/ShuLaPy/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/", controllers.PostsCreate)
	r.Run()
}
