package controller
//3

import (
	"net/http"

	"github.com/Dankei/conta-click/usecase"
	"github.com/gin-gonic/gin"
)

type ClickController struct {
	clickUseCase usecase.ClickUsecase
}

func NewClickController(usecase usecase.ClickUsecase) *ClickController {
	return &ClickController{
		clickUseCase: usecase,
	}
}

func (c *ClickController) GetClick(ctx *gin.Context) {
	click,_ := c.clickUseCase.GetClick()
	ctx.JSON(http.StatusOK,click)
}

func (c *ClickController) UpdateClick(ctx *gin.Context) {
	click,_ := c.clickUseCase.UpdateClick()
	ctx.JSON(http.StatusOK,click)
}

func (c *ClickController) RestartClick(ctx *gin.Context) {
	click,_ := c.clickUseCase.RestartClick()
	ctx.JSON(http.StatusOK,click)
}