package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/vansimke/restaurant-site/app/models"
)

type Accounts struct {
	*revel.Controller
}

func (c Accounts) Create() revel.Result {
	return c.Render()
}

func (c Accounts) CreatePost() revel.Result {
	var account models.Account
	c.Params.Bind(&account, "account")

	fmt.Printf("Account info: %v\n", account)
	return c.RenderTemplate("accounts/create.html")
}
