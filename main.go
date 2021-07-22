package main

import (
	"github.com/gin-gonic/gin"

	"go-api/controllers"
	"go-api/services"
)

func main() {
	r := gin.Default()

	services.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)

	r.POST("/books", controllers.CreateBook)

	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
