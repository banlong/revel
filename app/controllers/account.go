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

	c.Validation.Required(account.FirstName)
	c.Validation.Required(account.LastName)
	c.Validation.Required(account.Address1)
	c.Validation.Required(account.City)
	c.Validation.Required(account.State)
	c.Validation.Required(account.ZipCode)
	c.Validation.Length(account.ZipCode, 5)

	//Keep the data in the data context for later use
	c.Validation.Keep()
	//Make the flash params available to the client
	c.FlashParams()

	fmt.Printf("Has Error: %v\n", c.Validation.HasErrors())

	fmt.Printf("Account Info:, %v\n", account)
	return c.RenderTemplate("accounts/createAccount.html")
}