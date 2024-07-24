package repository
//5

import (
	"database/sql"

	"github.com/Dankei/conta-click/model"
)

type ClickRepository struct {
	connectuin *sql.DB
}

func NewClickRepository(db *sql.DB) *ClickRepository {
	return &ClickRepository{
		connectuin: db,
	}
}

func (c *ClickRepository) GetClick() (model.Click, error) {
	// Correção do nome da variável de "connectuin" para "connection"
	row := c.connectuin.QueryRow("SELECT * FROM click WHERE id = 1")

	var click model.Click

	// Correção na chamada de Scan, aplicada ao resultado de QueryRow
	err := row.Scan(&click.ID,&click.Name, &click.Value)
	if err != nil {
		return model.Click{}, err
	}

	return click, nil
}


func (c *ClickRepository) UpdateClick() (model.Click, error) {
	// Correção do nome da variável de "connectuin" para "connection"
	_, err := c.connectuin.Exec("UPDATE click SET value = value + 1 WHERE id = 1")
	if err != nil {
		return model.Click{}, err
	}

	// Correção na chamada de GetClick
	return c.GetClick()
}

func (c *ClickRepository) RestartClick() (model.Click, error) {
	// Correção do nome da variável de "connectuin" para "connection"
	_, err := c.connectuin.Exec("UPDATE click SET value = 0 WHERE id = 1")
	if err != nil {
		return model.Click{}, err
	}

	// Correção na chamada de GetClick
	return c.GetClick()
}