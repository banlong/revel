package controllers
import "github.com/revel/revel"


type Accounts struct{
	*revel.Controller
}

func (c Accounts) CreateAccount() revel.Result {
	return c.Render()
}

