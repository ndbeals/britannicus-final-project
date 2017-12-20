package controllers

import (
	"strconv"

	"github.com/ndbeals/brittanicus-final-project/models"

	"github.com/gin-gonic/gin"
)

//OrderController ...
type OrderController struct{}

var OrderModel = new(models.OrderModel)

//GetOne ...
func (ctrl OrderController) GetOne(c *gin.Context) {
	// OrderID := getOrderID(c)

	// if OrderID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	uid := c.Param("id")

	if uid, err := strconv.ParseInt(uid, 10, 64); err == nil {

		data, err := OrderModel.GetOne(uid)
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

//GetOne ...
func (ctrl OrderController) GetSet(c *gin.Context) {
	// OrderID := getOrderID(c)

	// if OrderID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	page := c.Param("page")
	amount, err := strconv.ParseInt(c.Param("amount"), 10, 64)

	if err != nil {
		amount = 100
	}

	if page, err := strconv.ParseInt(page, 10, 64); err == nil {
		data, err := OrderModel.GetSet(page, amount)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
