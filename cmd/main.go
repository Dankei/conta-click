package main
// fluxo das execuções 1

import (
	"github.com/Dankei/conta-click/controller"
	"github.com/Dankei/conta-click/db"
	"github.com/Dankei/conta-click/repository"
	"github.com/Dankei/conta-click/router"
	"github.com/Dankei/conta-click/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	ClickRepository := repository.NewClickRepository(dbConnection)
	ClickUsecase := usecase.NewClickUsecase(*ClickRepository)
	ClickController := controller.NewClickController(*ClickUsecase)
	router.InitRoutes(&server.RouterGroup, *ClickController)
	server.Run(":8080")
}
