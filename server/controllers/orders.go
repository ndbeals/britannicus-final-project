package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ndbeals/britannicus-final-project/forms"
	"github.com/ndbeals/britannicus-final-project/models"

	"github.com/gin-gonic/gin"
)

//OrderController ...
type OrderController struct{}

var orderModel = models.GetOrderModel()

//GetOne ...
func (ctrl OrderController) GetOne(c *gin.Context) {
	orderid := c.Param("id")

	if orderid, err := strconv.ParseInt(orderid, 10, 32); err == nil {
		orderid := int(orderid)

		data, err := orderModel.GetOne(orderid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, data)
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
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
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, data)
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//CreateOrder ...
func (ctrl OrderController) CreateOrder(c *gin.Context) {
	var createOrderForm forms.CreateOrderForm
	// fmt.Printf("\n\nOrder DATA: %s \n\n", c.PostForm("customer_id"))

	if c.BindJSON(&createOrderForm) != nil {
		c.IndentedJSON(406, gin.H{"message": "Invalid Post form", "form": createOrderForm})
		c.Abort()
		return
	}

	order, err := orderModel.Create(createOrderForm)

	// fmt.Printf("\n\nOrder POST DATA: %+v \n\n", order)
	if err != nil {
		c.IndentedJSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if order.ID > 0 {
		// session := sessions.Default(c)
		// session.Set("user_id", user.ID)
		// session.Set("user_email", user.Email)
		// session.Set("user_name", user.Name)
		// session.Save()
		c.IndentedJSON(200, gin.H{"message": "Successfully created order", "data": order})
	} else {
		// c.IndentedJSON(406, gin.H{"message": "Could not create order", "error": err.Error()})
	}

}

//Update ...
func (ctrl OrderController) Update(c *gin.Context) {
	orderid := c.Param("id")
	fmt.Println(orderid)

	if orderid, err := strconv.ParseInt(orderid, 10, 32); err == nil {
		orderid := int(orderid)

		var updateForm forms.UpdateOrderForm
		err := c.BindJSON(&updateForm)
		if err != nil {
			// panic(err)
			c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
			c.Abort()
			return
		}

		order, err := orderModel.GetOne(orderid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Order not found", "error": err.Error()})
			c.Abort()
			return
		}

		// order.Update(updateForm)

		c.IndentedJSON(200, gin.H{"data": order})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

// //Create ...
// func (ctrl OrderController) Create(c *gin.Context) {
// 	var updateForm forms.UpdateOrderForm
// 	err := c.BindJSON(&updateForm)
// 	if err != nil {
// 		// panic(err)
// 		c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
// 		c.Abort()
// 		return
// 	}

// 	product, _ := productModel.GetOne(updateForm.ProductID)

// 	order := models.Order{-1, product, 4, "Good", updateForm.Amount, updateForm.Price, updateForm.Note}
// 	if err != nil {
// 		c.IndentedJSON(404, gin.H{"Message": "Order not found", "error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	newid, _ := order.Create()

// 	c.IndentedJSON(200, gin.H{"data": order, "id": newid})
// }

//Delete ...
func (ctrl OrderController) Delete(c *gin.Context) {
	orderid := c.Param("id")
	fmt.Println("Delete")

	if orderid, err := strconv.ParseInt(orderid, 10, 32); err == nil {
		orderid := int(orderid)

		order, err := orderModel.GetOne(orderid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Order not found", "error": err.Error()})
			c.Abort()
			return
		}

		_, err = order.Delete()

		fmt.Println("deleted order from api")

		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Failed to delete", "error": err.Error()})
			c.Abort()
			return
		}

		c.IndentedJSON(200, gin.H{"Message": "Order Sucessfully deleted"})
	}
}

//OrderListingPage ...
func (ctrl OrderController) OrderListingPage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "orders.html", gin.H{
		"title": "Order Page",
		"route": "/orders",
		"user":  user,
	})
}

//OrderDetailPage ...
func (ctrl OrderController) OrderDetailPage(c *gin.Context) {
	orderid := c.Param("id")

	if orderid, err := strconv.ParseInt(orderid, 10, 32); err == nil {
		orderid := int(orderid)
		order, _ := orderModel.GetOne(orderid)

		user, _ := GetLoggedinUser(c)

		c.HTML(http.StatusOK, "order.html", gin.H{
			"title":   "Order Detail Page",
			"route":   "/order/get",
			"user":    user,
			"order":   order,
			"orderid": orderid,
		})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//OrderCreatePage ...
func (ctrl OrderController) OrderCreatePage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "neworder.html", gin.H{
		"title": "Create Order Page",
		"route": "/order/create",
		"user":  user,
	})
}
