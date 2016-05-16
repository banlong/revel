package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/revel/revel"
	"github.com/vansimke/restaurant-site/app/models"
)

type Orders struct {
	*revel.Controller
}

func (c Orders) Create() revel.Result {
	return c.Render()
}

func (c Orders) GetPayment(orderId int) revel.Result {
	println("The order ID: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	var order models.Order
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&order)
	fmt.Printf("The order data: %v\n", order)
	return c.RenderText("OK")
}
