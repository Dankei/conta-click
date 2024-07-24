package usecase
//4

import (
	"github.com/Dankei/conta-click/model"
	"github.com/Dankei/conta-click/repository"
)

type ClickUsecase struct {
	repository repository.ClickRepository
}

func NewClickUsecase(repo repository.ClickRepository) *ClickUsecase {
	return &ClickUsecase{
		repository: repo,
	}
}

func (c *ClickUsecase) GetClick() (model.Click,error) {
	return c.repository.GetClick()
}


func (c *ClickUsecase) UpdateClick() (model.Click,error) {
	return c.repository.UpdateClick()
}

func (c *ClickUsecase) RestartClick() (model.Click,error) {
	return c.repository.RestartClick()
}