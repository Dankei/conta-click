package main

import (
	"github.com/Dankei/conta-click/controller"
	"github.com/Dankei/conta-click/db"
	"github.com/Dankei/conta-click/repository"
	"github.com/Dankei/conta-click/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	dbConnection, err := db.ConnectDB()
	if(err != nil) {
		panic(err)
	}
	defer dbConnection.Close()

	ClickRepository := repository.NewClickRepository(dbConnection)
	ClickUsecase := usecase.NewClickUsecase(*ClickRepository)
	ClickController := controller.NewClickController(*ClickUsecase)
	server.GET("/click", ClickController.GetClick)

	server.Run(":8080")
}
