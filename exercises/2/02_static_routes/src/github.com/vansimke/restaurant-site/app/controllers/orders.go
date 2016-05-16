package controllers

import (
	"github.com/revel/revel"
)

type Orders struct {
	*revel.Controller
}

func (c Orders) Create() revel.Result {
	return c.Render()
}
