package main

import (
	"github.com/gin-gonic/gin"
	"mkdika.id/gin-rest-api/controllers"
	"mkdika.id/gin-rest-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
