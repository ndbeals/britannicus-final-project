package controllers

import (
	"fmt"
	"strconv"

	"github.com/ndbeals/brittanicus-final-project/forms"
	"github.com/ndbeals/brittanicus-final-project/models"

	"github.com/gin-gonic/gin"
)

//OrderController ...
type OrderController struct{}

var orderModel = models.GetOrderModel()

//CreateOrder ...
func (ctrl OrderController) CreateOrder(c *gin.Context) {
	var createOrderForm *forms.CreateOrder
	fmt.Printf("\n\nOrder DATA: %s \n\n", c.PostForm("customer_id"))

	if c.BindJSON(&createOrderForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid Post form", "form": createOrderForm})
		c.Abort()
		return
	}

	order, err = orderModel.Create(createOrderForm)

	fmt.Printf("\n\nOrder POST DATA: %+v \n\n", createOrderForm)
}

//GetOne ...
func (ctrl OrderController) GetOne(c *gin.Context) {
	orderid := c.Param("id")

	if orderid, err := strconv.ParseInt(orderid, 10, 32); err == nil {
		orderid := int(orderid)

		data, err := orderModel.GetOne(orderid)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//GetList ...
func (ctrl OrderController) GetList(c *gin.Context) {
	page := c.Param("page")
	amount, err := strconv.ParseInt(c.Param("amount"), 10, 32)

	if err != nil {
		amount = 100
	}

	if page, err := strconv.ParseInt(page, 10, 32); err == nil {
		page, amount := int(page), int(amount)
		data, err := orderModel.GetList(page, amount)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"Message": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
