package controllers

import (
	"github.com/revel/revel"
)

func (c App) Forums() revel.Result {
	return c.Render()
}

func ForumMessageIdPost(id string) revel.Result {
	return c.Render()
}
