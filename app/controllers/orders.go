package controllers
import (
	"github.com/revel/revel"
	"revel/app/models"
	"encoding/json"
	"fmt"
)

type Orders struct {
	*revel.Controller
}

func(c Orders) Create() revel.Result  {
	//look for a template  which is views/order/create.html
	return c.Render()
}

func (c Orders) GetPayment(orderId int) revel.Result  {
	println("The orderId:", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	var order models.Order
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&order)
	fmt.Printf("The order data: %v\n", order)
	return c.RenderText("OK")
}
