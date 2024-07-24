package router
//2


import (
	"github.com/gin-gonic/gin"
	"github.com/Dankei/conta-click/controller"
)

func InitRoutes(r *gin.RouterGroup,
	clickController controller.ClickController) {

	r.GET("/click", clickController.GetClick)
	r.PUT("/click", clickController.UpdateClick)
	r.PUT("/restart", clickController.RestartClick)
}