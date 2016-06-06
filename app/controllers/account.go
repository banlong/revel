package controllers
import (
	"github.com/revel/revel"
	"revel/app/models"
	"fmt"
)


type Accounts struct{
	*revel.Controller
}

func (c Accounts) CreateAccount() revel.Result {
	return c.Render()
}

func (c Accounts) CreatePost() revel.Result {
	var account models.Account
	c.Params.Bind(&account, "account")  //"account is the name of data coming in
	fmt.Printf("Account Info:, %v\n", account)
	return c.RenderTemplate("accounts/createAccount.html")
}