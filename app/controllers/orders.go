package controllers
import "github.com/revel/revel"

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