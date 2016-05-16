package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Order() revel.Result {
	return c.Render()
}

func (c App) CreateAccount() revel.Result {
	return c.RenderTemplate("App/create_account.html")
}

func (c App) Payment() revel.Result {
	return c.Render()
}
