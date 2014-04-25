package controllers

import (
	"github.com/revel/revel"
)

func (c App) Forums() revel.Result {
	return c.Render()
}
