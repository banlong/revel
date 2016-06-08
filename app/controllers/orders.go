package controllers
import (
	"github.com/revel/revel"
	"revel/app/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var rndGen = rand.New(rand.NewSource(time.Now().Unix()))
const chance  = 0.5

type Orders struct {
	*revel.Controller
	winner bool
}

func init()  {
	//Init the intercept
	revel.InterceptMethod((*Orders).randomDrawing, revel.BEFORE)
}

func(c Orders) Create() revel.Result  {
	//look for a template  which is views/order/create.html
	return c.Render()
}

func (c *Orders) randomDrawing() revel.Result {
	if(rndGen.Float32()< chance){
		c.winner = true
	}
	//must return nil in order this result not sending back to the client
	return nil
}

func (c Orders) GetPayment(orderId int) revel.Result  {
	println("The orderId:", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	var order models.Order
	if c.winner{
		println("You have won a discount")
	}
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&order)
	fmt.Printf("The order data: %v\n", order)
	return c.RenderText("OK")
}
