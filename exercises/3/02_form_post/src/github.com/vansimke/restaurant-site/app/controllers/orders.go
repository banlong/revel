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

func (c Orders) GetPayment(orderId int) revel.Result {
	println("The order ID: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}
